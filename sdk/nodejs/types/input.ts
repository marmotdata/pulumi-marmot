// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AssetEnvironmentArgs {
    metadata?: pulumi.Input<{[key: string]: any}>;
    name: pulumi.Input<string>;
    path: pulumi.Input<string>;
}

export interface AssetSourceArgs {
    name: pulumi.Input<string>;
    priority?: pulumi.Input<number>;
    properties?: pulumi.Input<{[key: string]: any}>;
}

export interface ExternalLinkArgs {
    icon?: pulumi.Input<string>;
    name: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
