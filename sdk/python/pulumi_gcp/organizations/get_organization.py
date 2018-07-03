# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetOrganizationResult(object):
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, create_time=None, directory_customer_id=None, domain=None, lifecycle_state=None, name=None, id=None):
        if create_time and not isinstance(create_time, basestring):
            raise TypeError('Expected argument create_time to be a basestring')
        __self__.create_time = create_time
        """
        Timestamp when the Organization was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        if directory_customer_id and not isinstance(directory_customer_id, basestring):
            raise TypeError('Expected argument directory_customer_id to be a basestring')
        __self__.directory_customer_id = directory_customer_id
        """
        The Google for Work customer ID of the Organization.
        """
        if domain and not isinstance(domain, basestring):
            raise TypeError('Expected argument domain to be a basestring')
        __self__.domain = domain
        if lifecycle_state and not isinstance(lifecycle_state, basestring):
            raise TypeError('Expected argument lifecycle_state to be a basestring')
        __self__.lifecycle_state = lifecycle_state
        """
        The Organization's current lifecycle state.
        """
        if name and not isinstance(name, basestring):
            raise TypeError('Expected argument name to be a basestring')
        __self__.name = name
        """
        The resource name of the Organization in the form `organizations/{organization_id}`.
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_organization(domain=None, organization=None):
    """
    Use this data source to get information about a Google Cloud Organization.
    
    ```hcl
    data "google_organization" "org" {
      domain = "example.com"
    }
    
    resource "google_folder" "sales" {
      display_name = "Sales"
      parent       = "${data.google_organization.org.name}"
    }
    ```
    """
    __args__ = dict()

    __args__['domain'] = domain
    __args__['organization'] = organization
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getOrganization:getOrganization', __args__)

    return GetOrganizationResult(
        create_time=__ret__.get('createTime'),
        directory_customer_id=__ret__.get('directoryCustomerId'),
        domain=__ret__.get('domain'),
        lifecycle_state=__ret__.get('lifecycleState'),
        name=__ret__.get('name'),
        id=__ret__.get('id'))
