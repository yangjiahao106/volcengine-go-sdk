// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDatabaseDescriptionCommon = "ModifyDatabaseDescription"

// ModifyDatabaseDescriptionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDatabaseDescriptionCommon operation. The "output" return
// value will be populated with the ModifyDatabaseDescriptionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDatabaseDescriptionCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDatabaseDescriptionCommon Send returns without error.
//
// See ModifyDatabaseDescriptionCommon for more information on using the ModifyDatabaseDescriptionCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyDatabaseDescriptionCommonRequest method.
//	req, resp := client.ModifyDatabaseDescriptionCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VEDBM) ModifyDatabaseDescriptionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDatabaseDescriptionCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDatabaseDescriptionCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation ModifyDatabaseDescriptionCommon for usage and error information.
func (c *VEDBM) ModifyDatabaseDescriptionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDatabaseDescriptionCommonRequest(input)
	return out, req.Send()
}

// ModifyDatabaseDescriptionCommonWithContext is the same as ModifyDatabaseDescriptionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDatabaseDescriptionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) ModifyDatabaseDescriptionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDatabaseDescriptionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDatabaseDescription = "ModifyDatabaseDescription"

// ModifyDatabaseDescriptionRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDatabaseDescription operation. The "output" return
// value will be populated with the ModifyDatabaseDescriptionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDatabaseDescriptionCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDatabaseDescriptionCommon Send returns without error.
//
// See ModifyDatabaseDescription for more information on using the ModifyDatabaseDescription
// API call, and error handling.
//
//	// Example sending a request using the ModifyDatabaseDescriptionRequest method.
//	req, resp := client.ModifyDatabaseDescriptionRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VEDBM) ModifyDatabaseDescriptionRequest(input *ModifyDatabaseDescriptionInput) (req *request.Request, output *ModifyDatabaseDescriptionOutput) {
	op := &request.Operation{
		Name:       opModifyDatabaseDescription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDatabaseDescriptionInput{}
	}

	output = &ModifyDatabaseDescriptionOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDatabaseDescription API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation ModifyDatabaseDescription for usage and error information.
func (c *VEDBM) ModifyDatabaseDescription(input *ModifyDatabaseDescriptionInput) (*ModifyDatabaseDescriptionOutput, error) {
	req, out := c.ModifyDatabaseDescriptionRequest(input)
	return out, req.Send()
}

// ModifyDatabaseDescriptionWithContext is the same as ModifyDatabaseDescription with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDatabaseDescription for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) ModifyDatabaseDescriptionWithContext(ctx volcengine.Context, input *ModifyDatabaseDescriptionInput, opts ...request.Option) (*ModifyDatabaseDescriptionOutput, error) {
	req, out := c.ModifyDatabaseDescriptionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDatabaseDescriptionInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// DBDesc is a required field
	DBDesc *string `type:"string" json:",omitempty" required:"true"`

	// DBName is a required field
	DBName *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyDatabaseDescriptionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDatabaseDescriptionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDatabaseDescriptionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDatabaseDescriptionInput"}
	if s.DBDesc == nil {
		invalidParams.Add(request.NewErrParamRequired("DBDesc"))
	}
	if s.DBName == nil {
		invalidParams.Add(request.NewErrParamRequired("DBName"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDBDesc sets the DBDesc field's value.
func (s *ModifyDatabaseDescriptionInput) SetDBDesc(v string) *ModifyDatabaseDescriptionInput {
	s.DBDesc = &v
	return s
}

// SetDBName sets the DBName field's value.
func (s *ModifyDatabaseDescriptionInput) SetDBName(v string) *ModifyDatabaseDescriptionInput {
	s.DBName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDatabaseDescriptionInput) SetInstanceId(v string) *ModifyDatabaseDescriptionInput {
	s.InstanceId = &v
	return s
}

type ModifyDatabaseDescriptionOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDatabaseDescriptionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDatabaseDescriptionOutput) GoString() string {
	return s.String()
}
