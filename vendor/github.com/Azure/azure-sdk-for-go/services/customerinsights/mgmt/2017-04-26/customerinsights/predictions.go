package customerinsights

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

// PredictionsClient is the the Azure Customer Insights management API provides a RESTful set of web services that
// interact with Azure Customer Insights service to manage your resources. The API has entities that capture the
// relationship between an end user and the Azure Customer Insights service.
type PredictionsClient struct {
	BaseClient
}

// NewPredictionsClient creates an instance of the PredictionsClient client.
func NewPredictionsClient(subscriptionID string) PredictionsClient {
	return NewPredictionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPredictionsClientWithBaseURI creates an instance of the PredictionsClient client.
func NewPredictionsClientWithBaseURI(baseURI string, subscriptionID string) PredictionsClient {
	return PredictionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a Prediction or updates an existing Prediction in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// predictionName - the name of the Prediction.
// parameters - parameters supplied to the create/update Prediction operation.
func (client PredictionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionResourceFormat) (result PredictionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: predictionName,
			Constraints: []validation.Constraint{{Target: "predictionName", Name: validation.MaxLength, Rule: 512, Chain: nil},
				{Target: "predictionName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Prediction", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Prediction.NegativeOutcomeExpression", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Prediction.PositiveOutcomeExpression", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Prediction.PrimaryProfileType", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Prediction.ScopeExpression", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Prediction.AutoAnalyze", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Prediction.Mappings", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.Prediction.Mappings.Score", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.Prediction.Mappings.Grade", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.Prediction.Mappings.Reason", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "parameters.Prediction.ScoreLabel", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("customerinsights.PredictionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hubName, predictionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PredictionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"predictionName":    autorest.Encode("path", predictionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionsClient) CreateOrUpdateSender(req *http.Request) (future PredictionsCreateOrUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PredictionsClient) CreateOrUpdateResponder(resp *http.Response) (result PredictionResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Prediction in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// predictionName - the name of the Prediction.
func (client PredictionsClient) Delete(ctx context.Context, resourceGroupName string, hubName string, predictionName string) (result PredictionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, hubName, predictionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PredictionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, hubName string, predictionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"predictionName":    autorest.Encode("path", predictionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionsClient) DeleteSender(req *http.Request) (future PredictionsDeleteFuture, err error) {
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
func (client PredictionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Prediction in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// predictionName - the name of the Prediction.
func (client PredictionsClient) Get(ctx context.Context, resourceGroupName string, hubName string, predictionName string) (result PredictionResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hubName, predictionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PredictionsClient) GetPreparer(ctx context.Context, resourceGroupName string, hubName string, predictionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"predictionName":    autorest.Encode("path", predictionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PredictionsClient) GetResponder(resp *http.Response) (result PredictionResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetModelStatus gets model status of the prediction.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// predictionName - the name of the Prediction.
func (client PredictionsClient) GetModelStatus(ctx context.Context, resourceGroupName string, hubName string, predictionName string) (result PredictionModelStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionsClient.GetModelStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetModelStatusPreparer(ctx, resourceGroupName, hubName, predictionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "GetModelStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetModelStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "GetModelStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetModelStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "GetModelStatus", resp, "Failure responding to request")
	}

	return
}

// GetModelStatusPreparer prepares the GetModelStatus request.
func (client PredictionsClient) GetModelStatusPreparer(ctx context.Context, resourceGroupName string, hubName string, predictionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"predictionName":    autorest.Encode("path", predictionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}/getModelStatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetModelStatusSender sends the GetModelStatus request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionsClient) GetModelStatusSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetModelStatusResponder handles the response to the GetModelStatus request. The method always
// closes the http.Response Body.
func (client PredictionsClient) GetModelStatusResponder(resp *http.Response) (result PredictionModelStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTrainingResults gets training results.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// predictionName - the name of the Prediction.
func (client PredictionsClient) GetTrainingResults(ctx context.Context, resourceGroupName string, hubName string, predictionName string) (result PredictionTrainingResults, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionsClient.GetTrainingResults")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetTrainingResultsPreparer(ctx, resourceGroupName, hubName, predictionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "GetTrainingResults", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTrainingResultsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "GetTrainingResults", resp, "Failure sending request")
		return
	}

	result, err = client.GetTrainingResultsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "GetTrainingResults", resp, "Failure responding to request")
	}

	return
}

// GetTrainingResultsPreparer prepares the GetTrainingResults request.
func (client PredictionsClient) GetTrainingResultsPreparer(ctx context.Context, resourceGroupName string, hubName string, predictionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"predictionName":    autorest.Encode("path", predictionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}/getTrainingResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTrainingResultsSender sends the GetTrainingResults request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionsClient) GetTrainingResultsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetTrainingResultsResponder handles the response to the GetTrainingResults request. The method always
// closes the http.Response Body.
func (client PredictionsClient) GetTrainingResultsResponder(resp *http.Response) (result PredictionTrainingResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all the predictions in the specified hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
func (client PredictionsClient) ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result PredictionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionsClient.ListByHub")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHubNextResults
	req, err := client.ListByHubPreparer(ctx, resourceGroupName, hubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "ListByHub", resp, "Failure responding to request")
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client PredictionsClient) ListByHubPreparer(ctx context.Context, resourceGroupName string, hubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionsClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client PredictionsClient) ListByHubResponder(resp *http.Response) (result PredictionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHubNextResults retrieves the next set of results, if any.
func (client PredictionsClient) listByHubNextResults(ctx context.Context, lastResults PredictionListResult) (result PredictionListResult, err error) {
	req, err := lastResults.predictionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "listByHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "listByHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "listByHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client PredictionsClient) ListByHubComplete(ctx context.Context, resourceGroupName string, hubName string) (result PredictionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionsClient.ListByHub")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHub(ctx, resourceGroupName, hubName)
	return
}

// ModelStatus creates or updates the model status of prediction.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// predictionName - the name of the Prediction.
// parameters - parameters supplied to the create/update prediction model status operation.
func (client PredictionsClient) ModelStatus(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionModelStatus) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PredictionsClient.ModelStatus")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ModelStatusPreparer(ctx, resourceGroupName, hubName, predictionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "ModelStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ModelStatusSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "ModelStatus", resp, "Failure sending request")
		return
	}

	result, err = client.ModelStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.PredictionsClient", "ModelStatus", resp, "Failure responding to request")
	}

	return
}

// ModelStatusPreparer prepares the ModelStatus request.
func (client PredictionsClient) ModelStatusPreparer(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionModelStatus) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"predictionName":    autorest.Encode("path", predictionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.TenantID = nil
	parameters.PredictionName = nil
	parameters.PredictionGUIDID = nil
	parameters.Message = nil
	parameters.TrainingSetCount = nil
	parameters.TestSetCount = nil
	parameters.ValidationSetCount = nil
	parameters.TrainingAccuracy = nil
	parameters.SignalsUsed = nil
	parameters.ModelVersion = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}/modelStatus", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ModelStatusSender sends the ModelStatus request. The method will close the
// http.Response Body if it receives an error.
func (client PredictionsClient) ModelStatusSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ModelStatusResponder handles the response to the ModelStatus request. The method always
// closes the http.Response Body.
func (client PredictionsClient) ModelStatusResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
