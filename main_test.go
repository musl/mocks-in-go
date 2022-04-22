package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

type MockFFClient struct {
	mock.Mock
}

func (mffc *MockFFClient) GetFlag(name string) string {
	mffc.Called(name)
	return "true"
}

func TestRun(t *testing.T) {
	m := new(MockFFClient)

	m.On("GetFlag", "wat").Return("true")

	r := Reconciler{FFClient: m}
	expected := "run"
	actual := r.Run()

	assert.Equal(t, expected, actual)
	m.AssertExpectations(t)
}
