//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ClientCreateResourcePoller provides polling facilities until the operation reaches a terminal state.
type ClientCreateResourcePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClientCreateResourcePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ClientCreateResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClientCreateResourceResponse will be returned.
func (p *ClientCreateResourcePoller) FinalResponse(ctx context.Context) (ClientCreateResourceResponse, error) {
	respType := ClientCreateResourceResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Resource)
	if err != nil {
		return ClientCreateResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClientCreateResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type ClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClientDeleteResponse will be returned.
func (p *ClientDeletePoller) FinalResponse(ctx context.Context) (ClientDeleteResponse, error) {
	respType := ClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ClientUpdateResourcePoller provides polling facilities until the operation reaches a terminal state.
type ClientUpdateResourcePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClientUpdateResourcePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ClientUpdateResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClientUpdateResourceResponse will be returned.
func (p *ClientUpdateResourcePoller) FinalResponse(ctx context.Context) (ClientUpdateResourceResponse, error) {
	respType := ClientUpdateResourceResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Resource)
	if err != nil {
		return ClientUpdateResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClientUpdateResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// OperationClientGetPoller provides polling facilities until the operation reaches a terminal state.
type OperationClientGetPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *OperationClientGetPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *OperationClientGetPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final OperationClientGetResponse will be returned.
func (p *OperationClientGetPoller) FinalResponse(ctx context.Context) (OperationClientGetResponse, error) {
	respType := OperationClientGetResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Resource)
	if err != nil {
		return OperationClientGetResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *OperationClientGetPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionLevelClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionLevelClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionLevelClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionLevelClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionLevelClientCreateOrUpdateResponse will be returned.
func (p *SubscriptionLevelClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (SubscriptionLevelClientCreateOrUpdateResponse, error) {
	respType := SubscriptionLevelClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Resource)
	if err != nil {
		return SubscriptionLevelClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionLevelClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionLevelClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionLevelClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionLevelClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionLevelClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionLevelClientDeleteResponse will be returned.
func (p *SubscriptionLevelClientDeletePoller) FinalResponse(ctx context.Context) (SubscriptionLevelClientDeleteResponse, error) {
	respType := SubscriptionLevelClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return SubscriptionLevelClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionLevelClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionLevelClientMoveResourcesPoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionLevelClientMoveResourcesPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionLevelClientMoveResourcesPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionLevelClientMoveResourcesPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionLevelClientMoveResourcesResponse will be returned.
func (p *SubscriptionLevelClientMoveResourcesPoller) FinalResponse(ctx context.Context) (SubscriptionLevelClientMoveResourcesResponse, error) {
	respType := SubscriptionLevelClientMoveResourcesResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return SubscriptionLevelClientMoveResourcesResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionLevelClientMoveResourcesPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionLevelClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionLevelClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionLevelClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionLevelClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionLevelClientUpdateResponse will be returned.
func (p *SubscriptionLevelClientUpdatePoller) FinalResponse(ctx context.Context) (SubscriptionLevelClientUpdateResponse, error) {
	respType := SubscriptionLevelClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Resource)
	if err != nil {
		return SubscriptionLevelClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionLevelClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionLevelClientUpdateToUnsubscribedPoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionLevelClientUpdateToUnsubscribedPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionLevelClientUpdateToUnsubscribedPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionLevelClientUpdateToUnsubscribedPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionLevelClientUpdateToUnsubscribedResponse will be returned.
func (p *SubscriptionLevelClientUpdateToUnsubscribedPoller) FinalResponse(ctx context.Context) (SubscriptionLevelClientUpdateToUnsubscribedResponse, error) {
	respType := SubscriptionLevelClientUpdateToUnsubscribedResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return SubscriptionLevelClientUpdateToUnsubscribedResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionLevelClientUpdateToUnsubscribedPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
