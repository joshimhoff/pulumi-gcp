// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Google Kubernetes Engine (GKE) cluster. For more information see
// [the official documentation](https://cloud.google.com/container-engine/docs/clusters)
// and [the API reference](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters).
// 
// > **Note:** All arguments and attributes, including basic auth username and
// passwords as well as certificate outputs will be stored in the raw state as
// plaintext. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
type Cluster struct {
	s *pulumi.ResourceState
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["additionalZones"] = nil
		inputs["addonsConfig"] = nil
		inputs["clusterAutoscaling"] = nil
		inputs["clusterIpv4Cidr"] = nil
		inputs["defaultMaxPodsPerNode"] = nil
		inputs["description"] = nil
		inputs["enableBinaryAuthorization"] = nil
		inputs["enableKubernetesAlpha"] = nil
		inputs["enableLegacyAbac"] = nil
		inputs["enableTpu"] = nil
		inputs["initialNodeCount"] = nil
		inputs["ipAllocationPolicy"] = nil
		inputs["location"] = nil
		inputs["loggingService"] = nil
		inputs["maintenancePolicy"] = nil
		inputs["masterAuth"] = nil
		inputs["masterAuthorizedNetworksConfig"] = nil
		inputs["minMasterVersion"] = nil
		inputs["monitoringService"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["networkPolicy"] = nil
		inputs["nodeConfig"] = nil
		inputs["nodeLocations"] = nil
		inputs["nodePools"] = nil
		inputs["nodeVersion"] = nil
		inputs["podSecurityPolicyConfig"] = nil
		inputs["privateClusterConfig"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["removeDefaultNodePool"] = nil
		inputs["resourceLabels"] = nil
		inputs["subnetwork"] = nil
		inputs["zone"] = nil
	} else {
		inputs["additionalZones"] = args.AdditionalZones
		inputs["addonsConfig"] = args.AddonsConfig
		inputs["clusterAutoscaling"] = args.ClusterAutoscaling
		inputs["clusterIpv4Cidr"] = args.ClusterIpv4Cidr
		inputs["defaultMaxPodsPerNode"] = args.DefaultMaxPodsPerNode
		inputs["description"] = args.Description
		inputs["enableBinaryAuthorization"] = args.EnableBinaryAuthorization
		inputs["enableKubernetesAlpha"] = args.EnableKubernetesAlpha
		inputs["enableLegacyAbac"] = args.EnableLegacyAbac
		inputs["enableTpu"] = args.EnableTpu
		inputs["initialNodeCount"] = args.InitialNodeCount
		inputs["ipAllocationPolicy"] = args.IpAllocationPolicy
		inputs["location"] = args.Location
		inputs["loggingService"] = args.LoggingService
		inputs["maintenancePolicy"] = args.MaintenancePolicy
		inputs["masterAuth"] = args.MasterAuth
		inputs["masterAuthorizedNetworksConfig"] = args.MasterAuthorizedNetworksConfig
		inputs["minMasterVersion"] = args.MinMasterVersion
		inputs["monitoringService"] = args.MonitoringService
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["networkPolicy"] = args.NetworkPolicy
		inputs["nodeConfig"] = args.NodeConfig
		inputs["nodeLocations"] = args.NodeLocations
		inputs["nodePools"] = args.NodePools
		inputs["nodeVersion"] = args.NodeVersion
		inputs["podSecurityPolicyConfig"] = args.PodSecurityPolicyConfig
		inputs["privateClusterConfig"] = args.PrivateClusterConfig
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["removeDefaultNodePool"] = args.RemoveDefaultNodePool
		inputs["resourceLabels"] = args.ResourceLabels
		inputs["subnetwork"] = args.Subnetwork
		inputs["zone"] = args.Zone
	}
	inputs["endpoint"] = nil
	inputs["instanceGroupUrls"] = nil
	inputs["masterVersion"] = nil
	inputs["tpuIpv4CidrBlock"] = nil
	s, err := ctx.RegisterResource("gcp:container/cluster:Cluster", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterState, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["additionalZones"] = state.AdditionalZones
		inputs["addonsConfig"] = state.AddonsConfig
		inputs["clusterAutoscaling"] = state.ClusterAutoscaling
		inputs["clusterIpv4Cidr"] = state.ClusterIpv4Cidr
		inputs["defaultMaxPodsPerNode"] = state.DefaultMaxPodsPerNode
		inputs["description"] = state.Description
		inputs["enableBinaryAuthorization"] = state.EnableBinaryAuthorization
		inputs["enableKubernetesAlpha"] = state.EnableKubernetesAlpha
		inputs["enableLegacyAbac"] = state.EnableLegacyAbac
		inputs["enableTpu"] = state.EnableTpu
		inputs["endpoint"] = state.Endpoint
		inputs["initialNodeCount"] = state.InitialNodeCount
		inputs["instanceGroupUrls"] = state.InstanceGroupUrls
		inputs["ipAllocationPolicy"] = state.IpAllocationPolicy
		inputs["location"] = state.Location
		inputs["loggingService"] = state.LoggingService
		inputs["maintenancePolicy"] = state.MaintenancePolicy
		inputs["masterAuth"] = state.MasterAuth
		inputs["masterAuthorizedNetworksConfig"] = state.MasterAuthorizedNetworksConfig
		inputs["masterVersion"] = state.MasterVersion
		inputs["minMasterVersion"] = state.MinMasterVersion
		inputs["monitoringService"] = state.MonitoringService
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["networkPolicy"] = state.NetworkPolicy
		inputs["nodeConfig"] = state.NodeConfig
		inputs["nodeLocations"] = state.NodeLocations
		inputs["nodePools"] = state.NodePools
		inputs["nodeVersion"] = state.NodeVersion
		inputs["podSecurityPolicyConfig"] = state.PodSecurityPolicyConfig
		inputs["privateClusterConfig"] = state.PrivateClusterConfig
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["removeDefaultNodePool"] = state.RemoveDefaultNodePool
		inputs["resourceLabels"] = state.ResourceLabels
		inputs["subnetwork"] = state.Subnetwork
		inputs["tpuIpv4CidrBlock"] = state.TpuIpv4CidrBlock
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:container/cluster:Cluster", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Cluster) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Cluster) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The list of zones in which the cluster's nodes
// should be located. These must be in the same region as the cluster zone for
// zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
// the number of nodes specified in `initial_node_count` is created in
// all specified zones as well as the primary zone. If specified for a regional
// cluster, nodes will only be created in these zones. `additional_zones` has been
// deprecated in favour of `node_locations`.
func (r *Cluster) AdditionalZones() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["additionalZones"])
}

// The configuration for addons supported by GKE.
// Structure is documented below.
func (r *Cluster) AddonsConfig() *pulumi.Output {
	return r.s.State["addonsConfig"]
}

// )
// Configuration for cluster autoscaling (also called autoprovisioning), as described in
// [the docs](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning).
// Structure is documented below.
func (r *Cluster) ClusterAutoscaling() *pulumi.Output {
	return r.s.State["clusterAutoscaling"]
}

// The IP address range of the kubernetes pods in
// this cluster. Default is an automatically assigned CIDR.
func (r *Cluster) ClusterIpv4Cidr() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterIpv4Cidr"])
}

