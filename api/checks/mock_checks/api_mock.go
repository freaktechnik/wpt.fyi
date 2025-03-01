// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/api/checks (interfaces: API)

// Package mock_checks is a generated GoMock package.
package mock_checks

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v42/github"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
)

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

// CancelRun mocks base method.
func (m *MockAPI) CancelRun(arg0, arg1, arg2 string, arg3 *github.CheckRun, arg4 *github.Installation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelRun", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRun indicates an expected call of CancelRun.
func (mr *MockAPIMockRecorder) CancelRun(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRun", reflect.TypeOf((*MockAPI)(nil).CancelRun), arg0, arg1, arg2, arg3, arg4)
}

// Context mocks base method.
func (m *MockAPI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAPIMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAPI)(nil).Context))
}

// CreateWPTCheckSuite mocks base method.
func (m *MockAPI) CreateWPTCheckSuite(arg0, arg1 int64, arg2 string, arg3 ...int) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWPTCheckSuite", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWPTCheckSuite indicates an expected call of CreateWPTCheckSuite.
func (mr *MockAPIMockRecorder) CreateWPTCheckSuite(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWPTCheckSuite", reflect.TypeOf((*MockAPI)(nil).CreateWPTCheckSuite), varargs...)
}

// GetGitHubClient mocks base method.
func (m *MockAPI) GetGitHubClient() (*github.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitHubClient")
	ret0, _ := ret[0].(*github.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGitHubClient indicates an expected call of GetGitHubClient.
func (mr *MockAPIMockRecorder) GetGitHubClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitHubClient", reflect.TypeOf((*MockAPI)(nil).GetGitHubClient))
}

// GetHTTPClient mocks base method.
func (m *MockAPI) GetHTTPClient() *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPClient")
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetHTTPClient indicates an expected call of GetHTTPClient.
func (mr *MockAPIMockRecorder) GetHTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPClient", reflect.TypeOf((*MockAPI)(nil).GetHTTPClient))
}

// GetHTTPClientWithTimeout mocks base method.
func (m *MockAPI) GetHTTPClientWithTimeout(arg0 time.Duration) *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPClientWithTimeout", arg0)
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetHTTPClientWithTimeout indicates an expected call of GetHTTPClientWithTimeout.
func (mr *MockAPIMockRecorder) GetHTTPClientWithTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPClientWithTimeout", reflect.TypeOf((*MockAPI)(nil).GetHTTPClientWithTimeout), arg0)
}

// GetHostname mocks base method.
func (m *MockAPI) GetHostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHostname indicates an expected call of GetHostname.
func (mr *MockAPIMockRecorder) GetHostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostname", reflect.TypeOf((*MockAPI)(nil).GetHostname))
}

// GetResultsURL mocks base method.
func (m *MockAPI) GetResultsURL(arg0 shared.TestRunFilter) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResultsURL", arg0)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetResultsURL indicates an expected call of GetResultsURL.
func (mr *MockAPIMockRecorder) GetResultsURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResultsURL", reflect.TypeOf((*MockAPI)(nil).GetResultsURL), arg0)
}

// GetResultsUploadURL mocks base method.
func (m *MockAPI) GetResultsUploadURL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResultsUploadURL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetResultsUploadURL indicates an expected call of GetResultsUploadURL.
func (mr *MockAPIMockRecorder) GetResultsUploadURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResultsUploadURL", reflect.TypeOf((*MockAPI)(nil).GetResultsUploadURL))
}

// GetRunsURL mocks base method.
func (m *MockAPI) GetRunsURL(arg0 shared.TestRunFilter) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunsURL", arg0)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetRunsURL indicates an expected call of GetRunsURL.
func (mr *MockAPIMockRecorder) GetRunsURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunsURL", reflect.TypeOf((*MockAPI)(nil).GetRunsURL), arg0)
}

