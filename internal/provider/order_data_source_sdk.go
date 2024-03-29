// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/test/terraform-provider-twst/internal/sdk/models/shared"
	"math/big"
)

func (r *OrderDataSourceModel) RefreshFromSharedOrder(resp *shared.Order) {
	if resp != nil {
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.Int64Value(resp.ID)
		r.Image = types.StringValue(resp.Image)
		r.Name = types.StringValue(resp.Name)
		r.Price = types.NumberValue(big.NewFloat(float64(resp.Price)))
		r.Teaser = types.StringValue(resp.Teaser)
	}
}
