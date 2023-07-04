// Code generated by MockGen. DO NOT EDIT.
// Source: ./../../x/transfer/types/expected_keepers.go

// Package mock_types is a generated GoMock package.
package mock_types

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	gomock "github.com/golang/mock/gomock"
	types2 "github.com/neutron-org/neutron/x/feerefunder/types"
)

// MockContractManagerKeeper is a mock of ContractManagerKeeper interface.
type MockContractManagerKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockContractManagerKeeperMockRecorder
}

// MockContractManagerKeeperMockRecorder is the mock recorder for MockContractManagerKeeper.
type MockContractManagerKeeperMockRecorder struct {
	mock *MockContractManagerKeeper
}

// NewMockContractManagerKeeper creates a new mock instance.
func NewMockContractManagerKeeper(ctrl *gomock.Controller) *MockContractManagerKeeper {
	mock := &MockContractManagerKeeper{ctrl: ctrl}
	mock.recorder = &MockContractManagerKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractManagerKeeper) EXPECT() *MockContractManagerKeeperMockRecorder {
	return m.recorder
}

// AddContractFailure mocks base method.
func (m *MockContractManagerKeeper) AddContractFailure(ctx types.Context, channelID, address string, ackID uint64, ackType string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddContractFailure", ctx, channelID, address, ackID, ackType)
}

// AddContractFailure indicates an expected call of AddContractFailure.
func (mr *MockContractManagerKeeperMockRecorder) AddContractFailure(ctx, channelID, address, ackID, ackType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddContractFailure", reflect.TypeOf((*MockContractManagerKeeper)(nil).AddContractFailure), ctx, channelID, address, ackID, ackType)
}

