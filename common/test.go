package common

import (
	"bytes"

	"github.com/spf13/cobra"
)

// ExecuteWithArgs allows to test commands properly and individually
// Great reference: https://gianarb.it/blog/golang-mockmania-cli-command-with-cobra
func ExecuteWithArgs(cmd *cobra.Command, args ...string) (out []byte, err error) {
	outBuf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	cmd.SetOut(outBuf)
	cmd.SetErr(errBuf)
	cmd.SetArgs(args)
	err = cmd.Execute()
	if err != nil {
		return errBuf.Bytes(), err
	}

	return outBuf.Bytes(), nil
}
