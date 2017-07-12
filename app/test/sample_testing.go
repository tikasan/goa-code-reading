// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "Sample": Sample TestHelpers
//
// Command:
// $ goagen
// --design=github.com/tikasan/goa-code-reading/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-code-reading
// --version=v1.2.0-dirty

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/tikasan/goa-code-reading/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// GETSampleBadRequest runs the method GET of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GETSampleBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{defaultType}
		query["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		query["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		query["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		query["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		query["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		query["stringType"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{defaultType}
		prms["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		prms["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		prms["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		prms["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		prms["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		prms["stringType"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	gETCtx, _err := app.NewGETSampleContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GET(gETCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// GETSampleInternalServerError runs the method GET of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GETSampleInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{defaultType}
		query["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		query["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		query["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		query["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		query["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		query["stringType"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{defaultType}
		prms["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		prms["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		prms["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		prms["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		prms["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		prms["stringType"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	gETCtx, _err := app.NewGETSampleContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GET(gETCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// GETSampleNotFound runs the method GET of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GETSampleNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{defaultType}
		query["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		query["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		query["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		query["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		query["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		query["stringType"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{defaultType}
		prms["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		prms["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		prms["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		prms["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		prms["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		prms["stringType"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	gETCtx, _err := app.NewGETSampleContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GET(gETCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// GETSampleOK runs the method GET of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GETSampleOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) (http.ResponseWriter, app.SampleCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{defaultType}
		query["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		query["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		query["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		query["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		query["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		query["stringType"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{defaultType}
		prms["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		prms["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		prms["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		prms["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		prms["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		prms["stringType"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	gETCtx, _err := app.NewGETSampleContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GET(gETCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.SampleCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.SampleCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.SampleCollection", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GETSampleUnauthorized runs the method GET of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GETSampleUnauthorized(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{defaultType}
		query["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		query["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		query["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		query["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		query["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		query["stringType"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{defaultType}
		prms["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		prms["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		prms["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		prms["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		prms["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		prms["stringType"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	gETCtx, _err := app.NewGETSampleContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.GET(gETCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 401 {
		t.Errorf("invalid response status code: got %+v, expected 401", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// POSTSampleBadRequest runs the method POST of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func POSTSampleBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, id int, payload *app.POSTSamplePayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/%v", id),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	pOSTCtx, __err := app.NewPOSTSampleContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	pOSTCtx.Payload = payload

	// Perform action
	__err = ctrl.POST(pOSTCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// POSTSampleInternalServerError runs the method POST of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func POSTSampleInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, id int, payload *app.POSTSamplePayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/%v", id),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	pOSTCtx, __err := app.NewPOSTSampleContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	pOSTCtx.Payload = payload

	// Perform action
	__err = ctrl.POST(pOSTCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// POSTSampleNotFound runs the method POST of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func POSTSampleNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, id int, payload *app.POSTSamplePayload) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/%v", id),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	pOSTCtx, __err := app.NewPOSTSampleContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	pOSTCtx.Payload = payload

	// Perform action
	__err = ctrl.POST(pOSTCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// POSTSampleOK runs the method POST of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func POSTSampleOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, id int, payload *app.POSTSamplePayload) (http.ResponseWriter, *app.Sample) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/%v", id),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	pOSTCtx, __err := app.NewPOSTSampleContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	pOSTCtx.Payload = payload

	// Perform action
	__err = ctrl.POST(pOSTCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Sample
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Sample)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Sample", resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// POSTSampleUnauthorized runs the method POST of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func POSTSampleUnauthorized(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.SampleController, id int, payload *app.POSTSamplePayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/%v", id),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "SampleTest"), rw, req, prms)
	pOSTCtx, __err := app.NewPOSTSampleContext(goaCtx, req, service)
	if __err != nil {
		panic("invalid test data " + __err.Error()) // bug
	}
	pOSTCtx.Payload = payload

	// Perform action
	__err = ctrl.POST(pOSTCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 401 {
		t.Errorf("invalid response status code: got %+v, expected 401", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}
