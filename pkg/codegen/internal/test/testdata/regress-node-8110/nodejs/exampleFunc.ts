// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function exampleFunc(args?: ExampleFuncArgs, opts?: pulumi.InvokeOptions): Promise<void> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("my8110::exampleFunc", {
        "enums": args.enums,
    }, opts);
}

export interface ExampleFuncArgs {
    enums?: (string | enums.MyEnum)[];
}
