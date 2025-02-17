// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	types "avito_test_assingment/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CheckAuthData mocks base method.
func (m *MockAuthorization) CheckAuthData(username, password string) (types.UserType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuthData", username, password)
	ret0, _ := ret[0].(types.UserType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAuthData indicates an expected call of CheckAuthData.
func (mr *MockAuthorizationMockRecorder) CheckAuthData(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuthData", reflect.TypeOf((*MockAuthorization)(nil).CheckAuthData), username, password)
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user types.UserType) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(user types.UserType) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), user)
}

// ParserToken mocks base method.
func (m *MockAuthorization) ParserToken(accessToken string) (*types.TokenClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParserToken", accessToken)
	ret0, _ := ret[0].(*types.TokenClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParserToken indicates an expected call of ParserToken.
func (mr *MockAuthorizationMockRecorder) ParserToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParserToken", reflect.TypeOf((*MockAuthorization)(nil).ParserToken), accessToken)
}

// MockBanner is a mock of Banner interface.
type MockBanner struct {
	ctrl     *gomock.Controller
	recorder *MockBannerMockRecorder
}

// MockBannerMockRecorder is the mock recorder for MockBanner.
type MockBannerMockRecorder struct {
	mock *MockBanner
}

// NewMockBanner creates a new mock instance.
func NewMockBanner(ctrl *gomock.Controller) *MockBanner {
	mock := &MockBanner{ctrl: ctrl}
	mock.recorder = &MockBannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBanner) EXPECT() *MockBannerMockRecorder {
	return m.recorder
}

// BannerGet mocks base method.
func (m *MockBanner) BannerGet(featureId int, tagId []int, limit, offset int) ([]types.BannerGet200ResponseInner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BannerGet", featureId, tagId, limit, offset)
	ret0, _ := ret[0].([]types.BannerGet200ResponseInner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BannerGet indicates an expected call of BannerGet.
func (mr *MockBannerMockRecorder) BannerGet(featureId, tagId, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BannerGet", reflect.TypeOf((*MockBanner)(nil).BannerGet), featureId, tagId, limit, offset)
}

// BannerIdDelete mocks base method.
func (m *MockBanner) BannerIdDelete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BannerIdDelete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// BannerIdDelete indicates an expected call of BannerIdDelete.
func (mr *MockBannerMockRecorder) BannerIdDelete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BannerIdDelete", reflect.TypeOf((*MockBanner)(nil).BannerIdDelete), id)
}

// BannerIdPatch mocks base method.
func (m *MockBanner) BannerIdPatch(id int, data types.BannerIdPatchRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BannerIdPatch", id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// BannerIdPatch indicates an expected call of BannerIdPatch.
func (mr *MockBannerMockRecorder) BannerIdPatch(id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BannerIdPatch", reflect.TypeOf((*MockBanner)(nil).BannerIdPatch), id, data)
}

// BannerPost mocks base method.
func (m *MockBanner) BannerPost(data types.BannerPostRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BannerPost", data)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BannerPost indicates an expected call of BannerPost.
func (mr *MockBannerMockRecorder) BannerPost(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BannerPost", reflect.TypeOf((*MockBanner)(nil).BannerPost), data)
}

// UserBannerGet mocks base method.
func (m *MockBanner) UserBannerGet(tagId []int, featureId int, useLastRevision bool) (types.BannerGet200ResponseInner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserBannerGet", tagId, featureId, useLastRevision)
	ret0, _ := ret[0].(types.BannerGet200ResponseInner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserBannerGet indicates an expected call of UserBannerGet.
func (mr *MockBannerMockRecorder) UserBannerGet(tagId, featureId, useLastRevision interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserBannerGet", reflect.TypeOf((*MockBanner)(nil).UserBannerGet), tagId, featureId, useLastRevision)
}
