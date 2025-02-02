//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

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

// AvailabilityStatusesClient contains the methods for the AvailabilityStatuses group.
// Don't use this type directly, use NewAvailabilityStatusesClient() instead.
type AvailabilityStatusesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailabilityStatusesClient creates a new instance of AvailabilityStatusesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailabilityStatusesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AvailabilityStatusesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AvailabilityStatusesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetByResource - Gets current availability status for a single resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The fully qualified ID of the resource, including the resource name and resource type. Currently the API
// support not nested and one nesting level resource types :
// /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
// and
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
// options - AvailabilityStatusesClientGetByResourceOptions contains the optional parameters for the AvailabilityStatusesClient.GetByResource
// method.
func (client *AvailabilityStatusesClient) GetByResource(ctx context.Context, resourceURI string, options *AvailabilityStatusesClientGetByResourceOptions) (AvailabilityStatusesClientGetByResourceResponse, error) {
	req, err := client.getByResourceCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return AvailabilityStatusesClientGetByResourceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AvailabilityStatusesClientGetByResourceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailabilityStatusesClientGetByResourceResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByResourceHandleResponse(resp)
}

// getByResourceCreateRequest creates the GetByResource request.
func (client *AvailabilityStatusesClient) getByResourceCreateRequest(ctx context.Context, resourceURI string, options *AvailabilityStatusesClientGetByResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses/current"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByResourceHandleResponse handles the GetByResource response.
func (client *AvailabilityStatusesClient) getByResourceHandleResponse(resp *http.Response) (AvailabilityStatusesClientGetByResourceResponse, error) {
	result := AvailabilityStatusesClientGetByResourceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityStatus); err != nil {
		return AvailabilityStatusesClientGetByResourceResponse{}, err
	}
	return result, nil
}

// List - Lists all historical availability transitions and impacting events for a single resource. Use the nextLink property
// in the response to get the next page of availability status
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The fully qualified ID of the resource, including the resource name and resource type. Currently the API
// support not nested and one nesting level resource types :
// /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
// and
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
// options - AvailabilityStatusesClientListOptions contains the optional parameters for the AvailabilityStatusesClient.List
// method.
func (client *AvailabilityStatusesClient) List(resourceURI string, options *AvailabilityStatusesClientListOptions) *AvailabilityStatusesClientListPager {
	return &AvailabilityStatusesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceURI, options)
		},
		advancer: func(ctx context.Context, resp AvailabilityStatusesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailabilityStatusListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AvailabilityStatusesClient) listCreateRequest(ctx context.Context, resourceURI string, options *AvailabilityStatusesClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailabilityStatusesClient) listHandleResponse(resp *http.Response) (AvailabilityStatusesClientListResponse, error) {
	result := AvailabilityStatusesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityStatusListResult); err != nil {
		return AvailabilityStatusesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists the current availability status for all the resources in the resource group. Use the nextLink
// property in the response to get the next page of availability statuses.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - AvailabilityStatusesClientListByResourceGroupOptions contains the optional parameters for the AvailabilityStatusesClient.ListByResourceGroup
// method.
func (client *AvailabilityStatusesClient) ListByResourceGroup(resourceGroupName string, options *AvailabilityStatusesClientListByResourceGroupOptions) *AvailabilityStatusesClientListByResourceGroupPager {
	return &AvailabilityStatusesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AvailabilityStatusesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailabilityStatusListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AvailabilityStatusesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AvailabilityStatusesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceHealth/availabilityStatuses"
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
	reqQP.Set("api-version", "2017-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AvailabilityStatusesClient) listByResourceGroupHandleResponse(resp *http.Response) (AvailabilityStatusesClientListByResourceGroupResponse, error) {
	result := AvailabilityStatusesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityStatusListResult); err != nil {
		return AvailabilityStatusesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscriptionID - Lists the current availability status for all the resources in the subscription. Use the nextLink
// property in the response to get the next page of availability statuses.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AvailabilityStatusesClientListBySubscriptionIDOptions contains the optional parameters for the AvailabilityStatusesClient.ListBySubscriptionID
// method.
func (client *AvailabilityStatusesClient) ListBySubscriptionID(options *AvailabilityStatusesClientListBySubscriptionIDOptions) *AvailabilityStatusesClientListBySubscriptionIDPager {
	return &AvailabilityStatusesClientListBySubscriptionIDPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionIDCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AvailabilityStatusesClientListBySubscriptionIDResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AvailabilityStatusListResult.NextLink)
		},
	}
}

// listBySubscriptionIDCreateRequest creates the ListBySubscriptionID request.
func (client *AvailabilityStatusesClient) listBySubscriptionIDCreateRequest(ctx context.Context, options *AvailabilityStatusesClientListBySubscriptionIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/availabilityStatuses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionIDHandleResponse handles the ListBySubscriptionID response.
func (client *AvailabilityStatusesClient) listBySubscriptionIDHandleResponse(resp *http.Response) (AvailabilityStatusesClientListBySubscriptionIDResponse, error) {
	result := AvailabilityStatusesClientListBySubscriptionIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityStatusListResult); err != nil {
		return AvailabilityStatusesClientListBySubscriptionIDResponse{}, err
	}
	return result, nil
}
