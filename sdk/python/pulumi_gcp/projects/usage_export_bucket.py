# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class UsageExportBucket(pulumi.CustomResource):
    bucket_name: pulumi.Output[str]
    prefix: pulumi.Output[str]
    project: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, bucket_name=None, prefix=None, project=None, __name__=None, __opts__=None):
        """
        Allows creation and management of a Google Cloud Platform project.
        
        Projects created with this resource must be associated with an Organization.
        See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
        
        The service account used to run Terraform when creating a `google_project`
        resource must have `roles/resourcemanager.projectCreator`. See the
        [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
        doc for more information.
        
        Note that prior to 0.8.5, `google_project` functioned like a data source,
        meaning any project referenced by it had to be created and managed outside
        Terraform. As of 0.8.5, `google_project` functions like any other Terraform
        resource, with Terraform creating and managing the project. To replicate the old
        behavior, either:
        
        * Use the project ID directly in whatever is referencing the project, using the
          [google_project_iam_policy](https://www.terraform.io/docs/providers/google/r/google_project_iam.html)
          to replace the old `policy_data` property.
        * Use the [import](https://www.terraform.io/docs/import/usage.html) functionality
          to import your pre-existing project into Terraform, where it can be referenced and
          used just like always, keeping in mind that Terraform will attempt to undo any changes
          made outside Terraform.
        
        > It's important to note that any project resources that were added to your Terraform config
        prior to 0.8.5 will continue to function as they always have, and will not be managed by
        Terraform. Only newly added projects are affected.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

        if bucket_name is None:
            raise TypeError('Missing required property bucket_name')
        __props__['bucket_name'] = bucket_name

        __props__['prefix'] = prefix

        __props__['project'] = project

        super(UsageExportBucket, __self__).__init__(
            'gcp:projects/usageExportBucket:UsageExportBucket',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

