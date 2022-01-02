package common

import (
	"bytes"

	"github.com/spf13/cobra"
)

// ExecuteWithArgs allows to test commands properly and individually
// Great reference: https://gianarb.it/blog/golang-mockmania-cli-command-with-cobra
func ExecuteWithArgs(cmd *cobra.Command, args ...string) (out []byte, err error) {
	outBuf := bytes.NewBuffer(nil)
	cmd.SetArgs(args)
	cmd.SetOutput(outBuf)
	cmd.SetOut(outBuf)
	err = cmd.Execute()
	return outBuf.Bytes(), err
}
