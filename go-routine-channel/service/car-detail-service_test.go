package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	carDetailService = NewCarDetailService()
)

func TestCarDetailService(t *testing.T) {
	carDetails := carDetailService.GetDetails()

	assert.NotNil(t, carDetails, "Car detail is empty")
	assert.Equal(t, 1, carDetails.ID)
	assert.Equal(t, "Mitsubishi", carDetails.Brand)
	assert.Equal(t, 2002, carDetails.Year)

}
