package search

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/satori/uuid"
	"net/http"
)

// AdminKeysClient is the client that can be used to manage Azure Search
// services and API keys.
type AdminKeysClient struct {
	ManagementClient
}

// NewAdminKeysClient creates an instance of the AdminKeysClient client.
func NewAdminKeysClient(subscriptionID string) AdminKeysClient {
	return NewAdminKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAdminKeysClientWithBaseURI creates an instance of the AdminKeysClient
// client.
func NewAdminKeysClientWithBaseURI(baseURI string, subscriptionID string) AdminKeysClient {
	return AdminKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the primary and secondary admin API keys for the specified Azure
// Search service.
//
// resourceGroupName is the name of the resource group within the current
// subscription. You can obtain this value from the Azure Resource Manager
// API or the portal. searchServiceName is the name of the Azure Search
// service associated with the specified resource group. xMsClientRequestID
// is a client-generated GUID value that identifies this request. If
// specified, this will be included in response information as a way to track
// the request.
func (client AdminKeysClient) Get(resourceGroupName string, searchServiceName string, xMsClientRequestID *uuid.UUID) (result AdminKeyResult, err error) {
	req, err := client.GetPreparer(resourceGroupName, searchServiceName, xMsClientRequestID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.AdminKeysClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search.AdminKeysClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.AdminKeysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AdminKeysClient) GetPreparer(resourceGroupName string, searchServiceName string, xMsClientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if xMsClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(xMsClientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AdminKeysClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AdminKeysClient) GetResponder(resp *http.Response) (result AdminKeyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Regenerate regenerates either the primary or secondary admin API key. You
// can only regenerate one key at a time.
//
// resourceGroupName is the name of the resource group within the current
// subscription. You can obtain this value from the Azure Resource Manager
// API or the portal. searchServiceName is the name of the Azure Search
// service associated with the specified resource group. keyKind is specifies
// which key to regenerate. Valid values include 'primary' and 'secondary'.
// Possible values include: 'primary', 'secondary' xMsClientRequestID is a
// client-generated GUID value that identifies this request. If specified,
// this will be included in response information as a way to track the
// request.
func (client AdminKeysClient) Regenerate(resourceGroupName string, searchServiceName string, keyKind AdminKeyKind, xMsClientRequestID *uuid.UUID) (result AdminKeyResult, err error) {
	req, err := client.RegeneratePreparer(resourceGroupName, searchServiceName, keyKind, xMsClientRequestID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.AdminKeysClient", "Regenerate", nil, "Failure preparing request")
	}

	resp, err := client.RegenerateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search.AdminKeysClient", "Regenerate", resp, "Failure sending request")
	}

	result, err = client.RegenerateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.AdminKeysClient", "Regenerate", resp, "Failure responding to request")
	}

	return
}

// RegeneratePreparer prepares the Regenerate request.
func (client AdminKeysClient) RegeneratePreparer(resourceGroupName string, searchServiceName string, keyKind AdminKeyKind, xMsClientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"keyKind":           autorest.Encode("path", keyKind),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/regenerateAdminKey/{keyKind}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if xMsClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(xMsClientRequestID)))
	}
	return preparer.Prepare(&http.Request{})
}

// RegenerateSender sends the Regenerate request. The method will close the
// http.Response Body if it receives an error.
func (client AdminKeysClient) RegenerateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RegenerateResponder handles the response to the Regenerate request. The method always
// closes the http.Response Body.
func (client AdminKeysClient) RegenerateResponder(resp *http.Response) (result AdminKeyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
