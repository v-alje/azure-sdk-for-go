//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

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
	"strconv"
	"strings"
)

// DNSResolversClient contains the methods for the DNSResolvers group.
// Don't use this type directly, use NewDNSResolversClient() instead.
type DNSResolversClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDNSResolversClient creates a new instance of DNSResolversClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDNSResolversClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DNSResolversClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &DNSResolversClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// parameters - Parameters supplied to the CreateOrUpdate operation.
// options - DNSResolversClientBeginCreateOrUpdateOptions contains the optional parameters for the DNSResolversClient.BeginCreateOrUpdate
// method.
func (client *DNSResolversClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, dnsResolverName string, parameters DNSResolver, options *DNSResolversClientBeginCreateOrUpdateOptions) (DNSResolversClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, dnsResolverName, parameters, options)
	if err != nil {
		return DNSResolversClientCreateOrUpdatePollerResponse{}, err
	}
	result := DNSResolversClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DNSResolversClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return DNSResolversClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DNSResolversClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DNSResolversClient) createOrUpdate(ctx context.Context, resourceGroupName string, dnsResolverName string, parameters DNSResolver, options *DNSResolversClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dnsResolverName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DNSResolversClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, parameters DNSResolver, options *DNSResolversClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a DNS resolver. WARNING: This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// options - DNSResolversClientBeginDeleteOptions contains the optional parameters for the DNSResolversClient.BeginDelete
// method.
func (client *DNSResolversClient) BeginDelete(ctx context.Context, resourceGroupName string, dnsResolverName string, options *DNSResolversClientBeginDeleteOptions) (DNSResolversClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, dnsResolverName, options)
	if err != nil {
		return DNSResolversClientDeletePollerResponse{}, err
	}
	result := DNSResolversClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DNSResolversClient.Delete", "", resp, client.pl)
	if err != nil {
		return DNSResolversClientDeletePollerResponse{}, err
	}
	result.Poller = &DNSResolversClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a DNS resolver. WARNING: This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DNSResolversClient) deleteOperation(ctx context.Context, resourceGroupName string, dnsResolverName string, options *DNSResolversClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dnsResolverName, options)
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
func (client *DNSResolversClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, options *DNSResolversClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets properties of a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// options - DNSResolversClientGetOptions contains the optional parameters for the DNSResolversClient.Get method.
func (client *DNSResolversClient) Get(ctx context.Context, resourceGroupName string, dnsResolverName string, options *DNSResolversClientGetOptions) (DNSResolversClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dnsResolverName, options)
	if err != nil {
		return DNSResolversClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DNSResolversClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DNSResolversClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DNSResolversClient) getCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, options *DNSResolversClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DNSResolversClient) getHandleResponse(resp *http.Response) (DNSResolversClientGetResponse, error) {
	result := DNSResolversClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DNSResolver); err != nil {
		return DNSResolversClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists DNS resolvers in all resource groups of a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DNSResolversClientListOptions contains the optional parameters for the DNSResolversClient.List method.
func (client *DNSResolversClient) List(options *DNSResolversClientListOptions) *DNSResolversClientListPager {
	return &DNSResolversClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DNSResolversClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DNSResolversClient) listCreateRequest(ctx context.Context, options *DNSResolversClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsResolvers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DNSResolversClient) listHandleResponse(resp *http.Response) (DNSResolversClientListResponse, error) {
	result := DNSResolversClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return DNSResolversClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists DNS resolvers within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - DNSResolversClientListByResourceGroupOptions contains the optional parameters for the DNSResolversClient.ListByResourceGroup
// method.
func (client *DNSResolversClient) ListByResourceGroup(resourceGroupName string, options *DNSResolversClientListByResourceGroupOptions) *DNSResolversClientListByResourceGroupPager {
	return &DNSResolversClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DNSResolversClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DNSResolversClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DNSResolversClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DNSResolversClient) listByResourceGroupHandleResponse(resp *http.Response) (DNSResolversClientListByResourceGroupResponse, error) {
	result := DNSResolversClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return DNSResolversClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListByVirtualNetwork - Lists DNS resolver resource IDs linked to a virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// virtualNetworkName - The name of the virtual network.
// options - DNSResolversClientListByVirtualNetworkOptions contains the optional parameters for the DNSResolversClient.ListByVirtualNetwork
// method.
func (client *DNSResolversClient) ListByVirtualNetwork(resourceGroupName string, virtualNetworkName string, options *DNSResolversClientListByVirtualNetworkOptions) *DNSResolversClientListByVirtualNetworkPager {
	return &DNSResolversClientListByVirtualNetworkPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVirtualNetworkCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
		},
		advancer: func(ctx context.Context, resp DNSResolversClientListByVirtualNetworkResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SubResourceListResult.NextLink)
		},
	}
}

// listByVirtualNetworkCreateRequest creates the ListByVirtualNetwork request.
func (client *DNSResolversClient) listByVirtualNetworkCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *DNSResolversClientListByVirtualNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/listDnsResolvers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByVirtualNetworkHandleResponse handles the ListByVirtualNetwork response.
func (client *DNSResolversClient) listByVirtualNetworkHandleResponse(resp *http.Response) (DNSResolversClientListByVirtualNetworkResponse, error) {
	result := DNSResolversClientListByVirtualNetworkResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubResourceListResult); err != nil {
		return DNSResolversClientListByVirtualNetworkResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// parameters - Parameters supplied to the Update operation.
// options - DNSResolversClientBeginUpdateOptions contains the optional parameters for the DNSResolversClient.BeginUpdate
// method.
func (client *DNSResolversClient) BeginUpdate(ctx context.Context, resourceGroupName string, dnsResolverName string, parameters Patch, options *DNSResolversClientBeginUpdateOptions) (DNSResolversClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, dnsResolverName, parameters, options)
	if err != nil {
		return DNSResolversClientUpdatePollerResponse{}, err
	}
	result := DNSResolversClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DNSResolversClient.Update", "", resp, client.pl)
	if err != nil {
		return DNSResolversClientUpdatePollerResponse{}, err
	}
	result.Poller = &DNSResolversClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DNSResolversClient) update(ctx context.Context, resourceGroupName string, dnsResolverName string, parameters Patch, options *DNSResolversClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dnsResolverName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DNSResolversClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, parameters Patch, options *DNSResolversClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