func (r *Cluster) DefaultMaxPodsPerNode() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["defaultMaxPodsPerNode"])
}

// Description of the cluster.
func (r *Cluster) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// ) Enable Binary Authorization for this cluster.
// If enabled, all container images will be validated by Google Binary Authorization.
func (r *Cluster) EnableBinaryAuthorization() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableBinaryAuthorization"])
}

// Whether to enable Kubernetes Alpha features for
// this cluster. Note that when this option is enabled, the cluster cannot be upgraded
// and will be automatically deleted after 30 days.
func (r *Cluster) EnableKubernetesAlpha() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableKubernetesAlpha"])
}

// Whether the ABAC authorizer is enabled for this cluster.
// When enabled, identities in the system, including service accounts, nodes, and controllers,
// will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
// Defaults to `false`
func (r *Cluster) EnableLegacyAbac() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableLegacyAbac"])
}

// ) Whether to enable Cloud TPU resources in this cluster.
// See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).
func (r *Cluster) EnableTpu() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableTpu"])
}

// The IP address of this cluster's Kubernetes master.
func (r *Cluster) Endpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpoint"])
}

// The number of nodes to create in this
// cluster's default node pool. Must be set if `node_pool` is not set. If
// you're using `google_container_node_pool` objects with no default node pool,
// you'll need to set this to a value of at least `1`, alongside setting
// `remove_default_node_pool` to `true`.
func (r *Cluster) InitialNodeCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["initialNodeCount"])
}

