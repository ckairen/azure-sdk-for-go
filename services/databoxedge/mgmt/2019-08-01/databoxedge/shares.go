package databoxedge

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// SharesClient is the client for the Shares methods of the Databoxedge service.
type SharesClient struct {
	BaseClient
}

// NewSharesClient creates an instance of the SharesClient client.
func NewSharesClient(subscriptionID string) SharesClient {
	return NewSharesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSharesClientWithBaseURI creates an instance of the SharesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSharesClientWithBaseURI(baseURI string, subscriptionID string) SharesClient {
	return SharesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// deviceName - the device name.
// name - the share name.
// share - the share properties.
// resourceGroupName - the resource group name.
func (client SharesClient) CreateOrUpdate(ctx context.Context, deviceName string, name string, share Share, resourceGroupName string) (result SharesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: share,
			Constraints: []validation.Constraint{{Target: "share.ShareProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "share.ShareProperties.AzureContainerInfo", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "share.ShareProperties.AzureContainerInfo.StorageAccountCredentialID", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "share.ShareProperties.AzureContainerInfo.ContainerName", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("databoxedge.SharesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, name, share, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SharesClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, name string, share Share, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", pathParameters),
		autorest.WithJSON(share),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) CreateOrUpdateSender(req *http.Request) (future SharesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SharesClient) CreateOrUpdateResponder(resp *http.Response) (result Share, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the share on the Data Box Edge/Data Box Gateway device.
// Parameters:
// deviceName - the device name.
// name - the share name.
// resourceGroupName - the resource group name.
func (client SharesClient) Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result SharesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, deviceName, name, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SharesClient) DeletePreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) DeleteSender(req *http.Request) (future SharesDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SharesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// deviceName - the device name.
// name - the share name.
// resourceGroupName - the resource group name.
func (client SharesClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result Share, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, deviceName, name, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SharesClient) GetPreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SharesClient) GetResponder(resp *http.Response) (result Share, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDataBoxEdgeDevice sends the list by data box edge device request.
// Parameters:
// deviceName - the device name.
// resourceGroupName - the resource group name.
func (client SharesClient) ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result ShareListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.ListByDataBoxEdgeDevice")
		defer func() {
			sc := -1
			if result.sl.Response.Response != nil {
				sc = result.sl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDataBoxEdgeDeviceNextResults
	req, err := client.ListByDataBoxEdgeDevicePreparer(ctx, deviceName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "ListByDataBoxEdgeDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataBoxEdgeDeviceSender(req)
	if err != nil {
		result.sl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "ListByDataBoxEdgeDevice", resp, "Failure sending request")
		return
	}

	result.sl, err = client.ListByDataBoxEdgeDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "ListByDataBoxEdgeDevice", resp, "Failure responding to request")
		return
	}
	if result.sl.hasNextLink() && result.sl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDataBoxEdgeDevicePreparer prepares the ListByDataBoxEdgeDevice request.
func (client SharesClient) ListByDataBoxEdgeDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataBoxEdgeDeviceSender sends the ListByDataBoxEdgeDevice request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) ListByDataBoxEdgeDeviceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDataBoxEdgeDeviceResponder handles the response to the ListByDataBoxEdgeDevice request. The method always
// closes the http.Response Body.
func (client SharesClient) ListByDataBoxEdgeDeviceResponder(resp *http.Response) (result ShareList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataBoxEdgeDeviceNextResults retrieves the next set of results, if any.
func (client SharesClient) listByDataBoxEdgeDeviceNextResults(ctx context.Context, lastResults ShareList) (result ShareList, err error) {
	req, err := lastResults.shareListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databoxedge.SharesClient", "listByDataBoxEdgeDeviceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataBoxEdgeDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databoxedge.SharesClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataBoxEdgeDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDataBoxEdgeDeviceComplete enumerates all values, automatically crossing page boundaries as required.
func (client SharesClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result ShareListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.ListByDataBoxEdgeDevice")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataBoxEdgeDevice(ctx, deviceName, resourceGroupName)
	return
}

// Refresh sends the refresh request.
// Parameters:
// deviceName - the device name.
// name - the share name.
// resourceGroupName - the resource group name.
func (client SharesClient) Refresh(ctx context.Context, deviceName string, name string, resourceGroupName string) (result SharesRefreshFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharesClient.Refresh")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RefreshPreparer(ctx, deviceName, name, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "Refresh", nil, "Failure preparing request")
		return
	}

	result, err = client.RefreshSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.SharesClient", "Refresh", result.Response(), "Failure sending request")
		return
	}

	return
}

// RefreshPreparer prepares the Refresh request.
func (client SharesClient) RefreshPreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}/refresh", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshSender sends the Refresh request. The method will close the
// http.Response Body if it receives an error.
func (client SharesClient) RefreshSender(req *http.Request) (future SharesRefreshFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RefreshResponder handles the response to the Refresh request. The method always
// closes the http.Response Body.
func (client SharesClient) RefreshResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
