package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcPagination(t *testing.T) {
	limit, offset, withPagination := CalcPagination(1, 10)

	assert.Equal(t, 10, limit, "limit equal 10")
	assert.Equal(t, 0, offset, "offset equal 0")
	assert.Equal(t, true, withPagination, "pagination equal true")
}
