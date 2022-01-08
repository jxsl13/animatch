package cmd

import (
	"os"
	"strings"

	"github.com/jxsl13/animatch/common"
	"github.com/jxsl13/animatch/filter"
	"github.com/spf13/cobra"
)

const (
	ErrPatchNotFound = common.Error("path not found")
)

func NewMatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "match filepath",
		Short: "allows to match a file path to an anime name.",
		RunE:  matchCmd,
		Args:  cobra.MinimumNArgs(1),
	}

	return cmd
}

func matchCmd(cmd *cobra.Command, args []string) error {

	filePath := strings.Join(args, " ")
	fileStat, err := os.Stat(filePath)

	filePaths := []string{}
	if err != nil || !fileStat.IsDir() {
		filePaths = append(filePaths, filePath)
	} else {
		files, err := common.Readdir(filePath)
		if err != nil {
			return err
		}
		filePaths = append(filePaths, files...)
	}

	filePaths = filter.VideoFilePaths(filePaths)

	mr, err := Match(cmd, filePaths)
	if err != nil {
		return err
	}

	for _, m := range mr {
		common.Println(cmd, m.String())
	}

	return nil
}
