//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// EmailEventsClient contains the methods for the EmailEvents group.
// Don't use this type directly, use NewEmailEventsClient() instead.
type EmailEventsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEmailEventsClient creates a new instance of EmailEventsClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEmailEventsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *EmailEventsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &EmailEventsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets a email event of a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// emailEventResourceName - The resource name of an email event.
// options - EmailEventsClientGetOptions contains the optional parameters for the EmailEventsClient.Get method.
func (client *EmailEventsClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, emailEventResourceName string, options *EmailEventsClientGetOptions) (EmailEventsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, emailEventResourceName, options)
	if err != nil {
		return EmailEventsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EmailEventsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmailEventsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EmailEventsClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, emailEventResourceName string, options *EmailEventsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/emailEvents/{emailEventResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if emailEventResourceName == "" {
		return nil, errors.New("parameter emailEventResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailEventResourceName}", url.PathEscape(emailEventResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EmailEventsClient) getHandleResponse(resp *http.Response) (EmailEventsClientGetResponse, error) {
	result := EmailEventsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailEventResource); err != nil {
		return EmailEventsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all the email events of a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource.
// testBaseAccountName - The resource name of the Test Base Account.
// options - EmailEventsClientListOptions contains the optional parameters for the EmailEventsClient.List method.
func (client *EmailEventsClient) List(resourceGroupName string, testBaseAccountName string, options *EmailEventsClientListOptions) *EmailEventsClientListPager {
	return &EmailEventsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
		},
		advancer: func(ctx context.Context, resp EmailEventsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EmailEventListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *EmailEventsClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *EmailEventsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/emailEvents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EmailEventsClient) listHandleResponse(resp *http.Response) (EmailEventsClientListResponse, error) {
	result := EmailEventsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailEventListResult); err != nil {
		return EmailEventsClientListResponse{}, err
	}
	return result, nil
}
