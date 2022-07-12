// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A toy for a dog
type Chew struct {
	Owner *Dog `pulumi:"owner"`
}

// ChewInput is an input type that accepts ChewArgs and ChewOutput values.
// You can construct a concrete instance of `ChewInput` via:
//
//          ChewArgs{...}
type ChewInput interface {
	pulumi.Input

	ToChewOutput() ChewOutput
	ToChewOutputWithContext(context.Context) ChewOutput
}

// A toy for a dog
type ChewArgs struct {
	Owner DogInput `pulumi:"owner"`
}

func (ChewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Chew)(nil)).Elem()
}

func (i ChewArgs) ToChewOutput() ChewOutput {
	return i.ToChewOutputWithContext(context.Background())
}

func (i ChewArgs) ToChewOutputWithContext(ctx context.Context) ChewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChewOutput)
}

func (i ChewArgs) ToChewPtrOutput() ChewPtrOutput {
	return i.ToChewPtrOutputWithContext(context.Background())
}

func (i ChewArgs) ToChewPtrOutputWithContext(ctx context.Context) ChewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChewOutput).ToChewPtrOutputWithContext(ctx)
}

// ChewPtrInput is an input type that accepts ChewArgs, ChewPtr and ChewPtrOutput values.
// You can construct a concrete instance of `ChewPtrInput` via:
//
//          ChewArgs{...}
//
//  or:
//
//          nil
type ChewPtrInput interface {
	pulumi.Input

	ToChewPtrOutput() ChewPtrOutput
	ToChewPtrOutputWithContext(context.Context) ChewPtrOutput
}

type chewPtrType ChewArgs

func ChewPtr(v *ChewArgs) ChewPtrInput {
	return (*chewPtrType)(v)
}

func (*chewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Chew)(nil)).Elem()
}

func (i *chewPtrType) ToChewPtrOutput() ChewPtrOutput {
	return i.ToChewPtrOutputWithContext(context.Background())
}

func (i *chewPtrType) ToChewPtrOutputWithContext(ctx context.Context) ChewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChewPtrOutput)
}

// A toy for a dog
type ChewOutput struct{ *pulumi.OutputState }

func (ChewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Chew)(nil)).Elem()
}

func (o ChewOutput) ToChewOutput() ChewOutput {
	return o
}

func (o ChewOutput) ToChewOutputWithContext(ctx context.Context) ChewOutput {
	return o
}

func (o ChewOutput) ToChewPtrOutput() ChewPtrOutput {
	return o.ToChewPtrOutputWithContext(context.Background())
}

func (o ChewOutput) ToChewPtrOutputWithContext(ctx context.Context) ChewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Chew) *Chew {
		return &v
	}).(ChewPtrOutput)
}

func (o ChewOutput) Owner() DogOutput {
	return o.ApplyT(func(v Chew) *Dog { return v.Owner }).(DogOutput)
}

type ChewPtrOutput struct{ *pulumi.OutputState }

func (ChewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Chew)(nil)).Elem()
}

func (o ChewPtrOutput) ToChewPtrOutput() ChewPtrOutput {
	return o
}

func (o ChewPtrOutput) ToChewPtrOutputWithContext(ctx context.Context) ChewPtrOutput {
	return o
}

func (o ChewPtrOutput) Elem() ChewOutput {
	return o.ApplyT(func(v *Chew) Chew {
		if v != nil {
			return *v
		}
		var ret Chew
		return ret
	}).(ChewOutput)
}

func (o ChewPtrOutput) Owner() DogOutput {
	return o.ApplyT(func(v *Chew) *Dog {
		if v == nil {
			return nil
		}
		return v.Owner
	}).(DogOutput)
}

// A Toy for a cat
type Laser struct {
	Animal    *Cat     `pulumi:"animal"`
	Batteries *bool    `pulumi:"batteries"`
	Light     *float64 `pulumi:"light"`
}

