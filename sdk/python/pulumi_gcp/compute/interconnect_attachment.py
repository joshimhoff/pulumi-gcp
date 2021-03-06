# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class InterconnectAttachment(pulumi.CustomResource):
    candidate_subnets: pulumi.Output[list]
    cloud_router_ip_address: pulumi.Output[str]
    creation_timestamp: pulumi.Output[str]
    customer_router_ip_address: pulumi.Output[str]
    description: pulumi.Output[str]
    edge_availability_domain: pulumi.Output[str]
    google_reference_id: pulumi.Output[str]
    interconnect: pulumi.Output[str]
    name: pulumi.Output[str]
    pairing_key: pulumi.Output[str]
    partner_asn: pulumi.Output[str]
    private_interconnect_info: pulumi.Output[dict]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    router: pulumi.Output[str]
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    state: pulumi.Output[str]
    type: pulumi.Output[str]
    vlan_tag8021q: pulumi.Output[float]
    def __init__(__self__, resource_name, opts=None, candidate_subnets=None, description=None, edge_availability_domain=None, interconnect=None, name=None, project=None, region=None, router=None, type=None, vlan_tag8021q=None, __name__=None, __opts__=None):
        """
        Represents an InterconnectAttachment (VLAN attachment) resource. For more
        information, see Creating VLAN Attachments.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

        __props__['candidate_subnets'] = candidate_subnets

        __props__['description'] = description

        __props__['edge_availability_domain'] = edge_availability_domain

        __props__['interconnect'] = interconnect

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        if router is None:
            raise TypeError('Missing required property router')
        __props__['router'] = router

        __props__['type'] = type

        __props__['vlan_tag8021q'] = vlan_tag8021q

        __props__['cloud_router_ip_address'] = None
        __props__['creation_timestamp'] = None
        __props__['customer_router_ip_address'] = None
        __props__['google_reference_id'] = None
        __props__['pairing_key'] = None
        __props__['partner_asn'] = None
        __props__['private_interconnect_info'] = None
        __props__['self_link'] = None
        __props__['state'] = None

        super(InterconnectAttachment, __self__).__init__(
            'gcp:compute/interconnectAttachment:InterconnectAttachment',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

