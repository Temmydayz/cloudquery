// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: SearchServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	search "github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
)

// MockSearchServiceClient is a mock of SearchServiceClient interface.
type MockSearchServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockSearchServiceClientMockRecorder
}

// MockSearchServiceClientMockRecorder is the mock recorder for MockSearchServiceClient.
type MockSearchServiceClientMockRecorder struct {
	mock *MockSearchServiceClient
}

// NewMockSearchServiceClient creates a new mock instance.
func NewMockSearchServiceClient(ctrl *gomock.Controller) *MockSearchServiceClient {
	mock := &MockSearchServiceClient{ctrl: ctrl}
	mock.recorder = &MockSearchServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchServiceClient) EXPECT() *MockSearchServiceClientMockRecorder {
	return m.recorder
}

// ListBySubscription mocks base method.
func (m *MockSearchServiceClient) ListBySubscription(arg0 context.Context, arg1 *uuid.UUID) (search.ServiceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(search.ServiceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySubscription indicates an expected call of ListBySubscription.
func (mr *MockSearchServiceClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockSearchServiceClient)(nil).ListBySubscription), arg0, arg1)
}