// LaserInput is an input type that accepts LaserArgs and LaserOutput values.
// You can construct a concrete instance of `LaserInput` via:
//
//          LaserArgs{...}
type LaserInput interface {
	pulumi.Input

	ToLaserOutput() LaserOutput
	ToLaserOutputWithContext(context.Context) LaserOutput
}

// A Toy for a cat
type LaserArgs struct {
	Animal    CatInput               `pulumi:"animal"`
	Batteries pulumi.BoolPtrInput    `pulumi:"batteries"`
	Light     pulumi.Float64PtrInput `pulumi:"light"`
}

func (LaserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Laser)(nil)).Elem()
}

func (i LaserArgs) ToLaserOutput() LaserOutput {
	return i.ToLaserOutputWithContext(context.Background())
}

func (i LaserArgs) ToLaserOutputWithContext(ctx context.Context) LaserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaserOutput)
}

func (i LaserArgs) ToLaserPtrOutput() LaserPtrOutput {
	return i.ToLaserPtrOutputWithContext(context.Background())
}

func (i LaserArgs) ToLaserPtrOutputWithContext(ctx context.Context) LaserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaserOutput).ToLaserPtrOutputWithContext(ctx)
}

// LaserPtrInput is an input type that accepts LaserArgs, LaserPtr and LaserPtrOutput values.
// You can construct a concrete instance of `LaserPtrInput` via:
//
//          LaserArgs{...}
//
//  or:
//
//          nil
type LaserPtrInput interface {
	pulumi.Input

	ToLaserPtrOutput() LaserPtrOutput
	ToLaserPtrOutputWithContext(context.Context) LaserPtrOutput
}

type laserPtrType LaserArgs

func LaserPtr(v *LaserArgs) LaserPtrInput {
	return (*laserPtrType)(v)
}

func (*laserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Laser)(nil)).Elem()
}

func (i *laserPtrType) ToLaserPtrOutput() LaserPtrOutput {
	return i.ToLaserPtrOutputWithContext(context.Background())
}

func (i *laserPtrType) ToLaserPtrOutputWithContext(ctx context.Context) LaserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaserPtrOutput)
}

// A Toy for a cat
type LaserOutput struct{ *pulumi.OutputState }

func (LaserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Laser)(nil)).Elem()
}

func (o LaserOutput) ToLaserOutput() LaserOutput {
	return o
}

func (o LaserOutput) ToLaserOutputWithContext(ctx context.Context) LaserOutput {
	return o
}

func (o LaserOutput) ToLaserPtrOutput() LaserPtrOutput {
	return o.ToLaserPtrOutputWithContext(context.Background())
}

func (o LaserOutput) ToLaserPtrOutputWithContext(ctx context.Context) LaserPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Laser) *Laser {
		return &v
	}).(LaserPtrOutput)
}

func (o LaserOutput) Animal() CatOutput {
	return o.ApplyT(func(v Laser) *Cat { return v.Animal }).(CatOutput)
}

func (o LaserOutput) Batteries() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Laser) *bool { return v.Batteries }).(pulumi.BoolPtrOutput)
}

func (o LaserOutput) Light() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Laser) *float64 { return v.Light }).(pulumi.Float64PtrOutput)
}

type LaserPtrOutput struct{ *pulumi.OutputState }

func (LaserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Laser)(nil)).Elem()
}

func (o LaserPtrOutput) ToLaserPtrOutput() LaserPtrOutput {
	return o
}

func (o LaserPtrOutput) ToLaserPtrOutputWithContext(ctx context.Context) LaserPtrOutput {
	return o
}

func (o LaserPtrOutput) Elem() LaserOutput {
	return o.ApplyT(func(v *Laser) Laser {
		if v != nil {
			return *v
		}
		var ret Laser
		return ret
	}).(LaserOutput)
}

func (o LaserPtrOutput) Animal() CatOutput {
	return o.ApplyT(func(v *Laser) *Cat {
		if v == nil {
			return nil
		}
		return v.Animal
	}).(CatOutput)
}

