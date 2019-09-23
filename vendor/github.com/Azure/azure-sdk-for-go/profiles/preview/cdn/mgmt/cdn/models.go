// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package cdn

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2019-04-15/cdn"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CacheBehavior = original.CacheBehavior

const (
	BypassCache  CacheBehavior = original.BypassCache
	Override     CacheBehavior = original.Override
	SetIfMissing CacheBehavior = original.SetIfMissing
)

type CertificateSource = original.CertificateSource

const (
	CertificateSourceAzureKeyVault               CertificateSource = original.CertificateSourceAzureKeyVault
	CertificateSourceCdn                         CertificateSource = original.CertificateSourceCdn
	CertificateSourceCustomDomainHTTPSParameters CertificateSource = original.CertificateSourceCustomDomainHTTPSParameters
)

type CertificateType = original.CertificateType

const (
	Dedicated CertificateType = original.Dedicated
	Shared    CertificateType = original.Shared
)

type CustomDomainResourceState = original.CustomDomainResourceState

const (
	Active   CustomDomainResourceState = original.Active
	Creating CustomDomainResourceState = original.Creating
	Deleting CustomDomainResourceState = original.Deleting
)

type CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningState

const (
	Disabled  CustomHTTPSProvisioningState = original.Disabled
	Disabling CustomHTTPSProvisioningState = original.Disabling
	Enabled   CustomHTTPSProvisioningState = original.Enabled
	Enabling  CustomHTTPSProvisioningState = original.Enabling
	Failed    CustomHTTPSProvisioningState = original.Failed
)

type CustomHTTPSProvisioningSubstate = original.CustomHTTPSProvisioningSubstate

const (
	CertificateDeleted                            CustomHTTPSProvisioningSubstate = original.CertificateDeleted
	CertificateDeployed                           CustomHTTPSProvisioningSubstate = original.CertificateDeployed
	DeletingCertificate                           CustomHTTPSProvisioningSubstate = original.DeletingCertificate
	DeployingCertificate                          CustomHTTPSProvisioningSubstate = original.DeployingCertificate
	DomainControlValidationRequestApproved        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestApproved
	DomainControlValidationRequestRejected        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestRejected
	DomainControlValidationRequestTimedOut        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestTimedOut
	IssuingCertificate                            CustomHTTPSProvisioningSubstate = original.IssuingCertificate
	PendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = original.PendingDomainControlValidationREquestApproval
	SubmittingDomainControlValidationRequest      CustomHTTPSProvisioningSubstate = original.SubmittingDomainControlValidationRequest
)

type DestinationProtocol = original.DestinationProtocol

const (
	HTTP         DestinationProtocol = original.HTTP
	HTTPS        DestinationProtocol = original.HTTPS
	MatchRequest DestinationProtocol = original.MatchRequest
)

type EndpointResourceState = original.EndpointResourceState

const (
	EndpointResourceStateCreating EndpointResourceState = original.EndpointResourceStateCreating
	EndpointResourceStateDeleting EndpointResourceState = original.EndpointResourceStateDeleting
	EndpointResourceStateRunning  EndpointResourceState = original.EndpointResourceStateRunning
	EndpointResourceStateStarting EndpointResourceState = original.EndpointResourceStateStarting
	EndpointResourceStateStopped  EndpointResourceState = original.EndpointResourceStateStopped
	EndpointResourceStateStopping EndpointResourceState = original.EndpointResourceStateStopping
)

type GeoFilterActions = original.GeoFilterActions

const (
	Allow GeoFilterActions = original.Allow
	Block GeoFilterActions = original.Block
)

type HeaderAction = original.HeaderAction

const (
	Append    HeaderAction = original.Append
	Delete    HeaderAction = original.Delete
	Overwrite HeaderAction = original.Overwrite
)

type Name = original.Name

