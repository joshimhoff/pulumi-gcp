// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class SubNetwork extends lumi.NamedResource implements SubNetworkArgs {
    public readonly description?: string;
    public /*out*/ readonly gatewayAddress: string;
    public readonly ipCidrRange: string;
    public readonly subNetworkName?: string;
    public readonly network: string;
    public readonly privateIpGoogleAccess?: boolean;
    public readonly project?: string;
    public readonly region?: string;
    public /*out*/ readonly selfLink: string;

    constructor(name: string, args: SubNetworkArgs) {
        super(name);
        this.description = args.description;
        if (lumirt.defaultIfComputed(args.ipCidrRange, "") === undefined) {
            throw new Error("Property argument 'ipCidrRange' is required, but was missing");
        }
        this.ipCidrRange = args.ipCidrRange;
        this.subNetworkName = args.subNetworkName;
        if (lumirt.defaultIfComputed(args.network, "") === undefined) {
            throw new Error("Property argument 'network' is required, but was missing");
        }
        this.network = args.network;
        this.privateIpGoogleAccess = args.privateIpGoogleAccess;
        this.project = args.project;
        this.region = args.region;
    }
}

export interface SubNetworkArgs {
    readonly description?: string;
    readonly ipCidrRange: string;
    readonly subNetworkName?: string;
    readonly network: string;
    readonly privateIpGoogleAccess?: boolean;
    readonly project?: string;
    readonly region?: string;
}