// List of instance group URLs which have been assigned
// to the cluster.
func (r *Cluster) InstanceGroupUrls() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["instanceGroupUrls"])
}

// Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
// This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
// Structure is documented below.
func (r *Cluster) IpAllocationPolicy() *pulumi.Output {
	return r.s.State["ipAllocationPolicy"]
}

// The location (region or zone) in which the cluster
// master will be created, as well as the default node location. If you specify a
// zone (such as `us-central1-a`), the cluster will be a zonal cluster with a
// single cluster master. If you specify a region (such as `us-west1`), the
// cluster will be a regional cluster with multiple masters spread across zones in
// the region, and with default node locations in those zones as well.
func (r *Cluster) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The logging service that the cluster should
// write logs to. Available options include `logging.googleapis.com`,
// `logging.googleapis.com/kubernetes` (beta), and `none`. Defaults to `logging.googleapis.com`
func (r *Cluster) LoggingService() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["loggingService"])
}

// The maintenance policy to use for the cluster. Structure is
// documented below.
func (r *Cluster) MaintenancePolicy() *pulumi.Output {
	return r.s.State["maintenancePolicy"]
}

// The authentication information for accessing the
// Kubernetes master. Structure is documented below.
func (r *Cluster) MasterAuth() *pulumi.Output {
	return r.s.State["masterAuth"]
}

// The desired configuration options
// for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
// external access (except the cluster node IPs, which GKE automatically whitelists).
func (r *Cluster) MasterAuthorizedNetworksConfig() *pulumi.Output {
	return r.s.State["masterAuthorizedNetworksConfig"]
}

// The current version of the master in the cluster. This may
// be different than the `min_master_version` set in the config if the master
// has been updated by GKE.
func (r *Cluster) MasterVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterVersion"])
}

// The minimum version of the master. GKE
// will auto-update the master to new versions, so this does not guarantee the
// current master version--use the read-only `master_version` field to obtain that.
// If unset, the cluster's version will be set by GKE to the version of the most recent
// official release (which is not necessarily the latest version).  Most users will find
// the `google_container_engine_versions` data source useful - it indicates which versions
// are available, and can be use to approximate fuzzy versions in a
// Terraform-compatible way. If you intend to specify versions manually,
// [the docs](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#specifying_cluster_version)
// describe the various acceptable formats for this field.
func (r *Cluster) MinMasterVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["minMasterVersion"])
}

// The monitoring service that the cluster
// should write metrics to.
// Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
// VM metrics will be collected by Google Compute Engine regardless of this setting
// Available options include
// `monitoring.googleapis.com`, `monitoring.googleapis.com/kubernetes` (beta) and `none`.
// Defaults to `monitoring.googleapis.com`
func (r *Cluster) MonitoringService() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["monitoringService"])
}

// The name of the cluster, unique within the project and
// location.
func (r *Cluster) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name or self_link of the Google Compute Engine
// network to which the cluster is connected. For Shared VPC, set this to the self link of the
// shared network.
func (r *Cluster) Network() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["network"])
}

// Configuration options for the
// [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
// feature. Structure is documented below.
func (r *Cluster) NetworkPolicy() *pulumi.Output {
	return r.s.State["networkPolicy"]
}

// Parameters used in creating the default node pool.
// Generally, this field should not be used at the same time as a
// `google_container_node_pool` or a `node_pool` block; this configuration
// manages the default node pool, which isn't recommended to be used with
// Terraform. Structure is documented below.
func (r *Cluster) NodeConfig() *pulumi.Output {
	return r.s.State["nodeConfig"]
}

