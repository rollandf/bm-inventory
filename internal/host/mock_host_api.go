// Code generated by MockGen. DO NOT EDIT.
// Source: host.go

// Package host is a generated GoMock package.
package host

import (
	context "context"
	models "github.com/filanov/bm-inventory/models"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
	reflect "reflect"
)

// MockStateAPI is a mock of StateAPI interface.
type MockStateAPI struct {
	ctrl     *gomock.Controller
	recorder *MockStateAPIMockRecorder
}

// MockStateAPIMockRecorder is the mock recorder for MockStateAPI.
type MockStateAPIMockRecorder struct {
	mock *MockStateAPI
}

// NewMockStateAPI creates a new mock instance.
func NewMockStateAPI(ctrl *gomock.Controller) *MockStateAPI {
	mock := &MockStateAPI{ctrl: ctrl}
	mock.recorder = &MockStateAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateAPI) EXPECT() *MockStateAPIMockRecorder {
	return m.recorder
}

// UpdateHwInfo mocks base method.
func (m *MockStateAPI) UpdateHwInfo(ctx context.Context, h *models.Host, hwInfo string) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHwInfo", ctx, h, hwInfo)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHwInfo indicates an expected call of UpdateHwInfo.
func (mr *MockStateAPIMockRecorder) UpdateHwInfo(ctx, h, hwInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHwInfo", reflect.TypeOf((*MockStateAPI)(nil).UpdateHwInfo), ctx, h, hwInfo)
}

// UpdateInventory mocks base method.
func (m *MockStateAPI) UpdateInventory(ctx context.Context, h *models.Host, inventory string) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInventory", ctx, h, inventory)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInventory indicates an expected call of UpdateInventory.
func (mr *MockStateAPIMockRecorder) UpdateInventory(ctx, h, inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockStateAPI)(nil).UpdateInventory), ctx, h, inventory)
}

// UpdateRole mocks base method.
func (m *MockStateAPI) UpdateRole(ctx context.Context, h *models.Host, role string, db *gorm.DB) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", ctx, h, role, db)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockStateAPIMockRecorder) UpdateRole(ctx, h, role, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockStateAPI)(nil).UpdateRole), ctx, h, role, db)
}

// RefreshStatus mocks base method.
func (m *MockStateAPI) RefreshStatus(ctx context.Context, h *models.Host) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshStatus", ctx, h)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshStatus indicates an expected call of RefreshStatus.
func (mr *MockStateAPIMockRecorder) RefreshStatus(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshStatus", reflect.TypeOf((*MockStateAPI)(nil).RefreshStatus), ctx, h)
}

// RefreshState mocks base method.
func (m *MockStateAPI) RefreshState(ctx context.Context, h *models.Host, db *gorm.DB) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshState", ctx, h, db)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshState indicates an expected call of RefreshState.
func (mr *MockStateAPIMockRecorder) RefreshState(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshState", reflect.TypeOf((*MockStateAPI)(nil).RefreshState), ctx, h, db)
}

// Install mocks base method.
func (m *MockStateAPI) Install(ctx context.Context, h *models.Host, db *gorm.DB) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ctx, h, db)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockStateAPIMockRecorder) Install(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockStateAPI)(nil).Install), ctx, h, db)
}

// EnableHost mocks base method.
func (m *MockStateAPI) EnableHost(ctx context.Context, h *models.Host) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableHost", ctx, h)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableHost indicates an expected call of EnableHost.
func (mr *MockStateAPIMockRecorder) EnableHost(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableHost", reflect.TypeOf((*MockStateAPI)(nil).EnableHost), ctx, h)
}

// DisableHost mocks base method.
func (m *MockStateAPI) DisableHost(ctx context.Context, h *models.Host) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableHost", ctx, h)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableHost indicates an expected call of DisableHost.
func (mr *MockStateAPIMockRecorder) DisableHost(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableHost", reflect.TypeOf((*MockStateAPI)(nil).DisableHost), ctx, h)
}

// MockSpecificHardwareParams is a mock of SpecificHardwareParams interface.
type MockSpecificHardwareParams struct {
	ctrl     *gomock.Controller
	recorder *MockSpecificHardwareParamsMockRecorder
}

// MockSpecificHardwareParamsMockRecorder is the mock recorder for MockSpecificHardwareParams.
type MockSpecificHardwareParamsMockRecorder struct {
	mock *MockSpecificHardwareParams
}

// NewMockSpecificHardwareParams creates a new mock instance.
func NewMockSpecificHardwareParams(ctrl *gomock.Controller) *MockSpecificHardwareParams {
	mock := &MockSpecificHardwareParams{ctrl: ctrl}
	mock.recorder = &MockSpecificHardwareParamsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpecificHardwareParams) EXPECT() *MockSpecificHardwareParamsMockRecorder {
	return m.recorder
}

