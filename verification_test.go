/*
Copyright © 2024 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckForStringInSlice(t *testing.T) {

	assert := assert.New(t)

	members := []string{"robbie", "marc", "gary"}
	starNotIn := "bernd"
	starIn := "robbie"

	assert.Equal(CheckForStringInSlice(members, starIn), true)
	assert.Equal(CheckForStringInSlice(members, starNotIn), false)

}
