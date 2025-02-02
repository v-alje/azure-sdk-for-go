//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// DNSForwardingRulesetsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DNSForwardingRulesetsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DNSForwardingRulesetsClientCreateOrUpdatePoller) Done() bool {
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
func (p *DNSForwardingRulesetsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DNSForwardingRulesetsClientCreateOrUpdateResponse will be returned.
func (p *DNSForwardingRulesetsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (DNSForwardingRulesetsClientCreateOrUpdateResponse, error) {
	respType := DNSForwardingRulesetsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DNSForwardingRuleset)
	if err != nil {
		return DNSForwardingRulesetsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DNSForwardingRulesetsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DNSForwardingRulesetsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type DNSForwardingRulesetsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DNSForwardingRulesetsClientDeletePoller) Done() bool {
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
func (p *DNSForwardingRulesetsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DNSForwardingRulesetsClientDeleteResponse will be returned.
func (p *DNSForwardingRulesetsClientDeletePoller) FinalResponse(ctx context.Context) (DNSForwardingRulesetsClientDeleteResponse, error) {
	respType := DNSForwardingRulesetsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return DNSForwardingRulesetsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DNSForwardingRulesetsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DNSForwardingRulesetsClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DNSForwardingRulesetsClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DNSForwardingRulesetsClientUpdatePoller) Done() bool {
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
func (p *DNSForwardingRulesetsClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DNSForwardingRulesetsClientUpdateResponse will be returned.
func (p *DNSForwardingRulesetsClientUpdatePoller) FinalResponse(ctx context.Context) (DNSForwardingRulesetsClientUpdateResponse, error) {
	respType := DNSForwardingRulesetsClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DNSForwardingRuleset)
	if err != nil {
		return DNSForwardingRulesetsClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DNSForwardingRulesetsClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DNSResolversClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DNSResolversClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DNSResolversClientCreateOrUpdatePoller) Done() bool {
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
func (p *DNSResolversClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DNSResolversClientCreateOrUpdateResponse will be returned.
func (p *DNSResolversClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (DNSResolversClientCreateOrUpdateResponse, error) {
	respType := DNSResolversClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DNSResolver)
	if err != nil {
		return DNSResolversClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DNSResolversClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DNSResolversClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type DNSResolversClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DNSResolversClientDeletePoller) Done() bool {
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
func (p *DNSResolversClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DNSResolversClientDeleteResponse will be returned.
func (p *DNSResolversClientDeletePoller) FinalResponse(ctx context.Context) (DNSResolversClientDeleteResponse, error) {
	respType := DNSResolversClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return DNSResolversClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DNSResolversClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DNSResolversClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DNSResolversClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DNSResolversClientUpdatePoller) Done() bool {
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
func (p *DNSResolversClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DNSResolversClientUpdateResponse will be returned.
func (p *DNSResolversClientUpdatePoller) FinalResponse(ctx context.Context) (DNSResolversClientUpdateResponse, error) {
	respType := DNSResolversClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DNSResolver)
	if err != nil {
		return DNSResolversClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DNSResolversClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// InboundEndpointsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type InboundEndpointsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *InboundEndpointsClientCreateOrUpdatePoller) Done() bool {
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
func (p *InboundEndpointsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final InboundEndpointsClientCreateOrUpdateResponse will be returned.
func (p *InboundEndpointsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (InboundEndpointsClientCreateOrUpdateResponse, error) {
	respType := InboundEndpointsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.InboundEndpoint)
	if err != nil {
		return InboundEndpointsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *InboundEndpointsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// InboundEndpointsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type InboundEndpointsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *InboundEndpointsClientDeletePoller) Done() bool {
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
func (p *InboundEndpointsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final InboundEndpointsClientDeleteResponse will be returned.
func (p *InboundEndpointsClientDeletePoller) FinalResponse(ctx context.Context) (InboundEndpointsClientDeleteResponse, error) {
	respType := InboundEndpointsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return InboundEndpointsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *InboundEndpointsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// InboundEndpointsClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type InboundEndpointsClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *InboundEndpointsClientUpdatePoller) Done() bool {
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
func (p *InboundEndpointsClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final InboundEndpointsClientUpdateResponse will be returned.
func (p *InboundEndpointsClientUpdatePoller) FinalResponse(ctx context.Context) (InboundEndpointsClientUpdateResponse, error) {
	respType := InboundEndpointsClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.InboundEndpoint)
	if err != nil {
		return InboundEndpointsClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *InboundEndpointsClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// OutboundEndpointsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type OutboundEndpointsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *OutboundEndpointsClientCreateOrUpdatePoller) Done() bool {
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
func (p *OutboundEndpointsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final OutboundEndpointsClientCreateOrUpdateResponse will be returned.
func (p *OutboundEndpointsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (OutboundEndpointsClientCreateOrUpdateResponse, error) {
	respType := OutboundEndpointsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OutboundEndpoint)
	if err != nil {
		return OutboundEndpointsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *OutboundEndpointsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// OutboundEndpointsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type OutboundEndpointsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *OutboundEndpointsClientDeletePoller) Done() bool {
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
func (p *OutboundEndpointsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final OutboundEndpointsClientDeleteResponse will be returned.
func (p *OutboundEndpointsClientDeletePoller) FinalResponse(ctx context.Context) (OutboundEndpointsClientDeleteResponse, error) {
	respType := OutboundEndpointsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return OutboundEndpointsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *OutboundEndpointsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// OutboundEndpointsClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type OutboundEndpointsClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *OutboundEndpointsClientUpdatePoller) Done() bool {
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
func (p *OutboundEndpointsClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final OutboundEndpointsClientUpdateResponse will be returned.
func (p *OutboundEndpointsClientUpdatePoller) FinalResponse(ctx context.Context) (OutboundEndpointsClientUpdateResponse, error) {
	respType := OutboundEndpointsClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.OutboundEndpoint)
	if err != nil {
		return OutboundEndpointsClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *OutboundEndpointsClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualNetworkLinksClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualNetworkLinksClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualNetworkLinksClientCreateOrUpdatePoller) Done() bool {
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
func (p *VirtualNetworkLinksClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualNetworkLinksClientCreateOrUpdateResponse will be returned.
func (p *VirtualNetworkLinksClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VirtualNetworkLinksClientCreateOrUpdateResponse, error) {
	respType := VirtualNetworkLinksClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VirtualNetworkLink)
	if err != nil {
		return VirtualNetworkLinksClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualNetworkLinksClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualNetworkLinksClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type VirtualNetworkLinksClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualNetworkLinksClientDeletePoller) Done() bool {
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
func (p *VirtualNetworkLinksClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualNetworkLinksClientDeleteResponse will be returned.
func (p *VirtualNetworkLinksClientDeletePoller) FinalResponse(ctx context.Context) (VirtualNetworkLinksClientDeleteResponse, error) {
	respType := VirtualNetworkLinksClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualNetworkLinksClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualNetworkLinksClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualNetworkLinksClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualNetworkLinksClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualNetworkLinksClientUpdatePoller) Done() bool {
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
func (p *VirtualNetworkLinksClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualNetworkLinksClientUpdateResponse will be returned.
func (p *VirtualNetworkLinksClientUpdatePoller) FinalResponse(ctx context.Context) (VirtualNetworkLinksClientUpdateResponse, error) {
	respType := VirtualNetworkLinksClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VirtualNetworkLink)
	if err != nil {
		return VirtualNetworkLinksClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualNetworkLinksClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
