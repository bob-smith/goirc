// Automatically generated by MockGen. DO NOT EDIT!
// Source: tracker.go

package state

import (
	gomock "gomock.googlecode.com/hg/gomock"
)

// Mock of StateTracker interface
type MockStateTracker struct {
	ctrl     *gomock.Controller
	recorder *_MockStateTrackerRecorder
}

// Recorder for MockStateTracker (not exported)
type _MockStateTrackerRecorder struct {
	mock *MockStateTracker
}

func NewMockStateTracker(ctrl *gomock.Controller) *MockStateTracker {
	mock := &MockStateTracker{ctrl: ctrl}
	mock.recorder = &_MockStateTrackerRecorder{mock}
	return mock
}

func (m *MockStateTracker) EXPECT() *_MockStateTrackerRecorder {
	return m.recorder
}

func (m *MockStateTracker) NewNick(nick string) *Nick {
	ret := m.ctrl.Call(m, "NewNick", nick)
	ret0, _ := ret[0].(*Nick)
	return ret0
}

func (mr *_MockStateTrackerRecorder) NewNick(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "NewNick", arg0)
}

func (m *MockStateTracker) GetNick(nick string) *Nick {
	ret := m.ctrl.Call(m, "GetNick", nick)
	ret0, _ := ret[0].(*Nick)
	return ret0
}

func (mr *_MockStateTrackerRecorder) GetNick(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "GetNick", arg0)
}

func (m *MockStateTracker) ReNick(old string, neu string) {
	m.ctrl.Call(m, "ReNick", old, neu)
}

func (mr *_MockStateTrackerRecorder) ReNick(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "ReNick", arg0, arg1)
}

func (m *MockStateTracker) DelNick(nick string) {
	m.ctrl.Call(m, "DelNick", nick)
}

func (mr *_MockStateTrackerRecorder) DelNick(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "DelNick", arg0)
}

func (m *MockStateTracker) NewChannel(channel string) *Channel {
	ret := m.ctrl.Call(m, "NewChannel", channel)
	ret0, _ := ret[0].(*Channel)
	return ret0
}

func (mr *_MockStateTrackerRecorder) NewChannel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "NewChannel", arg0)
}

func (m *MockStateTracker) GetChannel(channel string) *Channel {
	ret := m.ctrl.Call(m, "GetChannel", channel)
	ret0, _ := ret[0].(*Channel)
	return ret0
}

func (mr *_MockStateTrackerRecorder) GetChannel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "GetChannel", arg0)
}

func (m *MockStateTracker) DelChannel(channel string) {
	m.ctrl.Call(m, "DelChannel", channel)
}

func (mr *_MockStateTrackerRecorder) DelChannel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "DelChannel", arg0)
}

func (m *MockStateTracker) Me() *Nick {
	ret := m.ctrl.Call(m, "Me")
	ret0, _ := ret[0].(*Nick)
	return ret0
}

func (mr *_MockStateTrackerRecorder) Me() *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "Me")
}

func (m *MockStateTracker) IsOn(channel string, nick string) (*ChanPrivs, bool) {
	ret := m.ctrl.Call(m, "IsOn", channel, nick)
	ret0, _ := ret[0].(*ChanPrivs)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (mr *_MockStateTrackerRecorder) IsOn(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "IsOn", arg0, arg1)
}

func (m *MockStateTracker) Associate(channel *Channel, nick *Nick) *ChanPrivs {
	ret := m.ctrl.Call(m, "Associate", channel, nick)
	ret0, _ := ret[0].(*ChanPrivs)
	return ret0
}

func (mr *_MockStateTrackerRecorder) Associate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "Associate", arg0, arg1)
}

func (m *MockStateTracker) Dissociate(channel *Channel, nick *Nick) {
	m.ctrl.Call(m, "Dissociate", channel, nick)
}

func (mr *_MockStateTrackerRecorder) Dissociate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "Dissociate", arg0, arg1)
}

func (m *MockStateTracker) Wipe() {
	m.ctrl.Call(m, "Wipe")
}

func (mr *_MockStateTrackerRecorder) Wipe() *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "Wipe")
}

func (m *MockStateTracker) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *_MockStateTrackerRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCall(mr.mock, "String")
}