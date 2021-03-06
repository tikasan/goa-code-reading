// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Sample": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/tikasan/goa-code-reading/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-code-reading
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// GETSampleContext provides the Sample GET action context.
type GETSampleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	DefaultType string
	Email       string
	EnumType    string
	ID          int
	IntegerType int
	Reg         string
	StringType  string
}

// NewGETSampleContext parses the incoming request URL and body, performs validations and creates the
// context used by the Sample controller GET action.
func NewGETSampleContext(ctx context.Context, r *http.Request, service *goa.Service) (*GETSampleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GETSampleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramDefaultType := req.Params["defaultType"]
	if len(paramDefaultType) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("defaultType"))
	} else {
		rawDefaultType := paramDefaultType[0]
		rctx.DefaultType = rawDefaultType
	}
	paramEmail := req.Params["email"]
	if len(paramEmail) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("email"))
	} else {
		rawEmail := paramEmail[0]
		rctx.Email = rawEmail
	}
	paramEnumType := req.Params["enumType"]
	if len(paramEnumType) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("enumType"))
	} else {
		rawEnumType := paramEnumType[0]
		rctx.EnumType = rawEnumType
	}
	paramID := req.Params["id"]
	if len(paramID) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("id"))
	} else {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	paramIntegerType := req.Params["integerType"]
	if len(paramIntegerType) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("integerType"))
	} else {
		rawIntegerType := paramIntegerType[0]
		if integerType, err2 := strconv.Atoi(rawIntegerType); err2 == nil {
			rctx.IntegerType = integerType
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("integerType", rawIntegerType, "integer"))
		}
	}
	paramReg := req.Params["reg"]
	if len(paramReg) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("reg"))
	} else {
		rawReg := paramReg[0]
		rctx.Reg = rawReg
	}
	paramStringType := req.Params["stringType"]
	if len(paramStringType) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("stringType"))
	} else {
		rawStringType := paramStringType[0]
		rctx.StringType = rawStringType
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GETSampleContext) OK(r SampleCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.sample+json; type=collection")
	if r == nil {
		r = SampleCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GETSampleContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *GETSampleContext) Unauthorized(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GETSampleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GETSampleContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// POSTSampleContext provides the Sample POST action context.
type POSTSampleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *POSTSamplePayload
}

// NewPOSTSampleContext parses the incoming request URL and body, performs validations and creates the
// context used by the Sample controller POST action.
func NewPOSTSampleContext(ctx context.Context, r *http.Request, service *goa.Service) (*POSTSampleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := POSTSampleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// pOSTSamplePayload is the Sample POST action payload.
type pOSTSamplePayload struct {
	// デフォルト値
	DefaultType *string `form:"defaultType,omitempty" json:"defaultType,omitempty" xml:"defaultType,omitempty"`
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// 列挙型
	EnumType *string `form:"enumType,omitempty" json:"enumType,omitempty" xml:"enumType,omitempty"`
	// 数字（1〜10）
	IntegerType *int `form:"integerType,omitempty" json:"integerType,omitempty" xml:"integerType,omitempty"`
	// デフォルト値
	Reg *string `form:"reg,omitempty" json:"reg,omitempty" xml:"reg,omitempty"`
	// 文字（1~10文字）
	StringType *string `form:"stringType,omitempty" json:"stringType,omitempty" xml:"stringType,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *pOSTSamplePayload) Validate() (err error) {
	if payload.IntegerType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "integerType"))
	}
	if payload.StringType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "stringType"))
	}
	if payload.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.EnumType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "enumType"))
	}
	if payload.DefaultType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "defaultType"))
	}
	if payload.Reg == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "reg"))
	}
	return
}

// Publicize creates POSTSamplePayload from pOSTSamplePayload
func (payload *pOSTSamplePayload) Publicize() *POSTSamplePayload {
	var pub POSTSamplePayload
	if payload.DefaultType != nil {
		pub.DefaultType = *payload.DefaultType
	}
	if payload.Email != nil {
		pub.Email = *payload.Email
	}
	if payload.EnumType != nil {
		pub.EnumType = *payload.EnumType
	}
	if payload.IntegerType != nil {
		pub.IntegerType = *payload.IntegerType
	}
	if payload.Reg != nil {
		pub.Reg = *payload.Reg
	}
	if payload.StringType != nil {
		pub.StringType = *payload.StringType
	}
	return &pub
}

// POSTSamplePayload is the Sample POST action payload.
type POSTSamplePayload struct {
	// デフォルト値
	DefaultType string `form:"defaultType" json:"defaultType" xml:"defaultType"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 列挙型
	EnumType string `form:"enumType" json:"enumType" xml:"enumType"`
	// 数字（1〜10）
	IntegerType int `form:"integerType" json:"integerType" xml:"integerType"`
	// デフォルト値
	Reg string `form:"reg" json:"reg" xml:"reg"`
	// 文字（1~10文字）
	StringType string `form:"stringType" json:"stringType" xml:"stringType"`
}

// Validate runs the validation rules defined in the design.
func (payload *POSTSamplePayload) Validate() (err error) {

	if payload.StringType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "stringType"))
	}
	if payload.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "email"))
	}
	if payload.EnumType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "enumType"))
	}
	if payload.DefaultType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "defaultType"))
	}
	if payload.Reg == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "reg"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *POSTSampleContext) OK(r *Sample) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.sample+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *POSTSampleContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *POSTSampleContext) Unauthorized(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *POSTSampleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *POSTSampleContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
