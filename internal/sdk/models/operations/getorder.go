// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/test/terraform-provider-twst/internal/sdk/models/shared"
	"net/http"
)

type GetOrderRequest struct {
	// The ID of the order get.
	OrderID int64 `pathParam:"style=simple,explode=false,name=orderID"`
}

func (o *GetOrderRequest) GetOrderID() int64 {
	if o == nil {
		return 0
	}
	return o.OrderID
}

type GetOrderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	Order *shared.Order
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetOrderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOrderResponse) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *GetOrderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOrderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