// The list of zones in which the cluster's nodes
// should be located. These must be in the same region as the cluster zone for
// zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
// the number of nodes specified in `initial_node_count` is created in
// all specified zones as well as the primary zone. If specified for a regional
// cluster, nodes will be created in only these zones.
func (r *Cluster) NodeLocations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["nodeLocations"])
}

// List of node pools associated with this cluster.
// See google_container_node_pool for schema.
// **Warning:** node pools defined inside a cluster can't be changed (or added/removed) after
// cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability
// to say "these are the _only_ node pools associated with this cluster", use the
// google_container_node_pool resource instead of this property.
func (r *Cluster) NodePools() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["nodePools"])
}

// The Kubernetes version on the nodes. Must either be unset
// or set to the same value as `min_master_version` on create. Defaults to the default
// version set by GKE which is not necessarily the latest version. This only affects
// nodes in the default node pool. While a fuzzy version can be specified, it's
// recommended that you specify explicit versions as Terraform will see spurious diffs
// when fuzzy versions are used. See the `google_container_engine_versions` data source's
// `version_prefix` field to approximate fuzzy versions in a Terraform-compatible way.
// To update nodes in other node pools, use the `version` attribute on the node pool.
func (r *Cluster) NodeVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["nodeVersion"])
}

// ) Configuration for the
// [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
// Structure is documented below.
func (r *Cluster) PodSecurityPolicyConfig() *pulumi.Output {
	return r.s.State["podSecurityPolicyConfig"]
}

// A set of options for creating
// a private cluster. Structure is documented below.
func (r *Cluster) PrivateClusterConfig() *pulumi.Output {
	return r.s.State["privateClusterConfig"]
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *Cluster) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *Cluster) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// If `true`, deletes the default node
// pool upon cluster creation. If you're using `google_container_node_pool`
// resources with no default node pool, this should be set to `true`, alongside
// setting `initial_node_count` to at least `1`.
func (r *Cluster) RemoveDefaultNodePool() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["removeDefaultNodePool"])
}

// The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
func (r *Cluster) ResourceLabels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["resourceLabels"])
}

// The name or self_link of the Google Compute Engine subnetwork in
// which the cluster's instances are launched.
func (r *Cluster) Subnetwork() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetwork"])
}

// ([Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The IP address range of the Cloud TPUs in this cluster, in
// [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
// notation (e.g. `1.2.3.4/29`).
func (r *Cluster) TpuIpv4CidrBlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tpuIpv4CidrBlock"])
}

