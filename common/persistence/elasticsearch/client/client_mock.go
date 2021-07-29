// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v7 "github.com/olivere/elastic/v7"
	v1 "go.temporal.io/api/enums/v1"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockClient) Count(ctx context.Context, index, query string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, index, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockClientMockRecorder) Count(ctx, index, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockClient)(nil).Count), ctx, index, query)
}

// GetMapping mocks base method.
func (m *MockClient) GetMapping(ctx context.Context, index string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapping", ctx, index)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapping indicates an expected call of GetMapping.
func (mr *MockClientMockRecorder) GetMapping(ctx, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapping", reflect.TypeOf((*MockClient)(nil).GetMapping), ctx, index)
}

// PutMapping mocks base method.
func (m *MockClient) PutMapping(ctx context.Context, index string, mapping map[string]v1.IndexedValueType) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMapping", ctx, index, mapping)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutMapping indicates an expected call of PutMapping.
func (mr *MockClientMockRecorder) PutMapping(ctx, index, mapping interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMapping", reflect.TypeOf((*MockClient)(nil).PutMapping), ctx, index, mapping)
}

// RunBulkProcessor mocks base method.
func (m *MockClient) RunBulkProcessor(ctx context.Context, p *BulkProcessorParameters) (BulkProcessor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunBulkProcessor", ctx, p)
	ret0, _ := ret[0].(BulkProcessor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunBulkProcessor indicates an expected call of RunBulkProcessor.
func (mr *MockClientMockRecorder) RunBulkProcessor(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunBulkProcessor", reflect.TypeOf((*MockClient)(nil).RunBulkProcessor), ctx, p)
}

// Search mocks base method.
func (m *MockClient) Search(ctx context.Context, p *SearchParameters) (*v7.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, p)
	ret0, _ := ret[0].(*v7.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockClientMockRecorder) Search(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockClient)(nil).Search), ctx, p)
}

