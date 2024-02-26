// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"

	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// AsyncAgentService_ExecuteTaskSyncClient is an autogenerated mock type for the AsyncAgentService_ExecuteTaskSyncClient type
type AsyncAgentService_ExecuteTaskSyncClient struct {
	mock.Mock
}

type AsyncAgentService_ExecuteTaskSyncClient_CloseSend struct {
	*mock.Call
}

func (_m AsyncAgentService_ExecuteTaskSyncClient_CloseSend) Return(_a0 error) *AsyncAgentService_ExecuteTaskSyncClient_CloseSend {
	return &AsyncAgentService_ExecuteTaskSyncClient_CloseSend{Call: _m.Call.Return(_a0)}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnCloseSend() *AsyncAgentService_ExecuteTaskSyncClient_CloseSend {
	c_call := _m.On("CloseSend")
	return &AsyncAgentService_ExecuteTaskSyncClient_CloseSend{Call: c_call}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnCloseSendMatch(matchers ...interface{}) *AsyncAgentService_ExecuteTaskSyncClient_CloseSend {
	c_call := _m.On("CloseSend", matchers...)
	return &AsyncAgentService_ExecuteTaskSyncClient_CloseSend{Call: c_call}
}

// CloseSend provides a mock function with given fields:
func (_m *AsyncAgentService_ExecuteTaskSyncClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type AsyncAgentService_ExecuteTaskSyncClient_Context struct {
	*mock.Call
}

func (_m AsyncAgentService_ExecuteTaskSyncClient_Context) Return(_a0 context.Context) *AsyncAgentService_ExecuteTaskSyncClient_Context {
	return &AsyncAgentService_ExecuteTaskSyncClient_Context{Call: _m.Call.Return(_a0)}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnContext() *AsyncAgentService_ExecuteTaskSyncClient_Context {
	c_call := _m.On("Context")
	return &AsyncAgentService_ExecuteTaskSyncClient_Context{Call: c_call}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnContextMatch(matchers ...interface{}) *AsyncAgentService_ExecuteTaskSyncClient_Context {
	c_call := _m.On("Context", matchers...)
	return &AsyncAgentService_ExecuteTaskSyncClient_Context{Call: c_call}
}

// Context provides a mock function with given fields:
func (_m *AsyncAgentService_ExecuteTaskSyncClient) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

type AsyncAgentService_ExecuteTaskSyncClient_Header struct {
	*mock.Call
}

func (_m AsyncAgentService_ExecuteTaskSyncClient_Header) Return(_a0 metadata.MD, _a1 error) *AsyncAgentService_ExecuteTaskSyncClient_Header {
	return &AsyncAgentService_ExecuteTaskSyncClient_Header{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnHeader() *AsyncAgentService_ExecuteTaskSyncClient_Header {
	c_call := _m.On("Header")
	return &AsyncAgentService_ExecuteTaskSyncClient_Header{Call: c_call}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnHeaderMatch(matchers ...interface{}) *AsyncAgentService_ExecuteTaskSyncClient_Header {
	c_call := _m.On("Header", matchers...)
	return &AsyncAgentService_ExecuteTaskSyncClient_Header{Call: c_call}
}

// Header provides a mock function with given fields:
func (_m *AsyncAgentService_ExecuteTaskSyncClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AsyncAgentService_ExecuteTaskSyncClient_Recv struct {
	*mock.Call
}

func (_m AsyncAgentService_ExecuteTaskSyncClient_Recv) Return(_a0 *admin.ExecuteTaskSyncResponse, _a1 error) *AsyncAgentService_ExecuteTaskSyncClient_Recv {
	return &AsyncAgentService_ExecuteTaskSyncClient_Recv{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnRecv() *AsyncAgentService_ExecuteTaskSyncClient_Recv {
	c_call := _m.On("Recv")
	return &AsyncAgentService_ExecuteTaskSyncClient_Recv{Call: c_call}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnRecvMatch(matchers ...interface{}) *AsyncAgentService_ExecuteTaskSyncClient_Recv {
	c_call := _m.On("Recv", matchers...)
	return &AsyncAgentService_ExecuteTaskSyncClient_Recv{Call: c_call}
}

// Recv provides a mock function with given fields:
func (_m *AsyncAgentService_ExecuteTaskSyncClient) Recv() (*admin.ExecuteTaskSyncResponse, error) {
	ret := _m.Called()

	var r0 *admin.ExecuteTaskSyncResponse
	if rf, ok := ret.Get(0).(func() *admin.ExecuteTaskSyncResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ExecuteTaskSyncResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AsyncAgentService_ExecuteTaskSyncClient_RecvMsg struct {
	*mock.Call
}

func (_m AsyncAgentService_ExecuteTaskSyncClient_RecvMsg) Return(_a0 error) *AsyncAgentService_ExecuteTaskSyncClient_RecvMsg {
	return &AsyncAgentService_ExecuteTaskSyncClient_RecvMsg{Call: _m.Call.Return(_a0)}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnRecvMsg(m interface{}) *AsyncAgentService_ExecuteTaskSyncClient_RecvMsg {
	c_call := _m.On("RecvMsg", m)
	return &AsyncAgentService_ExecuteTaskSyncClient_RecvMsg{Call: c_call}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnRecvMsgMatch(matchers ...interface{}) *AsyncAgentService_ExecuteTaskSyncClient_RecvMsg {
	c_call := _m.On("RecvMsg", matchers...)
	return &AsyncAgentService_ExecuteTaskSyncClient_RecvMsg{Call: c_call}
}

// RecvMsg provides a mock function with given fields: m
func (_m *AsyncAgentService_ExecuteTaskSyncClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type AsyncAgentService_ExecuteTaskSyncClient_Send struct {
	*mock.Call
}

func (_m AsyncAgentService_ExecuteTaskSyncClient_Send) Return(_a0 error) *AsyncAgentService_ExecuteTaskSyncClient_Send {
	return &AsyncAgentService_ExecuteTaskSyncClient_Send{Call: _m.Call.Return(_a0)}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnSend(_a0 *admin.ExecuteTaskSyncRequest) *AsyncAgentService_ExecuteTaskSyncClient_Send {
	c_call := _m.On("Send", _a0)
	return &AsyncAgentService_ExecuteTaskSyncClient_Send{Call: c_call}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnSendMatch(matchers ...interface{}) *AsyncAgentService_ExecuteTaskSyncClient_Send {
	c_call := _m.On("Send", matchers...)
	return &AsyncAgentService_ExecuteTaskSyncClient_Send{Call: c_call}
}

// Send provides a mock function with given fields: _a0
func (_m *AsyncAgentService_ExecuteTaskSyncClient) Send(_a0 *admin.ExecuteTaskSyncRequest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*admin.ExecuteTaskSyncRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type AsyncAgentService_ExecuteTaskSyncClient_SendMsg struct {
	*mock.Call
}

func (_m AsyncAgentService_ExecuteTaskSyncClient_SendMsg) Return(_a0 error) *AsyncAgentService_ExecuteTaskSyncClient_SendMsg {
	return &AsyncAgentService_ExecuteTaskSyncClient_SendMsg{Call: _m.Call.Return(_a0)}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnSendMsg(m interface{}) *AsyncAgentService_ExecuteTaskSyncClient_SendMsg {
	c_call := _m.On("SendMsg", m)
	return &AsyncAgentService_ExecuteTaskSyncClient_SendMsg{Call: c_call}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnSendMsgMatch(matchers ...interface{}) *AsyncAgentService_ExecuteTaskSyncClient_SendMsg {
	c_call := _m.On("SendMsg", matchers...)
	return &AsyncAgentService_ExecuteTaskSyncClient_SendMsg{Call: c_call}
}

// SendMsg provides a mock function with given fields: m
func (_m *AsyncAgentService_ExecuteTaskSyncClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type AsyncAgentService_ExecuteTaskSyncClient_Trailer struct {
	*mock.Call
}

func (_m AsyncAgentService_ExecuteTaskSyncClient_Trailer) Return(_a0 metadata.MD) *AsyncAgentService_ExecuteTaskSyncClient_Trailer {
	return &AsyncAgentService_ExecuteTaskSyncClient_Trailer{Call: _m.Call.Return(_a0)}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnTrailer() *AsyncAgentService_ExecuteTaskSyncClient_Trailer {
	c_call := _m.On("Trailer")
	return &AsyncAgentService_ExecuteTaskSyncClient_Trailer{Call: c_call}
}

func (_m *AsyncAgentService_ExecuteTaskSyncClient) OnTrailerMatch(matchers ...interface{}) *AsyncAgentService_ExecuteTaskSyncClient_Trailer {
	c_call := _m.On("Trailer", matchers...)
	return &AsyncAgentService_ExecuteTaskSyncClient_Trailer{Call: c_call}
}

// Trailer provides a mock function with given fields:
func (_m *AsyncAgentService_ExecuteTaskSyncClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}