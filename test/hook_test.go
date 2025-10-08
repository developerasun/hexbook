package test

import (
	pkg "github.com/fatcat/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDummy(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3, pkg.Dummy())
}