// The zone that the cluster master and nodes
// should be created in. If specified, this cluster will be a zonal cluster. `zone`
// has been deprecated in favour of `location`.
func (r *Cluster) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering Cluster resources.
type ClusterState struct {
	// The list of zones in which the cluster's nodes
	// should be located. These must be in the same region as the cluster zone for
	// zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
	// the number of nodes specified in `initial_node_count` is created in
	// all specified zones as well as the primary zone. If specified for a regional
	// cluster, nodes will only be created in these zones. `additional_zones` has been
	// deprecated in favour of `node_locations`.
	AdditionalZones interface{}
	// The configuration for addons supported by GKE.
	// Structure is documented below.
	AddonsConfig interface{}
	// )
	// Configuration for cluster autoscaling (also called autoprovisioning), as described in
	// [the docs](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning).
	// Structure is documented below.
	ClusterAutoscaling interface{}
	// The IP address range of the kubernetes pods in
	// this cluster. Default is an automatically assigned CIDR.
	ClusterIpv4Cidr interface{}
	DefaultMaxPodsPerNode interface{}
	// Description of the cluster.
	Description interface{}
	// ) Enable Binary Authorization for this cluster.
	// If enabled, all container images will be validated by Google Binary Authorization.
	EnableBinaryAuthorization interface{}
	// Whether to enable Kubernetes Alpha features for
	// this cluster. Note that when this option is enabled, the cluster cannot be upgraded
	// and will be automatically deleted after 30 days.
	EnableKubernetesAlpha interface{}
	// Whether the ABAC authorizer is enabled for this cluster.
	// When enabled, identities in the system, including service accounts, nodes, and controllers,
	// will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
	// Defaults to `false`
	EnableLegacyAbac interface{}
	// ) Whether to enable Cloud TPU resources in this cluster.
	// See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).
	EnableTpu interface{}
	// The IP address of this cluster's Kubernetes master.
	Endpoint interface{}
	// The number of nodes to create in this
	// cluster's default node pool. Must be set if `node_pool` is not set. If
	// you're using `google_container_node_pool` objects with no default node pool,
	// you'll need to set this to a value of at least `1`, alongside setting
	// `remove_default_node_pool` to `true`.
	InitialNodeCount interface{}
	// List of instance group URLs which have been assigned
	// to the cluster.
	InstanceGroupUrls interface{}
	// Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
	// This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
	// Structure is documented below.
	IpAllocationPolicy interface{}
	// The location (region or zone) in which the cluster
	// master will be created, as well as the default node location. If you specify a
	// zone (such as `us-central1-a`), the cluster will be a zonal cluster with a
	// single cluster master. If you specify a region (such as `us-west1`), the
	// cluster will be a regional cluster with multiple masters spread across zones in
	// the region, and with default node locations in those zones as well.
	Location interface{}
	// The logging service that the cluster should
	// write logs to. Available options include `logging.googleapis.com`,
	// `logging.googleapis.com/kubernetes` (beta), and `none`. Defaults to `logging.googleapis.com`
	LoggingService interface{}
	// The maintenance policy to use for the cluster. Structure is
	// documented below.
	MaintenancePolicy interface{}
	// The authentication information for accessing the
	// Kubernetes master. Structure is documented below.
	MasterAuth interface{}
	// The desired configuration options
	// for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
	// external access (except the cluster node IPs, which GKE automatically whitelists).
	MasterAuthorizedNetworksConfig interface{}
	// The current version of the master in the cluster. This may
	// be different than the `min_master_version` set in the config if the master
	// has been updated by GKE.
	MasterVersion interface{}
	// The minimum version of the master. GKE
	// will auto-update the master to new versions, so this does not guarantee the
	// current master version--use the read-only `master_version` field to obtain that.
	// If unset, the cluster's version will be set by GKE to the version of the most recent
	// official release (which is not necessarily the latest version).  Most users will find
	// the `google_container_engine_versions` data source useful - it indicates which versions
	// are available, and can be use to approximate fuzzy versions in a
	// Terraform-compatible way. If you intend to specify versions manually,
	// [the docs](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#specifying_cluster_version)
	// describe the various acceptable formats for this field.
	MinMasterVersion interface{}
	// The monitoring service that the cluster
	// should write metrics to.
	// Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
	// VM metrics will be collected by Google Compute Engine regardless of this setting
	// Available options include
	// `monitoring.googleapis.com`, `monitoring.googleapis.com/kubernetes` (beta) and `none`.
	// Defaults to `monitoring.googleapis.com`
	MonitoringService interface{}
	// The name of the cluster, unique within the project and
	// location.
	Name interface{}
	// The name or self_link of the Google Compute Engine
	// network to which the cluster is connected. For Shared VPC, set this to the self link of the
	// shared network.
	Network interface{}
	// Configuration options for the
	// [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
	// feature. Structure is documented below.
	NetworkPolicy interface{}
	// Parameters used in creating the default node pool.
	// Generally, this field should not be used at the same time as a
	// `google_container_node_pool` or a `node_pool` block; this configuration
	// manages the default node pool, which isn't recommended to be used with
	// Terraform. Structure is documented below.
	NodeConfig interface{}
	// The list of zones in which the cluster's nodes
	// should be located. These must be in the same region as the cluster zone for
	// zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
	// the number of nodes specified in `initial_node_count` is created in
	// all specified zones as well as the primary zone. If specified for a regional
	// cluster, nodes will be created in only these zones.
	NodeLocations interface{}
	// List of node pools associated with this cluster.
	// See google_container_node_pool for schema.
	// **Warning:** node pools defined inside a cluster can't be changed (or added/removed) after
	// cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability
	// to say "these are the _only_ node pools associated with this cluster", use the
	// google_container_node_pool resource instead of this property.
	NodePools interface{}
	// The Kubernetes version on the nodes. Must either be unset
	// or set to the same value as `min_master_version` on create. Defaults to the default
	// version set by GKE which is not necessarily the latest version. This only affects
	// nodes in the default node pool. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as Terraform will see spurious diffs
	// when fuzzy versions are used. See the `google_container_engine_versions` data source's
	// `version_prefix` field to approximate fuzzy versions in a Terraform-compatible way.
	// To update nodes in other node pools, use the `version` attribute on the node pool.
	NodeVersion interface{}
	// ) Configuration for the
	// [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
	// Structure is documented below.
	PodSecurityPolicyConfig interface{}
	// A set of options for creating
	// a private cluster. Structure is documented below.
	PrivateClusterConfig interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	// If `true`, deletes the default node
	// pool upon cluster creation. If you're using `google_container_node_pool`
	// resources with no default node pool, this should be set to `true`, alongside
	// setting `initial_node_count` to at least `1`.
	RemoveDefaultNodePool interface{}
	// The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
	ResourceLabels interface{}
	// The name or self_link of the Google Compute Engine subnetwork in
	// which the cluster's instances are launched.
	Subnetwork interface{}
	// ([Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The IP address range of the Cloud TPUs in this cluster, in
	// [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	// notation (e.g. `1.2.3.4/29`).
	TpuIpv4CidrBlock interface{}
	// The zone that the cluster master and nodes
	// should be created in. If specified, this cluster will be a zonal cluster. `zone`
	// has been deprecated in favour of `location`.
	Zone interface{}
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The list of zones in which the cluster's nodes
	// should be located. These must be in the same region as the cluster zone for
	// zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
	// the number of nodes specified in `initial_node_count` is created in
	// all specified zones as well as the primary zone. If specified for a regional
	// cluster, nodes will only be created in these zones. `additional_zones` has been
	// deprecated in favour of `node_locations`.
	AdditionalZones interface{}
	// The configuration for addons supported by GKE.
	// Structure is documented below.
	AddonsConfig interface{}
	// )
	// Configuration for cluster autoscaling (also called autoprovisioning), as described in
	// [the docs](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning).
	// Structure is documented below.
	ClusterAutoscaling interface{}
	// The IP address range of the kubernetes pods in
	// this cluster. Default is an automatically assigned CIDR.
	ClusterIpv4Cidr interface{}
	DefaultMaxPodsPerNode interface{}
	// Description of the cluster.
	Description interface{}
	// ) Enable Binary Authorization for this cluster.
	// If enabled, all container images will be validated by Google Binary Authorization.
	EnableBinaryAuthorization interface{}
	// Whether to enable Kubernetes Alpha features for
	// this cluster. Note that when this option is enabled, the cluster cannot be upgraded
	// and will be automatically deleted after 30 days.
	EnableKubernetesAlpha interface{}
	// Whether the ABAC authorizer is enabled for this cluster.
	// When enabled, identities in the system, including service accounts, nodes, and controllers,
	// will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
	// Defaults to `false`
	EnableLegacyAbac interface{}
	// ) Whether to enable Cloud TPU resources in this cluster.
	// See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).
	EnableTpu interface{}
	// The number of nodes to create in this
	// cluster's default node pool. Must be set if `node_pool` is not set. If
	// you're using `google_container_node_pool` objects with no default node pool,
	// you'll need to set this to a value of at least `1`, alongside setting
	// `remove_default_node_pool` to `true`.
	InitialNodeCount interface{}
	// Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
	// This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
	// Structure is documented below.
	IpAllocationPolicy interface{}
	// The location (region or zone) in which the cluster
	// master will be created, as well as the default node location. If you specify a
	// zone (such as `us-central1-a`), the cluster will be a zonal cluster with a
	// single cluster master. If you specify a region (such as `us-west1`), the
	// cluster will be a regional cluster with multiple masters spread across zones in
	// the region, and with default node locations in those zones as well.
	Location interface{}
	// The logging service that the cluster should
	// write logs to. Available options include `logging.googleapis.com`,
	// `logging.googleapis.com/kubernetes` (beta), and `none`. Defaults to `logging.googleapis.com`
	LoggingService interface{}
	// The maintenance policy to use for the cluster. Structure is
	// documented below.
	MaintenancePolicy interface{}
	// The authentication information for accessing the
	// Kubernetes master. Structure is documented below.
	MasterAuth interface{}
	// The desired configuration options
	// for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
	// external access (except the cluster node IPs, which GKE automatically whitelists).
	MasterAuthorizedNetworksConfig interface{}
	// The minimum version of the master. GKE
	// will auto-update the master to new versions, so this does not guarantee the
	// current master version--use the read-only `master_version` field to obtain that.
	// If unset, the cluster's version will be set by GKE to the version of the most recent
	// official release (which is not necessarily the latest version).  Most users will find
	// the `google_container_engine_versions` data source useful - it indicates which versions
	// are available, and can be use to approximate fuzzy versions in a
	// Terraform-compatible way. If you intend to specify versions manually,
	// [the docs](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#specifying_cluster_version)
	// describe the various acceptable formats for this field.
	MinMasterVersion interface{}
	// The monitoring service that the cluster
	// should write metrics to.
	// Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
	// VM metrics will be collected by Google Compute Engine regardless of this setting
	// Available options include
	// `monitoring.googleapis.com`, `monitoring.googleapis.com/kubernetes` (beta) and `none`.
	// Defaults to `monitoring.googleapis.com`
	MonitoringService interface{}
	// The name of the cluster, unique within the project and
	// location.
	Name interface{}
	// The name or self_link of the Google Compute Engine
	// network to which the cluster is connected. For Shared VPC, set this to the self link of the
	// shared network.
	Network interface{}
	// Configuration options for the
	// [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
	// feature. Structure is documented below.
	NetworkPolicy interface{}
	// Parameters used in creating the default node pool.
	// Generally, this field should not be used at the same time as a
	// `google_container_node_pool` or a `node_pool` block; this configuration
	// manages the default node pool, which isn't recommended to be used with
	// Terraform. Structure is documented below.
	NodeConfig interface{}
	// The list of zones in which the cluster's nodes
	// should be located. These must be in the same region as the cluster zone for
	// zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
	// the number of nodes specified in `initial_node_count` is created in
	// all specified zones as well as the primary zone. If specified for a regional
	// cluster, nodes will be created in only these zones.
	NodeLocations interface{}
	// List of node pools associated with this cluster.
	// See google_container_node_pool for schema.
	// **Warning:** node pools defined inside a cluster can't be changed (or added/removed) after
	// cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability
	// to say "these are the _only_ node pools associated with this cluster", use the
	// google_container_node_pool resource instead of this property.
	NodePools interface{}
	// The Kubernetes version on the nodes. Must either be unset
	// or set to the same value as `min_master_version` on create. Defaults to the default
	// version set by GKE which is not necessarily the latest version. This only affects
	// nodes in the default node pool. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as Terraform will see spurious diffs
	// when fuzzy versions are used. See the `google_container_engine_versions` data source's
	// `version_prefix` field to approximate fuzzy versions in a Terraform-compatible way.
	// To update nodes in other node pools, use the `version` attribute on the node pool.
	NodeVersion interface{}
	// ) Configuration for the
	// [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
	// Structure is documented below.
	PodSecurityPolicyConfig interface{}
	// A set of options for creating
	// a private cluster. Structure is documented below.
	PrivateClusterConfig interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	// If `true`, deletes the default node
	// pool upon cluster creation. If you're using `google_container_node_pool`
	// resources with no default node pool, this should be set to `true`, alongside
	// setting `initial_node_count` to at least `1`.
	RemoveDefaultNodePool interface{}
	// The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
	ResourceLabels interface{}
	// The name or self_link of the Google Compute Engine subnetwork in
	// which the cluster's instances are launched.
	Subnetwork interface{}
	// The zone that the cluster master and nodes
	// should be created in. If specified, this cluster will be a zonal cluster. `zone`
	// has been deprecated in favour of `location`.
	Zone interface{}
}