// GetHostValidDisks mocks base method.
func (m *MockSpecificHardwareParams) GetHostValidDisks(h *models.Host) ([]*models.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostValidDisks", h)
	ret0, _ := ret[0].([]*models.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostValidDisks indicates an expected call of GetHostValidDisks.
func (mr *MockSpecificHardwareParamsMockRecorder) GetHostValidDisks(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostValidDisks", reflect.TypeOf((*MockSpecificHardwareParams)(nil).GetHostValidDisks), h)
}

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// RegisterHost mocks base method.
func (m *MockAPI) RegisterHost(ctx context.Context, h *models.Host) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterHost", ctx, h)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterHost indicates an expected call of RegisterHost.
func (mr *MockAPIMockRecorder) RegisterHost(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHost", reflect.TypeOf((*MockAPI)(nil).RegisterHost), ctx, h)
}

// UpdateHwInfo mocks base method.
func (m *MockAPI) UpdateHwInfo(ctx context.Context, h *models.Host, hwInfo string) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHwInfo", ctx, h, hwInfo)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHwInfo indicates an expected call of UpdateHwInfo.
func (mr *MockAPIMockRecorder) UpdateHwInfo(ctx, h, hwInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHwInfo", reflect.TypeOf((*MockAPI)(nil).UpdateHwInfo), ctx, h, hwInfo)
}

// UpdateInventory mocks base method.
func (m *MockAPI) UpdateInventory(ctx context.Context, h *models.Host, inventory string) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInventory", ctx, h, inventory)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInventory indicates an expected call of UpdateInventory.
func (mr *MockAPIMockRecorder) UpdateInventory(ctx, h, inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockAPI)(nil).UpdateInventory), ctx, h, inventory)
}

// UpdateRole mocks base method.
func (m *MockAPI) UpdateRole(ctx context.Context, h *models.Host, role string, db *gorm.DB) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", ctx, h, role, db)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockAPIMockRecorder) UpdateRole(ctx, h, role, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockAPI)(nil).UpdateRole), ctx, h, role, db)
}

// RefreshStatus mocks base method.
func (m *MockAPI) RefreshStatus(ctx context.Context, h *models.Host) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshStatus", ctx, h)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshStatus indicates an expected call of RefreshStatus.
func (mr *MockAPIMockRecorder) RefreshStatus(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshStatus", reflect.TypeOf((*MockAPI)(nil).RefreshStatus), ctx, h)
}

// RefreshState mocks base method.
func (m *MockAPI) RefreshState(ctx context.Context, h *models.Host, db *gorm.DB) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshState", ctx, h, db)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshState indicates an expected call of RefreshState.
func (mr *MockAPIMockRecorder) RefreshState(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshState", reflect.TypeOf((*MockAPI)(nil).RefreshState), ctx, h, db)
}

// Install mocks base method.
func (m *MockAPI) Install(ctx context.Context, h *models.Host, db *gorm.DB) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ctx, h, db)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockAPIMockRecorder) Install(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockAPI)(nil).Install), ctx, h, db)
}

// EnableHost mocks base method.
func (m *MockAPI) EnableHost(ctx context.Context, h *models.Host) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableHost", ctx, h)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableHost indicates an expected call of EnableHost.
func (mr *MockAPIMockRecorder) EnableHost(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableHost", reflect.TypeOf((*MockAPI)(nil).EnableHost), ctx, h)
}

// DisableHost mocks base method.
func (m *MockAPI) DisableHost(ctx context.Context, h *models.Host) (*UpdateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableHost", ctx, h)
	ret0, _ := ret[0].(*UpdateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableHost indicates an expected call of DisableHost.
func (mr *MockAPIMockRecorder) DisableHost(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableHost", reflect.TypeOf((*MockAPI)(nil).DisableHost), ctx, h)
}

// GetNextSteps mocks base method.
func (m *MockAPI) GetNextSteps(ctx context.Context, host *models.Host) (models.Steps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSteps", ctx, host)
	ret0, _ := ret[0].(models.Steps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextSteps indicates an expected call of GetNextSteps.
func (mr *MockAPIMockRecorder) GetNextSteps(ctx, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSteps", reflect.TypeOf((*MockAPI)(nil).GetNextSteps), ctx, host)
}

// GetHostValidDisks mocks base method.
func (m *MockAPI) GetHostValidDisks(h *models.Host) ([]*models.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostValidDisks", h)
	ret0, _ := ret[0].([]*models.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostValidDisks indicates an expected call of GetHostValidDisks.
func (mr *MockAPIMockRecorder) GetHostValidDisks(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostValidDisks", reflect.TypeOf((*MockAPI)(nil).GetHostValidDisks), h)
}

// UpdateInstallProgress mocks base method.
func (m *MockAPI) UpdateInstallProgress(ctx context.Context, h *models.Host, progress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstallProgress", ctx, h, progress)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstallProgress indicates an expected call of UpdateInstallProgress.
func (mr *MockAPIMockRecorder) UpdateInstallProgress(ctx, h, progress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstallProgress", reflect.TypeOf((*MockAPI)(nil).UpdateInstallProgress), ctx, h, progress)
}

// SetBootstrap mocks base method.
func (m *MockAPI) SetBootstrap(ctx context.Context, h *models.Host, isbootstrap bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBootstrap", ctx, h, isbootstrap)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBootstrap indicates an expected call of SetBootstrap.
func (mr *MockAPIMockRecorder) SetBootstrap(ctx, h, isbootstrap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBootstrap", reflect.TypeOf((*MockAPI)(nil).SetBootstrap), ctx, h, isbootstrap)
}

// UpdateConnectivityReport mocks base method.
func (m *MockAPI) UpdateConnectivityReport(ctx context.Context, h *models.Host, connectivityReport string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConnectivityReport", ctx, h, connectivityReport)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConnectivityReport indicates an expected call of UpdateConnectivityReport.
func (mr *MockAPIMockRecorder) UpdateConnectivityReport(ctx, h, connectivityReport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConnectivityReport", reflect.TypeOf((*MockAPI)(nil).UpdateConnectivityReport), ctx, h, connectivityReport)
}
