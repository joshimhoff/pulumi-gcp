# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetKMSKeyRingResult:
    """
    A collection of values returned by getKMSKeyRing.
    """
    def __init__(__self__, self_link=None, id=None):
        if self_link and not isinstance(self_link, str):
            raise TypeError('Expected argument self_link to be a str')
        __self__.self_link = self_link
        """
        The self link of the created KeyRing. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}`.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_kms_key_ring(location=None,name=None,project=None,opts=None):
    """
    Provides access to Google Cloud Platform KMS KeyRing. For more information see
    [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_ring)
    and
    [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings).
    
    A KeyRing is a grouping of CryptoKeys for organizational purposes. A KeyRing belongs to a Google Cloud Platform Project
    and resides in a specific location.
    """
    __args__ = dict()

    __args__['location'] = location
    __args__['name'] = name
    __args__['project'] = project
    __ret__ = await pulumi.runtime.invoke('gcp:kms/getKMSKeyRing:getKMSKeyRing', __args__, opts=opts)

    return GetKMSKeyRingResult(
        self_link=__ret__.get('selfLink'),
        id=__ret__.get('id'))
