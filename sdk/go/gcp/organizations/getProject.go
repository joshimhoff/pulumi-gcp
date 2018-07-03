// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get project details.
// For more information see 
// [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects#Project)
func LookupProject(ctx *pulumi.Context, args *GetProjectArgs) (*GetProjectResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["projectId"] = args.ProjectId
	}
	outputs, err := ctx.Invoke("gcp:organizations/getProject:getProject", inputs)
	if err != nil {
		return nil, err
	}
	return &GetProjectResult{
		AppEngines: outputs["appEngines"],
		AutoCreateNetwork: outputs["autoCreateNetwork"],
		BillingAccount: outputs["billingAccount"],
		FolderId: outputs["folderId"],
		Labels: outputs["labels"],
		Name: outputs["name"],
		Number: outputs["number"],
		OrgId: outputs["orgId"],
		PolicyData: outputs["policyData"],
		PolicyEtag: outputs["policyEtag"],
		SkipDelete: outputs["skipDelete"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getProject.
type GetProjectArgs struct {
	// The project ID. If it is not provided, the provider project is used.
	ProjectId interface{}
}

// A collection of values returned by getProject.
type GetProjectResult struct {
	AppEngines interface{}
	AutoCreateNetwork interface{}
	BillingAccount interface{}
	FolderId interface{}
	Labels interface{}
	Name interface{}
	Number interface{}
	OrgId interface{}
	PolicyData interface{}
	PolicyEtag interface{}
	SkipDelete interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
