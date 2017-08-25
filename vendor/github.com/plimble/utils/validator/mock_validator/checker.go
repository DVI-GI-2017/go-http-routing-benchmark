package mock_validator

import "github.com/stretchr/testify/mock"

import "time"

type MockChecker struct {
	mock.Mock
}

func NewMockChecker() *MockChecker {
	return &MockChecker{}
}

func (m *MockChecker) GetError() error {
	ret := m.Called()

	r0 := ret.Error(0)

	return r0
}
func (m *MockChecker) Clear() {
	m.Called()
}
func (m *MockChecker) CheckErr(err error) {
	m.Called(err)
}
func (m *MockChecker) RequiredString(val string, err error) {
	m.Called(val, err)
}
func (m *MockChecker) RequiredInt(val int, err error) {
	m.Called(val, err)
}
func (m *MockChecker) RequiredFloat64(val float64, err error) {
	m.Called(val, err)
}
func (m *MockChecker) RequiredBool(val bool, err error) {
	m.Called(val, err)
}
func (m *MockChecker) RequiredEmail(val string, err error) {
	m.Called(val, err)
}
func (m *MockChecker) NotNil(val interface{}, err error) {
	m.Called(val, err)
}
func (m *MockChecker) RequiredTime(val time.Time, err error) {
	m.Called(val, err)
}
func (m *MockChecker) MinInt(val int, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockChecker) MaxInt(val int, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockChecker) MinFloat64(val float64, n float64, err error) {
	m.Called(val, n, err)
}
func (m *MockChecker) MaxFloat64(val float64, n float64, err error) {
	m.Called(val, n, err)
}
func (m *MockChecker) MinChar(val string, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockChecker) MaxChar(val string, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockChecker) Email(val string, err error) {
	m.Called(val, err)
}
func (m *MockChecker) Gender(val string, err error) {
	m.Called(val, err)
}
func (m *MockChecker) Confirm(val string, confirm string, err error) {
	m.Called(val, confirm, err)
}
func (m *MockChecker) ISO8601DataTime(val string, err error) {
	m.Called(val, err)
}
func (m *MockChecker) InString(val string, in []string, err error) {
	m.Called(val, in, err)
}
func (m *MockChecker) Length(val int, atleast int, err error) {
	m.Called(val, atleast, err)
}
