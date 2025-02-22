package hdinsight

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

// VirtualMachinesClient is the hDInsight Management Client
type VirtualMachinesClient struct {
	BaseClient
}

// NewVirtualMachinesClient creates an instance of the VirtualMachinesClient client.
func NewVirtualMachinesClient(subscriptionID string) VirtualMachinesClient {
	return NewVirtualMachinesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachinesClientWithBaseURI creates an instance of the VirtualMachinesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVirtualMachinesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachinesClient {
	return VirtualMachinesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetAsyncOperationStatus gets the async operation status.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// operationID - the long running operation id.
func (client VirtualMachinesClient) GetAsyncOperationStatus(ctx context.Context, resourceGroupName string, clusterName string, operationID string) (result AsyncOperationResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.GetAsyncOperationStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAsyncOperationStatusPreparer(ctx, resourceGroupName, clusterName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.VirtualMachinesClient", "GetAsyncOperationStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAsyncOperationStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.VirtualMachinesClient", "GetAsyncOperationStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetAsyncOperationStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.VirtualMachinesClient", "GetAsyncOperationStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetAsyncOperationStatusPreparer prepares the GetAsyncOperationStatus request.
func (client VirtualMachinesClient) GetAsyncOperationStatusPreparer(ctx context.Context, resourceGroupName string, clusterName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/restartHosts/azureasyncoperations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAsyncOperationStatusSender sends the GetAsyncOperationStatus request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) GetAsyncOperationStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAsyncOperationStatusResponder handles the response to the GetAsyncOperationStatus request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) GetAsyncOperationStatusResponder(resp *http.Response) (result AsyncOperationResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListHosts lists the HDInsight clusters hosts
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
func (client VirtualMachinesClient) ListHosts(ctx context.Context, resourceGroupName string, clusterName string) (result ListHostInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.ListHosts")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListHostsPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.VirtualMachinesClient", "ListHosts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListHostsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.VirtualMachinesClient", "ListHosts", resp, "Failure sending request")
		return
	}

	result, err = client.ListHostsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.VirtualMachinesClient", "ListHosts", resp, "Failure responding to request")
		return
	}

	return
}

// ListHostsPreparer prepares the ListHosts request.
func (client VirtualMachinesClient) ListHostsPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/listHosts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListHostsSender sends the ListHosts request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) ListHostsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListHostsResponder handles the response to the ListHosts request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) ListHostsResponder(resp *http.Response) (result ListHostInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RestartHosts restarts the specified HDInsight cluster hosts.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// hosts - the list of hosts to restart
func (client VirtualMachinesClient) RestartHosts(ctx context.Context, resourceGroupName string, clusterName string, hosts []string) (result VirtualMachinesRestartHostsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachinesClient.RestartHosts")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: hosts,
			Constraints: []validation.Constraint{{Target: "hosts", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hdinsight.VirtualMachinesClient", "RestartHosts", err.Error())
	}

	req, err := client.RestartHostsPreparer(ctx, resourceGroupName, clusterName, hosts)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.VirtualMachinesClient", "RestartHosts", nil, "Failure preparing request")
		return
	}

	result, err = client.RestartHostsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.VirtualMachinesClient", "RestartHosts", result.Response(), "Failure sending request")
		return
	}

	return
}

// RestartHostsPreparer prepares the RestartHosts request.
func (client VirtualMachinesClient) RestartHostsPreparer(ctx context.Context, resourceGroupName string, clusterName string, hosts []string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/restartHosts", pathParameters),
		autorest.WithJSON(hosts),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartHostsSender sends the RestartHosts request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachinesClient) RestartHostsSender(req *http.Request) (future VirtualMachinesRestartHostsFuture, err error) {
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

// RestartHostsResponder handles the response to the RestartHosts request. The method always
// closes the http.Response Body.
func (client VirtualMachinesClient) RestartHostsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
