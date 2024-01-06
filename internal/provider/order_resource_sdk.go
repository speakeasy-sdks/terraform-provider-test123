// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/test/terraform-provider-twst/internal/sdk/pkg/models/shared"
	"math/big"
)

func (r *OrderResourceModel) ToSharedCreateOrderInput() *shared.CreateOrderInput {
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	image := r.Image.ValueString()
	name := r.Name.ValueString()
	price, _ := r.Price.ValueBigFloat().Float64()
	teaser := r.Teaser.ValueString()
	out := shared.CreateOrderInput{
		Description: description,
		Image:       image,
		Name:        name,
		Price:       price,
		Teaser:      teaser,
	}
	return &out
}

func (r *OrderResourceModel) RefreshFromSharedOrder(resp *shared.Order) {
	if resp.Description != nil {
		r.Description = types.StringValue(*resp.Description)
	} else {
		r.Description = types.StringNull()
	}
	r.ID = types.Int64Value(resp.ID)
	r.Image = types.StringValue(resp.Image)
	r.Name = types.StringValue(resp.Name)
	r.Price = types.NumberValue(big.NewFloat(float64(resp.Price)))
	r.Teaser = types.StringValue(resp.Teaser)
}
