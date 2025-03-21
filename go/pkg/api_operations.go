/*
H3Net Protocol API

H3Net Protocol schema with various messages and configurations.

API version: 1.0.0
Contact: h3netprotocol@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package h3netclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type OperationsAPI interface {

	/*
	GetDeviceInterfaceStates Get device interface states

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deviceId
	@return OperationsAPIGetDeviceInterfaceStatesRequest
	*/
	GetDeviceInterfaceStates(ctx context.Context, deviceId string) OperationsAPIGetDeviceInterfaceStatesRequest

	// GetDeviceInterfaceStatesExecute executes the request
	//  @return []InterfaceOperationalState
	GetDeviceInterfaceStatesExecute(r OperationsAPIGetDeviceInterfaceStatesRequest) ([]InterfaceOperationalState, *http.Response, error)

	/*
	GetDeviceOperationalState Get device operational state

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deviceId
	@return OperationsAPIGetDeviceOperationalStateRequest
	*/
	GetDeviceOperationalState(ctx context.Context, deviceId string) OperationsAPIGetDeviceOperationalStateRequest

	// GetDeviceOperationalStateExecute executes the request
	//  @return DeviceOperationalState
	GetDeviceOperationalStateExecute(r OperationsAPIGetDeviceOperationalStateRequest) (*DeviceOperationalState, *http.Response, error)

	/*
	GetOperationalState Get operational state

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OperationsAPIGetOperationalStateRequest
	*/
	GetOperationalState(ctx context.Context) OperationsAPIGetOperationalStateRequest

	// GetOperationalStateExecute executes the request
	//  @return OperationalState
	GetOperationalStateExecute(r OperationsAPIGetOperationalStateRequest) (*OperationalState, *http.Response, error)

	/*
	GetTableEntries Get table entries

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deviceId
	@return OperationsAPIGetTableEntriesRequest
	*/
	GetTableEntries(ctx context.Context, deviceId string) OperationsAPIGetTableEntriesRequest

	// GetTableEntriesExecute executes the request
	//  @return StateQueryResponseMessage
	GetTableEntriesExecute(r OperationsAPIGetTableEntriesRequest) (*StateQueryResponseMessage, *http.Response, error)

	/*
	QueryDeviceState Query Device State

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OperationsAPIQueryDeviceStateRequest
	*/
	QueryDeviceState(ctx context.Context) OperationsAPIQueryDeviceStateRequest

	// QueryDeviceStateExecute executes the request
	//  @return StateQueryResponseMessage
	QueryDeviceStateExecute(r OperationsAPIQueryDeviceStateRequest) (*StateQueryResponseMessage, *http.Response, error)

	/*
	QuerySpecificDeviceState Query specific device state

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deviceId
	@return OperationsAPIQuerySpecificDeviceStateRequest
	*/
	QuerySpecificDeviceState(ctx context.Context, deviceId string) OperationsAPIQuerySpecificDeviceStateRequest

	// QuerySpecificDeviceStateExecute executes the request
	//  @return StateQueryResponseMessage
	QuerySpecificDeviceStateExecute(r OperationsAPIQuerySpecificDeviceStateRequest) (*StateQueryResponseMessage, *http.Response, error)

	/*
	RebootDevice Reboot device

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deviceId
	@return OperationsAPIRebootDeviceRequest
	*/
	RebootDevice(ctx context.Context, deviceId string) OperationsAPIRebootDeviceRequest

	// RebootDeviceExecute executes the request
	//  @return ResponseMessage
	RebootDeviceExecute(r OperationsAPIRebootDeviceRequest) (*ResponseMessage, *http.Response, error)
}

// OperationsAPIService OperationsAPI service
type OperationsAPIService service

type OperationsAPIGetDeviceInterfaceStatesRequest struct {
	ctx context.Context
	ApiService OperationsAPI
	deviceId string
}

