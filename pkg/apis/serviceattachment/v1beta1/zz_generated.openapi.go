//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ConsumerForwardingRule":  schema_pkg_apis_serviceattachment_v1beta1_ConsumerForwardingRule(ref),
		"k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ConsumerProject":         schema_pkg_apis_serviceattachment_v1beta1_ConsumerProject(ref),
		"k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ServiceAttachment":       schema_pkg_apis_serviceattachment_v1beta1_ServiceAttachment(ref),
		"k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ServiceAttachmentSpec":   schema_pkg_apis_serviceattachment_v1beta1_ServiceAttachmentSpec(ref),
		"k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ServiceAttachmentStatus": schema_pkg_apis_serviceattachment_v1beta1_ServiceAttachmentStatus(ref),
	}
}

func schema_pkg_apis_serviceattachment_v1beta1_ConsumerForwardingRule(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConsumerForwardingRule is a reference to the PSC consumer forwarding rule",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"forwardingRuleURL": {
						SchemaProps: spec.SchemaProps{
							Description: "Forwarding rule consumer created to use ServiceAttachment",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of consumer forwarding rule",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_serviceattachment_v1beta1_ConsumerProject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConsumerProject is the consumer project and project level configuration",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"connectionLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "ConnectionLimit is the connection limit for this Consumer project",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"project": {
						SchemaProps: spec.SchemaProps{
							Description: "Project is the project id or number for the project to set the limit for.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"forceSendFields": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "ForceSendFields is a list of field names (e.g. \"ConnectionLimit\") to unconditionally include in API requests. By default, fields with empty values are omitted from API requests. However, any non-pointer, non-interface field appearing in ForceSendFields will be sent to the server regardless of whether the field is empty or not. This may be used to include empty fields in Patch requests.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"nullFields": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "NullFields is a list of field names (e.g. \"ConnectionLimit\") to include in API requests with the JSON null value. By default, fields with empty values are omitted from API requests. However, any field with an empty value appearing in NullFields will be sent to the server as null. It is an error if a field in this list has a non-empty value. This may be used to include null fields in Patch requests.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_serviceattachment_v1beta1_ServiceAttachment(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceAttachment represents a Service Attachment associated with a service/ingress/gateway class",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ServiceAttachmentSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ServiceAttachmentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ServiceAttachmentSpec", "k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ServiceAttachmentStatus"},
	}
}

func schema_pkg_apis_serviceattachment_v1beta1_ServiceAttachmentSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceAttachmentSpec is the spec for a ServiceAttachment resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"connectionPreference": {
						SchemaProps: spec.SchemaProps{
							Description: "ConnectionPreference determines how consumers are accepted.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"natSubnets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "NATSubnets contains the list of subnet names for PSC or nat subnet resource urls",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"resourceRef": {
						SchemaProps: spec.SchemaProps{
							Description: "ResourceRef is the reference to the K8s resource that created the forwarding rule Only Services can be used as a reference",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/api/core/v1.TypedLocalObjectReference"),
						},
					},
					"proxyProtocol": {
						SchemaProps: spec.SchemaProps{
							Description: "ProxyProtocol when set will expose client information TCP/IP information",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"consumerAllowList": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "ConsumerAllowList is list of consumer projects that should be allow listed for this ServiceAttachment",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ConsumerProject"),
									},
								},
							},
						},
					},
					"consumerRejectList": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "ConsumerRejectList is the list of Consumer Project IDs or Numbers that should be rejected for this ServiceAttachment",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.TypedLocalObjectReference", "k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ConsumerProject"},
	}
}

func schema_pkg_apis_serviceattachment_v1beta1_ServiceAttachmentStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceAttachmentStatus is the status for a ServiceAttachment resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"serviceAttachmentURL": {
						SchemaProps: spec.SchemaProps{
							Description: "ServiceAttachmentURL is the URL for the GCE Service Attachment resource",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"forwardingRuleURL": {
						SchemaProps: spec.SchemaProps{
							Description: "ForwardingRuleURL is the URL to the GCE Forwarding Rule resource the Service Attachment points to",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"consumerForwardingRules": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Consumer Forwarding Rules using ts Service Attachment",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ConsumerForwardingRule"),
									},
								},
							},
						},
					},
					"lastModifiedTimestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "LastModifiedTimestamp tracks last time Status was updated",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time", "k8s.io/ingress-gce/pkg/apis/serviceattachment/v1beta1.ConsumerForwardingRule"},
	}
}
