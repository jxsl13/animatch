package cmd

import (
	"testing"

	"github.com/jxsl13/animatch/common"
	"github.com/stretchr/testify/assert"
)

func TestNewMatchCmd(t *testing.T) {
	assert := assert.New(t)
	out, err := common.ExecuteWithArgs(NewMatchCmd(), "/Users/jxsl13/Desktop/filebot/input/one", "piece/One Piece - 922.mkv")
	assert.NoError(err)
	t.Log(string(out))
}
