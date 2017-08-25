package mock_validator

import "github.com/stretchr/testify/mock"

import "time"

type MockFormer struct {
	mock.Mock
}

func NewMockFormer() *MockFormer {
	return &MockFormer{}
}

func (m *MockFormer) Close() {
	m.Called()
}
func (m *MockFormer) Clear() {
	m.Called()
}
func (m *MockFormer) HasError() bool {
	ret := m.Called()

	r0 := ret.Get(0).(bool)

	return r0
}
func (m *MockFormer) AddError(field string, msg string) {
	m.Called(field, msg)
}
func (m *MockFormer) Messages() map[string]string {
	ret := m.Called()

	var r0 map[string]string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(map[string]string)
	}

	return r0
}
func (m *MockFormer) Message() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockFormer) Error() error {
	ret := m.Called()

	r0 := ret.Error(0)

	return r0
}
func (m *MockFormer) RequiredString(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) RequiredInt(val int, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) RequiredFloat64(val float64, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) RequiredBool(val bool, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) RequiredEmail(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) NotNil(val interface{}, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) RequiredTime(val time.Time, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) MinInt(val int, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormer) MaxInt(val int, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormer) MaxFloat64(val float64, n float64, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormer) MinFloat64(val float64, n float64, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormer) MinChar(val string, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormer) MaxChar(val string, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormer) Email(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) Gender(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) Confirm(val string, confirm string, field string, msg ...string) {
	m.Called(val, confirm, field, msg)
}
func (m *MockFormer) ISO8601DataTime(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormer) InString(val string, in []string, field string, msg ...string) {
	m.Called(val, in, field, msg)
}
func (m *MockFormer) Length(val int, atleast int, field string, msg ...string) {
	m.Called(val, atleast, field, msg)
}
