// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package spanner

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
// 
// * `google_spanner_database_iam_policy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
// 
// ~> **Warning:** It's entirely possibly to lock yourself out of your database using `google_spanner_database_iam_policy`. Any permissions granted by default will be removed unless you include them in your config.
// 
// * `google_spanner_database_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
// * `google_spanner_database_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
// 
// ~> **Note:** `google_spanner_database_iam_policy` **cannot** be used in conjunction with `google_spanner_database_iam_binding` and `google_spanner_database_iam_member` or they will fight over what your policy should be.
// 
// ~> **Note:** `google_spanner_database_iam_binding` resources **can be** used in conjunction with `google_spanner_database_iam_member` resources **only if** they do not grant privilege to the same role.
type DatabaseIAMMember struct {
	s *pulumi.ResourceState
}

// NewDatabaseIAMMember registers a new resource with the given unique name, arguments, and options.
func NewDatabaseIAMMember(ctx *pulumi.Context,
	name string, args *DatabaseIAMMemberArgs, opts ...pulumi.ResourceOpt) (*DatabaseIAMMember, error) {
	if args == nil || args.Database == nil {
		return nil, errors.New("missing required argument 'Database'")
	}
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["database"] = nil
		inputs["instance"] = nil
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
	} else {
		inputs["database"] = args.Database
		inputs["instance"] = args.Instance
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:spanner/databaseIAMMember:DatabaseIAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatabaseIAMMember{s: s}, nil
}

// GetDatabaseIAMMember gets an existing DatabaseIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DatabaseIAMMemberState, opts ...pulumi.ResourceOpt) (*DatabaseIAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["database"] = state.Database
		inputs["etag"] = state.Etag
		inputs["instance"] = state.Instance
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:spanner/databaseIAMMember:DatabaseIAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatabaseIAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DatabaseIAMMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DatabaseIAMMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the Spanner database.
func (r *DatabaseIAMMember) Database() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["database"])
}

// (Computed) The etag of the database's IAM policy.
func (r *DatabaseIAMMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The name of the Spanner instance the database belongs to.
func (r *DatabaseIAMMember) Instance() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instance"])
}

func (r *DatabaseIAMMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *DatabaseIAMMember) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `google_spanner_database_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *DatabaseIAMMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering DatabaseIAMMember resources.
type DatabaseIAMMemberState struct {
	// The name of the Spanner database.
	Database interface{}
	// (Computed) The etag of the database's IAM policy.
	Etag interface{}
	// The name of the Spanner instance the database belongs to.
	Instance interface{}
	Member interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `google_spanner_database_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a DatabaseIAMMember resource.
type DatabaseIAMMemberArgs struct {
	// The name of the Spanner database.
	Database interface{}
	// The name of the Spanner instance the database belongs to.
	Instance interface{}
	Member interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `google_spanner_database_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