func (o LaserPtrOutput) Batteries() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Laser) *bool {
		if v == nil {
			return nil
		}
		return v.Batteries
	}).(pulumi.BoolPtrOutput)
}

func (o LaserPtrOutput) Light() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Laser) *float64 {
		if v == nil {
			return nil
		}
		return v.Light
	}).(pulumi.Float64PtrOutput)
}

type Rec struct {
	Rec1 *Rec `pulumi:"rec1"`
}

// RecInput is an input type that accepts RecArgs and RecOutput values.
// You can construct a concrete instance of `RecInput` via:
//
//          RecArgs{...}
type RecInput interface {
	pulumi.Input

	ToRecOutput() RecOutput
	ToRecOutputWithContext(context.Context) RecOutput
}

type RecArgs struct {
	Rec1 RecPtrInput `pulumi:"rec1"`
}

func (RecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Rec)(nil)).Elem()
}

func (i RecArgs) ToRecOutput() RecOutput {
	return i.ToRecOutputWithContext(context.Background())
}

func (i RecArgs) ToRecOutputWithContext(ctx context.Context) RecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecOutput)
}

func (i RecArgs) ToRecPtrOutput() RecPtrOutput {
	return i.ToRecPtrOutputWithContext(context.Background())
}

func (i RecArgs) ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecOutput).ToRecPtrOutputWithContext(ctx)
}

// RecPtrInput is an input type that accepts RecArgs, RecPtr and RecPtrOutput values.
// You can construct a concrete instance of `RecPtrInput` via:
//
//          RecArgs{...}
//
//  or:
//
//          nil
type RecPtrInput interface {
	pulumi.Input

	ToRecPtrOutput() RecPtrOutput
	ToRecPtrOutputWithContext(context.Context) RecPtrOutput
}

type recPtrType RecArgs

func RecPtr(v *RecArgs) RecPtrInput {
	return (*recPtrType)(v)
}

func (*recPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Rec)(nil)).Elem()
}

func (i *recPtrType) ToRecPtrOutput() RecPtrOutput {
	return i.ToRecPtrOutputWithContext(context.Background())
}

func (i *recPtrType) ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecPtrOutput)
}

type RecOutput struct{ *pulumi.OutputState }

func (RecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rec)(nil)).Elem()
}

func (o RecOutput) ToRecOutput() RecOutput {
	return o
}

func (o RecOutput) ToRecOutputWithContext(ctx context.Context) RecOutput {
	return o
}

func (o RecOutput) ToRecPtrOutput() RecPtrOutput {
	return o.ToRecPtrOutputWithContext(context.Background())
}

func (o RecOutput) ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Rec) *Rec {
		return &v
	}).(RecPtrOutput)
}

func (o RecOutput) Rec1() RecPtrOutput {
	return o.ApplyT(func(v Rec) *Rec { return v.Rec1 }).(RecPtrOutput)
}

type RecPtrOutput struct{ *pulumi.OutputState }

func (RecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rec)(nil)).Elem()
}

func (o RecPtrOutput) ToRecPtrOutput() RecPtrOutput {
	return o
}

func (o RecPtrOutput) ToRecPtrOutputWithContext(ctx context.Context) RecPtrOutput {
	return o
}

func (o RecPtrOutput) Elem() RecOutput {
	return o.ApplyT(func(v *Rec) Rec {
		if v != nil {
			return *v
		}
		var ret Rec
		return ret
	}).(RecOutput)
}

func (o RecPtrOutput) Rec1() RecPtrOutput {
	return o.ApplyT(func(v *Rec) *Rec {
		if v == nil {
			return nil
		}
		return v.Rec1
	}).(RecPtrOutput)
}

// This is a toy
type Toy struct {
	Associated            *Toy     `pulumi:"associated"`
	Color                 *string  `pulumi:"color"`
	HasHazardousChemicals *bool    `pulumi:"hasHazardousChemicals"`
	Wear                  *float64 `pulumi:"wear"`
}

