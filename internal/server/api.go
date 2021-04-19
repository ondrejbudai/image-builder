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

// AWSUploadRequest defines model for AWSUploadRequest.
type AWSUploadRequest struct {
	Options AWSUploadRequestOptions `json:"options"`
	Type    string                  `json:"type"`
}

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// AzureUploadRequest defines model for AzureUploadRequest.
type AzureUploadRequest struct {
	Options AzureUploadRequestOptions `json:"options"`
	Type    string                    `json:"type"`
}

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

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
	ImageStatus ImageStatus `json:"image_status"`
}

// ComposesResponse defines model for ComposesResponse.
type ComposesResponse struct {
	Data  []ComposesResponseItem `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// ComposesResponseItem defines model for ComposesResponseItem.
type ComposesResponseItem struct {
	CreatedAt string      `json:"created_at"`
	Id        string      `json:"id"`
	Request   interface{} `json:"request"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions []DistributionItem

// GCPUploadRequest defines model for GCPUploadRequest.
type GCPUploadRequest struct {
	Options GCPUploadRequestOptions `json:"options"`
	Type    string                  `json:"type"`
}

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	//     If not specified, the imported Compute Node image is not shared with any
	//     account.
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

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
	Architecture  string        `json:"architecture"`
	ImageType     string        `json:"image_type"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       string        `json:"status"`
	UploadStatus *UploadStatus `json:"upload_status,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Version string `json:"version"`
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

// Readiness defines model for Readiness.
type Readiness struct {
	Readiness string `json:"readiness"`
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
type UploadRequest interface{}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{} `json:"options"`
	Status  string      `json:"status"`
	Type    UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// List of UploadTypes
const (
	UploadTypes_aws   UploadTypes = "aws"
	UploadTypes_azure UploadTypes = "azure"
	UploadTypes_gcp   UploadTypes = "gcp"
)

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetComposesParams defines parameters for GetComposes.
type GetComposesParams struct {

	// max amount of composes, default 100
	Limit *int `json:"limit,omitempty"`

	// composes page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

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
	// get a collection of previous compose requests for the logged in user
	// (GET /composes)
	GetComposes(ctx echo.Context, params GetComposesParams) error
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
	// return the readiness
	// (GET /ready)
	GetReadiness(ctx echo.Context) error
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

// GetComposes converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetComposesParams
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
	err = w.Handler.GetComposes(ctx, params)
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

// GetReadiness converts echo context to params.
func (w *ServerInterfaceWrapper) GetReadiness(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReadiness(ctx)
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
	router.GET("/composes", wrapper.GetComposes)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/ready", wrapper.GetReadiness)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"TPlKhNNvfYju5Dik2m7/FwAA//8xMJhBqS0AAA==",
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
