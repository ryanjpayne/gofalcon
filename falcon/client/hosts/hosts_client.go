// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hosts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hosts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDeviceDetails(params *GetDeviceDetailsParams, opts ...ClientOption) (*GetDeviceDetailsOK, error)

	GetDeviceDetailsV2(params *GetDeviceDetailsV2Params, opts ...ClientOption) (*GetDeviceDetailsV2OK, error)

	GetOnlineStateV1(params *GetOnlineStateV1Params, opts ...ClientOption) (*GetOnlineStateV1OK, error)

	PerformActionV2(params *PerformActionV2Params, opts ...ClientOption) (*PerformActionV2Accepted, error)

	PostDeviceDetailsV2(params *PostDeviceDetailsV2Params, opts ...ClientOption) (*PostDeviceDetailsV2OK, error)

	QueryDeviceLoginHistory(params *QueryDeviceLoginHistoryParams, opts ...ClientOption) (*QueryDeviceLoginHistoryOK, error)

	QueryDevicesByFilter(params *QueryDevicesByFilterParams, opts ...ClientOption) (*QueryDevicesByFilterOK, error)

	QueryDevicesByFilterScroll(params *QueryDevicesByFilterScrollParams, opts ...ClientOption) (*QueryDevicesByFilterScrollOK, error)

	QueryGetNetworkAddressHistoryV1(params *QueryGetNetworkAddressHistoryV1Params, opts ...ClientOption) (*QueryGetNetworkAddressHistoryV1OK, error)

	QueryHiddenDevices(params *QueryHiddenDevicesParams, opts ...ClientOption) (*QueryHiddenDevicesOK, error)

	UpdateDeviceTags(params *UpdateDeviceTagsParams, opts ...ClientOption) (*UpdateDeviceTagsOK, *UpdateDeviceTagsAccepted, error)

	EntitiesPerformAction(params *EntitiesPerformActionParams, opts ...ClientOption) (*EntitiesPerformActionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDeviceDetails deprecateds please use new methods get device details v2 or post device details v2 this method now redirects to get device details v2 the original API endpoint will be removed on or sometime after february 9 2023
*/
func (a *Client) GetDeviceDetails(params *GetDeviceDetailsParams, opts ...ClientOption) (*GetDeviceDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceDetails",
		Method:             "GET",
		PathPattern:        "/devices/entities/devices//v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceDetailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceDetailsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDeviceDetailsV2 gets details on one or more hosts by providing host i ds as a query parameter supports up to a maximum 100 i ds
*/
func (a *Client) GetDeviceDetailsV2(params *GetDeviceDetailsV2Params, opts ...ClientOption) (*GetDeviceDetailsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceDetailsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeviceDetailsV2",
		Method:             "GET",
		PathPattern:        "/devices/entities/devices/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceDetailsV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceDetailsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceDetailsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetOnlineStateV1 gets the online status for one or more hosts by specifying each host s unique ID successful requests return an HTTP 200 response and the status for each host identified by a state of online offline or unknown for each host identified by host id make a g e t request to devices queries devices v1 to get a list of host i ds
*/
func (a *Client) GetOnlineStateV1(params *GetOnlineStateV1Params, opts ...ClientOption) (*GetOnlineStateV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnlineStateV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOnlineState.V1",
		Method:             "GET",
		PathPattern:        "/devices/entities/online-state/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOnlineStateV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOnlineStateV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOnlineStateV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PerformActionV2 takes various actions on the hosts in your environment contain or lift containment on a host delete or restore a host
*/
func (a *Client) PerformActionV2(params *PerformActionV2Params, opts ...ClientOption) (*PerformActionV2Accepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformActionV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PerformActionV2",
		Method:             "POST",
		PathPattern:        "/devices/entities/devices-actions/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformActionV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PerformActionV2Accepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PerformActionV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostDeviceDetailsV2 gets details on one or more hosts by providing host i ds in a p o s t body supports up to a maximum 5000 i ds
*/
func (a *Client) PostDeviceDetailsV2(params *PostDeviceDetailsV2Params, opts ...ClientOption) (*PostDeviceDetailsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDeviceDetailsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostDeviceDetailsV2",
		Method:             "POST",
		PathPattern:        "/devices/entities/devices/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDeviceDetailsV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDeviceDetailsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostDeviceDetailsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDeviceLoginHistory retrieves details about recent login sessions for a set of devices
*/
func (a *Client) QueryDeviceLoginHistory(params *QueryDeviceLoginHistoryParams, opts ...ClientOption) (*QueryDeviceLoginHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDeviceLoginHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryDeviceLoginHistory",
		Method:             "POST",
		PathPattern:        "/devices/combined/devices/login-history/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDeviceLoginHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDeviceLoginHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDeviceLoginHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDevicesByFilter searches for hosts in your environment by platform hostname IP and other criteria
*/
func (a *Client) QueryDevicesByFilter(params *QueryDevicesByFilterParams, opts ...ClientOption) (*QueryDevicesByFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDevicesByFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryDevicesByFilter",
		Method:             "GET",
		PathPattern:        "/devices/queries/devices/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDevicesByFilterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDevicesByFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDevicesByFilterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryDevicesByFilterScroll searches for hosts in your environment by platform hostname IP and other criteria with continuous pagination capability based on offset pointer which expires after 2 minutes with no maximum limit
*/
func (a *Client) QueryDevicesByFilterScroll(params *QueryDevicesByFilterScrollParams, opts ...ClientOption) (*QueryDevicesByFilterScrollOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDevicesByFilterScrollParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryDevicesByFilterScroll",
		Method:             "GET",
		PathPattern:        "/devices/queries/devices-scroll/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDevicesByFilterScrollReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryDevicesByFilterScrollOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryDevicesByFilterScrollDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryGetNetworkAddressHistoryV1 retrieves history of IP and m a c addresses of devices
*/
func (a *Client) QueryGetNetworkAddressHistoryV1(params *QueryGetNetworkAddressHistoryV1Params, opts ...ClientOption) (*QueryGetNetworkAddressHistoryV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryGetNetworkAddressHistoryV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryGetNetworkAddressHistoryV1",
		Method:             "POST",
		PathPattern:        "/devices/combined/devices/network-address-history/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryGetNetworkAddressHistoryV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryGetNetworkAddressHistoryV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryGetNetworkAddressHistoryV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryHiddenDevices retrieves hidden hosts that match the provided filter criteria
*/
func (a *Client) QueryHiddenDevices(params *QueryHiddenDevicesParams, opts ...ClientOption) (*QueryHiddenDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryHiddenDevicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryHiddenDevices",
		Method:             "GET",
		PathPattern:        "/devices/queries/devices-hidden/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryHiddenDevicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryHiddenDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryHiddenDevicesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateDeviceTags appends or remove one or more falcon grouping tags on one or more hosts tags must be of the form falcon grouping tags
*/
func (a *Client) UpdateDeviceTags(params *UpdateDeviceTagsParams, opts ...ClientOption) (*UpdateDeviceTagsOK, *UpdateDeviceTagsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeviceTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateDeviceTags",
		Method:             "PATCH",
		PathPattern:        "/devices/entities/devices/tags/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeviceTagsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateDeviceTagsOK:
		return value, nil, nil
	case *UpdateDeviceTagsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for hosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EntitiesPerformAction performs the specified action on the provided prevention policy i ds
*/
func (a *Client) EntitiesPerformAction(params *EntitiesPerformActionParams, opts ...ClientOption) (*EntitiesPerformActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEntitiesPerformActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "entities.perform_action",
		Method:             "POST",
		PathPattern:        "/devices/entities/group-actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EntitiesPerformActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EntitiesPerformActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EntitiesPerformActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