// GetServiceHostname mocks base method.
func (m *MockAPI) GetServiceHostname(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceHostname", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceHostname indicates an expected call of GetServiceHostname.
func (mr *MockAPIMockRecorder) GetServiceHostname(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceHostname", reflect.TypeOf((*MockAPI)(nil).GetServiceHostname), arg0)
}

// GetSuitesForSHA mocks base method.
func (m *MockAPI) GetSuitesForSHA(arg0 string) ([]shared.CheckSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSuitesForSHA", arg0)
	ret0, _ := ret[0].([]shared.CheckSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuitesForSHA indicates an expected call of GetSuitesForSHA.
func (mr *MockAPIMockRecorder) GetSuitesForSHA(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuitesForSHA", reflect.TypeOf((*MockAPI)(nil).GetSuitesForSHA), arg0)
}

// GetUploader mocks base method.
func (m *MockAPI) GetUploader(arg0 string) (shared.Uploader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUploader", arg0)
	ret0, _ := ret[0].(shared.Uploader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploader indicates an expected call of GetUploader.
func (mr *MockAPIMockRecorder) GetUploader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploader", reflect.TypeOf((*MockAPI)(nil).GetUploader), arg0)
}

// GetVersion mocks base method.
func (m *MockAPI) GetVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockAPIMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockAPI)(nil).GetVersion))
}

// GetVersionedHostname mocks base method.
func (m *MockAPI) GetVersionedHostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionedHostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersionedHostname indicates an expected call of GetVersionedHostname.
func (mr *MockAPIMockRecorder) GetVersionedHostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionedHostname", reflect.TypeOf((*MockAPI)(nil).GetVersionedHostname))
}

// GetWPTRepoAppInstallationIDs mocks base method.
func (m *MockAPI) GetWPTRepoAppInstallationIDs() (int64, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWPTRepoAppInstallationIDs")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// GetWPTRepoAppInstallationIDs indicates an expected call of GetWPTRepoAppInstallationIDs.
func (mr *MockAPIMockRecorder) GetWPTRepoAppInstallationIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWPTRepoAppInstallationIDs", reflect.TypeOf((*MockAPI)(nil).GetWPTRepoAppInstallationIDs))
}

// IgnoreFailure mocks base method.
func (m *MockAPI) IgnoreFailure(arg0, arg1, arg2 string, arg3 *github.CheckRun, arg4 *github.Installation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IgnoreFailure", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// IgnoreFailure indicates an expected call of IgnoreFailure.
func (mr *MockAPIMockRecorder) IgnoreFailure(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IgnoreFailure", reflect.TypeOf((*MockAPI)(nil).IgnoreFailure), arg0, arg1, arg2, arg3, arg4)
}

// IsFeatureEnabled mocks base method.
func (m *MockAPI) IsFeatureEnabled(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFeatureEnabled", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFeatureEnabled indicates an expected call of IsFeatureEnabled.
func (mr *MockAPIMockRecorder) IsFeatureEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFeatureEnabled", reflect.TypeOf((*MockAPI)(nil).IsFeatureEnabled), arg0)
}

// ScheduleResultsProcessing mocks base method.
func (m *MockAPI) ScheduleResultsProcessing(arg0 string, arg1 shared.ProductSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleResultsProcessing", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleResultsProcessing indicates an expected call of ScheduleResultsProcessing.
func (mr *MockAPIMockRecorder) ScheduleResultsProcessing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleResultsProcessing", reflect.TypeOf((*MockAPI)(nil).ScheduleResultsProcessing), arg0, arg1)
}

// ScheduleTask mocks base method.
func (m *MockAPI) ScheduleTask(arg0, arg1, arg2 string, arg3 url.Values) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleTask indicates an expected call of ScheduleTask.
func (mr *MockAPIMockRecorder) ScheduleTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleTask", reflect.TypeOf((*MockAPI)(nil).ScheduleTask), arg0, arg1, arg2, arg3)
}
