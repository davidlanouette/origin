// Package eventgrid implements the Azure ARM Eventgrid service API version 2018-01-01.
//
// EventGrid Client
package eventgrid

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

// BaseClient is the base client for Eventgrid.
type BaseClient struct {
	autorest.Client
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithoutDefaults()
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults() BaseClient {
	return BaseClient{
		Client: autorest.NewClientWithUserAgent(UserAgent()),
	}
}

// PublishEvents publishes a batch of events to an Azure Event Grid topic.
// Parameters:
// topicHostname - the host name of the topic, e.g. topic1.westus2-1.eventgrid.azure.net
// events - an array of events to be published to Event Grid.
func (client BaseClient) PublishEvents(ctx context.Context, topicHostname string, events []Event) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PublishEvents")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: events,
			Constraints: []validation.Constraint{{Target: "events", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("eventgrid.BaseClient", "PublishEvents", err.Error())
	}

	req, err := client.PublishEventsPreparer(ctx, topicHostname, events)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishEvents", nil, "Failure preparing request")
		return
	}

	resp, err := client.PublishEventsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishEvents", resp, "Failure sending request")
		return
	}

	result, err = client.PublishEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.BaseClient", "PublishEvents", resp, "Failure responding to request")
	}

	return
}

// PublishEventsPreparer prepares the PublishEvents request.
func (client BaseClient) PublishEventsPreparer(ctx context.Context, topicHostname string, events []Event) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"topicHostname": topicHostname,
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{topicHostname}", urlParameters),
		autorest.WithPath("/api/events"),
		autorest.WithJSON(events),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PublishEventsSender sends the PublishEvents request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PublishEventsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// PublishEventsResponder handles the response to the PublishEvents request. The method always
// closes the http.Response Body.
func (client BaseClient) PublishEventsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
