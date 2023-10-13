/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConstraintsObservation struct {

	// List of acceptable instance CPU architectures, the default is amd64. Allowed values: amd64, arm64.
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// Compute optimized instance constraint - will only pick compute optimized nodes if true.
	ComputeOptimized *bool `json:"computeOptimized,omitempty" tf:"compute_optimized,omitempty"`

	// Fallback restore rate in seconds: defines how much time should pass before spot fallback should be attempted to be restored to real spot.
	FallbackRestoreRateSeconds *float64 `json:"fallbackRestoreRateSeconds,omitempty" tf:"fallback_restore_rate_seconds,omitempty"`

	Gpu []GpuObservation `json:"gpu,omitempty" tf:"gpu,omitempty"`

	InstanceFamilies []InstanceFamiliesObservation `json:"instanceFamilies,omitempty" tf:"instance_families,omitempty"`

	// Max CPU cores per node.
	MaxCPU *float64 `json:"maxCpu,omitempty" tf:"max_cpu,omitempty"`

	// Max Memory (Mib) per node.
	MaxMemory *float64 `json:"maxMemory,omitempty" tf:"max_memory,omitempty"`

	// Min CPU cores per node.
	MinCPU *float64 `json:"minCpu,omitempty" tf:"min_cpu,omitempty"`

	// Min Memory (Mib) per node.
	MinMemory *float64 `json:"minMemory,omitempty" tf:"min_memory,omitempty"`

	// Spot instance constraint - true only spot, false only on-demand.
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`

	// Storage optimized instance constraint - will only pick storage optimized nodes if true
	StorageOptimized *bool `json:"storageOptimized,omitempty" tf:"storage_optimized,omitempty"`

	// Spot instance fallback constraint - when true, on-demand instances will be created, when spots are unavailable.
	UseSpotFallbacks *bool `json:"useSpotFallbacks,omitempty" tf:"use_spot_fallbacks,omitempty"`
}

type ConstraintsParameters struct {

	// List of acceptable instance CPU architectures, the default is amd64. Allowed values: amd64, arm64.
	// +kubebuilder:validation:Optional
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// Compute optimized instance constraint - will only pick compute optimized nodes if true.
	// +kubebuilder:validation:Optional
	ComputeOptimized *bool `json:"computeOptimized,omitempty" tf:"compute_optimized,omitempty"`

	// Fallback restore rate in seconds: defines how much time should pass before spot fallback should be attempted to be restored to real spot.
	// +kubebuilder:validation:Optional
	FallbackRestoreRateSeconds *float64 `json:"fallbackRestoreRateSeconds,omitempty" tf:"fallback_restore_rate_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	Gpu []GpuParameters `json:"gpu,omitempty" tf:"gpu,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceFamilies []InstanceFamiliesParameters `json:"instanceFamilies,omitempty" tf:"instance_families,omitempty"`

	// Max CPU cores per node.
	// +kubebuilder:validation:Optional
	MaxCPU *float64 `json:"maxCpu,omitempty" tf:"max_cpu,omitempty"`

	// Max Memory (Mib) per node.
	// +kubebuilder:validation:Optional
	MaxMemory *float64 `json:"maxMemory,omitempty" tf:"max_memory,omitempty"`

	// Min CPU cores per node.
	// +kubebuilder:validation:Optional
	MinCPU *float64 `json:"minCpu,omitempty" tf:"min_cpu,omitempty"`

	// Min Memory (Mib) per node.
	// +kubebuilder:validation:Optional
	MinMemory *float64 `json:"minMemory,omitempty" tf:"min_memory,omitempty"`

	// Spot instance constraint - true only spot, false only on-demand.
	// +kubebuilder:validation:Optional
	Spot *bool `json:"spot,omitempty" tf:"spot,omitempty"`

	// Storage optimized instance constraint - will only pick storage optimized nodes if true
	// +kubebuilder:validation:Optional
	StorageOptimized *bool `json:"storageOptimized,omitempty" tf:"storage_optimized,omitempty"`

	// Spot instance fallback constraint - when true, on-demand instances will be created, when spots are unavailable.
	// +kubebuilder:validation:Optional
	UseSpotFallbacks *bool `json:"useSpotFallbacks,omitempty" tf:"use_spot_fallbacks,omitempty"`
}

