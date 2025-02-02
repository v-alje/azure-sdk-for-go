# Release History

## 0.2.1 (2022-02-22)

### Other Changes

- Remove the go_mod_tidy_hack.go file.

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*IotDpsResourceClient.Get` parameter(s) have been changed from `(context.Context, string, string, *IotDpsResourceGetOptions)` to `(context.Context, string, string, *IotDpsResourceClientGetOptions)`
- Function `*IotDpsResourceClient.Get` return value(s) have been changed from `(IotDpsResourceGetResponse, error)` to `(IotDpsResourceClientGetResponse, error)`
- Function `*IotDpsResourceClient.GetPrivateEndpointConnection` parameter(s) have been changed from `(context.Context, string, string, string, *IotDpsResourceGetPrivateEndpointConnectionOptions)` to `(context.Context, string, string, string, *IotDpsResourceClientGetPrivateEndpointConnectionOptions)`
- Function `*IotDpsResourceClient.GetPrivateEndpointConnection` return value(s) have been changed from `(IotDpsResourceGetPrivateEndpointConnectionResponse, error)` to `(IotDpsResourceClientGetPrivateEndpointConnectionResponse, error)`
- Function `*DpsCertificateClient.Delete` parameter(s) have been changed from `(context.Context, string, string, string, string, *DpsCertificateDeleteOptions)` to `(context.Context, string, string, string, string, *DpsCertificateClientDeleteOptions)`
- Function `*DpsCertificateClient.Delete` return value(s) have been changed from `(DpsCertificateDeleteResponse, error)` to `(DpsCertificateClientDeleteResponse, error)`
- Function `*IotDpsResourceClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, ProvisioningServiceDescription, *IotDpsResourceBeginCreateOrUpdateOptions)` to `(context.Context, string, string, ProvisioningServiceDescription, *IotDpsResourceClientBeginCreateOrUpdateOptions)`
- Function `*IotDpsResourceClient.BeginCreateOrUpdate` return value(s) have been changed from `(IotDpsResourceCreateOrUpdatePollerResponse, error)` to `(IotDpsResourceClientCreateOrUpdatePollerResponse, error)`
- Function `*IotDpsResourceClient.ListByResourceGroup` parameter(s) have been changed from `(string, *IotDpsResourceListByResourceGroupOptions)` to `(string, *IotDpsResourceClientListByResourceGroupOptions)`
- Function `*IotDpsResourceClient.ListByResourceGroup` return value(s) have been changed from `(*IotDpsResourceListByResourceGroupPager)` to `(*IotDpsResourceClientListByResourceGroupPager)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(*OperationsListOptions)` to `(*OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(*OperationsListPager)` to `(*OperationsClientListPager)`
- Function `*DpsCertificateClient.List` parameter(s) have been changed from `(context.Context, string, string, *DpsCertificateListOptions)` to `(context.Context, string, string, *DpsCertificateClientListOptions)`
- Function `*DpsCertificateClient.List` return value(s) have been changed from `(DpsCertificateListResponse, error)` to `(DpsCertificateClientListResponse, error)`
- Function `*DpsCertificateClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *DpsCertificateGetOptions)` to `(context.Context, string, string, string, *DpsCertificateClientGetOptions)`
- Function `*DpsCertificateClient.Get` return value(s) have been changed from `(DpsCertificateGetResponse, error)` to `(DpsCertificateClientGetResponse, error)`
- Function `*DpsCertificateClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, CertificateBodyDescription, *DpsCertificateCreateOrUpdateOptions)` to `(context.Context, string, string, string, CertificateBodyDescription, *DpsCertificateClientCreateOrUpdateOptions)`
- Function `*DpsCertificateClient.CreateOrUpdate` return value(s) have been changed from `(DpsCertificateCreateOrUpdateResponse, error)` to `(DpsCertificateClientCreateOrUpdateResponse, error)`
- Function `*IotDpsResourceClient.ListValidSKUs` parameter(s) have been changed from `(string, string, *IotDpsResourceListValidSKUsOptions)` to `(string, string, *IotDpsResourceClientListValidSKUsOptions)`
- Function `*IotDpsResourceClient.ListValidSKUs` return value(s) have been changed from `(*IotDpsResourceListValidSKUsPager)` to `(*IotDpsResourceClientListValidSKUsPager)`
- Function `*IotDpsResourceClient.ListKeys` parameter(s) have been changed from `(string, string, *IotDpsResourceListKeysOptions)` to `(string, string, *IotDpsResourceClientListKeysOptions)`
- Function `*IotDpsResourceClient.ListKeys` return value(s) have been changed from `(*IotDpsResourceListKeysPager)` to `(*IotDpsResourceClientListKeysPager)`
- Function `*IotDpsResourceClient.BeginCreateOrUpdatePrivateEndpointConnection` parameter(s) have been changed from `(context.Context, string, string, string, PrivateEndpointConnection, *IotDpsResourceBeginCreateOrUpdatePrivateEndpointConnectionOptions)` to `(context.Context, string, string, string, PrivateEndpointConnection, *IotDpsResourceClientBeginCreateOrUpdatePrivateEndpointConnectionOptions)`
- Function `*IotDpsResourceClient.BeginCreateOrUpdatePrivateEndpointConnection` return value(s) have been changed from `(IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPollerResponse, error)` to `(IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPollerResponse, error)`
- Function `*IotDpsResourceClient.GetOperationResult` parameter(s) have been changed from `(context.Context, string, string, string, string, *IotDpsResourceGetOperationResultOptions)` to `(context.Context, string, string, string, string, *IotDpsResourceClientGetOperationResultOptions)`
- Function `*IotDpsResourceClient.GetOperationResult` return value(s) have been changed from `(IotDpsResourceGetOperationResultResponse, error)` to `(IotDpsResourceClientGetOperationResultResponse, error)`
- Function `*IotDpsResourceClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, TagsResource, *IotDpsResourceBeginUpdateOptions)` to `(context.Context, string, string, TagsResource, *IotDpsResourceClientBeginUpdateOptions)`
- Function `*IotDpsResourceClient.BeginUpdate` return value(s) have been changed from `(IotDpsResourceUpdatePollerResponse, error)` to `(IotDpsResourceClientUpdatePollerResponse, error)`
- Function `*IotDpsResourceClient.ListBySubscription` parameter(s) have been changed from `(*IotDpsResourceListBySubscriptionOptions)` to `(*IotDpsResourceClientListBySubscriptionOptions)`
- Function `*IotDpsResourceClient.ListBySubscription` return value(s) have been changed from `(*IotDpsResourceListBySubscriptionPager)` to `(*IotDpsResourceClientListBySubscriptionPager)`
- Function `*IotDpsResourceClient.GetPrivateLinkResources` parameter(s) have been changed from `(context.Context, string, string, string, *IotDpsResourceGetPrivateLinkResourcesOptions)` to `(context.Context, string, string, string, *IotDpsResourceClientGetPrivateLinkResourcesOptions)`
- Function `*IotDpsResourceClient.GetPrivateLinkResources` return value(s) have been changed from `(IotDpsResourceGetPrivateLinkResourcesResponse, error)` to `(IotDpsResourceClientGetPrivateLinkResourcesResponse, error)`
- Function `*IotDpsResourceClient.BeginDeletePrivateEndpointConnection` parameter(s) have been changed from `(context.Context, string, string, string, *IotDpsResourceBeginDeletePrivateEndpointConnectionOptions)` to `(context.Context, string, string, string, *IotDpsResourceClientBeginDeletePrivateEndpointConnectionOptions)`
- Function `*IotDpsResourceClient.BeginDeletePrivateEndpointConnection` return value(s) have been changed from `(IotDpsResourceDeletePrivateEndpointConnectionPollerResponse, error)` to `(IotDpsResourceClientDeletePrivateEndpointConnectionPollerResponse, error)`
- Function `*IotDpsResourceClient.ListPrivateEndpointConnections` parameter(s) have been changed from `(context.Context, string, string, *IotDpsResourceListPrivateEndpointConnectionsOptions)` to `(context.Context, string, string, *IotDpsResourceClientListPrivateEndpointConnectionsOptions)`
- Function `*IotDpsResourceClient.ListPrivateEndpointConnections` return value(s) have been changed from `(IotDpsResourceListPrivateEndpointConnectionsResponse, error)` to `(IotDpsResourceClientListPrivateEndpointConnectionsResponse, error)`
- Function `*IotDpsResourceClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *IotDpsResourceBeginDeleteOptions)` to `(context.Context, string, string, *IotDpsResourceClientBeginDeleteOptions)`
- Function `*IotDpsResourceClient.BeginDelete` return value(s) have been changed from `(IotDpsResourceDeletePollerResponse, error)` to `(IotDpsResourceClientDeletePollerResponse, error)`
- Function `*DpsCertificateClient.GenerateVerificationCode` parameter(s) have been changed from `(context.Context, string, string, string, string, *DpsCertificateGenerateVerificationCodeOptions)` to `(context.Context, string, string, string, string, *DpsCertificateClientGenerateVerificationCodeOptions)`
- Function `*DpsCertificateClient.GenerateVerificationCode` return value(s) have been changed from `(DpsCertificateGenerateVerificationCodeResponse, error)` to `(DpsCertificateClientGenerateVerificationCodeResponse, error)`
- Function `*DpsCertificateClient.VerifyCertificate` parameter(s) have been changed from `(context.Context, string, string, string, string, VerificationCodeRequest, *DpsCertificateVerifyCertificateOptions)` to `(context.Context, string, string, string, string, VerificationCodeRequest, *DpsCertificateClientVerifyCertificateOptions)`
- Function `*DpsCertificateClient.VerifyCertificate` return value(s) have been changed from `(DpsCertificateVerifyCertificateResponse, error)` to `(DpsCertificateClientVerifyCertificateResponse, error)`
- Function `*IotDpsResourceClient.ListPrivateLinkResources` parameter(s) have been changed from `(context.Context, string, string, *IotDpsResourceListPrivateLinkResourcesOptions)` to `(context.Context, string, string, *IotDpsResourceClientListPrivateLinkResourcesOptions)`
- Function `*IotDpsResourceClient.ListPrivateLinkResources` return value(s) have been changed from `(IotDpsResourceListPrivateLinkResourcesResponse, error)` to `(IotDpsResourceClientListPrivateLinkResourcesResponse, error)`
- Function `*IotDpsResourceClient.ListKeysForKeyName` parameter(s) have been changed from `(context.Context, string, string, string, *IotDpsResourceListKeysForKeyNameOptions)` to `(context.Context, string, string, string, *IotDpsResourceClientListKeysForKeyNameOptions)`
- Function `*IotDpsResourceClient.ListKeysForKeyName` return value(s) have been changed from `(IotDpsResourceListKeysForKeyNameResponse, error)` to `(IotDpsResourceClientListKeysForKeyNameResponse, error)`
- Function `*IotDpsResourceClient.CheckProvisioningServiceNameAvailability` parameter(s) have been changed from `(context.Context, OperationInputs, *IotDpsResourceCheckProvisioningServiceNameAvailabilityOptions)` to `(context.Context, OperationInputs, *IotDpsResourceClientCheckProvisioningServiceNameAvailabilityOptions)`
- Function `*IotDpsResourceClient.CheckProvisioningServiceNameAvailability` return value(s) have been changed from `(IotDpsResourceCheckProvisioningServiceNameAvailabilityResponse, error)` to `(IotDpsResourceClientCheckProvisioningServiceNameAvailabilityResponse, error)`
- Function `*IotDpsResourceListByResourceGroupPager.Err` has been removed
- Function `*IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller.ResumeToken` has been removed
- Function `*IotDpsResourceUpdatePoller.FinalResponse` has been removed
- Function `*IotDpsResourceDeletePoller.Poll` has been removed
- Function `IotDpsResourceCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `*IotDpsResourceListValidSKUsPager.Err` has been removed
- Function `*IotDpsResourceListValidSKUsPager.PageResponse` has been removed
- Function `*IotDpsResourceListByResourceGroupPager.NextPage` has been removed
- Function `*IotDpsResourceDeletePrivateEndpointConnectionPoller.Done` has been removed
- Function `IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPollerResponse.PollUntilDone` has been removed
- Function `*IotDpsResourceListBySubscriptionPager.Err` has been removed
- Function `ErrorDetails.Error` has been removed
- Function `IotDpsResourceDeletePollerResponse.PollUntilDone` has been removed
- Function `*IotDpsResourceListValidSKUsPager.NextPage` has been removed
- Function `*IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller.FinalResponse` has been removed
- Function `*IotDpsResourceCreateOrUpdatePollerResponse.Resume` has been removed
- Function `*IotDpsResourceListBySubscriptionPager.PageResponse` has been removed
- Function `*IotDpsResourceUpdatePoller.ResumeToken` has been removed
- Function `*IotDpsResourceListBySubscriptionPager.NextPage` has been removed
- Function `*IotDpsResourceDeletePollerResponse.Resume` has been removed
- Function `*IotDpsResourceDeletePrivateEndpointConnectionPoller.FinalResponse` has been removed
- Function `*IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller.Poll` has been removed
- Function `*IotDpsResourceCreateOrUpdatePoller.ResumeToken` has been removed
- Function `*IotDpsResourceUpdatePoller.Poll` has been removed
- Function `*OperationsListPager.PageResponse` has been removed
- Function `*OperationsListPager.Err` has been removed
- Function `*IotDpsResourceDeletePrivateEndpointConnectionPoller.ResumeToken` has been removed
- Function `*IotDpsResourceDeletePrivateEndpointConnectionPoller.Poll` has been removed
- Function `*IotDpsResourceDeletePrivateEndpointConnectionPollerResponse.Resume` has been removed
- Function `*IotDpsResourceCreateOrUpdatePoller.Poll` has been removed
- Function `*IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller.Done` has been removed
- Function `*IotDpsResourceDeletePoller.FinalResponse` has been removed
- Function `*IotDpsResourceDeletePoller.Done` has been removed
- Function `IotDpsResourceDeletePrivateEndpointConnectionPollerResponse.PollUntilDone` has been removed
- Function `*IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPollerResponse.Resume` has been removed
- Function `*IotDpsResourceDeletePoller.ResumeToken` has been removed
- Function `*IotDpsResourceUpdatePollerResponse.Resume` has been removed
- Function `*IotDpsResourceCreateOrUpdatePoller.Done` has been removed
- Function `IotDpsResourceUpdatePollerResponse.PollUntilDone` has been removed
- Function `*IotDpsResourceListKeysPager.NextPage` has been removed
- Function `*OperationsListPager.NextPage` has been removed
- Function `*IotDpsResourceUpdatePoller.Done` has been removed
- Function `*IotDpsResourceListByResourceGroupPager.PageResponse` has been removed
- Function `*IotDpsResourceListKeysPager.Err` has been removed
- Function `*IotDpsResourceCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*IotDpsResourceListKeysPager.PageResponse` has been removed
- Struct `DpsCertificateCreateOrUpdateOptions` has been removed
- Struct `DpsCertificateCreateOrUpdateResponse` has been removed
- Struct `DpsCertificateCreateOrUpdateResult` has been removed
- Struct `DpsCertificateDeleteOptions` has been removed
- Struct `DpsCertificateDeleteResponse` has been removed
- Struct `DpsCertificateGenerateVerificationCodeOptions` has been removed
- Struct `DpsCertificateGenerateVerificationCodeResponse` has been removed
- Struct `DpsCertificateGenerateVerificationCodeResult` has been removed
- Struct `DpsCertificateGetOptions` has been removed
- Struct `DpsCertificateGetResponse` has been removed
- Struct `DpsCertificateGetResult` has been removed
- Struct `DpsCertificateListOptions` has been removed
- Struct `DpsCertificateListResponse` has been removed
- Struct `DpsCertificateListResult` has been removed
- Struct `DpsCertificateVerifyCertificateOptions` has been removed
- Struct `DpsCertificateVerifyCertificateResponse` has been removed
- Struct `DpsCertificateVerifyCertificateResult` has been removed
- Struct `IotDpsResourceBeginCreateOrUpdateOptions` has been removed
- Struct `IotDpsResourceBeginCreateOrUpdatePrivateEndpointConnectionOptions` has been removed
- Struct `IotDpsResourceBeginDeleteOptions` has been removed
- Struct `IotDpsResourceBeginDeletePrivateEndpointConnectionOptions` has been removed
- Struct `IotDpsResourceBeginUpdateOptions` has been removed
- Struct `IotDpsResourceCheckProvisioningServiceNameAvailabilityOptions` has been removed
- Struct `IotDpsResourceCheckProvisioningServiceNameAvailabilityResponse` has been removed
- Struct `IotDpsResourceCheckProvisioningServiceNameAvailabilityResult` has been removed
- Struct `IotDpsResourceCreateOrUpdatePoller` has been removed
- Struct `IotDpsResourceCreateOrUpdatePollerResponse` has been removed
- Struct `IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPoller` has been removed
- Struct `IotDpsResourceCreateOrUpdatePrivateEndpointConnectionPollerResponse` has been removed
- Struct `IotDpsResourceCreateOrUpdatePrivateEndpointConnectionResponse` has been removed
- Struct `IotDpsResourceCreateOrUpdatePrivateEndpointConnectionResult` has been removed
- Struct `IotDpsResourceCreateOrUpdateResponse` has been removed
- Struct `IotDpsResourceCreateOrUpdateResult` has been removed
- Struct `IotDpsResourceDeletePoller` has been removed
- Struct `IotDpsResourceDeletePollerResponse` has been removed
- Struct `IotDpsResourceDeletePrivateEndpointConnectionPoller` has been removed
- Struct `IotDpsResourceDeletePrivateEndpointConnectionPollerResponse` has been removed
- Struct `IotDpsResourceDeletePrivateEndpointConnectionResponse` has been removed
- Struct `IotDpsResourceDeletePrivateEndpointConnectionResult` has been removed
- Struct `IotDpsResourceDeleteResponse` has been removed
- Struct `IotDpsResourceGetOperationResultOptions` has been removed
- Struct `IotDpsResourceGetOperationResultResponse` has been removed
- Struct `IotDpsResourceGetOperationResultResult` has been removed
- Struct `IotDpsResourceGetOptions` has been removed
- Struct `IotDpsResourceGetPrivateEndpointConnectionOptions` has been removed
- Struct `IotDpsResourceGetPrivateEndpointConnectionResponse` has been removed
- Struct `IotDpsResourceGetPrivateEndpointConnectionResult` has been removed
- Struct `IotDpsResourceGetPrivateLinkResourcesOptions` has been removed
- Struct `IotDpsResourceGetPrivateLinkResourcesResponse` has been removed
- Struct `IotDpsResourceGetPrivateLinkResourcesResult` has been removed
- Struct `IotDpsResourceGetResponse` has been removed
- Struct `IotDpsResourceGetResult` has been removed
- Struct `IotDpsResourceListByResourceGroupOptions` has been removed
- Struct `IotDpsResourceListByResourceGroupPager` has been removed
- Struct `IotDpsResourceListByResourceGroupResponse` has been removed
- Struct `IotDpsResourceListByResourceGroupResult` has been removed
- Struct `IotDpsResourceListBySubscriptionOptions` has been removed
- Struct `IotDpsResourceListBySubscriptionPager` has been removed
- Struct `IotDpsResourceListBySubscriptionResponse` has been removed
- Struct `IotDpsResourceListBySubscriptionResult` has been removed
- Struct `IotDpsResourceListKeysForKeyNameOptions` has been removed
- Struct `IotDpsResourceListKeysForKeyNameResponse` has been removed
- Struct `IotDpsResourceListKeysForKeyNameResult` has been removed
- Struct `IotDpsResourceListKeysOptions` has been removed
- Struct `IotDpsResourceListKeysPager` has been removed
- Struct `IotDpsResourceListKeysResponse` has been removed
- Struct `IotDpsResourceListKeysResult` has been removed
- Struct `IotDpsResourceListPrivateEndpointConnectionsOptions` has been removed
- Struct `IotDpsResourceListPrivateEndpointConnectionsResponse` has been removed
- Struct `IotDpsResourceListPrivateEndpointConnectionsResult` has been removed
- Struct `IotDpsResourceListPrivateLinkResourcesOptions` has been removed
- Struct `IotDpsResourceListPrivateLinkResourcesResponse` has been removed
- Struct `IotDpsResourceListPrivateLinkResourcesResult` has been removed
- Struct `IotDpsResourceListValidSKUsOptions` has been removed
- Struct `IotDpsResourceListValidSKUsPager` has been removed
- Struct `IotDpsResourceListValidSKUsResponse` has been removed
- Struct `IotDpsResourceListValidSKUsResult` has been removed
- Struct `IotDpsResourceUpdatePoller` has been removed
- Struct `IotDpsResourceUpdatePollerResponse` has been removed
- Struct `IotDpsResourceUpdateResponse` has been removed
- Struct `IotDpsResourceUpdateResult` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListPager` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Field `Resource` of struct `ProvisioningServiceDescription` has been removed

### Features Added

- New function `*IotDpsResourceClientListKeysPager.NextPage(context.Context) bool`
- New function `*IotDpsResourceClientListValidSKUsPager.PageResponse() IotDpsResourceClientListValidSKUsResponse`
- New function `IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPollerResponse.PollUntilDone(context.Context, time.Duration) (IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionResponse, error)`
- New function `*OperationsClientListPager.Err() error`
- New function `*IotDpsResourceClientListByResourceGroupPager.NextPage(context.Context) bool`
- New function `*IotDpsResourceClientUpdatePoller.Done() bool`
- New function `*IotDpsResourceClientUpdatePoller.FinalResponse(context.Context) (IotDpsResourceClientUpdateResponse, error)`
- New function `*IotDpsResourceClientListValidSKUsPager.NextPage(context.Context) bool`
- New function `*IotDpsResourceClientListBySubscriptionPager.Err() error`
- New function `*OperationsClientListPager.NextPage(context.Context) bool`
- New function `*IotDpsResourceClientListValidSKUsPager.Err() error`
- New function `*IotDpsResourceClientListKeysPager.Err() error`
- New function `*IotDpsResourceClientCreateOrUpdatePoller.FinalResponse(context.Context) (IotDpsResourceClientCreateOrUpdateResponse, error)`
- New function `*IotDpsResourceClientDeletePrivateEndpointConnectionPoller.Done() bool`
- New function `*IotDpsResourceClientUpdatePollerResponse.Resume(context.Context, *IotDpsResourceClient, string) error`
- New function `IotDpsResourceClientDeletePrivateEndpointConnectionPollerResponse.PollUntilDone(context.Context, time.Duration) (IotDpsResourceClientDeletePrivateEndpointConnectionResponse, error)`
- New function `*IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPoller.FinalResponse(context.Context) (IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionResponse, error)`
- New function `*IotDpsResourceClientDeletePrivateEndpointConnectionPoller.Poll(context.Context) (*http.Response, error)`
- New function `*IotDpsResourceClientDeletePrivateEndpointConnectionPoller.FinalResponse(context.Context) (IotDpsResourceClientDeletePrivateEndpointConnectionResponse, error)`
- New function `*IotDpsResourceClientCreateOrUpdatePoller.Done() bool`
- New function `*IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPollerResponse.Resume(context.Context, *IotDpsResourceClient, string) error`
- New function `*IotDpsResourceClientDeletePoller.ResumeToken() (string, error)`
- New function `*IotDpsResourceClientUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*IotDpsResourceClientListKeysPager.PageResponse() IotDpsResourceClientListKeysResponse`
- New function `IotDpsResourceClientCreateOrUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (IotDpsResourceClientCreateOrUpdateResponse, error)`
- New function `*IotDpsResourceClientDeletePollerResponse.Resume(context.Context, *IotDpsResourceClient, string) error`
- New function `*IotDpsResourceClientListBySubscriptionPager.NextPage(context.Context) bool`
- New function `*IotDpsResourceClientCreateOrUpdatePoller.Poll(context.Context) (*http.Response, error)`
- New function `*IotDpsResourceClientDeletePrivateEndpointConnectionPollerResponse.Resume(context.Context, *IotDpsResourceClient, string) error`
- New function `*IotDpsResourceClientDeletePrivateEndpointConnectionPoller.ResumeToken() (string, error)`
- New function `*IotDpsResourceClientDeletePoller.Poll(context.Context) (*http.Response, error)`
- New function `*IotDpsResourceClientListByResourceGroupPager.PageResponse() IotDpsResourceClientListByResourceGroupResponse`
- New function `*IotDpsResourceClientCreateOrUpdatePoller.ResumeToken() (string, error)`
- New function `*IotDpsResourceClientDeletePoller.FinalResponse(context.Context) (IotDpsResourceClientDeleteResponse, error)`
- New function `*IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPoller.Poll(context.Context) (*http.Response, error)`
- New function `IotDpsResourceClientDeletePollerResponse.PollUntilDone(context.Context, time.Duration) (IotDpsResourceClientDeleteResponse, error)`
- New function `*IotDpsResourceClientListByResourceGroupPager.Err() error`
- New function `*IotDpsResourceClientCreateOrUpdatePollerResponse.Resume(context.Context, *IotDpsResourceClient, string) error`
- New function `*IotDpsResourceClientListBySubscriptionPager.PageResponse() IotDpsResourceClientListBySubscriptionResponse`
- New function `*IotDpsResourceClientDeletePoller.Done() bool`
- New function `IotDpsResourceClientUpdatePollerResponse.PollUntilDone(context.Context, time.Duration) (IotDpsResourceClientUpdateResponse, error)`
- New function `*IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPoller.ResumeToken() (string, error)`
- New function `*OperationsClientListPager.PageResponse() OperationsClientListResponse`
- New function `*IotDpsResourceClientUpdatePoller.ResumeToken() (string, error)`
- New function `*IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPoller.Done() bool`
- New struct `DpsCertificateClientCreateOrUpdateOptions`
- New struct `DpsCertificateClientCreateOrUpdateResponse`
- New struct `DpsCertificateClientCreateOrUpdateResult`
- New struct `DpsCertificateClientDeleteOptions`
- New struct `DpsCertificateClientDeleteResponse`
- New struct `DpsCertificateClientGenerateVerificationCodeOptions`
- New struct `DpsCertificateClientGenerateVerificationCodeResponse`
- New struct `DpsCertificateClientGenerateVerificationCodeResult`
- New struct `DpsCertificateClientGetOptions`
- New struct `DpsCertificateClientGetResponse`
- New struct `DpsCertificateClientGetResult`
- New struct `DpsCertificateClientListOptions`
- New struct `DpsCertificateClientListResponse`
- New struct `DpsCertificateClientListResult`
- New struct `DpsCertificateClientVerifyCertificateOptions`
- New struct `DpsCertificateClientVerifyCertificateResponse`
- New struct `DpsCertificateClientVerifyCertificateResult`
- New struct `IotDpsResourceClientBeginCreateOrUpdateOptions`
- New struct `IotDpsResourceClientBeginCreateOrUpdatePrivateEndpointConnectionOptions`
- New struct `IotDpsResourceClientBeginDeleteOptions`
- New struct `IotDpsResourceClientBeginDeletePrivateEndpointConnectionOptions`
- New struct `IotDpsResourceClientBeginUpdateOptions`
- New struct `IotDpsResourceClientCheckProvisioningServiceNameAvailabilityOptions`
- New struct `IotDpsResourceClientCheckProvisioningServiceNameAvailabilityResponse`
- New struct `IotDpsResourceClientCheckProvisioningServiceNameAvailabilityResult`
- New struct `IotDpsResourceClientCreateOrUpdatePoller`
- New struct `IotDpsResourceClientCreateOrUpdatePollerResponse`
- New struct `IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPoller`
- New struct `IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionPollerResponse`
- New struct `IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionResponse`
- New struct `IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionResult`
- New struct `IotDpsResourceClientCreateOrUpdateResponse`
- New struct `IotDpsResourceClientCreateOrUpdateResult`
- New struct `IotDpsResourceClientDeletePoller`
- New struct `IotDpsResourceClientDeletePollerResponse`
- New struct `IotDpsResourceClientDeletePrivateEndpointConnectionPoller`
- New struct `IotDpsResourceClientDeletePrivateEndpointConnectionPollerResponse`
- New struct `IotDpsResourceClientDeletePrivateEndpointConnectionResponse`
- New struct `IotDpsResourceClientDeletePrivateEndpointConnectionResult`
- New struct `IotDpsResourceClientDeleteResponse`
- New struct `IotDpsResourceClientGetOperationResultOptions`
- New struct `IotDpsResourceClientGetOperationResultResponse`
- New struct `IotDpsResourceClientGetOperationResultResult`
- New struct `IotDpsResourceClientGetOptions`
- New struct `IotDpsResourceClientGetPrivateEndpointConnectionOptions`
- New struct `IotDpsResourceClientGetPrivateEndpointConnectionResponse`
- New struct `IotDpsResourceClientGetPrivateEndpointConnectionResult`
- New struct `IotDpsResourceClientGetPrivateLinkResourcesOptions`
- New struct `IotDpsResourceClientGetPrivateLinkResourcesResponse`
- New struct `IotDpsResourceClientGetPrivateLinkResourcesResult`
- New struct `IotDpsResourceClientGetResponse`
- New struct `IotDpsResourceClientGetResult`
- New struct `IotDpsResourceClientListByResourceGroupOptions`
- New struct `IotDpsResourceClientListByResourceGroupPager`
- New struct `IotDpsResourceClientListByResourceGroupResponse`
- New struct `IotDpsResourceClientListByResourceGroupResult`
- New struct `IotDpsResourceClientListBySubscriptionOptions`
- New struct `IotDpsResourceClientListBySubscriptionPager`
- New struct `IotDpsResourceClientListBySubscriptionResponse`
- New struct `IotDpsResourceClientListBySubscriptionResult`
- New struct `IotDpsResourceClientListKeysForKeyNameOptions`
- New struct `IotDpsResourceClientListKeysForKeyNameResponse`
- New struct `IotDpsResourceClientListKeysForKeyNameResult`
- New struct `IotDpsResourceClientListKeysOptions`
- New struct `IotDpsResourceClientListKeysPager`
- New struct `IotDpsResourceClientListKeysResponse`
- New struct `IotDpsResourceClientListKeysResult`
- New struct `IotDpsResourceClientListPrivateEndpointConnectionsOptions`
- New struct `IotDpsResourceClientListPrivateEndpointConnectionsResponse`
- New struct `IotDpsResourceClientListPrivateEndpointConnectionsResult`
- New struct `IotDpsResourceClientListPrivateLinkResourcesOptions`
- New struct `IotDpsResourceClientListPrivateLinkResourcesResponse`
- New struct `IotDpsResourceClientListPrivateLinkResourcesResult`
- New struct `IotDpsResourceClientListValidSKUsOptions`
- New struct `IotDpsResourceClientListValidSKUsPager`
- New struct `IotDpsResourceClientListValidSKUsResponse`
- New struct `IotDpsResourceClientListValidSKUsResult`
- New struct `IotDpsResourceClientUpdatePoller`
- New struct `IotDpsResourceClientUpdatePollerResponse`
- New struct `IotDpsResourceClientUpdateResponse`
- New struct `IotDpsResourceClientUpdateResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListPager`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New field `Type` in struct `ProvisioningServiceDescription`
- New field `Location` in struct `ProvisioningServiceDescription`
- New field `Tags` in struct `ProvisioningServiceDescription`
- New field `ID` in struct `ProvisioningServiceDescription`
- New field `Name` in struct `ProvisioningServiceDescription`


## 0.1.0 (2021-12-07)

- Initial preview release.
