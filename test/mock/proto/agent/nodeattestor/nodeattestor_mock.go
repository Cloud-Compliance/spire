// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/agent/nodeattestor (interfaces: NodeAttestor,Plugin,FetchAttestationData_Stream)

// Package mock_nodeattestor is a generated GoMock package.
package mock_nodeattestor

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	nodeattestor "github.com/spiffe/spire/proto/agent/nodeattestor"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	reflect "reflect"
)

// MockNodeAttestor is a mock of NodeAttestor interface
type MockNodeAttestor struct {
	ctrl     *gomock.Controller
	recorder *MockNodeAttestorMockRecorder
}

// MockNodeAttestorMockRecorder is the mock recorder for MockNodeAttestor
type MockNodeAttestorMockRecorder struct {
	mock *MockNodeAttestor
}

// NewMockNodeAttestor creates a new mock instance
func NewMockNodeAttestor(ctrl *gomock.Controller) *MockNodeAttestor {
	mock := &MockNodeAttestor{ctrl: ctrl}
	mock.recorder = &MockNodeAttestorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeAttestor) EXPECT() *MockNodeAttestorMockRecorder {
	return m.recorder
}

// FetchAttestationData mocks base method
func (m *MockNodeAttestor) FetchAttestationData(arg0 context.Context) (nodeattestor.FetchAttestationData_Stream, error) {
	ret := m.ctrl.Call(m, "FetchAttestationData", arg0)
	ret0, _ := ret[0].(nodeattestor.FetchAttestationData_Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAttestationData indicates an expected call of FetchAttestationData
func (mr *MockNodeAttestorMockRecorder) FetchAttestationData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttestationData", reflect.TypeOf((*MockNodeAttestor)(nil).FetchAttestationData), arg0)
}

// MockPlugin is a mock of Plugin interface
type MockPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockPluginMockRecorder
}

// MockPluginMockRecorder is the mock recorder for MockPlugin
type MockPluginMockRecorder struct {
	mock *MockPlugin
}

// NewMockPlugin creates a new mock instance
func NewMockPlugin(ctrl *gomock.Controller) *MockPlugin {
	mock := &MockPlugin{ctrl: ctrl}
	mock.recorder = &MockPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlugin) EXPECT() *MockPluginMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockPlugin) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockPluginMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockPlugin)(nil).Configure), arg0, arg1)
}

// FetchAttestationData mocks base method
func (m *MockPlugin) FetchAttestationData(arg0 nodeattestor.FetchAttestationData_PluginStream) error {
	ret := m.ctrl.Call(m, "FetchAttestationData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchAttestationData indicates an expected call of FetchAttestationData
func (mr *MockPluginMockRecorder) FetchAttestationData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAttestationData", reflect.TypeOf((*MockPlugin)(nil).FetchAttestationData), arg0)
}

// GetPluginInfo mocks base method
func (m *MockPlugin) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockPluginMockRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockPlugin)(nil).GetPluginInfo), arg0, arg1)
}

// MockFetchAttestationData_Stream is a mock of FetchAttestationData_Stream interface
type MockFetchAttestationData_Stream struct {
	ctrl     *gomock.Controller
	recorder *MockFetchAttestationData_StreamMockRecorder
}

// MockFetchAttestationData_StreamMockRecorder is the mock recorder for MockFetchAttestationData_Stream
type MockFetchAttestationData_StreamMockRecorder struct {
	mock *MockFetchAttestationData_Stream
}

// NewMockFetchAttestationData_Stream creates a new mock instance
func NewMockFetchAttestationData_Stream(ctrl *gomock.Controller) *MockFetchAttestationData_Stream {
	mock := &MockFetchAttestationData_Stream{ctrl: ctrl}
	mock.recorder = &MockFetchAttestationData_StreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFetchAttestationData_Stream) EXPECT() *MockFetchAttestationData_StreamMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockFetchAttestationData_Stream) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockFetchAttestationData_StreamMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockFetchAttestationData_Stream)(nil).CloseSend))
}

// Context mocks base method
func (m *MockFetchAttestationData_Stream) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockFetchAttestationData_StreamMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockFetchAttestationData_Stream)(nil).Context))
}

// Recv mocks base method
func (m *MockFetchAttestationData_Stream) Recv() (*nodeattestor.FetchAttestationDataResponse, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*nodeattestor.FetchAttestationDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockFetchAttestationData_StreamMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockFetchAttestationData_Stream)(nil).Recv))
}

// Send mocks base method
func (m *MockFetchAttestationData_Stream) Send(arg0 *nodeattestor.FetchAttestationDataRequest) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockFetchAttestationData_StreamMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockFetchAttestationData_Stream)(nil).Send), arg0)
}
