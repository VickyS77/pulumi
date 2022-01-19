// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NoRecursive extends pulumi.CustomResource {
    /**
     * Get an existing NoRecursive resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NoRecursive {
        return new NoRecursive(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'example::NoRecursive';

    /**
     * Returns true if the given object is an instance of NoRecursive.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NoRecursive {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NoRecursive.__pulumiType;
    }

    public /*out*/ readonly rec!: pulumi.Output<outputs.Rec | undefined>;
    public /*out*/ readonly replace!: pulumi.Output<string | undefined>;

    /**
     * Create a NoRecursive resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NoRecursiveArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["rec"] = undefined /*out*/;
            resourceInputs["replace"] = undefined /*out*/;
        } else {
            resourceInputs["rec"] = undefined /*out*/;
            resourceInputs["replace"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["replace"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(NoRecursive.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NoRecursive resource.
 */
export interface NoRecursiveArgs {
}