type CustomLabelObservation struct {

	// Label key to be added to nodes created from this template.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Label value to be added to nodes created from this template.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomLabelParameters struct {

	// Label key to be added to nodes created from this template.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Label value to be added to nodes created from this template.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type CustomTaintsObservation struct {

	// Effect of a taint to be added to nodes created from this template. The effect must always be NoSchedule.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// Key of a taint to be added to nodes created from this template.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Value of a taint to be added to nodes created from this template.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomTaintsParameters struct {

	// Effect of a taint to be added to nodes created from this template. The effect must always be NoSchedule.
	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// Key of a taint to be added to nodes created from this template.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Value of a taint to be added to nodes created from this template.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type GpuObservation struct {

	// Names of the GPUs to exclude.
	ExcludeNames []*string `json:"excludeNames,omitempty" tf:"exclude_names,omitempty"`

	// Instance families to include when filtering (excludes all other families).
	IncludeNames []*string `json:"includeNames,omitempty" tf:"include_names,omitempty"`

	// Manufacturers of the gpus to select - NVIDIA, AMD.
	Manufacturers []*string `json:"manufacturers,omitempty" tf:"manufacturers,omitempty"`

	// Max GPU count for the instance type to have.
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// Min GPU count for the instance type to have.
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`
}

type GpuParameters struct {

	// Names of the GPUs to exclude.
	// +kubebuilder:validation:Optional
	ExcludeNames []*string `json:"excludeNames,omitempty" tf:"exclude_names,omitempty"`

	// Instance families to include when filtering (excludes all other families).
	// +kubebuilder:validation:Optional
	IncludeNames []*string `json:"includeNames,omitempty" tf:"include_names,omitempty"`

	// Manufacturers of the gpus to select - NVIDIA, AMD.
	// +kubebuilder:validation:Optional
	Manufacturers []*string `json:"manufacturers,omitempty" tf:"manufacturers,omitempty"`

	// Max GPU count for the instance type to have.
	// +kubebuilder:validation:Optional
	MaxCount *float64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// Min GPU count for the instance type to have.
	// +kubebuilder:validation:Optional
	MinCount *float64 `json:"minCount,omitempty" tf:"min_count,omitempty"`
}

type InstanceFamiliesObservation struct {

	// Instance families to include when filtering (excludes all other families).
	Exclude []*string `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// Instance families to exclude when filtering (includes all other families).
	Include []*string `json:"include,omitempty" tf:"include,omitempty"`
}

