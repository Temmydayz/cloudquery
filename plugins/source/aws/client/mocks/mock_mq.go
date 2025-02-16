// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/aws/client (interfaces: MQClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	mq "github.com/aws/aws-sdk-go-v2/service/mq"
	gomock "github.com/golang/mock/gomock"
)

// MockMQClient is a mock of MQClient interface.
type MockMQClient struct {
	ctrl     *gomock.Controller
	recorder *MockMQClientMockRecorder
}

// MockMQClientMockRecorder is the mock recorder for MockMQClient.
type MockMQClientMockRecorder struct {
	mock *MockMQClient
}

// NewMockMQClient creates a new mock instance.
func NewMockMQClient(ctrl *gomock.Controller) *MockMQClient {
	mock := &MockMQClient{ctrl: ctrl}
	mock.recorder = &MockMQClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMQClient) EXPECT() *MockMQClientMockRecorder {
	return m.recorder
}

// DescribeBroker mocks base method.
func (m *MockMQClient) DescribeBroker(arg0 context.Context, arg1 *mq.DescribeBrokerInput, arg2 ...func(*mq.Options)) (*mq.DescribeBrokerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBroker", varargs...)
	ret0, _ := ret[0].(*mq.DescribeBrokerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBroker indicates an expected call of DescribeBroker.
func (mr *MockMQClientMockRecorder) DescribeBroker(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBroker", reflect.TypeOf((*MockMQClient)(nil).DescribeBroker), varargs...)
}

// DescribeConfiguration mocks base method.
func (m *MockMQClient) DescribeConfiguration(arg0 context.Context, arg1 *mq.DescribeConfigurationInput, arg2 ...func(*mq.Options)) (*mq.DescribeConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfiguration", varargs...)
	ret0, _ := ret[0].(*mq.DescribeConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfiguration indicates an expected call of DescribeConfiguration.
func (mr *MockMQClientMockRecorder) DescribeConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfiguration", reflect.TypeOf((*MockMQClient)(nil).DescribeConfiguration), varargs...)
}

// DescribeConfigurationRevision mocks base method.
func (m *MockMQClient) DescribeConfigurationRevision(arg0 context.Context, arg1 *mq.DescribeConfigurationRevisionInput, arg2 ...func(*mq.Options)) (*mq.DescribeConfigurationRevisionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfigurationRevision", varargs...)
	ret0, _ := ret[0].(*mq.DescribeConfigurationRevisionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfigurationRevision indicates an expected call of DescribeConfigurationRevision.
func (mr *MockMQClientMockRecorder) DescribeConfigurationRevision(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfigurationRevision", reflect.TypeOf((*MockMQClient)(nil).DescribeConfigurationRevision), varargs...)
}

// DescribeUser mocks base method.
func (m *MockMQClient) DescribeUser(arg0 context.Context, arg1 *mq.DescribeUserInput, arg2 ...func(*mq.Options)) (*mq.DescribeUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUser", varargs...)
	ret0, _ := ret[0].(*mq.DescribeUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUser indicates an expected call of DescribeUser.
func (mr *MockMQClientMockRecorder) DescribeUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUser", reflect.TypeOf((*MockMQClient)(nil).DescribeUser), varargs...)
}

// ListBrokers mocks base method.
func (m *MockMQClient) ListBrokers(arg0 context.Context, arg1 *mq.ListBrokersInput, arg2 ...func(*mq.Options)) (*mq.ListBrokersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBrokers", varargs...)
	ret0, _ := ret[0].(*mq.ListBrokersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBrokers indicates an expected call of ListBrokers.
func (mr *MockMQClientMockRecorder) ListBrokers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBrokers", reflect.TypeOf((*MockMQClient)(nil).ListBrokers), varargs...)
}

// ListConfigurationRevisions mocks base method.
func (m *MockMQClient) ListConfigurationRevisions(arg0 context.Context, arg1 *mq.ListConfigurationRevisionsInput, arg2 ...func(*mq.Options)) (*mq.ListConfigurationRevisionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConfigurationRevisions", varargs...)
	ret0, _ := ret[0].(*mq.ListConfigurationRevisionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurationRevisions indicates an expected call of ListConfigurationRevisions.
func (mr *MockMQClientMockRecorder) ListConfigurationRevisions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurationRevisions", reflect.TypeOf((*MockMQClient)(nil).ListConfigurationRevisions), varargs...)
}