func (r OperationsAPIGetDeviceInterfaceStatesRequest) Execute() ([]InterfaceOperationalState, *http.Response, error) {
	return r.ApiService.GetDeviceInterfaceStatesExecute(r)
}

/*
GetDeviceInterfaceStates Get device interface states

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return OperationsAPIGetDeviceInterfaceStatesRequest
*/
func (a *OperationsAPIService) GetDeviceInterfaceStates(ctx context.Context, deviceId string) OperationsAPIGetDeviceInterfaceStatesRequest {
	return OperationsAPIGetDeviceInterfaceStatesRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return []InterfaceOperationalState
func (a *OperationsAPIService) GetDeviceInterfaceStatesExecute(r OperationsAPIGetDeviceInterfaceStatesRequest) ([]InterfaceOperationalState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InterfaceOperationalState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsAPIService.GetDeviceInterfaceStates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operations/state/{deviceId}/interfaces"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["operationsAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OperationsAPIGetDeviceOperationalStateRequest struct {
	ctx context.Context
	ApiService OperationsAPI
	deviceId string
}

func (r OperationsAPIGetDeviceOperationalStateRequest) Execute() (*DeviceOperationalState, *http.Response, error) {
	return r.ApiService.GetDeviceOperationalStateExecute(r)
}

/*
GetDeviceOperationalState Get device operational state

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return OperationsAPIGetDeviceOperationalStateRequest
*/
func (a *OperationsAPIService) GetDeviceOperationalState(ctx context.Context, deviceId string) OperationsAPIGetDeviceOperationalStateRequest {
	return OperationsAPIGetDeviceOperationalStateRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return DeviceOperationalState
func (a *OperationsAPIService) GetDeviceOperationalStateExecute(r OperationsAPIGetDeviceOperationalStateRequest) (*DeviceOperationalState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeviceOperationalState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsAPIService.GetDeviceOperationalState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operations/state/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["operationsAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OperationsAPIGetOperationalStateRequest struct {
	ctx context.Context
	ApiService OperationsAPI
}

func (r OperationsAPIGetOperationalStateRequest) Execute() (*OperationalState, *http.Response, error) {
	return r.ApiService.GetOperationalStateExecute(r)
}

/*
GetOperationalState Get operational state

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OperationsAPIGetOperationalStateRequest
*/
func (a *OperationsAPIService) GetOperationalState(ctx context.Context) OperationsAPIGetOperationalStateRequest {
	return OperationsAPIGetOperationalStateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OperationalState
func (a *OperationsAPIService) GetOperationalStateExecute(r OperationsAPIGetOperationalStateRequest) (*OperationalState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OperationalState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsAPIService.GetOperationalState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operations/state"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["operationsAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OperationsAPIGetTableEntriesRequest struct {
	ctx context.Context
	ApiService OperationsAPI
	deviceId string
	table *string
}

func (r OperationsAPIGetTableEntriesRequest) Table(table string) OperationsAPIGetTableEntriesRequest {
	r.table = &table
	return r
}

func (r OperationsAPIGetTableEntriesRequest) Execute() (*StateQueryResponseMessage, *http.Response, error) {
	return r.ApiService.GetTableEntriesExecute(r)
}

/*
GetTableEntries Get table entries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return OperationsAPIGetTableEntriesRequest
*/
func (a *OperationsAPIService) GetTableEntries(ctx context.Context, deviceId string) OperationsAPIGetTableEntriesRequest {
	return OperationsAPIGetTableEntriesRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return StateQueryResponseMessage
func (a *OperationsAPIService) GetTableEntriesExecute(r OperationsAPIGetTableEntriesRequest) (*StateQueryResponseMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StateQueryResponseMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsAPIService.GetTableEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operations/state/{deviceId}/tables"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.table == nil {
		return localVarReturnValue, nil, reportError("table is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "table", r.table, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["operationsAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OperationsAPIQueryDeviceStateRequest struct {
	ctx context.Context
	ApiService OperationsAPI
	stateQueryMessage *StateQueryMessage
}

// State query parameters
func (r OperationsAPIQueryDeviceStateRequest) StateQueryMessage(stateQueryMessage StateQueryMessage) OperationsAPIQueryDeviceStateRequest {
	r.stateQueryMessage = &stateQueryMessage
	return r
}

func (r OperationsAPIQueryDeviceStateRequest) Execute() (*StateQueryResponseMessage, *http.Response, error) {
	return r.ApiService.QueryDeviceStateExecute(r)
}

/*
QueryDeviceState Query Device State

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OperationsAPIQueryDeviceStateRequest
*/
func (a *OperationsAPIService) QueryDeviceState(ctx context.Context) OperationsAPIQueryDeviceStateRequest {
	return OperationsAPIQueryDeviceStateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StateQueryResponseMessage
func (a *OperationsAPIService) QueryDeviceStateExecute(r OperationsAPIQueryDeviceStateRequest) (*StateQueryResponseMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StateQueryResponseMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsAPIService.QueryDeviceState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operations/state"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stateQueryMessage == nil {
		return localVarReturnValue, nil, reportError("stateQueryMessage is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.stateQueryMessage
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["operationsAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OperationsAPIQuerySpecificDeviceStateRequest struct {
	ctx context.Context
	ApiService OperationsAPI
	deviceId string
	stateQueryMessage *StateQueryMessage
}

func (r OperationsAPIQuerySpecificDeviceStateRequest) StateQueryMessage(stateQueryMessage StateQueryMessage) OperationsAPIQuerySpecificDeviceStateRequest {
	r.stateQueryMessage = &stateQueryMessage
	return r
}

func (r OperationsAPIQuerySpecificDeviceStateRequest) Execute() (*StateQueryResponseMessage, *http.Response, error) {
	return r.ApiService.QuerySpecificDeviceStateExecute(r)
}

/*
QuerySpecificDeviceState Query specific device state

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return OperationsAPIQuerySpecificDeviceStateRequest
*/
func (a *OperationsAPIService) QuerySpecificDeviceState(ctx context.Context, deviceId string) OperationsAPIQuerySpecificDeviceStateRequest {
	return OperationsAPIQuerySpecificDeviceStateRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return StateQueryResponseMessage
func (a *OperationsAPIService) QuerySpecificDeviceStateExecute(r OperationsAPIQuerySpecificDeviceStateRequest) (*StateQueryResponseMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StateQueryResponseMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsAPIService.QuerySpecificDeviceState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operations/state/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stateQueryMessage == nil {
		return localVarReturnValue, nil, reportError("stateQueryMessage is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.stateQueryMessage
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["operationsAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OperationsAPIRebootDeviceRequest struct {
	ctx context.Context
	ApiService OperationsAPI
	deviceId string
	rebootCommandMessage *RebootCommandMessage
}

func (r OperationsAPIRebootDeviceRequest) RebootCommandMessage(rebootCommandMessage RebootCommandMessage) OperationsAPIRebootDeviceRequest {
	r.rebootCommandMessage = &rebootCommandMessage
	return r
}

func (r OperationsAPIRebootDeviceRequest) Execute() (*ResponseMessage, *http.Response, error) {
	return r.ApiService.RebootDeviceExecute(r)
}

/*
RebootDevice Reboot device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return OperationsAPIRebootDeviceRequest
*/
func (a *OperationsAPIService) RebootDevice(ctx context.Context, deviceId string) OperationsAPIRebootDeviceRequest {
	return OperationsAPIRebootDeviceRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return ResponseMessage
func (a *OperationsAPIService) RebootDeviceExecute(r OperationsAPIRebootDeviceRequest) (*ResponseMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperationsAPIService.RebootDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/operations/state/{deviceId}/reboot"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rebootCommandMessage == nil {
		return localVarReturnValue, nil, reportError("rebootCommandMessage is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.rebootCommandMessage
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["operationsAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
