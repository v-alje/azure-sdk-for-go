//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// APIPortalCustomDomainsClient contains the methods for the APIPortalCustomDomains group.
// Don't use this type directly, use NewAPIPortalCustomDomainsClient() instead.
type APIPortalCustomDomainsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPIPortalCustomDomainsClient creates a new instance of APIPortalCustomDomainsClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIPortalCustomDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *APIPortalCustomDomainsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &APIPortalCustomDomainsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create or update the API portal custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// apiPortalName - The name of API portal.
// domainName - The name of the API portal custom domain.
// apiPortalCustomDomainResource - The API portal custom domain for the create or update operation
// options - APIPortalCustomDomainsClientBeginCreateOrUpdateOptions contains the optional parameters for the APIPortalCustomDomainsClient.BeginCreateOrUpdate
// method.
func (client *APIPortalCustomDomainsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, domainName string, apiPortalCustomDomainResource APIPortalCustomDomainResource, options *APIPortalCustomDomainsClientBeginCreateOrUpdateOptions) (APIPortalCustomDomainsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, apiPortalName, domainName, apiPortalCustomDomainResource, options)
	if err != nil {
		return APIPortalCustomDomainsClientCreateOrUpdatePollerResponse{}, err
	}
	result := APIPortalCustomDomainsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("APIPortalCustomDomainsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return APIPortalCustomDomainsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &APIPortalCustomDomainsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update the API portal custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *APIPortalCustomDomainsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, domainName string, apiPortalCustomDomainResource APIPortalCustomDomainResource, options *APIPortalCustomDomainsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiPortalName, domainName, apiPortalCustomDomainResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIPortalCustomDomainsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, domainName string, apiPortalCustomDomainResource APIPortalCustomDomainResource, options *APIPortalCustomDomainsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals/{apiPortalName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiPortalName == "" {
		return nil, errors.New("parameter apiPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiPortalName}", url.PathEscape(apiPortalName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, apiPortalCustomDomainResource)
}

// BeginDelete - Delete the API portal custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// apiPortalName - The name of API portal.
// domainName - The name of the API portal custom domain.
// options - APIPortalCustomDomainsClientBeginDeleteOptions contains the optional parameters for the APIPortalCustomDomainsClient.BeginDelete
// method.
func (client *APIPortalCustomDomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, domainName string, options *APIPortalCustomDomainsClientBeginDeleteOptions) (APIPortalCustomDomainsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, apiPortalName, domainName, options)
	if err != nil {
		return APIPortalCustomDomainsClientDeletePollerResponse{}, err
	}
	result := APIPortalCustomDomainsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("APIPortalCustomDomainsClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return APIPortalCustomDomainsClientDeletePollerResponse{}, err
	}
	result.Poller = &APIPortalCustomDomainsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete the API portal custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *APIPortalCustomDomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, domainName string, options *APIPortalCustomDomainsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiPortalName, domainName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIPortalCustomDomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, domainName string, options *APIPortalCustomDomainsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals/{apiPortalName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiPortalName == "" {
		return nil, errors.New("parameter apiPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiPortalName}", url.PathEscape(apiPortalName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the API portal custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// apiPortalName - The name of API portal.
// domainName - The name of the API portal custom domain.
// options - APIPortalCustomDomainsClientGetOptions contains the optional parameters for the APIPortalCustomDomainsClient.Get
// method.
func (client *APIPortalCustomDomainsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, domainName string, options *APIPortalCustomDomainsClientGetOptions) (APIPortalCustomDomainsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiPortalName, domainName, options)
	if err != nil {
		return APIPortalCustomDomainsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIPortalCustomDomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIPortalCustomDomainsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIPortalCustomDomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, domainName string, options *APIPortalCustomDomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals/{apiPortalName}/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiPortalName == "" {
		return nil, errors.New("parameter apiPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiPortalName}", url.PathEscape(apiPortalName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIPortalCustomDomainsClient) getHandleResponse(resp *http.Response) (APIPortalCustomDomainsClientGetResponse, error) {
	result := APIPortalCustomDomainsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIPortalCustomDomainResource); err != nil {
		return APIPortalCustomDomainsClientGetResponse{}, err
	}
	return result, nil
}

// List - Handle requests to list all API portal custom domains.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// apiPortalName - The name of API portal.
// options - APIPortalCustomDomainsClientListOptions contains the optional parameters for the APIPortalCustomDomainsClient.List
// method.
func (client *APIPortalCustomDomainsClient) List(resourceGroupName string, serviceName string, apiPortalName string, options *APIPortalCustomDomainsClientListOptions) *APIPortalCustomDomainsClientListPager {
	return &APIPortalCustomDomainsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, serviceName, apiPortalName, options)
		},
		advancer: func(ctx context.Context, resp APIPortalCustomDomainsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.APIPortalCustomDomainResourceCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *APIPortalCustomDomainsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiPortalName string, options *APIPortalCustomDomainsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apiPortals/{apiPortalName}/domains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiPortalName == "" {
		return nil, errors.New("parameter apiPortalName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiPortalName}", url.PathEscape(apiPortalName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *APIPortalCustomDomainsClient) listHandleResponse(resp *http.Response) (APIPortalCustomDomainsClientListResponse, error) {
	result := APIPortalCustomDomainsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIPortalCustomDomainResourceCollection); err != nil {
		return APIPortalCustomDomainsClientListResponse{}, err
	}
	return result, nil
}
