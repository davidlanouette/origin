package appconfiguration

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ConfigurationStoresClient is the client for the ConfigurationStores methods of the Appconfiguration service.
type ConfigurationStoresClient struct {
	BaseClient
}

// NewConfigurationStoresClient creates an instance of the ConfigurationStoresClient client.
func NewConfigurationStoresClient(subscriptionID string) ConfigurationStoresClient {
	return NewConfigurationStoresClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationStoresClientWithBaseURI creates an instance of the ConfigurationStoresClient client.
func NewConfigurationStoresClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationStoresClient {
	return ConfigurationStoresClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a configuration store with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
// configStoreCreationParameters - the parameters for creating a configuration store.
func (client ConfigurationStoresClient) Create(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore) (result ConfigurationStoresCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.ConfigurationStoresClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, configStoreName, configStoreCreationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ConfigurationStoresClient) CreatePreparer(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters ConfigurationStore) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", pathParameters),
		autorest.WithJSON(configStoreCreationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationStoresClient) CreateSender(req *http.Request) (future ConfigurationStoresCreateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ConfigurationStoresClient) CreateResponder(resp *http.Response) (result ConfigurationStore, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a configuration store.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
func (client ConfigurationStoresClient) Delete(ctx context.Context, resourceGroupName string, configStoreName string) (result ConfigurationStoresDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.ConfigurationStoresClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, configStoreName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConfigurationStoresClient) DeletePreparer(ctx context.Context, resourceGroupName string, configStoreName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationStoresClient) DeleteSender(req *http.Request) (future ConfigurationStoresDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConfigurationStoresClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified configuration store.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
func (client ConfigurationStoresClient) Get(ctx context.Context, resourceGroupName string, configStoreName string) (result ConfigurationStore, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.ConfigurationStoresClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, configStoreName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationStoresClient) GetPreparer(ctx context.Context, resourceGroupName string, configStoreName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationStoresClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationStoresClient) GetResponder(resp *http.Response) (result ConfigurationStore, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the configuration stores for a given subscription.
// Parameters:
// skipToken - a skip token is used to continue retrieving items after an operation returns a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a
// skipToken parameter that specifies a starting point to use for subsequent calls.
func (client ConfigurationStoresClient) List(ctx context.Context, skipToken string) (result ConfigurationStoreListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.List")
		defer func() {
			sc := -1
			if result.cslr.Response.Response != nil {
				sc = result.cslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "List", resp, "Failure sending request")
		return
	}

	result.cslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ConfigurationStoresClient) ListPreparer(ctx context.Context, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/configurationStores", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationStoresClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ConfigurationStoresClient) ListResponder(resp *http.Response) (result ConfigurationStoreListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ConfigurationStoresClient) listNextResults(ctx context.Context, lastResults ConfigurationStoreListResult) (result ConfigurationStoreListResult, err error) {
	req, err := lastResults.configurationStoreListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConfigurationStoresClient) ListComplete(ctx context.Context, skipToken string) (result ConfigurationStoreListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, skipToken)
	return
}

// ListByResourceGroup lists the configuration stores for a given resource group.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// skipToken - a skip token is used to continue retrieving items after an operation returns a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a
// skipToken parameter that specifies a starting point to use for subsequent calls.
func (client ConfigurationStoresClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, skipToken string) (result ConfigurationStoreListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.cslr.Response.Response != nil {
				sc = result.cslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.cslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.cslr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ConfigurationStoresClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationStoresClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ConfigurationStoresClient) ListByResourceGroupResponder(resp *http.Response) (result ConfigurationStoreListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ConfigurationStoresClient) listByResourceGroupNextResults(ctx context.Context, lastResults ConfigurationStoreListResult) (result ConfigurationStoreListResult, err error) {
	req, err := lastResults.configurationStoreListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConfigurationStoresClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skipToken string) (result ConfigurationStoreListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, skipToken)
	return
}

// ListKeys lists the access key for the specified configuration store.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
// skipToken - a skip token is used to continue retrieving items after an operation returns a partial result.
// If a previous response contains a nextLink element, the value of the nextLink element will include a
// skipToken parameter that specifies a starting point to use for subsequent calls.
func (client ConfigurationStoresClient) ListKeys(ctx context.Context, resourceGroupName string, configStoreName string, skipToken string) (result APIKeyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.ListKeys")
		defer func() {
			sc := -1
			if result.aklr.Response.Response != nil {
				sc = result.aklr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.ConfigurationStoresClient", "ListKeys", err.Error())
	}

	result.fn = client.listKeysNextResults
	req, err := client.ListKeysPreparer(ctx, resourceGroupName, configStoreName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.aklr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "ListKeys", resp, "Failure sending request")
		return
	}

	result.aklr, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client ConfigurationStoresClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, configStoreName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/ListKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationStoresClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client ConfigurationStoresClient) ListKeysResponder(resp *http.Response) (result APIKeyListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listKeysNextResults retrieves the next set of results, if any.
func (client ConfigurationStoresClient) listKeysNextResults(ctx context.Context, lastResults APIKeyListResult) (result APIKeyListResult, err error) {
	req, err := lastResults.aPIKeyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listKeysNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listKeysNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "listKeysNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListKeysComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConfigurationStoresClient) ListKeysComplete(ctx context.Context, resourceGroupName string, configStoreName string, skipToken string) (result APIKeyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.ListKeys")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListKeys(ctx, resourceGroupName, configStoreName, skipToken)
	return
}

// RegenerateKey regenerates an access key for the specified configuration store.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
// regenerateKeyParameters - the parameters for regenerating an access key.
func (client ConfigurationStoresClient) RegenerateKey(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters) (result APIKey, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.RegenerateKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.ConfigurationStoresClient", "RegenerateKey", err.Error())
	}

	req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, configStoreName, regenerateKeyParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "RegenerateKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "RegenerateKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "RegenerateKey", resp, "Failure responding to request")
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client ConfigurationStoresClient) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters RegenerateKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/RegenerateKey", pathParameters),
		autorest.WithJSON(regenerateKeyParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationStoresClient) RegenerateKeySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client ConfigurationStoresClient) RegenerateKeyResponder(resp *http.Response) (result APIKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a configuration store with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// configStoreName - the name of the configuration store.
// configStoreUpdateParameters - the parameters for updating a configuration store.
func (client ConfigurationStoresClient) Update(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters) (result ConfigurationStoresUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoresClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: configStoreName,
			Constraints: []validation.Constraint{{Target: "configStoreName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "configStoreName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "configStoreName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appconfiguration.ConfigurationStoresClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, configStoreName, configStoreUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ConfigurationStoresClient) UpdatePreparer(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters ConfigurationStoreUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"configStoreName":   autorest.Encode("path", configStoreName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", pathParameters),
		autorest.WithJSON(configStoreUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationStoresClient) UpdateSender(req *http.Request) (future ConfigurationStoresUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ConfigurationStoresClient) UpdateResponder(resp *http.Response) (result ConfigurationStore, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
