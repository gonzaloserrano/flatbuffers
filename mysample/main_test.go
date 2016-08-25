package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)

	for _, bin := range []string{
		"monster_out.bin",
		//"monster2_out.bin",
	} {
		t.Log(bin)
		assert.NoError(check(bin))
	}
}
