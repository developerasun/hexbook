package test

import (
	"testing"

	pkg "github.com/hexbook/pkg"
	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3, pkg.Dummy())
}
