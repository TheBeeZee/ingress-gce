/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"math"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

const (
	isMultinetwork                = "true"
	notMultinetwork               = "false"
	enabledStrongSessionAffinity  = "true"
	disabledStrongSessionAffinity = "false"
	isWeightedLBPodsPerNode       = "true"
	notWeightedLBPodsPerNode      = "false"
)

func TestExportILBMetric(t *testing.T) {
	newMetrics := FakeControllerMetrics()
	pastPersistentErrorThresholdTime := time.Now().Add(-1*persistentErrorThresholdTime - time.Minute)
	notExceedingPersistentErrorThresholdTime := time.Now().Add(-1*persistentErrorThresholdTime + 5*time.Minute)

	newMetrics.SetL4ILBService("svc-success-all-labels", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: true, WeightedLBPodsPerNode: true},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4ILBService("svc-success-multinet-1", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: true, WeightedLBPodsPerNode: false},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4ILBService("svc-success-multinet-2", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: true, WeightedLBPodsPerNode: false},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4ILBService("svc-success-weightedlb", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, WeightedLBPodsPerNode: true},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4ILBService("svc-user-error", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, WeightedLBPodsPerNode: false},
		Status:                  StatusUserError,
	})
	newMetrics.SetL4ILBService("svc-error", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, WeightedLBPodsPerNode: false},
		Status:                  StatusError,
		FirstSyncErrorTime:      &notExceedingPersistentErrorThresholdTime,
	})
	newMetrics.SetL4ILBService("svc-threshold-check", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, WeightedLBPodsPerNode: false},
		Status:                  StatusError,
		FirstSyncErrorTime:      &pastPersistentErrorThresholdTime,
	})
	// check that updating later does not move FirstSyncErrorTime
	newMetrics.SetL4ILBService("svc-threshold-check", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, WeightedLBPodsPerNode: false},
		Status:                  StatusError,
		FirstSyncErrorTime:      &notExceedingPersistentErrorThresholdTime,
	})
	newMetrics.SetL4ILBService("svc-single-stack", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, WeightedLBPodsPerNode: false},
		L4DualStackServiceLabels: L4DualStackServiceLabels{
			IPFamilies:     "IPv4",
			IPFamilyPolicy: "SingleStack",
		},
		Status:             StatusError,
		FirstSyncErrorTime: &notExceedingPersistentErrorThresholdTime,
	})

	newMetrics.exportL4ILBsMetrics()

	verifyL4ILBMetric(t, 1, StatusSuccess, isMultinetwork, isWeightedLBPodsPerNode)
	verifyL4ILBMetric(t, 2, StatusSuccess, isMultinetwork, notWeightedLBPodsPerNode)
	verifyL4ILBMetric(t, 1, StatusSuccess, notMultinetwork, isWeightedLBPodsPerNode)
	verifyL4ILBMetric(t, 1, StatusUserError, notMultinetwork, notWeightedLBPodsPerNode)
	verifyL4ILBMetric(t, 2, StatusError, notMultinetwork, notWeightedLBPodsPerNode)
	verifyL4ILBMetric(t, 1, StatusPersistentError, notMultinetwork, notWeightedLBPodsPerNode)
}

func verifyL4ILBMetric(t *testing.T, expectedCount int, status L4ServiceStatus, multinet string, weightedLBPodsPerNode string) {
	countFloat := testutil.ToFloat64(l4ILBCount.With(prometheus.Labels{l4LabelStatus: string(status), l4LabelMultinet: multinet, l4LabelWeightedLBPodsPerNode: weightedLBPodsPerNode}))
	actualCount := int(math.Round(countFloat))
	if expectedCount != actualCount {
		t.Errorf("expected value %d but got %d", expectedCount, actualCount)
	}
}