const (
	NameDeliveryRuleCondition Name = original.NameDeliveryRuleCondition
	NameIsDevice              Name = original.NameIsDevice
	NamePostArgs              Name = original.NamePostArgs
	NameQueryString           Name = original.NameQueryString
	NameRemoteAddress         Name = original.NameRemoteAddress
	NameRequestBody           Name = original.NameRequestBody
	NameRequestHeader         Name = original.NameRequestHeader
	NameRequestMethod         Name = original.NameRequestMethod
	NameRequestScheme         Name = original.NameRequestScheme
	NameRequestURI            Name = original.NameRequestURI
	NameURLFileExtension      Name = original.NameURLFileExtension
	NameURLFileName           Name = original.NameURLFileName
	NameURLPath               Name = original.NameURLPath
)

type NameBasicDeliveryRuleAction = original.NameBasicDeliveryRuleAction

const (
	NameCacheExpiration      NameBasicDeliveryRuleAction = original.NameCacheExpiration
	NameCacheKeyQueryString  NameBasicDeliveryRuleAction = original.NameCacheKeyQueryString
	NameDeliveryRuleAction   NameBasicDeliveryRuleAction = original.NameDeliveryRuleAction
	NameModifyRequestHeader  NameBasicDeliveryRuleAction = original.NameModifyRequestHeader
	NameModifyResponseHeader NameBasicDeliveryRuleAction = original.NameModifyResponseHeader
	NameURLRedirect          NameBasicDeliveryRuleAction = original.NameURLRedirect
	NameURLRewrite           NameBasicDeliveryRuleAction = original.NameURLRewrite
)

type OptimizationType = original.OptimizationType

const (
	DynamicSiteAcceleration     OptimizationType = original.DynamicSiteAcceleration
	GeneralMediaStreaming       OptimizationType = original.GeneralMediaStreaming
	GeneralWebDelivery          OptimizationType = original.GeneralWebDelivery
	LargeFileDownload           OptimizationType = original.LargeFileDownload
	VideoOnDemandMediaStreaming OptimizationType = original.VideoOnDemandMediaStreaming
)

type OriginResourceState = original.OriginResourceState

const (
	OriginResourceStateActive   OriginResourceState = original.OriginResourceStateActive
	OriginResourceStateCreating OriginResourceState = original.OriginResourceStateCreating
	OriginResourceStateDeleting OriginResourceState = original.OriginResourceStateDeleting
)

type PostArgsOperator = original.PostArgsOperator

const (
	Any                PostArgsOperator = original.Any
	BeginsWith         PostArgsOperator = original.BeginsWith
	Contains           PostArgsOperator = original.Contains
	EndsWith           PostArgsOperator = original.EndsWith
	Equal              PostArgsOperator = original.Equal
	GreaterThan        PostArgsOperator = original.GreaterThan
	GreaterThanOrEqual PostArgsOperator = original.GreaterThanOrEqual
	LessThan           PostArgsOperator = original.LessThan
	LessThanOrEqual    PostArgsOperator = original.LessThanOrEqual
)

type ProfileResourceState = original.ProfileResourceState

const (
	ProfileResourceStateActive   ProfileResourceState = original.ProfileResourceStateActive
	ProfileResourceStateCreating ProfileResourceState = original.ProfileResourceStateCreating
	ProfileResourceStateDeleting ProfileResourceState = original.ProfileResourceStateDeleting
	ProfileResourceStateDisabled ProfileResourceState = original.ProfileResourceStateDisabled
)

type ProtocolType = original.ProtocolType

const (
	IPBased              ProtocolType = original.IPBased
	ServerNameIndication ProtocolType = original.ServerNameIndication
)

type QueryStringBehavior = original.QueryStringBehavior

const (
	Exclude    QueryStringBehavior = original.Exclude
	ExcludeAll QueryStringBehavior = original.ExcludeAll
	Include    QueryStringBehavior = original.Include
	IncludeAll QueryStringBehavior = original.IncludeAll
)

type QueryStringCachingBehavior = original.QueryStringCachingBehavior

const (
	BypassCaching     QueryStringCachingBehavior = original.BypassCaching
	IgnoreQueryString QueryStringCachingBehavior = original.IgnoreQueryString
	NotSet            QueryStringCachingBehavior = original.NotSet
	UseQueryString    QueryStringCachingBehavior = original.UseQueryString
)

type QueryStringOperator = original.QueryStringOperator

