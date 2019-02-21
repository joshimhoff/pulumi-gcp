# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    additional_zones: pulumi.Output[list]
    """
    The list of additional Google Compute Engine
    locations in which the cluster's nodes should be located. If additional zones are
    configured, the number of nodes specified in `initial_node_count` is created in
    all specified zones.
    """
    addons_config: pulumi.Output[dict]
    """
    The configuration for addons supported by GKE.
    Structure is documented below.
    """
    cluster_ipv4_cidr: pulumi.Output[str]
    """
    The IP address range of the kubernetes pods in
    this cluster. Default is an automatically assigned CIDR.
    """
    description: pulumi.Output[str]
    """
    Description of the cluster.
    """
    enable_binary_authorization: pulumi.Output[bool]
    """
    Enable Binary Authorization for this cluster.
    If enabled, all container images will be validated by Google Binary Authorization.
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    """
    enable_kubernetes_alpha: pulumi.Output[bool]
    """
    Whether to enable Kubernetes Alpha features for
    this cluster. Note that when this option is enabled, the cluster cannot be upgraded
    and will be automatically deleted after 30 days.
    """
    enable_legacy_abac: pulumi.Output[bool]
    """
    Whether the ABAC authorizer is enabled for this cluster.
    When enabled, identities in the system, including service accounts, nodes, and controllers,
    will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
    Defaults to `false`
    """
    enable_tpu: pulumi.Output[bool]
    """
    Whether to enable Cloud TPU resources in this cluster.
    See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    """
    endpoint: pulumi.Output[str]
    """
    The IP address of this cluster's Kubernetes master.
    """
    initial_node_count: pulumi.Output[int]
    """
    The number of nodes to create in this
    cluster (not including the Kubernetes master). Must be set if `node_pool` is not set.
    """
    instance_group_urls: pulumi.Output[list]
    """
    List of instance group URLs which have been assigned
    to the cluster.
    """
    ip_allocation_policy: pulumi.Output[dict]
    """
    Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
    This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
    Structure is documented below.
    """
    logging_service: pulumi.Output[str]
    """
    The logging service that the cluster should
    write logs to. Available options include `logging.googleapis.com`,
    `logging.googleapis.com/kubernetes` (beta), and `none`. Defaults to `logging.googleapis.com`
    """
    maintenance_policy: pulumi.Output[dict]
    """
    The maintenance policy to use for the cluster. Structure is
    documented below.
    """
    master_auth: pulumi.Output[dict]
    """
    The authentication information for accessing the
    Kubernetes master. Structure is documented below.
    """
    master_authorized_networks_config: pulumi.Output[dict]
    """
    The desired configuration options
    for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
    external access (except the cluster node IPs, which GKE automatically whitelists).
    """
    master_ipv4_cidr_block: pulumi.Output[str]
    """
    Specifies a private
    [RFC1918](https://tools.ietf.org/html/rfc1918) block for the master's VPC. The master range must not overlap with any subnet in your cluster's VPC.
    The master and your cluster use VPC peering. Must be specified in CIDR notation and must be `/28` subnet.
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    This field is deprecated, use `private_cluster_config.master_ipv4_cidr_block` instead.
    """
    master_version: pulumi.Output[str]
    """
    The current version of the master in the cluster. This may
    be different than the `min_master_version` set in the config if the master
    has been updated by GKE.
    """
    min_master_version: pulumi.Output[str]
    """
    The minimum version of the master. GKE
    will auto-update the master to new versions, so this does not guarantee the
    current master version--use the read-only `master_version` field to obtain that.
    If unset, the cluster's version will be set by GKE to the version of the most recent
    official release (which is not necessarily the latest version).
    """
    monitoring_service: pulumi.Output[str]
    """
    The monitoring service that the cluster
    should write metrics to.
    Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
    VM metrics will be collected by Google Compute Engine regardless of this setting
    Available options include
    `monitoring.googleapis.com`, `monitoring.googleapis.com/kubernetes` (beta) and `none`.
    Defaults to `monitoring.googleapis.com`
    """
    name: pulumi.Output[str]
    """
    The name of the cluster, unique within the project and
    zone.
    """
    network: pulumi.Output[str]
    """
    The name or self_link of the Google Compute Engine
    network to which the cluster is connected. For Shared VPC, set this to the self link of the
    shared network.
    """
    network_policy: pulumi.Output[dict]
    """
    Configuration options for the
    [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
    feature. Structure is documented below.
    """
    node_config: pulumi.Output[dict]
    """
    Parameters used in creating the cluster's nodes.
    Structure is documented below.
    """
    node_pools: pulumi.Output[list]
    """
    List of node pools associated with this cluster.
    See google_container_node_pool for schema.
    **Warning:** node pools defined inside a cluster can't be changed (or added/removed) after
    cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability
    to say "these are the _only_ node pools associated with this cluster", use the
    google_container_node_pool resource instead of this property.
    """
    node_version: pulumi.Output[str]
    """
    The Kubernetes version on the nodes. Must either be unset
    or set to the same value as `min_master_version` on create. Defaults to the default
    version set by GKE which is not necessarily the latest version.
    """
    pod_security_policy_config: pulumi.Output[dict]
    """
    Configuration for the
    [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
    Structure is documented below.
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    """
    private_cluster: pulumi.Output[bool]
    """
    If true, a
    [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters) will be created, meaning
    nodes do not get public IP addresses. It is mandatory to specify `master_ipv4_cidr_block` and
    `ip_allocation_policy` with this option.
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    This field is deprecated, use `private_cluster_config.enable_private_nodes` instead.
    """
    private_cluster_config: pulumi.Output[dict]
    """
    A set of options for creating
    a private cluster. Structure is documented below.
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    remove_default_node_pool: pulumi.Output[bool]
    """
    If true, deletes the default node pool upon cluster creation.
    """
    resource_labels: pulumi.Output[dict]
    """
    The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
    """
    subnetwork: pulumi.Output[str]
    """
    The name or self_link of the Google Compute Engine subnetwork in
    which the cluster's instances are launched.
    """
    zone: pulumi.Output[str]
    """
    The zone that the master and the number of nodes specified
    in `initial_node_count` should be created in. Only one of `zone` and `region`
    may be set. If neither zone nor region are set, the provider zone is used.
    """
    def __init__(__self__, resource_name, opts=None, additional_zones=None, addons_config=None, cluster_ipv4_cidr=None, description=None, enable_binary_authorization=None, enable_kubernetes_alpha=None, enable_legacy_abac=None, enable_tpu=None, initial_node_count=None, ip_allocation_policy=None, logging_service=None, maintenance_policy=None, master_auth=None, master_authorized_networks_config=None, master_ipv4_cidr_block=None, min_master_version=None, monitoring_service=None, name=None, network=None, network_policy=None, node_config=None, node_pools=None, node_version=None, pod_security_policy_config=None, private_cluster=None, private_cluster_config=None, project=None, region=None, remove_default_node_pool=None, resource_labels=None, subnetwork=None, zone=None, __name__=None, __opts__=None):
        """
        Creates a Google Kubernetes Engine (GKE) cluster. For more information see
        [the official documentation](https://cloud.google.com/container-engine/docs/clusters)
        and
        [API](https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters).
        
        > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] additional_zones: The list of additional Google Compute Engine
               locations in which the cluster's nodes should be located. If additional zones are
               configured, the number of nodes specified in `initial_node_count` is created in
               all specified zones.
        :param pulumi.Input[dict] addons_config: The configuration for addons supported by GKE.
               Structure is documented below.
        :param pulumi.Input[str] cluster_ipv4_cidr: The IP address range of the kubernetes pods in
               this cluster. Default is an automatically assigned CIDR.
        :param pulumi.Input[str] description: Description of the cluster.
        :param pulumi.Input[bool] enable_binary_authorization: Enable Binary Authorization for this cluster.
               If enabled, all container images will be validated by Google Binary Authorization.
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
        :param pulumi.Input[bool] enable_kubernetes_alpha: Whether to enable Kubernetes Alpha features for
               this cluster. Note that when this option is enabled, the cluster cannot be upgraded
               and will be automatically deleted after 30 days.
        :param pulumi.Input[bool] enable_legacy_abac: Whether the ABAC authorizer is enabled for this cluster.
               When enabled, identities in the system, including service accounts, nodes, and controllers,
               will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
               Defaults to `false`
        :param pulumi.Input[bool] enable_tpu: Whether to enable Cloud TPU resources in this cluster.
               See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
        :param pulumi.Input[int] initial_node_count: The number of nodes to create in this
               cluster (not including the Kubernetes master). Must be set if `node_pool` is not set.
        :param pulumi.Input[dict] ip_allocation_policy: Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
               This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
               Structure is documented below.
        :param pulumi.Input[str] logging_service: The logging service that the cluster should
               write logs to. Available options include `logging.googleapis.com`,
               `logging.googleapis.com/kubernetes` (beta), and `none`. Defaults to `logging.googleapis.com`
        :param pulumi.Input[dict] maintenance_policy: The maintenance policy to use for the cluster. Structure is
               documented below.
        :param pulumi.Input[dict] master_auth: The authentication information for accessing the
               Kubernetes master. Structure is documented below.
        :param pulumi.Input[dict] master_authorized_networks_config: The desired configuration options
               for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
               external access (except the cluster node IPs, which GKE automatically whitelists).
        :param pulumi.Input[str] master_ipv4_cidr_block: Specifies a private
               [RFC1918](https://tools.ietf.org/html/rfc1918) block for the master's VPC. The master range must not overlap with any subnet in your cluster's VPC.
               The master and your cluster use VPC peering. Must be specified in CIDR notation and must be `/28` subnet.
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
               This field is deprecated, use `private_cluster_config.master_ipv4_cidr_block` instead.
        :param pulumi.Input[str] min_master_version: The minimum version of the master. GKE
               will auto-update the master to new versions, so this does not guarantee the
               current master version--use the read-only `master_version` field to obtain that.
               If unset, the cluster's version will be set by GKE to the version of the most recent
               official release (which is not necessarily the latest version).
        :param pulumi.Input[str] monitoring_service: The monitoring service that the cluster
               should write metrics to.
               Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
               VM metrics will be collected by Google Compute Engine regardless of this setting
               Available options include
               `monitoring.googleapis.com`, `monitoring.googleapis.com/kubernetes` (beta) and `none`.
               Defaults to `monitoring.googleapis.com`
        :param pulumi.Input[str] name: The name of the cluster, unique within the project and
               zone.
        :param pulumi.Input[str] network: The name or self_link of the Google Compute Engine
               network to which the cluster is connected. For Shared VPC, set this to the self link of the
               shared network.
        :param pulumi.Input[dict] network_policy: Configuration options for the
               [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
               feature. Structure is documented below.
        :param pulumi.Input[dict] node_config: Parameters used in creating the cluster's nodes.
               Structure is documented below.
        :param pulumi.Input[list] node_pools: List of node pools associated with this cluster.
               See google_container_node_pool for schema.
               **Warning:** node pools defined inside a cluster can't be changed (or added/removed) after
               cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability
               to say "these are the _only_ node pools associated with this cluster", use the
               google_container_node_pool resource instead of this property.
        :param pulumi.Input[str] node_version: The Kubernetes version on the nodes. Must either be unset
               or set to the same value as `min_master_version` on create. Defaults to the default
               version set by GKE which is not necessarily the latest version.
        :param pulumi.Input[dict] pod_security_policy_config: Configuration for the
               [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
               Structure is documented below.
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
        :param pulumi.Input[bool] private_cluster: If true, a
               [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters) will be created, meaning
               nodes do not get public IP addresses. It is mandatory to specify `master_ipv4_cidr_block` and
               `ip_allocation_policy` with this option.
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
               This field is deprecated, use `private_cluster_config.enable_private_nodes` instead.
        :param pulumi.Input[dict] private_cluster_config: A set of options for creating
               a private cluster. Structure is documented below.
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[bool] remove_default_node_pool: If true, deletes the default node pool upon cluster creation.
        :param pulumi.Input[dict] resource_labels: The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
        :param pulumi.Input[str] subnetwork: The name or self_link of the Google Compute Engine subnetwork in
               which the cluster's instances are launched.
        :param pulumi.Input[str] zone: The zone that the master and the number of nodes specified
               in `initial_node_count` should be created in. Only one of `zone` and `region`
               may be set. If neither zone nor region are set, the provider zone is used.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['additional_zones'] = additional_zones

        __props__['addons_config'] = addons_config

        __props__['cluster_ipv4_cidr'] = cluster_ipv4_cidr

        __props__['description'] = description

        __props__['enable_binary_authorization'] = enable_binary_authorization

        __props__['enable_kubernetes_alpha'] = enable_kubernetes_alpha

        __props__['enable_legacy_abac'] = enable_legacy_abac

        __props__['enable_tpu'] = enable_tpu

        __props__['initial_node_count'] = initial_node_count

        __props__['ip_allocation_policy'] = ip_allocation_policy

        __props__['logging_service'] = logging_service

        __props__['maintenance_policy'] = maintenance_policy

        __props__['master_auth'] = master_auth

        __props__['master_authorized_networks_config'] = master_authorized_networks_config

        __props__['master_ipv4_cidr_block'] = master_ipv4_cidr_block

        __props__['min_master_version'] = min_master_version

        __props__['monitoring_service'] = monitoring_service

        __props__['name'] = name

        __props__['network'] = network

        __props__['network_policy'] = network_policy

        __props__['node_config'] = node_config

        __props__['node_pools'] = node_pools

        __props__['node_version'] = node_version

        __props__['pod_security_policy_config'] = pod_security_policy_config

        __props__['private_cluster'] = private_cluster

        __props__['private_cluster_config'] = private_cluster_config

        __props__['project'] = project

        __props__['region'] = region

        __props__['remove_default_node_pool'] = remove_default_node_pool

        __props__['resource_labels'] = resource_labels

        __props__['subnetwork'] = subnetwork

        __props__['zone'] = zone

        __props__['endpoint'] = None
        __props__['instance_group_urls'] = None
        __props__['master_version'] = None

        super(Cluster, __self__).__init__(
            'gcp:container/cluster:Cluster',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

