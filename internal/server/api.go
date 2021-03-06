// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	Status string `json:"status"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions []DistributionItem

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture   string          `json:"architecture"`
	ImageType      string          `json:"image_type"`
	UploadRequests []UploadRequest `json:"upload_requests"`
}

// Package defines model for Package.
type Package struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options AWSUploadRequestOptions `json:"options"`
	Type    string                  `json:"type"`
}

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {

	// distribution to look up packages for
	Distribution string `json:"distribution"`

	// architecture to look up packages for
	Architecture string `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// ComposeImageRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get the architectures and their image types available for a given distribution
	// (GET /architectures/{distribution})
	GetArchitectures(ctx echo.Context, distribution string) error
	// compose image
	// (POST /compose)
	ComposeImage(ctx echo.Context) error
	// get status of an image compose
	// (GET /composes/{composeId})
	GetComposeStatus(ctx echo.Context, composeId string) error
	// get the available distributions
	// (GET /distributions)
	GetDistributions(ctx echo.Context) error
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(ctx echo.Context) error

	// (GET /packages)
	GetPackages(ctx echo.Context, params GetPackagesParams) error
	// get the service version
	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArchitectures converts echo context to params.
func (w *ServerInterfaceWrapper) GetArchitectures(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "distribution" -------------
	var distribution string

	err = runtime.BindStyledParameter("simple", false, "distribution", ctx.Param("distribution"), &distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArchitectures(ctx, distribution)
	return err
}

// ComposeImage converts echo context to params.
func (w *ServerInterfaceWrapper) ComposeImage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ComposeImage(ctx)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, composeId)
	return err
}

// GetDistributions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDistributions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDistributions(ctx)
	return err
}

// GetOpenapiJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapiJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapiJson(ctx)
	return err
}

// GetPackages converts echo context to params.
func (w *ServerInterfaceWrapper) GetPackages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPackagesParams
	// ------------- Required query parameter "distribution" -------------

	err = runtime.BindQueryParameter("form", true, true, "distribution", ctx.QueryParams(), &params.Distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// ------------- Required query parameter "architecture" -------------

	err = runtime.BindQueryParameter("form", true, true, "architecture", ctx.QueryParams(), &params.Architecture)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter architecture: %s", err))
	}

	// ------------- Required query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, true, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPackages(ctx, params)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/architectures/:distribution", wrapper.GetArchitectures)
	router.POST("/compose", wrapper.ComposeImage)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
