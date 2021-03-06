# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class CryptoKeyIAMMember(pulumi.CustomResource):
    crypto_key_id: pulumi.Output[str]
    """
    The key ring ID, in the form
    `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
    `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
    the provider's project setting will be used as a fallback.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the project's IAM policy.
    """
    member: pulumi.Output[str]
    """
    The user that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, crypto_key_id=None, member=None, role=None, __name__=None, __opts__=None):
        """
        Allows creation and management of a single member for a single binding within
        the IAM policy for an existing Google Cloud KMS crypto key.
        
        > **Note:** This resource _must not_ be used in conjunction with
           `google_kms_crypto_key_iam_policy` or they will fight over what your policy
           should be. Similarly, roles controlled by `google_kms_crypto_key_iam_binding`
           should not be assigned to using `google_kms_crypto_key_iam_member`.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_key_id: The key ring ID, in the form
               `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
               `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
               the provider's project setting will be used as a fallback.
        :param pulumi.Input[str] member: The user that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        :param pulumi.Input[str] role: The role that should be applied. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
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

        if crypto_key_id is None:
            raise TypeError('Missing required property crypto_key_id')
        __props__['crypto_key_id'] = crypto_key_id

        if member is None:
            raise TypeError('Missing required property member')
        __props__['member'] = member

        if role is None:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        __props__['etag'] = None

        super(CryptoKeyIAMMember, __self__).__init__(
            'gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