// HasContractInfo mocks base method.
func (m *MockContractManagerKeeper) HasContractInfo(ctx types.Context, contractAddress types.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasContractInfo", ctx, contractAddress)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasContractInfo indicates an expected call of HasContractInfo.
func (mr *MockContractManagerKeeperMockRecorder) HasContractInfo(ctx, contractAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasContractInfo", reflect.TypeOf((*MockContractManagerKeeper)(nil).HasContractInfo), ctx, contractAddress)
}

// SudoError mocks base method.
func (m *MockContractManagerKeeper) SudoError(ctx types.Context, senderAddress types.AccAddress, request types1.Packet, details string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoError", ctx, senderAddress, request, details)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoError indicates an expected call of SudoError.
func (mr *MockContractManagerKeeperMockRecorder) SudoError(ctx, senderAddress, request, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoError", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoError), ctx, senderAddress, request, details)
}

// SudoResponse mocks base method.
func (m *MockContractManagerKeeper) SudoResponse(ctx types.Context, senderAddress types.AccAddress, request types1.Packet, msg []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoResponse", ctx, senderAddress, request, msg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoResponse indicates an expected call of SudoResponse.
func (mr *MockContractManagerKeeperMockRecorder) SudoResponse(ctx, senderAddress, request, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoResponse", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoResponse), ctx, senderAddress, request, msg)
}

// SudoTimeout mocks base method.
func (m *MockContractManagerKeeper) SudoTimeout(ctx types.Context, senderAddress types.AccAddress, request types1.Packet) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoTimeout", ctx, senderAddress, request)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoTimeout indicates an expected call of SudoTimeout.
func (mr *MockContractManagerKeeperMockRecorder) SudoTimeout(ctx, senderAddress, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoTimeout", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoTimeout), ctx, senderAddress, request)
}

// MockFeeRefunderKeeper is a mock of FeeRefunderKeeper interface.
type MockFeeRefunderKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockFeeRefunderKeeperMockRecorder
}

// MockFeeRefunderKeeperMockRecorder is the mock recorder for MockFeeRefunderKeeper.
type MockFeeRefunderKeeperMockRecorder struct {
	mock *MockFeeRefunderKeeper
}

// NewMockFeeRefunderKeeper creates a new mock instance.
func NewMockFeeRefunderKeeper(ctrl *gomock.Controller) *MockFeeRefunderKeeper {
	mock := &MockFeeRefunderKeeper{ctrl: ctrl}
	mock.recorder = &MockFeeRefunderKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeeRefunderKeeper) EXPECT() *MockFeeRefunderKeeperMockRecorder {
	return m.recorder
}

// DistributeAcknowledgementFee mocks base method.
func (m *MockFeeRefunderKeeper) DistributeAcknowledgementFee(ctx types.Context, receiver types.AccAddress, packetID types2.PacketID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DistributeAcknowledgementFee", ctx, receiver, packetID)
}

// DistributeAcknowledgementFee indicates an expected call of DistributeAcknowledgementFee.
func (mr *MockFeeRefunderKeeperMockRecorder) DistributeAcknowledgementFee(ctx, receiver, packetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeAcknowledgementFee", reflect.TypeOf((*MockFeeRefunderKeeper)(nil).DistributeAcknowledgementFee), ctx, receiver, packetID)
}

// DistributeTimeoutFee mocks base method.
func (m *MockFeeRefunderKeeper) DistributeTimeoutFee(ctx types.Context, receiver types.AccAddress, packetID types2.PacketID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DistributeTimeoutFee", ctx, receiver, packetID)
}

// DistributeTimeoutFee indicates an expected call of DistributeTimeoutFee.
func (mr *MockFeeRefunderKeeperMockRecorder) DistributeTimeoutFee(ctx, receiver, packetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTimeoutFee", reflect.TypeOf((*MockFeeRefunderKeeper)(nil).DistributeTimeoutFee), ctx, receiver, packetID)
}

// LockFees mocks base method.
func (m *MockFeeRefunderKeeper) LockFees(ctx types.Context, payer types.AccAddress, packetID types2.PacketID, fee types2.Fee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockFees", ctx, payer, packetID, fee)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockFees indicates an expected call of LockFees.
func (mr *MockFeeRefunderKeeperMockRecorder) LockFees(ctx, payer, packetID, fee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockFees", reflect.TypeOf((*MockFeeRefunderKeeper)(nil).LockFees), ctx, payer, packetID, fee)
}

// MockChannelKeeper is a mock of ChannelKeeper interface.
type MockChannelKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockChannelKeeperMockRecorder
}

// MockChannelKeeperMockRecorder is the mock recorder for MockChannelKeeper.
type MockChannelKeeperMockRecorder struct {
	mock *MockChannelKeeper
}

// NewMockChannelKeeper creates a new mock instance.
func NewMockChannelKeeper(ctrl *gomock.Controller) *MockChannelKeeper {
	mock := &MockChannelKeeper{ctrl: ctrl}
	mock.recorder = &MockChannelKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelKeeper) EXPECT() *MockChannelKeeperMockRecorder {
	return m.recorder
}

// GetAllChannelsWithPortPrefix mocks base method.
func (m *MockChannelKeeper) GetAllChannelsWithPortPrefix(ctx types.Context, portPrefix string) []types1.IdentifiedChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChannelsWithPortPrefix", ctx, portPrefix)
	ret0, _ := ret[0].([]types1.IdentifiedChannel)
	return ret0
}

// GetAllChannelsWithPortPrefix indicates an expected call of GetAllChannelsWithPortPrefix.
func (mr *MockChannelKeeperMockRecorder) GetAllChannelsWithPortPrefix(ctx, portPrefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChannelsWithPortPrefix", reflect.TypeOf((*MockChannelKeeper)(nil).GetAllChannelsWithPortPrefix), ctx, portPrefix)
}

// GetChannel mocks base method.
func (m *MockChannelKeeper) GetChannel(ctx types.Context, srcPort, srcChan string) (types1.Channel, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannel", ctx, srcPort, srcChan)
	ret0, _ := ret[0].(types1.Channel)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel.
func (mr *MockChannelKeeperMockRecorder) GetChannel(ctx, srcPort, srcChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockChannelKeeper)(nil).GetChannel), ctx, srcPort, srcChan)
}

// GetNextSequenceSend mocks base method.
func (m *MockChannelKeeper) GetNextSequenceSend(ctx types.Context, portID, channelID string) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSequenceSend", ctx, portID, channelID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNextSequenceSend indicates an expected call of GetNextSequenceSend.
func (mr *MockChannelKeeperMockRecorder) GetNextSequenceSend(ctx, portID, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSequenceSend", reflect.TypeOf((*MockChannelKeeper)(nil).GetNextSequenceSend), ctx, portID, channelID)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx types.Context, name string) types0.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, name)
	ret0, _ := ret[0].(types0.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, name)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}
