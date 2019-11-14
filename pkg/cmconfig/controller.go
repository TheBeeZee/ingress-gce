package cmconfig

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/ingress-gce/pkg/utils"
	"k8s.io/klog"
)

// ConfigMapConfigController is the ConfigMap based config controller.
// If cmConfigModeEnabled set to true, it will load the config from configmap: configMapNamespace/configMapName and restart ingress controller if the config has any ligeal changes.
// If cmConfigModeEnabled set to false, it will return the default values for the configs.
type ConfigMapConfigController struct {
	configMapNamespace     string
	configMapName          string
	currentConfig          *Config
	currentConfigMapObject *v1.ConfigMap
	kubeClient             kubernetes.Interface
	recorder               record.EventRecorder
}

// NewConfigMapConfigController creates a new ConfigMapConfigController, it will load the config from the target configmap
func NewConfigMapConfigController(kubeClient kubernetes.Interface, recorder record.EventRecorder, configMapNamespace, configMapName string) *ConfigMapConfigController {

	currentConfig := NewConfig()
	cm, err := kubeClient.CoreV1().ConfigMaps(configMapNamespace).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			klog.Infof("ConfigMapConfigController: Not found the configmap based config, using default config: %v", currentConfig)
		} else {
			klog.Warningf("ConfigMapConfigController failed to load config from api server, using the defualt config. Error: %v", err)
		}
	} else {
		if err := currentConfig.LoadValue(cm.Data); err != nil {
			if recorder != nil {
				recorder.Event(cm, "Warning", "LoadValueError", err.Error())
			}
			klog.Warningf("LoadValue error: %s", err.Error())
		}
		klog.Infof("ConfigMapConfigController: loaded config from configmap, config %v", currentConfig)
	}

	c := &ConfigMapConfigController{
		configMapNamespace:     configMapNamespace,
		configMapName:          configMapName,
		currentConfig:          &currentConfig,
		kubeClient:             kubeClient,
		recorder:               recorder,
		currentConfigMapObject: cm,
	}
	return c
}

// GetConfig returns the internal Config
func (c *ConfigMapConfigController) GetConfig() Config {
	return *c.currentConfig
}

// DisableASMMode disables the ASM Mode by updating the ConfigMap and setting the internal flag.
func (c *ConfigMapConfigController) DisableASMMode() {
	patchBytes, err := utils.StrategicMergePatchBytes(v1.ConfigMap{Data: map[string]string{enableASM: trueValue}},
		v1.ConfigMap{Data: map[string]string{enableASM: falseValue}}, v1.ConfigMap{})
	if err != nil {
		c.RecordEvent("Warning", "FailedDisableASMMode", fmt.Sprintf("Failed to disable ASM Mode, failed to create patch for ASM ConfigMap, error: %s", err))
		return
	}
	cm, err := c.kubeClient.CoreV1().ConfigMaps(c.configMapNamespace).Patch(c.configMapName, apimachinerytypes.MergePatchType, patchBytes, "")
	if err != nil {
		c.RecordEvent("Warning", "FailedDisableASMMode", fmt.Sprintf("Failed to patch ASM ConfigMap, error: %s", err))
		return
	}
	c.currentConfigMapObject = cm
	c.currentConfig.EnableASM = false
}

// RecordEvent records a event to the ASMConfigmap
func (c *ConfigMapConfigController) RecordEvent(eventtype, reason, message string) bool {
	if c.recorder == nil || c.currentConfigMapObject == nil {
		return false
	}
	c.recorder.Event(c.currentConfigMapObject, eventtype, reason, message)
	return true
}

// RegisterInformer regjister the configmap based config controller handler to the configapInformer which will watch the target
// configmap and send stop message to the stopCh if any valid change detected.
func (c *ConfigMapConfigController) RegisterInformer(configMapInformer cache.SharedIndexInformer, cancel func()) {
	configMapInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.processItem(obj, cancel)
		},
		DeleteFunc: func(obj interface{}) {
			c.processItem(obj, cancel)
		},
		UpdateFunc: func(_, cur interface{}) {
			c.processItem(cur, cancel)
		},
	})

}

func (c *ConfigMapConfigController) processItem(obj interface{}, cancel func()) {
	configMap, ok := obj.(*v1.ConfigMap)
	if !ok {
		klog.Errorf("ConfigMapConfigController: failed to convert informer object to ConfigMap.")
	}
	if configMap.Namespace != c.configMapNamespace || configMap.Name != c.configMapName {
		return
	}

	config := NewConfig()
	cm, err := c.kubeClient.CoreV1().ConfigMaps(c.configMapNamespace).Get(c.configMapName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			klog.Infof("ConfigMapConfigController: Not found the configmap based config, using default config: %v", config)
		} else {
			klog.Warningf("ConfigMapConfigController failed to load config from api server, using the defualt config. Error: %v", err)
		}
	} else {
		c.currentConfigMapObject = cm
		if err := config.LoadValue(cm.Data); err != nil {
			c.RecordEvent("Warning", "LoadValueError", err.Error())
		}
	}

	if !config.Equals(c.currentConfig) {
		c.RecordEvent("Normal", "ASMConfigMapTiggerRestart", "ConfigMapConfigController: Get a update on the ConfigMapConfig, Restarting Ingress controller")
		cancel()
	}
}
