package test

import (
	"github.com/stretchr/testify/assert"
	"qiaosong03.com/simple02_grpc/api"
	"testing"
)

func Test_apiAdd_01(t *testing.T) {
	res, err := api.Add(3, 4)
	assert.Nil(t, err)
	assert.Equal(t, int32(7), res.Res)
}

func Test_apiSubtract_01(t *testing.T) {
	res, err := api.Subtract(3, 4)
	assert.Nil(t, err)
	assert.Equal(t, int32(-1), res.Res)
}

func Test_apiMultiply_01(t *testing.T) {
	res, err := api.Multiply(3, 4)
	assert.Nil(t, err)
	assert.Equal(t, int32(12), res.Res)
}

func Test_apiDivide_01(t *testing.T) {
	res, err := api.Divide(3, 4)
	assert.Nil(t, err)
	assert.Equal(t, int32(0), res.Res)
}