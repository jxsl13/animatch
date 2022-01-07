package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jxsl13/animatch/common"
	"github.com/jxsl13/animatch/filter"
	"github.com/pkg/xattr"
	"github.com/spf13/cobra"
)

func NewMXattrCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xattr filepath",
		Short: "allows to show the xattr values of a given file",
		RunE:  xattrCmd,
		Args:  cobra.MinimumNArgs(1),
	}

	return cmd
}

func xattrCmd(cmd *cobra.Command, args []string) error {

	filePath, err := filepath.Abs(strings.Join(args, " "))
	if err != nil {
		return err
	}
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	files := make([]string, 0, 1)
	if info.IsDir() {
		de, err := os.ReadDir(filePath)
		if err != nil {
			return err
		}

		for _, info := range de {
			files = append(files, filepath.Join(filePath, info.Name()))
		}
	} else {
		files = append(files, filePath)
	}

	files = filter.VideoFilePaths(files)

	sb := strings.Builder{}
	sb.Grow(len(files) * 32)

	for _, filePath := range files {
		attributes, err := xattr.List(filePath)
		if err != nil {
			return err
		}

		sb.WriteString(fmt.Sprintf("Found %d xargs attributes for: \n%s\n", len(attributes), filePath))
		for _, attr := range attributes {
			sb.WriteString("\t")
			sb.WriteString(attr)
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	return common.Println(cmd, sb.String())
}
