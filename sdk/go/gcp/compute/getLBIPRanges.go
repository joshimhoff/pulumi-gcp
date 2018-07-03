// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access IP ranges in your firewall rules.
// 
// https://cloud.google.com/compute/docs/load-balancing/health-checks#health_check_source_ips_and_firewall_rules
func LookupLBIPRanges(ctx *pulumi.Context) (*GetLBIPRangesResult, error) {
	outputs, err := ctx.Invoke("gcp:compute/getLBIPRanges:getLBIPRanges", nil)
	if err != nil {
		return nil, err
	}
	return &GetLBIPRangesResult{
		HttpSslTcpInternals: outputs["httpSslTcpInternals"],
		Networks: outputs["networks"],
		Id: outputs["id"],
	}, nil
}

// A collection of values returned by getLBIPRanges.
type GetLBIPRangesResult struct {
	// The IP ranges used for health checks when **HTTP(S), SSL proxy, TCP proxy, and Internal load balancing** is used
	HttpSslTcpInternals interface{}
	// The IP ranges used for health checks when **Network load balancing** is used
	Networks interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
