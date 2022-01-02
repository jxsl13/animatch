package cmd

import (
	"testing"

	"github.com/jxsl13/animatch/common"
	"github.com/stretchr/testify/assert"
)

func TestNewSearchCmd(t *testing.T) {
	assert := assert.New(t)
	out, err := common.ExecuteWithArgs(NewSearchCmd(), "one", "piece")
	assert.NoError(err)
	t.Log(string(out))
}