const (
	QueryStringOperatorAny                QueryStringOperator = original.QueryStringOperatorAny
	QueryStringOperatorBeginsWith         QueryStringOperator = original.QueryStringOperatorBeginsWith
	QueryStringOperatorContains           QueryStringOperator = original.QueryStringOperatorContains
	QueryStringOperatorEndsWith           QueryStringOperator = original.QueryStringOperatorEndsWith
	QueryStringOperatorEqual              QueryStringOperator = original.QueryStringOperatorEqual
	QueryStringOperatorGreaterThan        QueryStringOperator = original.QueryStringOperatorGreaterThan
	QueryStringOperatorGreaterThanOrEqual QueryStringOperator = original.QueryStringOperatorGreaterThanOrEqual
	QueryStringOperatorLessThan           QueryStringOperator = original.QueryStringOperatorLessThan
	QueryStringOperatorLessThanOrEqual    QueryStringOperator = original.QueryStringOperatorLessThanOrEqual
)

type RedirectType = original.RedirectType

const (
	Found             RedirectType = original.Found
	Moved             RedirectType = original.Moved
	PermanentRedirect RedirectType = original.PermanentRedirect
	TemporaryRedirect RedirectType = original.TemporaryRedirect
)

type RemoteAddressOperator = original.RemoteAddressOperator

const (
	RemoteAddressOperatorAny      RemoteAddressOperator = original.RemoteAddressOperatorAny
	RemoteAddressOperatorGeoMatch RemoteAddressOperator = original.RemoteAddressOperatorGeoMatch
	RemoteAddressOperatorIPMatch  RemoteAddressOperator = original.RemoteAddressOperatorIPMatch
)

type RequestBodyOperator = original.RequestBodyOperator

const (
	RequestBodyOperatorAny                RequestBodyOperator = original.RequestBodyOperatorAny
	RequestBodyOperatorBeginsWith         RequestBodyOperator = original.RequestBodyOperatorBeginsWith
	RequestBodyOperatorContains           RequestBodyOperator = original.RequestBodyOperatorContains
	RequestBodyOperatorEndsWith           RequestBodyOperator = original.RequestBodyOperatorEndsWith
	RequestBodyOperatorEqual              RequestBodyOperator = original.RequestBodyOperatorEqual
	RequestBodyOperatorGreaterThan        RequestBodyOperator = original.RequestBodyOperatorGreaterThan
	RequestBodyOperatorGreaterThanOrEqual RequestBodyOperator = original.RequestBodyOperatorGreaterThanOrEqual
	RequestBodyOperatorLessThan           RequestBodyOperator = original.RequestBodyOperatorLessThan
	RequestBodyOperatorLessThanOrEqual    RequestBodyOperator = original.RequestBodyOperatorLessThanOrEqual
)

type RequestHeaderOperator = original.RequestHeaderOperator

const (
	RequestHeaderOperatorAny                RequestHeaderOperator = original.RequestHeaderOperatorAny
	RequestHeaderOperatorBeginsWith         RequestHeaderOperator = original.RequestHeaderOperatorBeginsWith
	RequestHeaderOperatorContains           RequestHeaderOperator = original.RequestHeaderOperatorContains
	RequestHeaderOperatorEndsWith           RequestHeaderOperator = original.RequestHeaderOperatorEndsWith
	RequestHeaderOperatorEqual              RequestHeaderOperator = original.RequestHeaderOperatorEqual
	RequestHeaderOperatorGreaterThan        RequestHeaderOperator = original.RequestHeaderOperatorGreaterThan
	RequestHeaderOperatorGreaterThanOrEqual RequestHeaderOperator = original.RequestHeaderOperatorGreaterThanOrEqual
	RequestHeaderOperatorLessThan           RequestHeaderOperator = original.RequestHeaderOperatorLessThan
	RequestHeaderOperatorLessThanOrEqual    RequestHeaderOperator = original.RequestHeaderOperatorLessThanOrEqual
)

type RequestURIOperator = original.RequestURIOperator

