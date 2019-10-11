// Test the ability to invoke provider functions via RPC.

let assert = require("assert");
let pulumi = require("../../../../../");

class Provider extends pulumi.ProviderResource {
	constructor(name, opts) {
		super("test", name, {}, opts);
	}
}

const provider = new Provider("p");

let args = {
    a: "hello",
    b: true,
    c: [0.99, 42, { z: "x" }],
    id: "some-id",
    urn: "some-urn",
};

let result1 = pulumi.runtime.invoke("invoke:index:echo", args, { provider });
assert.deepEqual(result1, args);

let result2 = pulumi.runtime.invoke("invoke:index:echo", args, { provider });
result2.then((v) => {
    assert.deepEqual(v, args);
});

let result3 = pulumi.runtime.invoke("invoke:index:echo", args, { provider, async: true });
result3.then((v) => {
    assert.deepEqual(v, args);
});
