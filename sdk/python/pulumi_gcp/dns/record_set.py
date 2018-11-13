# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class RecordSet(pulumi.CustomResource):
    """
    Manages a set of DNS records within Google Cloud DNS. For more information see [the official documentation](https://cloud.google.com/dns/records/) and
    [API](https://cloud.google.com/dns/api/v1/resourceRecordSets).
    
    ~> **Note:** The Google Cloud DNS API requires NS records be present at all
    times. To accommodate this, when creating NS records, the default records
    Google automatically creates will be silently overwritten.  Also, when
    destroying NS records, Terraform will not actually remove NS records, but will
    report that it did.
    """
    def __init__(__self__, __name__, __opts__=None, managed_zone=None, name=None, project=None, rrdatas=None, ttl=None, type=None):
        """Create a RecordSet resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not managed_zone:
            raise TypeError('Missing required property managed_zone')
        __props__['managed_zone'] = managed_zone

        __props__['name'] = name

        __props__['project'] = project

        if not rrdatas:
            raise TypeError('Missing required property rrdatas')
        __props__['rrdatas'] = rrdatas

        if not ttl:
            raise TypeError('Missing required property ttl')
        __props__['ttl'] = ttl

        if not type:
            raise TypeError('Missing required property type')
        __props__['type'] = type

        super(RecordSet, __self__).__init__(
            'gcp:dns/recordSet:RecordSet',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

