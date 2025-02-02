//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// WorkflowRunActionsClient contains the methods for the WorkflowRunActions group.
// Don't use this type directly, use NewWorkflowRunActionsClient() instead.
type WorkflowRunActionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowRunActionsClient creates a new instance of WorkflowRunActionsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowRunActionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkflowRunActionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &WorkflowRunActionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets a workflow run action.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// actionName - The workflow action name.
// options - WorkflowRunActionsClientGetOptions contains the optional parameters for the WorkflowRunActionsClient.Get method.
func (client *WorkflowRunActionsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionsClientGetOptions) (WorkflowRunActionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, options)
	if err != nil {
		return WorkflowRunActionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunActionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowRunActionsClient) getHandleResponse(resp *http.Response) (WorkflowRunActionsClientGetResponse, error) {
	result := WorkflowRunActionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunAction); err != nil {
		return WorkflowRunActionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of workflow run actions.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// options - WorkflowRunActionsClientListOptions contains the optional parameters for the WorkflowRunActionsClient.List method.
func (client *WorkflowRunActionsClient) List(resourceGroupName string, workflowName string, runName string, options *WorkflowRunActionsClientListOptions) *WorkflowRunActionsClientListPager {
	return &WorkflowRunActionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workflowName, runName, options)
		},
		advancer: func(ctx context.Context, resp WorkflowRunActionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkflowRunActionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkflowRunActionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunActionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowRunActionsClient) listHandleResponse(resp *http.Response) (WorkflowRunActionsClientListResponse, error) {
	result := WorkflowRunActionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunActionListResult); err != nil {
		return WorkflowRunActionsClientListResponse{}, err
	}
	return result, nil
}

// ListExpressionTraces - Lists a workflow run expression trace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// runName - The workflow run name.
// actionName - The workflow action name.
// options - WorkflowRunActionsClientListExpressionTracesOptions contains the optional parameters for the WorkflowRunActionsClient.ListExpressionTraces
// method.
func (client *WorkflowRunActionsClient) ListExpressionTraces(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionsClientListExpressionTracesOptions) (WorkflowRunActionsClientListExpressionTracesResponse, error) {
	req, err := client.listExpressionTracesCreateRequest(ctx, resourceGroupName, workflowName, runName, actionName, options)
	if err != nil {
		return WorkflowRunActionsClientListExpressionTracesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunActionsClientListExpressionTracesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunActionsClientListExpressionTracesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listExpressionTracesHandleResponse(resp)
}

// listExpressionTracesCreateRequest creates the ListExpressionTraces request.
func (client *WorkflowRunActionsClient) listExpressionTracesCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, options *WorkflowRunActionsClientListExpressionTracesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/listExpressionTraces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	if actionName == "" {
		return nil, errors.New("parameter actionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionName}", url.PathEscape(actionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listExpressionTracesHandleResponse handles the ListExpressionTraces response.
func (client *WorkflowRunActionsClient) listExpressionTracesHandleResponse(resp *http.Response) (WorkflowRunActionsClientListExpressionTracesResponse, error) {
	result := WorkflowRunActionsClientListExpressionTracesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressionTraces); err != nil {
		return WorkflowRunActionsClientListExpressionTracesResponse{}, err
	}
	return result, nil
}