const (
	RequestURIOperatorAny                RequestURIOperator = original.RequestURIOperatorAny
	RequestURIOperatorBeginsWith         RequestURIOperator = original.RequestURIOperatorBeginsWith
	RequestURIOperatorContains           RequestURIOperator = original.RequestURIOperatorContains
	RequestURIOperatorEndsWith           RequestURIOperator = original.RequestURIOperatorEndsWith
	RequestURIOperatorEqual              RequestURIOperator = original.RequestURIOperatorEqual
	RequestURIOperatorGreaterThan        RequestURIOperator = original.RequestURIOperatorGreaterThan
	RequestURIOperatorGreaterThanOrEqual RequestURIOperator = original.RequestURIOperatorGreaterThanOrEqual
	RequestURIOperatorLessThan           RequestURIOperator = original.RequestURIOperatorLessThan
	RequestURIOperatorLessThanOrEqual    RequestURIOperator = original.RequestURIOperatorLessThanOrEqual
)

type ResourceType = original.ResourceType

const (
	MicrosoftCdnProfilesEndpoints ResourceType = original.MicrosoftCdnProfilesEndpoints
)

type SkuName = original.SkuName

const (
	CustomVerizon     SkuName = original.CustomVerizon
	PremiumChinaCdn   SkuName = original.PremiumChinaCdn
	PremiumVerizon    SkuName = original.PremiumVerizon
	StandardAkamai    SkuName = original.StandardAkamai
	StandardChinaCdn  SkuName = original.StandardChinaCdn
	StandardMicrosoft SkuName = original.StandardMicrosoft
	StandardVerizon   SkuName = original.StandardVerizon
)

type Transform = original.Transform

const (
	Lowercase Transform = original.Lowercase
	Uppercase Transform = original.Uppercase
)

type URLFileExtensionOperator = original.URLFileExtensionOperator

const (
	URLFileExtensionOperatorAny                URLFileExtensionOperator = original.URLFileExtensionOperatorAny
	URLFileExtensionOperatorBeginsWith         URLFileExtensionOperator = original.URLFileExtensionOperatorBeginsWith
	URLFileExtensionOperatorContains           URLFileExtensionOperator = original.URLFileExtensionOperatorContains
	URLFileExtensionOperatorEndsWith           URLFileExtensionOperator = original.URLFileExtensionOperatorEndsWith
	URLFileExtensionOperatorEqual              URLFileExtensionOperator = original.URLFileExtensionOperatorEqual
	URLFileExtensionOperatorGreaterThan        URLFileExtensionOperator = original.URLFileExtensionOperatorGreaterThan
	URLFileExtensionOperatorGreaterThanOrEqual URLFileExtensionOperator = original.URLFileExtensionOperatorGreaterThanOrEqual
	URLFileExtensionOperatorLessThan           URLFileExtensionOperator = original.URLFileExtensionOperatorLessThan
	URLFileExtensionOperatorLessThanOrEqual    URLFileExtensionOperator = original.URLFileExtensionOperatorLessThanOrEqual
)

type URLFileNameOperator = original.URLFileNameOperator

const (
	URLFileNameOperatorAny                URLFileNameOperator = original.URLFileNameOperatorAny
	URLFileNameOperatorBeginsWith         URLFileNameOperator = original.URLFileNameOperatorBeginsWith
	URLFileNameOperatorContains           URLFileNameOperator = original.URLFileNameOperatorContains
	URLFileNameOperatorEndsWith           URLFileNameOperator = original.URLFileNameOperatorEndsWith
	URLFileNameOperatorEqual              URLFileNameOperator = original.URLFileNameOperatorEqual
	URLFileNameOperatorGreaterThan        URLFileNameOperator = original.URLFileNameOperatorGreaterThan
	URLFileNameOperatorGreaterThanOrEqual URLFileNameOperator = original.URLFileNameOperatorGreaterThanOrEqual
	URLFileNameOperatorLessThan           URLFileNameOperator = original.URLFileNameOperatorLessThan
	URLFileNameOperatorLessThanOrEqual    URLFileNameOperator = original.URLFileNameOperatorLessThanOrEqual
)

type URLPathOperator = original.URLPathOperator

