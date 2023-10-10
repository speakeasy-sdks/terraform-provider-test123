// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"net/http"
	"net/http/httputil"
	"reflect"
	tfReflect "twst/internal/provider/reflect"
)

func debugResponse(response *http.Response) string {
	dumpReq, err := httputil.DumpRequest(response.Request, true)
	if err != nil {
		dumpReq, err = httputil.DumpRequest(response.Request, false)
		if err != nil {
			return err.Error()
		}
	}
	dumpRes, err := httputil.DumpResponse(response, true)
	if err != nil {
		dumpRes, err = httputil.DumpResponse(response, false)
		if err != nil {
			return err.Error()
		}
	}
	return fmt.Sprintf("**Request**:\n%s\n**Response**:\n%s", string(dumpReq), string(dumpRes))
}

func reflectJSONKey(data any, key string) reflect.Value {
	jsonIfied, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("failed to marshal data: %w", err))
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonIfied, &jsonMap)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal data: %w", err))
	}
	return reflect.ValueOf(jsonMap[key])
}

func merge(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse, target interface{}) {
	var plan types.Object
	var state types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(state.As(ctx, target, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// we need a tftypes.Value for this Object to be able to use it with
	// our reflection code
	obj := types.ObjectType{AttrTypes: plan.AttributeTypes(ctx)}
	val, err := plan.ToTerraformValue(ctx)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("Object Conversion Error", "An unexpected error was encountered trying to convert object. This is always an error in the provider. Please report the following to the provider developer:\n\n"+err.Error()))
		return
	}
	resp.Diagnostics.Append(tfReflect.Into(ctx, obj, val, target, tfReflect.Options{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	}, path.Empty())...)
}
