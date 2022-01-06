//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsiot

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
func NewServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ServicesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckDeviceServiceNameAvailability - Check if a Windows IoT Device Service name is available.
// If the operation fails it returns the *ErrorDetails error type.
func (client *ServicesClient) CheckDeviceServiceNameAvailability(ctx context.Context, deviceServiceCheckNameAvailabilityParameters DeviceServiceCheckNameAvailabilityParameters, options *ServicesCheckDeviceServiceNameAvailabilityOptions) (ServicesCheckDeviceServiceNameAvailabilityResponse, error) {
	req, err := client.checkDeviceServiceNameAvailabilityCreateRequest(ctx, deviceServiceCheckNameAvailabilityParameters, options)
	if err != nil {
		return ServicesCheckDeviceServiceNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesCheckDeviceServiceNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesCheckDeviceServiceNameAvailabilityResponse{}, client.checkDeviceServiceNameAvailabilityHandleError(resp)
	}
	return client.checkDeviceServiceNameAvailabilityHandleResponse(resp)
}

// checkDeviceServiceNameAvailabilityCreateRequest creates the CheckDeviceServiceNameAvailability request.
func (client *ServicesClient) checkDeviceServiceNameAvailabilityCreateRequest(ctx context.Context, deviceServiceCheckNameAvailabilityParameters DeviceServiceCheckNameAvailabilityParameters, options *ServicesCheckDeviceServiceNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.WindowsIoT/checkDeviceServiceNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, deviceServiceCheckNameAvailabilityParameters)
}

// checkDeviceServiceNameAvailabilityHandleResponse handles the CheckDeviceServiceNameAvailability response.
func (client *ServicesClient) checkDeviceServiceNameAvailabilityHandleResponse(resp *http.Response) (ServicesCheckDeviceServiceNameAvailabilityResponse, error) {
	result := ServicesCheckDeviceServiceNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceServiceNameAvailabilityInfo); err != nil {
		return ServicesCheckDeviceServiceNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkDeviceServiceNameAvailabilityHandleError handles the CheckDeviceServiceNameAvailability error response.
func (client *ServicesClient) checkDeviceServiceNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CreateOrUpdate - Create or update the metadata of a Windows IoT Device Service. The usual pattern to modify a property is to retrieve the Windows IoT
// Device Service metadata and security metadata, and then combine
// them with the modified values in a new body to update the Windows IoT Device Service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *ServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, deviceName string, deviceService DeviceService, options *ServicesCreateOrUpdateOptions) (ServicesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, deviceName, deviceService, options)
	if err != nil {
		return ServicesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ServicesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, deviceService DeviceService, options *ServicesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, deviceService)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServicesClient) createOrUpdateHandleResponse(resp *http.Response) (ServicesCreateOrUpdateResponse, error) {
	result := ServicesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceService); err != nil {
		return ServicesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServicesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete a Windows IoT Device Service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *ServicesClient) Delete(ctx context.Context, resourceGroupName string, deviceName string, options *ServicesDeleteOptions) (ServicesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, deviceName, options)
	if err != nil {
		return ServicesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ServicesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, options *ServicesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ServicesClient) deleteHandleResponse(resp *http.Response) (ServicesDeleteResponse, error) {
	result := ServicesDeleteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceService); err != nil {
		return ServicesDeleteResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServicesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the non-security related metadata of a Windows IoT Device Service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *ServicesClient) Get(ctx context.Context, resourceGroupName string, deviceName string, options *ServicesGetOptions) (ServicesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, deviceName, options)
	if err != nil {
		return ServicesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, options *ServicesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesGetResponse, error) {
	result := ServicesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceService); err != nil {
		return ServicesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServicesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get all the IoT hubs in a subscription.
// If the operation fails it returns the *ErrorDetails error type.
func (client *ServicesClient) List(options *ServicesListOptions) *ServicesListPager {
	return &ServicesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ServicesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DeviceServiceDescriptionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServicesClient) listCreateRequest(ctx context.Context, options *ServicesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.WindowsIoT/deviceServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServicesClient) listHandleResponse(resp *http.Response) (ServicesListResponse, error) {
	result := ServicesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceServiceDescriptionListResult); err != nil {
		return ServicesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ServicesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Get all the IoT hubs in a resource group.
// If the operation fails it returns the *ErrorDetails error type.
func (client *ServicesClient) ListByResourceGroup(resourceGroupName string, options *ServicesListByResourceGroupOptions) *ServicesListByResourceGroupPager {
	return &ServicesListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ServicesListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DeviceServiceDescriptionListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServicesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ServicesListByResourceGroupResponse, error) {
	result := ServicesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceServiceDescriptionListResult); err != nil {
		return ServicesListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ServicesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates the metadata of a Windows IoT Device Service. The usual pattern to modify a property is to retrieve the Windows IoT Device Service metadata
// and security metadata, and then combine them with
// the modified values in a new body to update the Windows IoT Device Service.
// If the operation fails it returns the *ErrorDetails error type.
func (client *ServicesClient) Update(ctx context.Context, resourceGroupName string, deviceName string, deviceService DeviceService, options *ServicesUpdateOptions) (ServicesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, deviceName, deviceService, options)
	if err != nil {
		return ServicesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, deviceService DeviceService, options *ServicesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, deviceService)
}

// updateHandleResponse handles the Update response.
func (client *ServicesClient) updateHandleResponse(resp *http.Response) (ServicesUpdateResponse, error) {
	result := ServicesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceService); err != nil {
		return ServicesUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *ServicesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}