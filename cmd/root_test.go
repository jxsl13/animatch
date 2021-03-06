package cmd

import (
	"testing"

	"github.com/jxsl13/animatch/common"
	"github.com/stretchr/testify/assert"
)

func TestNewRootCmd(t *testing.T) {
	assert := assert.New(t)
	out, err := common.ExecuteWithArgs(NewRootCmd(), "--help")
	assert.NoError(err)
	t.Log(string(out))
}