// ToyInput is an input type that accepts ToyArgs and ToyOutput values.
// You can construct a concrete instance of `ToyInput` via:
//
//          ToyArgs{...}
type ToyInput interface {
	pulumi.Input

	ToToyOutput() ToyOutput
	ToToyOutputWithContext(context.Context) ToyOutput
}

// This is a toy
type ToyArgs struct {
	Associated            ToyPtrInput            `pulumi:"associated"`
	Color                 pulumi.StringPtrInput  `pulumi:"color"`
	HasHazardousChemicals pulumi.BoolPtrInput    `pulumi:"hasHazardousChemicals"`
	Wear                  pulumi.Float64PtrInput `pulumi:"wear"`
}

func (ToyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Toy)(nil)).Elem()
}

func (i ToyArgs) ToToyOutput() ToyOutput {
	return i.ToToyOutputWithContext(context.Background())
}

func (i ToyArgs) ToToyOutputWithContext(ctx context.Context) ToyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ToyOutput)
}

func (i ToyArgs) ToToyPtrOutput() ToyPtrOutput {
	return i.ToToyPtrOutputWithContext(context.Background())
}

func (i ToyArgs) ToToyPtrOutputWithContext(ctx context.Context) ToyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ToyOutput).ToToyPtrOutputWithContext(ctx)
}

// ToyPtrInput is an input type that accepts ToyArgs, ToyPtr and ToyPtrOutput values.
// You can construct a concrete instance of `ToyPtrInput` via:
//
//          ToyArgs{...}
//
//  or:
//
//          nil
type ToyPtrInput interface {
	pulumi.Input

	ToToyPtrOutput() ToyPtrOutput
	ToToyPtrOutputWithContext(context.Context) ToyPtrOutput
}

type toyPtrType ToyArgs

func ToyPtr(v *ToyArgs) ToyPtrInput {
	return (*toyPtrType)(v)
}

func (*toyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Toy)(nil)).Elem()
}

func (i *toyPtrType) ToToyPtrOutput() ToyPtrOutput {
	return i.ToToyPtrOutputWithContext(context.Background())
}

func (i *toyPtrType) ToToyPtrOutputWithContext(ctx context.Context) ToyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ToyPtrOutput)
}

// ToyArrayInput is an input type that accepts ToyArray and ToyArrayOutput values.
// You can construct a concrete instance of `ToyArrayInput` via:
//
//          ToyArray{ ToyArgs{...} }
type ToyArrayInput interface {
	pulumi.Input

	ToToyArrayOutput() ToyArrayOutput
	ToToyArrayOutputWithContext(context.Context) ToyArrayOutput
}

type ToyArray []ToyInput

func (ToyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Toy)(nil)).Elem()
}

func (i ToyArray) ToToyArrayOutput() ToyArrayOutput {
	return i.ToToyArrayOutputWithContext(context.Background())
}

func (i ToyArray) ToToyArrayOutputWithContext(ctx context.Context) ToyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ToyArrayOutput)
}

// ToyMapInput is an input type that accepts ToyMap and ToyMapOutput values.
// You can construct a concrete instance of `ToyMapInput` via:
//
//          ToyMap{ "key": ToyArgs{...} }
type ToyMapInput interface {
	pulumi.Input

	ToToyMapOutput() ToyMapOutput
	ToToyMapOutputWithContext(context.Context) ToyMapOutput
}

type ToyMap map[string]ToyInput

func (ToyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Toy)(nil)).Elem()
}

func (i ToyMap) ToToyMapOutput() ToyMapOutput {
	return i.ToToyMapOutputWithContext(context.Background())
}

func (i ToyMap) ToToyMapOutputWithContext(ctx context.Context) ToyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ToyMapOutput)
}

// This is a toy
type ToyOutput struct{ *pulumi.OutputState }

func (ToyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Toy)(nil)).Elem()
}

func (o ToyOutput) ToToyOutput() ToyOutput {
	return o
}

func (o ToyOutput) ToToyOutputWithContext(ctx context.Context) ToyOutput {
	return o
}

