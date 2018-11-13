// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of an App Engine application.
// 
// ~> App Engine applications cannot be deleted once they're created; you have to delete the
//    entire project to delete the application. Terraform will report the application has been
//    successfully deleted; this is a limitation of Terraform, and will go away in the future.
//    Terraform is not able to delete App Engine applications.
type Application struct {
	s *pulumi.ResourceState
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOpt) (*Application, error) {
	if args == nil || args.LocationId == nil {
		return nil, errors.New("missing required argument 'LocationId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["authDomain"] = nil
		inputs["featureSettings"] = nil
		inputs["locationId"] = nil
		inputs["project"] = nil
		inputs["servingStatus"] = nil
	} else {
		inputs["authDomain"] = args.AuthDomain
		inputs["featureSettings"] = args.FeatureSettings
		inputs["locationId"] = args.LocationId
		inputs["project"] = args.Project
		inputs["servingStatus"] = args.ServingStatus
	}
	inputs["codeBucket"] = nil
	inputs["defaultBucket"] = nil
	inputs["defaultHostname"] = nil
	inputs["gcrDomain"] = nil
	inputs["name"] = nil
	inputs["urlDispatchRules"] = nil
	s, err := ctx.RegisterResource("gcp:appengine/application:Application", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Application{s: s}, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ApplicationState, opts ...pulumi.ResourceOpt) (*Application, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["authDomain"] = state.AuthDomain
		inputs["codeBucket"] = state.CodeBucket
		inputs["defaultBucket"] = state.DefaultBucket
		inputs["defaultHostname"] = state.DefaultHostname
		inputs["featureSettings"] = state.FeatureSettings
		inputs["gcrDomain"] = state.GcrDomain
		inputs["locationId"] = state.LocationId
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["servingStatus"] = state.ServingStatus
		inputs["urlDispatchRules"] = state.UrlDispatchRules
	}
	s, err := ctx.ReadResource("gcp:appengine/application:Application", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Application{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Application) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Application) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The domain to authenticate users with when using App Engine's User API.
func (r *Application) AuthDomain() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["authDomain"])
}

// The GCS bucket code is being stored in for this app.
func (r *Application) CodeBucket() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["codeBucket"])
}

// The GCS bucket content is being stored in for this app.
func (r *Application) DefaultBucket() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultBucket"])
}

// The default hostname for this app.
func (r *Application) DefaultHostname() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultHostname"])
}

// A block of optional settings to configure specific App Engine features:
func (r *Application) FeatureSettings() *pulumi.Output {
	return r.s.State["featureSettings"]
}

// The GCR domain used for storing managed Docker images for this app.
func (r *Application) GcrDomain() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["gcrDomain"])
}

// The [location](https://cloud.google.com/appengine/docs/locations)
// to serve the app from.
func (r *Application) LocationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["locationId"])
}

// Unique name of the app, usually `apps/{PROJECT_ID}`
func (r *Application) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Application) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The serving status of the app.
func (r *Application) ServingStatus() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["servingStatus"])
}

// A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
func (r *Application) UrlDispatchRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["urlDispatchRules"])
}

// Input properties used for looking up and filtering Application resources.
type ApplicationState struct {
	// The domain to authenticate users with when using App Engine's User API.
	AuthDomain interface{}
	// The GCS bucket code is being stored in for this app.
	CodeBucket interface{}
	// The GCS bucket content is being stored in for this app.
	DefaultBucket interface{}
	// The default hostname for this app.
	DefaultHostname interface{}
	// A block of optional settings to configure specific App Engine features:
	FeatureSettings interface{}
	// The GCR domain used for storing managed Docker images for this app.
	GcrDomain interface{}
	// The [location](https://cloud.google.com/appengine/docs/locations)
	// to serve the app from.
	LocationId interface{}
	// Unique name of the app, usually `apps/{PROJECT_ID}`
	Name interface{}
	Project interface{}
	// The serving status of the app.
	ServingStatus interface{}
	// A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
	UrlDispatchRules interface{}
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The domain to authenticate users with when using App Engine's User API.
	AuthDomain interface{}
	// A block of optional settings to configure specific App Engine features:
	FeatureSettings interface{}
	// The [location](https://cloud.google.com/appengine/docs/locations)
	// to serve the app from.
	LocationId interface{}
	Project interface{}
	// The serving status of the app.
	ServingStatus interface{}
}