type InstanceFamiliesParameters struct {

	// Instance families to include when filtering (excludes all other families).
	// +kubebuilder:validation:Optional
	Exclude []*string `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// Instance families to exclude when filtering (includes all other families).
	// +kubebuilder:validation:Optional
	Include []*string `json:"include,omitempty" tf:"include,omitempty"`
}

type NodeTemplateObservation struct {

	// CAST AI cluster id.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// CAST AI node configuration id to be used for node template.
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	Constraints []ConstraintsObservation `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Marks whether custom instances should be used when deciding which parts of inventory are available. Custom instances are only supported in GCP.
	CustomInstancesEnabled *bool `json:"customInstancesEnabled,omitempty" tf:"custom_instances_enabled,omitempty"`

	// Custom label key/value to be added to nodes created from this template.
	CustomLabel []CustomLabelObservation `json:"customLabel,omitempty" tf:"custom_label,omitempty"`

	// Custom labels to be added to nodes created from this template. If the field `custom_label` is present, the value of `custom_labels` will be ignored.
	CustomLabels map[string]*string `json:"customLabels,omitempty" tf:"custom_labels,omitempty"`

	// Custom taints to be added to the nodes created from this template. `shouldTaint` has to be `true` in order to create/update the node template with custom taints. If `shouldTaint` is `true`, but no custom taints are provided, the nodes will be tainted with the default node template taint.
	CustomTaints []CustomTaintsObservation `json:"customTaints,omitempty" tf:"custom_taints,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the node template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Minimum nodes that will be kept when rebalancing nodes using this node template.
	RebalancingConfigMinNodes *float64 `json:"rebalancingConfigMinNodes,omitempty" tf:"rebalancing_config_min_nodes,omitempty"`

	// Marks whether the templated nodes will have a taint.
	ShouldTaint *bool `json:"shouldTaint,omitempty" tf:"should_taint,omitempty"`
}

type NodeTemplateParameters struct {

	// CAST AI cluster id.
	// +crossplane:generate:reference:type=github.com/castai/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// CAST AI node configuration id to be used for node template.
	// +crossplane:generate:reference:type=github.com/castai/crossplane-provider-castai/apis/castai/v1alpha1.NodeConfiguration
	// +kubebuilder:validation:Optional
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	// Reference to a NodeConfiguration in castai to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDRef *v1.Reference `json:"configurationIdRef,omitempty" tf:"-"`

	// Selector for a NodeConfiguration in castai to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDSelector *v1.Selector `json:"configurationIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Constraints []ConstraintsParameters `json:"constraints,omitempty" tf:"constraints,omitempty"`

	// Marks whether custom instances should be used when deciding which parts of inventory are available. Custom instances are only supported in GCP.
	// +kubebuilder:validation:Optional
	CustomInstancesEnabled *bool `json:"customInstancesEnabled,omitempty" tf:"custom_instances_enabled,omitempty"`

	// Custom label key/value to be added to nodes created from this template.
	// +kubebuilder:validation:Optional
	CustomLabel []CustomLabelParameters `json:"customLabel,omitempty" tf:"custom_label,omitempty"`

	// Custom labels to be added to nodes created from this template. If the field `custom_label` is present, the value of `custom_labels` will be ignored.
	// +kubebuilder:validation:Optional
	CustomLabels map[string]*string `json:"customLabels,omitempty" tf:"custom_labels,omitempty"`

	// Custom taints to be added to the nodes created from this template. `shouldTaint` has to be `true` in order to create/update the node template with custom taints. If `shouldTaint` is `true`, but no custom taints are provided, the nodes will be tainted with the default node template taint.
	// +kubebuilder:validation:Optional
	CustomTaints []CustomTaintsParameters `json:"customTaints,omitempty" tf:"custom_taints,omitempty"`

	// Name of the node template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Minimum nodes that will be kept when rebalancing nodes using this node template.
	// +kubebuilder:validation:Optional
	RebalancingConfigMinNodes *float64 `json:"rebalancingConfigMinNodes,omitempty" tf:"rebalancing_config_min_nodes,omitempty"`

	// Marks whether the templated nodes will have a taint.
	// +kubebuilder:validation:Optional
	ShouldTaint *bool `json:"shouldTaint,omitempty" tf:"should_taint,omitempty"`
}

// NodeTemplateSpec defines the desired state of NodeTemplate
type NodeTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeTemplateParameters `json:"forProvider"`
}

// NodeTemplateStatus defines the observed state of NodeTemplate.
type NodeTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodeTemplate is the Schema for the NodeTemplates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type NodeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   NodeTemplateSpec   `json:"spec"`
	Status NodeTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeTemplateList contains a list of NodeTemplates
type NodeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeTemplate `json:"items"`
}

// Repository type metadata.
var (
	NodeTemplate_Kind             = "NodeTemplate"
	NodeTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeTemplate_Kind}.String()
	NodeTemplate_KindAPIVersion   = NodeTemplate_Kind + "." + CRDGroupVersion.String()
	NodeTemplate_GroupVersionKind = CRDGroupVersion.WithKind(NodeTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeTemplate{}, &NodeTemplateList{})
}
