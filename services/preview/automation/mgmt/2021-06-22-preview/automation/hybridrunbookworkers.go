package automation

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

// HybridRunbookWorkersClient is the automation Client
type HybridRunbookWorkersClient struct {
	BaseClient
}

// NewHybridRunbookWorkersClient creates an instance of the HybridRunbookWorkersClient client.
func NewHybridRunbookWorkersClient(subscriptionID string) HybridRunbookWorkersClient {
	return NewHybridRunbookWorkersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewHybridRunbookWorkersClientWithBaseURI creates an instance of the HybridRunbookWorkersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewHybridRunbookWorkersClientWithBaseURI(baseURI string, subscriptionID string) HybridRunbookWorkersClient {
	return HybridRunbookWorkersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a hybrid runbook worker.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// hybridRunbookWorkerGroupName - the hybrid runbook worker group name
// hybridRunbookWorkerID - the hybrid runbook worker id
// hybridRunbookWorkerCreationParameters - the create or update parameters for hybrid runbook worker.
func (client HybridRunbookWorkersClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerCreationParameters HybridRunbookWorkerCreateParameters) (result HybridRunbookWorker, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HybridRunbookWorkersClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: hybridRunbookWorkerCreationParameters,
			Constraints: []validation.Constraint{{Target: "hybridRunbookWorkerCreationParameters.HybridRunbookWorkerCreateOrUpdateParameters", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.HybridRunbookWorkersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, hybridRunbookWorkerCreationParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client HybridRunbookWorkersClient) CreatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerCreationParameters HybridRunbookWorkerCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":        autorest.Encode("path", automationAccountName),
		"hybridRunbookWorkerGroupName": autorest.Encode("path", hybridRunbookWorkerGroupName),
		"hybridRunbookWorkerId":        autorest.Encode("path", hybridRunbookWorkerID),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-22"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}", pathParameters),
		autorest.WithJSON(hybridRunbookWorkerCreationParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client HybridRunbookWorkersClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client HybridRunbookWorkersClient) CreateResponder(resp *http.Response) (result HybridRunbookWorker, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a hybrid runbook worker.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// hybridRunbookWorkerGroupName - the hybrid runbook worker group name
// hybridRunbookWorkerID - the hybrid runbook worker id
func (client HybridRunbookWorkersClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HybridRunbookWorkersClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.HybridRunbookWorkersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client HybridRunbookWorkersClient) DeletePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":        autorest.Encode("path", automationAccountName),
		"hybridRunbookWorkerGroupName": autorest.Encode("path", hybridRunbookWorkerGroupName),
		"hybridRunbookWorkerId":        autorest.Encode("path", hybridRunbookWorkerID),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-22"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client HybridRunbookWorkersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client HybridRunbookWorkersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve a hybrid runbook worker.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// hybridRunbookWorkerGroupName - the hybrid runbook worker group name
// hybridRunbookWorkerID - the hybrid runbook worker id
func (client HybridRunbookWorkersClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string) (result HybridRunbookWorker, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HybridRunbookWorkersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.HybridRunbookWorkersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client HybridRunbookWorkersClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":        autorest.Encode("path", automationAccountName),
		"hybridRunbookWorkerGroupName": autorest.Encode("path", hybridRunbookWorkerGroupName),
		"hybridRunbookWorkerId":        autorest.Encode("path", hybridRunbookWorkerID),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-22"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client HybridRunbookWorkersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client HybridRunbookWorkersClient) GetResponder(resp *http.Response) (result HybridRunbookWorker, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHybridRunbookWorkerGroup retrieve a list of hybrid runbook workers.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// hybridRunbookWorkerGroupName - the hybrid runbook worker group name
// filter - the filter to apply on the operation.
func (client HybridRunbookWorkersClient) ListByHybridRunbookWorkerGroup(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, filter string) (result HybridRunbookWorkersListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HybridRunbookWorkersClient.ListByHybridRunbookWorkerGroup")
		defer func() {
			sc := -1
			if result.hrwlr.Response.Response != nil {
				sc = result.hrwlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.HybridRunbookWorkersClient", "ListByHybridRunbookWorkerGroup", err.Error())
	}

	result.fn = client.listByHybridRunbookWorkerGroupNextResults
	req, err := client.ListByHybridRunbookWorkerGroupPreparer(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "ListByHybridRunbookWorkerGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHybridRunbookWorkerGroupSender(req)
	if err != nil {
		result.hrwlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "ListByHybridRunbookWorkerGroup", resp, "Failure sending request")
		return
	}

	result.hrwlr, err = client.ListByHybridRunbookWorkerGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "ListByHybridRunbookWorkerGroup", resp, "Failure responding to request")
		return
	}
	if result.hrwlr.hasNextLink() && result.hrwlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByHybridRunbookWorkerGroupPreparer prepares the ListByHybridRunbookWorkerGroup request.
func (client HybridRunbookWorkersClient) ListByHybridRunbookWorkerGroupPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":        autorest.Encode("path", automationAccountName),
		"hybridRunbookWorkerGroupName": autorest.Encode("path", hybridRunbookWorkerGroupName),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-22"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHybridRunbookWorkerGroupSender sends the ListByHybridRunbookWorkerGroup request. The method will close the
// http.Response Body if it receives an error.
func (client HybridRunbookWorkersClient) ListByHybridRunbookWorkerGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHybridRunbookWorkerGroupResponder handles the response to the ListByHybridRunbookWorkerGroup request. The method always
// closes the http.Response Body.
func (client HybridRunbookWorkersClient) ListByHybridRunbookWorkerGroupResponder(resp *http.Response) (result HybridRunbookWorkersListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHybridRunbookWorkerGroupNextResults retrieves the next set of results, if any.
func (client HybridRunbookWorkersClient) listByHybridRunbookWorkerGroupNextResults(ctx context.Context, lastResults HybridRunbookWorkersListResult) (result HybridRunbookWorkersListResult, err error) {
	req, err := lastResults.hybridRunbookWorkersListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "listByHybridRunbookWorkerGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHybridRunbookWorkerGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "listByHybridRunbookWorkerGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHybridRunbookWorkerGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "listByHybridRunbookWorkerGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHybridRunbookWorkerGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client HybridRunbookWorkersClient) ListByHybridRunbookWorkerGroupComplete(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, filter string) (result HybridRunbookWorkersListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HybridRunbookWorkersClient.ListByHybridRunbookWorkerGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHybridRunbookWorkerGroup(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, filter)
	return
}

// Move move a hybrid worker to a different group.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// hybridRunbookWorkerGroupName - the hybrid runbook worker group name
// hybridRunbookWorkerID - the hybrid runbook worker id
// hybridRunbookWorkerMoveParameters - the hybrid runbook worker move parameters
func (client HybridRunbookWorkersClient) Move(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerMoveParameters HybridRunbookWorkerMoveParameters) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HybridRunbookWorkersClient.Move")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.HybridRunbookWorkersClient", "Move", err.Error())
	}

	req, err := client.MovePreparer(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, hybridRunbookWorkerMoveParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Move", nil, "Failure preparing request")
		return
	}

	resp, err := client.MoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Move", resp, "Failure sending request")
		return
	}

	result, err = client.MoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.HybridRunbookWorkersClient", "Move", resp, "Failure responding to request")
		return
	}

	return
}

// MovePreparer prepares the Move request.
func (client HybridRunbookWorkersClient) MovePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerMoveParameters HybridRunbookWorkerMoveParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName":        autorest.Encode("path", automationAccountName),
		"hybridRunbookWorkerGroupName": autorest.Encode("path", hybridRunbookWorkerGroupName),
		"hybridRunbookWorkerId":        autorest.Encode("path", hybridRunbookWorkerID),
		"resourceGroupName":            autorest.Encode("path", resourceGroupName),
		"subscriptionId":               autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-22"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}/move", pathParameters),
		autorest.WithJSON(hybridRunbookWorkerMoveParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MoveSender sends the Move request. The method will close the
// http.Response Body if it receives an error.
func (client HybridRunbookWorkersClient) MoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// MoveResponder handles the response to the Move request. The method always
// closes the http.Response Body.
func (client HybridRunbookWorkersClient) MoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
