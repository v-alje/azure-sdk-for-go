// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ServerAdministratorsClient contains the methods for the ServerAdministrators group.
// Don't use this type directly, use NewServerAdministratorsClient() instead.
type ServerAdministratorsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewServerAdministratorsClient creates a new instance of ServerAdministratorsClient with the specified values.
func NewServerAdministratorsClient(con *armcore.Connection, subscriptionID string) *ServerAdministratorsClient {
	return &ServerAdministratorsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or update active directory administrator on an existing server. The update action will overwrite the existing administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsBeginCreateOrUpdateOptions) (ServerAdministratorsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, properties, options)
	if err != nil {
		return ServerAdministratorsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerAdministratorsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServerAdministratorsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ServerAdministratorsCreateOrUpdatePollerResponse{}, err
	}
	poller := &serverAdministratorsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerAdministratorsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ServerAdministratorsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ServerAdministratorsCreateOrUpdatePoller.ResumeToken().
func (client *ServerAdministratorsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ServerAdministratorsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServerAdministratorsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ServerAdministratorsCreateOrUpdatePollerResponse{}, err
	}
	poller := &serverAdministratorsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServerAdministratorsCreateOrUpdatePollerResponse{}, err
	}
	result := ServerAdministratorsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerAdministratorsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or update active directory administrator on an existing server. The update action will overwrite the existing administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerAdministratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(properties)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServerAdministratorsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes server active directory administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsBeginDeleteOptions) (ServerAdministratorsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdministratorsDeletePollerResponse{}, err
	}
	result := ServerAdministratorsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServerAdministratorsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ServerAdministratorsDeletePollerResponse{}, err
	}
	poller := &serverAdministratorsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerAdministratorsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new ServerAdministratorsDeletePoller from the specified resume token.
// token - The value must come from a previous call to ServerAdministratorsDeletePoller.ResumeToken().
func (client *ServerAdministratorsClient) ResumeDelete(ctx context.Context, token string) (ServerAdministratorsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServerAdministratorsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ServerAdministratorsDeletePollerResponse{}, err
	}
	poller := &serverAdministratorsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServerAdministratorsDeletePollerResponse{}, err
	}
	result := ServerAdministratorsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServerAdministratorsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes server active directory administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServerAdministratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServerAdministratorsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets information about a AAD server administrator.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsGetOptions) (ServerAdministratorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdministratorsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServerAdministratorsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServerAdministratorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAdministratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAdministratorsClient) getHandleResponse(resp *azcore.Response) (ServerAdministratorsGetResponse, error) {
	result := ServerAdministratorsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ServerAdministratorResource); err != nil {
		return ServerAdministratorsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerAdministratorsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Returns a list of server Administrators.
// If the operation fails it returns the *CloudError error type.
func (client *ServerAdministratorsClient) List(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsListOptions) (ServerAdministratorsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdministratorsListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServerAdministratorsListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServerAdministratorsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServerAdministratorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/administrators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServerAdministratorsClient) listHandleResponse(resp *azcore.Response) (ServerAdministratorsListResponse, error) {
	result := ServerAdministratorsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ServerAdministratorResourceListResult); err != nil {
		return ServerAdministratorsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ServerAdministratorsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}