// SearchWithDSL mocks base method.
func (m *MockClient) SearchWithDSL(ctx context.Context, index, query string) (*v7.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithDSL", ctx, index, query)
	ret0, _ := ret[0].(*v7.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithDSL indicates an expected call of SearchWithDSL.
func (mr *MockClientMockRecorder) SearchWithDSL(ctx, index, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithDSL", reflect.TypeOf((*MockClient)(nil).SearchWithDSL), ctx, index, query)
}

// WaitForYellowStatus mocks base method.
func (m *MockClient) WaitForYellowStatus(ctx context.Context, index string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForYellowStatus", ctx, index)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForYellowStatus indicates an expected call of WaitForYellowStatus.
func (mr *MockClientMockRecorder) WaitForYellowStatus(ctx, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForYellowStatus", reflect.TypeOf((*MockClient)(nil).WaitForYellowStatus), ctx, index)
}

// MockClientV7 is a mock of ClientV7 interface.
type MockClientV7 struct {
	ctrl     *gomock.Controller
	recorder *MockClientV7MockRecorder
}

// MockClientV7MockRecorder is the mock recorder for MockClientV7.
type MockClientV7MockRecorder struct {
	mock *MockClientV7
}

// NewMockClientV7 creates a new mock instance.
func NewMockClientV7(ctrl *gomock.Controller) *MockClientV7 {
	mock := &MockClientV7{ctrl: ctrl}
	mock.recorder = &MockClientV7MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientV7) EXPECT() *MockClientV7MockRecorder {
	return m.recorder
}

// ClosePointInTime mocks base method.
func (m *MockClientV7) ClosePointInTime(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClosePointInTime", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClosePointInTime indicates an expected call of ClosePointInTime.
func (mr *MockClientV7MockRecorder) ClosePointInTime(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosePointInTime", reflect.TypeOf((*MockClientV7)(nil).ClosePointInTime), ctx, id)
}

// OpenPointInTime mocks base method.
func (m *MockClientV7) OpenPointInTime(ctx context.Context, index, keepAliveInterval string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenPointInTime", ctx, index, keepAliveInterval)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenPointInTime indicates an expected call of OpenPointInTime.
func (mr *MockClientV7MockRecorder) OpenPointInTime(ctx, index, keepAliveInterval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPointInTime", reflect.TypeOf((*MockClientV7)(nil).OpenPointInTime), ctx, index, keepAliveInterval)
}

// SearchWithDSLWithPIT mocks base method.
func (m *MockClientV7) SearchWithDSLWithPIT(ctx context.Context, query string) (*v7.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithDSLWithPIT", ctx, query)
	ret0, _ := ret[0].(*v7.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithDSLWithPIT indicates an expected call of SearchWithDSLWithPIT.
func (mr *MockClientV7MockRecorder) SearchWithDSLWithPIT(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithDSLWithPIT", reflect.TypeOf((*MockClientV7)(nil).SearchWithDSLWithPIT), ctx, query)
}

// MockClientV6 is a mock of ClientV6 interface.
type MockClientV6 struct {
	ctrl     *gomock.Controller
	recorder *MockClientV6MockRecorder
}

// MockClientV6MockRecorder is the mock recorder for MockClientV6.
type MockClientV6MockRecorder struct {
	mock *MockClientV6
}

// NewMockClientV6 creates a new mock instance.
func NewMockClientV6(ctrl *gomock.Controller) *MockClientV6 {
	mock := &MockClientV6{ctrl: ctrl}
	mock.recorder = &MockClientV6MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientV6) EXPECT() *MockClientV6MockRecorder {
	return m.recorder
}

// Scroll mocks base method.
func (m *MockClientV6) Scroll(ctx context.Context, scrollID string) (*v7.SearchResult, ScrollService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scroll", ctx, scrollID)
	ret0, _ := ret[0].(*v7.SearchResult)
	ret1, _ := ret[1].(ScrollService)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Scroll indicates an expected call of Scroll.
func (mr *MockClientV6MockRecorder) Scroll(ctx, scrollID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scroll", reflect.TypeOf((*MockClientV6)(nil).Scroll), ctx, scrollID)
}

// ScrollFirstPage mocks base method.
func (m *MockClientV6) ScrollFirstPage(ctx context.Context, index, query string) (*v7.SearchResult, ScrollService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScrollFirstPage", ctx, index, query)
	ret0, _ := ret[0].(*v7.SearchResult)
	ret1, _ := ret[1].(ScrollService)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ScrollFirstPage indicates an expected call of ScrollFirstPage.
func (mr *MockClientV6MockRecorder) ScrollFirstPage(ctx, index, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScrollFirstPage", reflect.TypeOf((*MockClientV6)(nil).ScrollFirstPage), ctx, index, query)
}

// MockCLIClient is a mock of CLIClient interface.
type MockCLIClient struct {
	ctrl     *gomock.Controller
	recorder *MockCLIClientMockRecorder
}

// MockCLIClientMockRecorder is the mock recorder for MockCLIClient.
type MockCLIClientMockRecorder struct {
	mock *MockCLIClient
}

// NewMockCLIClient creates a new mock instance.
func NewMockCLIClient(ctrl *gomock.Controller) *MockCLIClient {
	mock := &MockCLIClient{ctrl: ctrl}
	mock.recorder = &MockCLIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCLIClient) EXPECT() *MockCLIClientMockRecorder {
	return m.recorder
}

// Bulk mocks base method.
func (m *MockCLIClient) Bulk() BulkService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bulk")
	ret0, _ := ret[0].(BulkService)
	return ret0
}

// Bulk indicates an expected call of Bulk.
func (mr *MockCLIClientMockRecorder) Bulk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bulk", reflect.TypeOf((*MockCLIClient)(nil).Bulk))
}

// CatIndices mocks base method.
func (m *MockCLIClient) CatIndices(ctx context.Context) (v7.CatIndicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CatIndices", ctx)
	ret0, _ := ret[0].(v7.CatIndicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CatIndices indicates an expected call of CatIndices.
func (mr *MockCLIClientMockRecorder) CatIndices(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CatIndices", reflect.TypeOf((*MockCLIClient)(nil).CatIndices), ctx)
}

// Count mocks base method.
func (m *MockCLIClient) Count(ctx context.Context, index, query string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, index, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockCLIClientMockRecorder) Count(ctx, index, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCLIClient)(nil).Count), ctx, index, query)
}

// Delete mocks base method.
func (m *MockCLIClient) Delete(ctx context.Context, indexName, docID string, version int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, indexName, docID, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCLIClientMockRecorder) Delete(ctx, indexName, docID, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCLIClient)(nil).Delete), ctx, indexName, docID, version)
}

// GetMapping mocks base method.
func (m *MockCLIClient) GetMapping(ctx context.Context, index string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapping", ctx, index)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapping indicates an expected call of GetMapping.
func (mr *MockCLIClientMockRecorder) GetMapping(ctx, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapping", reflect.TypeOf((*MockCLIClient)(nil).GetMapping), ctx, index)
}

// PutMapping mocks base method.
func (m *MockCLIClient) PutMapping(ctx context.Context, index string, mapping map[string]v1.IndexedValueType) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMapping", ctx, index, mapping)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutMapping indicates an expected call of PutMapping.
func (mr *MockCLIClientMockRecorder) PutMapping(ctx, index, mapping interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMapping", reflect.TypeOf((*MockCLIClient)(nil).PutMapping), ctx, index, mapping)
}

// RunBulkProcessor mocks base method.
func (m *MockCLIClient) RunBulkProcessor(ctx context.Context, p *BulkProcessorParameters) (BulkProcessor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunBulkProcessor", ctx, p)
	ret0, _ := ret[0].(BulkProcessor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunBulkProcessor indicates an expected call of RunBulkProcessor.
func (mr *MockCLIClientMockRecorder) RunBulkProcessor(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunBulkProcessor", reflect.TypeOf((*MockCLIClient)(nil).RunBulkProcessor), ctx, p)
}

// Search mocks base method.
func (m *MockCLIClient) Search(ctx context.Context, p *SearchParameters) (*v7.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, p)
	ret0, _ := ret[0].(*v7.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockCLIClientMockRecorder) Search(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockCLIClient)(nil).Search), ctx, p)
}

// SearchWithDSL mocks base method.
func (m *MockCLIClient) SearchWithDSL(ctx context.Context, index, query string) (*v7.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithDSL", ctx, index, query)
	ret0, _ := ret[0].(*v7.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithDSL indicates an expected call of SearchWithDSL.
func (mr *MockCLIClientMockRecorder) SearchWithDSL(ctx, index, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithDSL", reflect.TypeOf((*MockCLIClient)(nil).SearchWithDSL), ctx, index, query)
}

// WaitForYellowStatus mocks base method.
func (m *MockCLIClient) WaitForYellowStatus(ctx context.Context, index string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForYellowStatus", ctx, index)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForYellowStatus indicates an expected call of WaitForYellowStatus.
func (mr *MockCLIClientMockRecorder) WaitForYellowStatus(ctx, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForYellowStatus", reflect.TypeOf((*MockCLIClient)(nil).WaitForYellowStatus), ctx, index)
}

// MockIntegrationTestsClient is a mock of IntegrationTestsClient interface.
type MockIntegrationTestsClient struct {
	ctrl     *gomock.Controller
	recorder *MockIntegrationTestsClientMockRecorder
}

// MockIntegrationTestsClientMockRecorder is the mock recorder for MockIntegrationTestsClient.
type MockIntegrationTestsClientMockRecorder struct {
	mock *MockIntegrationTestsClient
}

// NewMockIntegrationTestsClient creates a new mock instance.
func NewMockIntegrationTestsClient(ctrl *gomock.Controller) *MockIntegrationTestsClient {
	mock := &MockIntegrationTestsClient{ctrl: ctrl}
	mock.recorder = &MockIntegrationTestsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntegrationTestsClient) EXPECT() *MockIntegrationTestsClientMockRecorder {
	return m.recorder
}

// CreateIndex mocks base method.
func (m *MockIntegrationTestsClient) CreateIndex(ctx context.Context, index string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndex", ctx, index)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIndex indicates an expected call of CreateIndex.
func (mr *MockIntegrationTestsClientMockRecorder) CreateIndex(ctx, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndex", reflect.TypeOf((*MockIntegrationTestsClient)(nil).CreateIndex), ctx, index)
}

// DeleteIndex mocks base method.
func (m *MockIntegrationTestsClient) DeleteIndex(ctx context.Context, indexName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIndex", ctx, indexName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIndex indicates an expected call of DeleteIndex.
func (mr *MockIntegrationTestsClientMockRecorder) DeleteIndex(ctx, indexName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIndex", reflect.TypeOf((*MockIntegrationTestsClient)(nil).DeleteIndex), ctx, indexName)
}

// IndexExists mocks base method.
func (m *MockIntegrationTestsClient) IndexExists(ctx context.Context, indexName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexExists", ctx, indexName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexExists indicates an expected call of IndexExists.
func (mr *MockIntegrationTestsClientMockRecorder) IndexExists(ctx, indexName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexExists", reflect.TypeOf((*MockIntegrationTestsClient)(nil).IndexExists), ctx, indexName)
}

// IndexGetSettings mocks base method.
func (m *MockIntegrationTestsClient) IndexGetSettings(ctx context.Context, indexName string) (map[string]*v7.IndicesGetSettingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexGetSettings", ctx, indexName)
	ret0, _ := ret[0].(map[string]*v7.IndicesGetSettingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexGetSettings indicates an expected call of IndexGetSettings.
func (mr *MockIntegrationTestsClientMockRecorder) IndexGetSettings(ctx, indexName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexGetSettings", reflect.TypeOf((*MockIntegrationTestsClient)(nil).IndexGetSettings), ctx, indexName)
}

// IndexPutSettings mocks base method.
func (m *MockIntegrationTestsClient) IndexPutSettings(ctx context.Context, indexName, bodyString string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexPutSettings", ctx, indexName, bodyString)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexPutSettings indicates an expected call of IndexPutSettings.
func (mr *MockIntegrationTestsClientMockRecorder) IndexPutSettings(ctx, indexName, bodyString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexPutSettings", reflect.TypeOf((*MockIntegrationTestsClient)(nil).IndexPutSettings), ctx, indexName, bodyString)
}

// IndexPutTemplate mocks base method.
func (m *MockIntegrationTestsClient) IndexPutTemplate(ctx context.Context, templateName, bodyString string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexPutTemplate", ctx, templateName, bodyString)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexPutTemplate indicates an expected call of IndexPutTemplate.
func (mr *MockIntegrationTestsClientMockRecorder) IndexPutTemplate(ctx, templateName, bodyString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexPutTemplate", reflect.TypeOf((*MockIntegrationTestsClient)(nil).IndexPutTemplate), ctx, templateName, bodyString)
}

// MockScrollService is a mock of ScrollService interface.
type MockScrollService struct {
	ctrl     *gomock.Controller
	recorder *MockScrollServiceMockRecorder
}

// MockScrollServiceMockRecorder is the mock recorder for MockScrollService.
type MockScrollServiceMockRecorder struct {
	mock *MockScrollService
}

// NewMockScrollService creates a new mock instance.
func NewMockScrollService(ctrl *gomock.Controller) *MockScrollService {
	mock := &MockScrollService{ctrl: ctrl}
	mock.recorder = &MockScrollServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScrollService) EXPECT() *MockScrollServiceMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockScrollService) Clear(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockScrollServiceMockRecorder) Clear(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockScrollService)(nil).Clear), ctx)
}
