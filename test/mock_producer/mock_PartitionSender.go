// Code generated by MockGen. DO NOT EDIT.
// Source: producer/PartitionSender.go

// Package mock_producer is a generated GoMock package.
package mock_producer

import (
	reflect "reflect"

	producer "github.com/XiaoMi/talos-sdk-golang/producer"
	gomock "github.com/golang/mock/gomock"
)

// MockSender is a mock of Sender interface.
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender.
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance.
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// AddMessage mocks base method.
func (m *MockSender) AddMessage(userMessageList []*producer.UserMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddMessage", userMessageList)
}

// AddMessage indicates an expected call of AddMessage.
func (mr *MockSenderMockRecorder) AddMessage(userMessageList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessage", reflect.TypeOf((*MockSender)(nil).AddMessage), userMessageList)
}

// MessageCallbackTask mocks base method.
func (m *MockSender) MessageCallbackTask(userMessageResult *producer.UserMessageResult) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MessageCallbackTask", userMessageResult)
}

// MessageCallbackTask indicates an expected call of MessageCallbackTask.
func (mr *MockSenderMockRecorder) MessageCallbackTask(userMessageResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageCallbackTask", reflect.TypeOf((*MockSender)(nil).MessageCallbackTask), userMessageResult)
}

// MessageWriterTask mocks base method.
func (m *MockSender) MessageWriterTask() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MessageWriterTask")
}

// MessageWriterTask indicates an expected call of MessageWriterTask.
func (mr *MockSenderMockRecorder) MessageWriterTask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageWriterTask", reflect.TypeOf((*MockSender)(nil).MessageWriterTask))
}

// Shutdown mocks base method.
func (m *MockSender) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockSenderMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockSender)(nil).Shutdown))
}
