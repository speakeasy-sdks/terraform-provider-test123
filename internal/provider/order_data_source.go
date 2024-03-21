// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/test/terraform-provider-twst/internal/sdk"
	"github.com/test/terraform-provider-twst/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OrderDataSource{}
var _ datasource.DataSourceWithConfigure = &OrderDataSource{}

func NewOrderDataSource() datasource.DataSource {
	return &OrderDataSource{}
}

// OrderDataSource is the data source implementation.
type OrderDataSource struct {
	client *sdk.Twst
}

// OrderDataSourceModel describes the data model.
type OrderDataSourceModel struct {
	Description types.String `tfsdk:"description"`
	ID          types.Int64  `tfsdk:"id"`
	Image       types.String `tfsdk:"image"`
	Name        types.String `tfsdk:"name"`
	Price       types.Number `tfsdk:"price"`
	Teaser      types.String `tfsdk:"teaser"`
}

// Metadata returns the data source type name.
func (r *OrderDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_order"
}

// Schema defines the schema for the data source.
func (r *OrderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Order DataSource",

		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `Product description of the coffee.`,
			},
			"id": schema.Int64Attribute{
				Required:    true,
				Description: `The ID of the order get.`,
			},
			"image": schema.StringAttribute{
				Computed:    true,
				Description: `URI for an image of the coffee.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `Product name of the coffee.`,
			},
			"price": schema.NumberAttribute{
				Computed:    true,
				Description: `Suggested cost of the coffee.`,
			},
			"teaser": schema.StringAttribute{
				Computed:    true,
				Description: `Fun tagline for the coffee.`,
			},
		},
	}
}

func (r *OrderDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Twst)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Twst, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *OrderDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *OrderDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	orderID := data.ID.ValueInt64()
	request := operations.GetOrderRequest{
		OrderID: orderID,
	}
	res, err := r.client.Order.GetOrder(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Order == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedOrder(res.Order)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