const (
	URLPathOperatorAny                URLPathOperator = original.URLPathOperatorAny
	URLPathOperatorBeginsWith         URLPathOperator = original.URLPathOperatorBeginsWith
	URLPathOperatorContains           URLPathOperator = original.URLPathOperatorContains
	URLPathOperatorEndsWith           URLPathOperator = original.URLPathOperatorEndsWith
	URLPathOperatorEqual              URLPathOperator = original.URLPathOperatorEqual
	URLPathOperatorGreaterThan        URLPathOperator = original.URLPathOperatorGreaterThan
	URLPathOperatorGreaterThanOrEqual URLPathOperator = original.URLPathOperatorGreaterThanOrEqual
	URLPathOperatorLessThan           URLPathOperator = original.URLPathOperatorLessThan
	URLPathOperatorLessThanOrEqual    URLPathOperator = original.URLPathOperatorLessThanOrEqual
	URLPathOperatorWildcard           URLPathOperator = original.URLPathOperatorWildcard
)

type BaseClient = original.BaseClient
type BasicCustomDomainHTTPSParameters = original.BasicCustomDomainHTTPSParameters
type BasicDeliveryRuleAction = original.BasicDeliveryRuleAction
type BasicDeliveryRuleCondition = original.BasicDeliveryRuleCondition
type CacheExpirationActionParameters = original.CacheExpirationActionParameters
type CacheKeyQueryStringActionParameters = original.CacheKeyQueryStringActionParameters
type CertificateSourceParameters = original.CertificateSourceParameters
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CidrIPAddress = original.CidrIPAddress
type CustomDomain = original.CustomDomain
type CustomDomainHTTPSParameters = original.CustomDomainHTTPSParameters
type CustomDomainListResult = original.CustomDomainListResult
type CustomDomainListResultIterator = original.CustomDomainListResultIterator
type CustomDomainListResultPage = original.CustomDomainListResultPage
type CustomDomainParameters = original.CustomDomainParameters
type CustomDomainProperties = original.CustomDomainProperties
type CustomDomainPropertiesParameters = original.CustomDomainPropertiesParameters
type CustomDomainsClient = original.CustomDomainsClient
type CustomDomainsCreateFuture = original.CustomDomainsCreateFuture
type CustomDomainsDeleteFuture = original.CustomDomainsDeleteFuture
type DeepCreatedOrigin = original.DeepCreatedOrigin
type DeepCreatedOriginProperties = original.DeepCreatedOriginProperties
type DeliveryRule = original.DeliveryRule
type DeliveryRuleAction = original.DeliveryRuleAction
type DeliveryRuleCacheExpirationAction = original.DeliveryRuleCacheExpirationAction
type DeliveryRuleCacheKeyQueryStringAction = original.DeliveryRuleCacheKeyQueryStringAction
type DeliveryRuleCondition = original.DeliveryRuleCondition
type DeliveryRuleIsDeviceCondition = original.DeliveryRuleIsDeviceCondition
type DeliveryRulePostArgsCondition = original.DeliveryRulePostArgsCondition
type DeliveryRuleQueryStringCondition = original.DeliveryRuleQueryStringCondition
type DeliveryRuleRemoteAddressCondition = original.DeliveryRuleRemoteAddressCondition
type DeliveryRuleRequestBodyCondition = original.DeliveryRuleRequestBodyCondition
type DeliveryRuleRequestHeaderAction = original.DeliveryRuleRequestHeaderAction
type DeliveryRuleRequestHeaderCondition = original.DeliveryRuleRequestHeaderCondition
type DeliveryRuleRequestMethodCondition = original.DeliveryRuleRequestMethodCondition
type DeliveryRuleRequestSchemeCondition = original.DeliveryRuleRequestSchemeCondition
type DeliveryRuleRequestURICondition = original.DeliveryRuleRequestURICondition
type DeliveryRuleResponseHeaderAction = original.DeliveryRuleResponseHeaderAction
type DeliveryRuleURLFileExtensionCondition = original.DeliveryRuleURLFileExtensionCondition
type DeliveryRuleURLFileNameCondition = original.DeliveryRuleURLFileNameCondition
type DeliveryRuleURLPathCondition = original.DeliveryRuleURLPathCondition
type EdgeNode = original.EdgeNode
type EdgeNodeProperties = original.EdgeNodeProperties
type EdgeNodesClient = original.EdgeNodesClient
type EdgenodeResult = original.EdgenodeResult
type EdgenodeResultIterator = original.EdgenodeResultIterator
type EdgenodeResultPage = original.EdgenodeResultPage
type Endpoint = original.Endpoint
type EndpointListResult = original.EndpointListResult
type EndpointListResultIterator = original.EndpointListResultIterator
type EndpointListResultPage = original.EndpointListResultPage
type EndpointProperties = original.EndpointProperties
type EndpointPropertiesUpdateParameters = original.EndpointPropertiesUpdateParameters
type EndpointPropertiesUpdateParametersDeliveryPolicy = original.EndpointPropertiesUpdateParametersDeliveryPolicy
type EndpointUpdateParameters = original.EndpointUpdateParameters
type EndpointsClient = original.EndpointsClient
type EndpointsCreateFuture = original.EndpointsCreateFuture
type EndpointsDeleteFuture = original.EndpointsDeleteFuture
type EndpointsLoadContentFuture = original.EndpointsLoadContentFuture
type EndpointsPurgeContentFuture = original.EndpointsPurgeContentFuture
type EndpointsStartFuture = original.EndpointsStartFuture
type EndpointsStopFuture = original.EndpointsStopFuture
type EndpointsUpdateFuture = original.EndpointsUpdateFuture
type ErrorResponse = original.ErrorResponse
type GeoFilter = original.GeoFilter
type HeaderActionParameters = original.HeaderActionParameters
type IPAddressGroup = original.IPAddressGroup
type IsDeviceMatchConditionParameters = original.IsDeviceMatchConditionParameters
type KeyVaultCertificateSourceParameters = original.KeyVaultCertificateSourceParameters
type LoadParameters = original.LoadParameters
type ManagedHTTPSParameters = original.ManagedHTTPSParameters
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsListResult = original.OperationsListResult
type OperationsListResultIterator = original.OperationsListResultIterator
type OperationsListResultPage = original.OperationsListResultPage
type Origin = original.Origin
type OriginListResult = original.OriginListResult
type OriginListResultIterator = original.OriginListResultIterator
type OriginListResultPage = original.OriginListResultPage
type OriginProperties = original.OriginProperties
type OriginPropertiesParameters = original.OriginPropertiesParameters
type OriginUpdateParameters = original.OriginUpdateParameters
type OriginsClient = original.OriginsClient
type OriginsUpdateFuture = original.OriginsUpdateFuture
type PostArgsMatchConditionParameters = original.PostArgsMatchConditionParameters
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileListResultIterator = original.ProfileListResultIterator
type ProfileListResultPage = original.ProfileListResultPage
type ProfileProperties = original.ProfileProperties
type ProfileUpdateParameters = original.ProfileUpdateParameters
type ProfilesClient = original.ProfilesClient
type ProfilesCreateFuture = original.ProfilesCreateFuture
type ProfilesDeleteFuture = original.ProfilesDeleteFuture
type ProfilesUpdateFuture = original.ProfilesUpdateFuture
type ProxyResource = original.ProxyResource
type PurgeParameters = original.PurgeParameters
type QueryStringMatchConditionParameters = original.QueryStringMatchConditionParameters
type RemoteAddressMatchConditionParameters = original.RemoteAddressMatchConditionParameters
type RequestBodyMatchConditionParameters = original.RequestBodyMatchConditionParameters
type RequestHeaderMatchConditionParameters = original.RequestHeaderMatchConditionParameters
type RequestMethodMatchConditionParameters = original.RequestMethodMatchConditionParameters
type RequestSchemeMatchConditionParameters = original.RequestSchemeMatchConditionParameters
type RequestURIMatchConditionParameters = original.RequestURIMatchConditionParameters
type Resource = original.Resource
type ResourceUsage = original.ResourceUsage
type ResourceUsageClient = original.ResourceUsageClient
type ResourceUsageListResult = original.ResourceUsageListResult
type ResourceUsageListResultIterator = original.ResourceUsageListResultIterator
type ResourceUsageListResultPage = original.ResourceUsageListResultPage
type Sku = original.Sku
type SsoURI = original.SsoURI
type SupportedOptimizationTypesListResult = original.SupportedOptimizationTypesListResult
type TrackedResource = original.TrackedResource
type URLFileExtensionMatchConditionParameters = original.URLFileExtensionMatchConditionParameters
type URLFileNameMatchConditionParameters = original.URLFileNameMatchConditionParameters
type URLPathMatchConditionParameters = original.URLPathMatchConditionParameters
type URLRedirectAction = original.URLRedirectAction
type URLRedirectActionParameters = original.URLRedirectActionParameters
type URLRewriteAction = original.URLRewriteAction
type URLRewriteActionParameters = original.URLRewriteActionParameters
type UserManagedHTTPSParameters = original.UserManagedHTTPSParameters
type ValidateCustomDomainInput = original.ValidateCustomDomainInput
type ValidateCustomDomainOutput = original.ValidateCustomDomainOutput
type ValidateProbeInput = original.ValidateProbeInput
type ValidateProbeOutput = original.ValidateProbeOutput

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCustomDomainListResultIterator(page CustomDomainListResultPage) CustomDomainListResultIterator {
	return original.NewCustomDomainListResultIterator(page)
}
func NewCustomDomainListResultPage(getNextPage func(context.Context, CustomDomainListResult) (CustomDomainListResult, error)) CustomDomainListResultPage {
	return original.NewCustomDomainListResultPage(getNextPage)
}
func NewCustomDomainsClient(subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClient(subscriptionID)
}
func NewCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string) CustomDomainsClient {
	return original.NewCustomDomainsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEdgeNodesClient(subscriptionID string) EdgeNodesClient {
	return original.NewEdgeNodesClient(subscriptionID)
}
func NewEdgeNodesClientWithBaseURI(baseURI string, subscriptionID string) EdgeNodesClient {
	return original.NewEdgeNodesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEdgenodeResultIterator(page EdgenodeResultPage) EdgenodeResultIterator {
	return original.NewEdgenodeResultIterator(page)
}
func NewEdgenodeResultPage(getNextPage func(context.Context, EdgenodeResult) (EdgenodeResult, error)) EdgenodeResultPage {
	return original.NewEdgenodeResultPage(getNextPage)
}
func NewEndpointListResultIterator(page EndpointListResultPage) EndpointListResultIterator {
	return original.NewEndpointListResultIterator(page)
}
func NewEndpointListResultPage(getNextPage func(context.Context, EndpointListResult) (EndpointListResult, error)) EndpointListResultPage {
	return original.NewEndpointListResultPage(getNextPage)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsListResultIterator(page OperationsListResultPage) OperationsListResultIterator {
	return original.NewOperationsListResultIterator(page)
}
func NewOperationsListResultPage(getNextPage func(context.Context, OperationsListResult) (OperationsListResult, error)) OperationsListResultPage {
	return original.NewOperationsListResultPage(getNextPage)
}
func NewOriginListResultIterator(page OriginListResultPage) OriginListResultIterator {
	return original.NewOriginListResultIterator(page)
}
func NewOriginListResultPage(getNextPage func(context.Context, OriginListResult) (OriginListResult, error)) OriginListResultPage {
	return original.NewOriginListResultPage(getNextPage)
}
func NewOriginsClient(subscriptionID string) OriginsClient {
	return original.NewOriginsClient(subscriptionID)
}
func NewOriginsClientWithBaseURI(baseURI string, subscriptionID string) OriginsClient {
	return original.NewOriginsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProfileListResultIterator(page ProfileListResultPage) ProfileListResultIterator {
	return original.NewProfileListResultIterator(page)
}
func NewProfileListResultPage(getNextPage func(context.Context, ProfileListResult) (ProfileListResult, error)) ProfileListResultPage {
	return original.NewProfileListResultPage(getNextPage)
}
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return original.NewProfilesClient(subscriptionID)
}
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return original.NewProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceUsageClient(subscriptionID string) ResourceUsageClient {
	return original.NewResourceUsageClient(subscriptionID)
}
func NewResourceUsageClientWithBaseURI(baseURI string, subscriptionID string) ResourceUsageClient {
	return original.NewResourceUsageClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceUsageListResultIterator(page ResourceUsageListResultPage) ResourceUsageListResultIterator {
	return original.NewResourceUsageListResultIterator(page)
}
func NewResourceUsageListResultPage(getNextPage func(context.Context, ResourceUsageListResult) (ResourceUsageListResult, error)) ResourceUsageListResultPage {
	return original.NewResourceUsageListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCacheBehaviorValues() []CacheBehavior {
	return original.PossibleCacheBehaviorValues()
}
func PossibleCertificateSourceValues() []CertificateSource {
	return original.PossibleCertificateSourceValues()
}
func PossibleCertificateTypeValues() []CertificateType {
	return original.PossibleCertificateTypeValues()
}
func PossibleCustomDomainResourceStateValues() []CustomDomainResourceState {
	return original.PossibleCustomDomainResourceStateValues()
}
func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return original.PossibleCustomHTTPSProvisioningStateValues()
}
func PossibleCustomHTTPSProvisioningSubstateValues() []CustomHTTPSProvisioningSubstate {
	return original.PossibleCustomHTTPSProvisioningSubstateValues()
}
func PossibleDestinationProtocolValues() []DestinationProtocol {
	return original.PossibleDestinationProtocolValues()
}
func PossibleEndpointResourceStateValues() []EndpointResourceState {
	return original.PossibleEndpointResourceStateValues()
}
func PossibleGeoFilterActionsValues() []GeoFilterActions {
	return original.PossibleGeoFilterActionsValues()
}
func PossibleHeaderActionValues() []HeaderAction {
	return original.PossibleHeaderActionValues()
}
func PossibleNameBasicDeliveryRuleActionValues() []NameBasicDeliveryRuleAction {
	return original.PossibleNameBasicDeliveryRuleActionValues()
}
func PossibleNameValues() []Name {
	return original.PossibleNameValues()
}
func PossibleOptimizationTypeValues() []OptimizationType {
	return original.PossibleOptimizationTypeValues()
}
func PossibleOriginResourceStateValues() []OriginResourceState {
	return original.PossibleOriginResourceStateValues()
}
func PossiblePostArgsOperatorValues() []PostArgsOperator {
	return original.PossiblePostArgsOperatorValues()
}
func PossibleProfileResourceStateValues() []ProfileResourceState {
	return original.PossibleProfileResourceStateValues()
}
func PossibleProtocolTypeValues() []ProtocolType {
	return original.PossibleProtocolTypeValues()
}
func PossibleQueryStringBehaviorValues() []QueryStringBehavior {
	return original.PossibleQueryStringBehaviorValues()
}
func PossibleQueryStringCachingBehaviorValues() []QueryStringCachingBehavior {
	return original.PossibleQueryStringCachingBehaviorValues()
}
func PossibleQueryStringOperatorValues() []QueryStringOperator {
	return original.PossibleQueryStringOperatorValues()
}
func PossibleRedirectTypeValues() []RedirectType {
	return original.PossibleRedirectTypeValues()
}
func PossibleRemoteAddressOperatorValues() []RemoteAddressOperator {
	return original.PossibleRemoteAddressOperatorValues()
}
func PossibleRequestBodyOperatorValues() []RequestBodyOperator {
	return original.PossibleRequestBodyOperatorValues()
}
func PossibleRequestHeaderOperatorValues() []RequestHeaderOperator {
	return original.PossibleRequestHeaderOperatorValues()
}
func PossibleRequestURIOperatorValues() []RequestURIOperator {
	return original.PossibleRequestURIOperatorValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleTransformValues() []Transform {
	return original.PossibleTransformValues()
}
func PossibleURLFileExtensionOperatorValues() []URLFileExtensionOperator {
	return original.PossibleURLFileExtensionOperatorValues()
}
func PossibleURLFileNameOperatorValues() []URLFileNameOperator {
	return original.PossibleURLFileNameOperatorValues()
}
func PossibleURLPathOperatorValues() []URLPathOperator {
	return original.PossibleURLPathOperatorValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