func (o ToyOutput) ToToyPtrOutput() ToyPtrOutput {
	return o.ToToyPtrOutputWithContext(context.Background())
}

func (o ToyOutput) ToToyPtrOutputWithContext(ctx context.Context) ToyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Toy) *Toy {
		return &v
	}).(ToyPtrOutput)
}

func (o ToyOutput) Associated() ToyPtrOutput {
	return o.ApplyT(func(v Toy) *Toy { return v.Associated }).(ToyPtrOutput)
}

func (o ToyOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Toy) *string { return v.Color }).(pulumi.StringPtrOutput)
}

func (o ToyOutput) HasHazardousChemicals() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Toy) *bool { return v.HasHazardousChemicals }).(pulumi.BoolPtrOutput)
}

func (o ToyOutput) Wear() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Toy) *float64 { return v.Wear }).(pulumi.Float64PtrOutput)
}

type ToyPtrOutput struct{ *pulumi.OutputState }

func (ToyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Toy)(nil)).Elem()
}

func (o ToyPtrOutput) ToToyPtrOutput() ToyPtrOutput {
	return o
}

func (o ToyPtrOutput) ToToyPtrOutputWithContext(ctx context.Context) ToyPtrOutput {
	return o
}

func (o ToyPtrOutput) Elem() ToyOutput {
	return o.ApplyT(func(v *Toy) Toy {
		if v != nil {
			return *v
		}
		var ret Toy
		return ret
	}).(ToyOutput)
}

func (o ToyPtrOutput) Associated() ToyPtrOutput {
	return o.ApplyT(func(v *Toy) *Toy {
		if v == nil {
			return nil
		}
		return v.Associated
	}).(ToyPtrOutput)
}

func (o ToyPtrOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Toy) *string {
		if v == nil {
			return nil
		}
		return v.Color
	}).(pulumi.StringPtrOutput)
}

func (o ToyPtrOutput) HasHazardousChemicals() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Toy) *bool {
		if v == nil {
			return nil
		}
		return v.HasHazardousChemicals
	}).(pulumi.BoolPtrOutput)
}

func (o ToyPtrOutput) Wear() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Toy) *float64 {
		if v == nil {
			return nil
		}
		return v.Wear
	}).(pulumi.Float64PtrOutput)
}

type ToyArrayOutput struct{ *pulumi.OutputState }

func (ToyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Toy)(nil)).Elem()
}

func (o ToyArrayOutput) ToToyArrayOutput() ToyArrayOutput {
	return o
}

func (o ToyArrayOutput) ToToyArrayOutputWithContext(ctx context.Context) ToyArrayOutput {
	return o
}

func (o ToyArrayOutput) Index(i pulumi.IntInput) ToyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Toy {
		return vs[0].([]Toy)[vs[1].(int)]
	}).(ToyOutput)
}

type ToyMapOutput struct{ *pulumi.OutputState }

func (ToyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Toy)(nil)).Elem()
}

func (o ToyMapOutput) ToToyMapOutput() ToyMapOutput {
	return o
}

func (o ToyMapOutput) ToToyMapOutputWithContext(ctx context.Context) ToyMapOutput {
	return o
}

func (o ToyMapOutput) MapIndex(k pulumi.StringInput) ToyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Toy {
		return vs[0].(map[string]Toy)[vs[1].(string)]
	}).(ToyOutput)
}

func init() {
	pulumi.RegisterOutputType(ChewOutput{})
	pulumi.RegisterOutputType(ChewPtrOutput{})
	pulumi.RegisterOutputType(LaserOutput{})
	pulumi.RegisterOutputType(LaserPtrOutput{})
	pulumi.RegisterOutputType(RecOutput{})
	pulumi.RegisterOutputType(RecPtrOutput{})
	pulumi.RegisterOutputType(ToyOutput{})
	pulumi.RegisterOutputType(ToyPtrOutput{})
	pulumi.RegisterOutputType(ToyArrayOutput{})
	pulumi.RegisterOutputType(ToyMapOutput{})
}
