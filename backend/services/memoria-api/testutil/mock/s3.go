// Code generated by MockGen. DO NOT EDIT.
// Source: domain/interfaces/s3.go
//
// Generated by this command:
//
//	mockgen -source=domain/interfaces/s3.go -destination=testutil/mock/s3.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	interfaces "memoria-api/domain/interfaces"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockS3Client is a mock of S3Client interface.
type MockS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientMockRecorder
}

// MockS3ClientMockRecorder is the mock recorder for MockS3Client.
type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

// NewMockS3Client creates a new mock instance.
func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

// DeleteFolder mocks base method.
func (m *MockS3Client) DeleteFolder(dto interfaces.S3ClientDeleteFolderDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFolder", dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFolder indicates an expected call of DeleteFolder.
func (mr *MockS3ClientMockRecorder) DeleteFolder(dto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFolder", reflect.TypeOf((*MockS3Client)(nil).DeleteFolder), dto)
}

// GetPresignedPutObjectURL mocks base method.
func (m *MockS3Client) GetPresignedPutObjectURL(dto interfaces.S3ClientGetPresignedPutObjectURLDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPresignedPutObjectURL", dto)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPresignedPutObjectURL indicates an expected call of GetPresignedPutObjectURL.
func (mr *MockS3ClientMockRecorder) GetPresignedPutObjectURL(dto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPresignedPutObjectURL", reflect.TypeOf((*MockS3Client)(nil).GetPresignedPutObjectURL), dto)
}

// PutObject mocks base method.
func (m *MockS3Client) PutObject(dto interfaces.S3ClientPutObjectDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", dto)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObject indicates an expected call of PutObject.
func (mr *MockS3ClientMockRecorder) PutObject(dto any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3Client)(nil).PutObject), dto)
}