func TestExportNetLBMetric(t *testing.T) {
	newMetrics := FakeControllerMetrics()
	pastPersistentErrorThresholdTime := time.Now().Add(-1*persistentErrorThresholdTime - time.Minute)
	notExceedingPersistentErrorThresholdTime := time.Now().Add(-1*persistentErrorThresholdTime + 5*time.Minute)

	newMetrics.SetL4NetLBService("svc-success-multinet-1", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: true, StrongSessionAffinity: false, WeightedLBPodsPerNode: false},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4NetLBService("svc-success-all-labels", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: true, StrongSessionAffinity: true, WeightedLBPodsPerNode: true},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4NetLBService("svc-success-multinet-2", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: true, StrongSessionAffinity: false, WeightedLBPodsPerNode: false},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4NetLBService("svc-success-ssa", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, StrongSessionAffinity: true, WeightedLBPodsPerNode: false},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4NetLBService("svc-success-weightedlb", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, StrongSessionAffinity: false, WeightedLBPodsPerNode: true},
		Status:                  StatusSuccess,
	})
	newMetrics.SetL4NetLBService("svc-user-error-ssa", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, StrongSessionAffinity: true, WeightedLBPodsPerNode: false},
		Status:                  StatusUserError,
	})
	newMetrics.SetL4NetLBService("svc-error", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, StrongSessionAffinity: false, WeightedLBPodsPerNode: false},
		Status:                  StatusError,
		FirstSyncErrorTime:      &notExceedingPersistentErrorThresholdTime,
	})
	newMetrics.SetL4NetLBService("svc-error", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, StrongSessionAffinity: false, WeightedLBPodsPerNode: false},
		Status:                  StatusError,
		FirstSyncErrorTime:      &notExceedingPersistentErrorThresholdTime,
	})
	newMetrics.SetL4NetLBService("svc-threshold-check", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, StrongSessionAffinity: false, WeightedLBPodsPerNode: false},
		Status:                  StatusError,
		FirstSyncErrorTime:      &pastPersistentErrorThresholdTime,
	})
	// check that updating later does not move FirstSyncErrorTime
	newMetrics.SetL4NetLBService("svc-threshold-check", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, StrongSessionAffinity: false, WeightedLBPodsPerNode: false},
		Status:                  StatusError,
		FirstSyncErrorTime:      &notExceedingPersistentErrorThresholdTime,
	})
	newMetrics.SetL4NetLBService("svc-single-stack", L4ServiceState{
		L4FeaturesServiceLabels: L4FeaturesServiceLabels{Multinetwork: false, StrongSessionAffinity: false, WeightedLBPodsPerNode: false},
		L4DualStackServiceLabels: L4DualStackServiceLabels{
			IPFamilies:     "IPv4",
			IPFamilyPolicy: "SingleStack",
		},
		Status:             StatusError,
		FirstSyncErrorTime: &notExceedingPersistentErrorThresholdTime,
	})

	newMetrics.exportL4NetLBsMetrics()

	verifyL4NetLBMetric(t, 2, StatusSuccess, isMultinetwork, disabledStrongSessionAffinity, notWeightedLBPodsPerNode)
	verifyL4NetLBMetric(t, 1, StatusSuccess, isMultinetwork, enabledStrongSessionAffinity, isWeightedLBPodsPerNode)
	verifyL4NetLBMetric(t, 1, StatusSuccess, notMultinetwork, enabledStrongSessionAffinity, notWeightedLBPodsPerNode)
	verifyL4NetLBMetric(t, 1, StatusSuccess, notMultinetwork, disabledStrongSessionAffinity, isWeightedLBPodsPerNode)
	verifyL4NetLBMetric(t, 1, StatusUserError, notMultinetwork, enabledStrongSessionAffinity, notWeightedLBPodsPerNode)
	verifyL4NetLBMetric(t, 2, StatusError, notMultinetwork, disabledStrongSessionAffinity, notWeightedLBPodsPerNode)
	verifyL4NetLBMetric(t, 1, StatusPersistentError, notMultinetwork, disabledStrongSessionAffinity, notWeightedLBPodsPerNode)
}

func verifyL4NetLBMetric(t *testing.T, expectedCount int, status L4ServiceStatus, multinet string, strongSessionAffinity string, weightedLBPodsPerNode string) {
	countFloat := testutil.ToFloat64(l4NetLBCount.With(prometheus.Labels{l4LabelStatus: string(status), l4LabelMultinet: multinet, l4LabelStrongSessionAffinity: strongSessionAffinity, l4LabelWeightedLBPodsPerNode: weightedLBPodsPerNode}))
	actualCount := int(math.Round(countFloat))
	if expectedCount != actualCount {
		t.Errorf("expected value %d but got %d", expectedCount, actualCount)
	}
}
