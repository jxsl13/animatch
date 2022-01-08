package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jxsl13/animatch/common"
	"github.com/jxsl13/animatch/filter"
	"github.com/spf13/cobra"
)

const (
	FlagApply          = "apply"
	FlagApplyShorthand = "a"
)

func NewTagCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag filepath",
		Short: "matches your anime and adds an AniDB suffix [a12345] to all of your matched files.",
		RunE:  tagCmd,
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.Flags().BoolP(
		FlagApply,
		FlagApplyShorthand,
		false,
		"use this flag in order to apply the action to the file system.",
	)

	return cmd
}

func tagCmd(cmd *cobra.Command, args []string) error {
	apply, err := common.LookupFlagBool(cmd, FlagApply)
	if err != nil {
		return err
	}

	filePath := strings.Join(args, " ")
	fileStat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	filePaths := []string{}
	if !fileStat.IsDir() {
		filePaths = append(filePaths, filePath)
	} else {
		files, err := common.Readdir(filePath)
		if err != nil {
			return err
		}
		filePaths = append(filePaths, filter.VideoFilePaths(files)...)
	}

	mr, err := Match(cmd, filePaths)
	if err != nil {
		return err
	}

	var (
		longestTerm       = mr.LongestMatchTerm()
		longestTaggedPath = mr.LongestTaggedPath()
		format            = fmt.Sprintf("Renamed: %%-%ds -> %%-%ds\n", longestTerm, longestTaggedPath)
		failFormat        = fmt.Sprintf("FAILED: %%-%ds -> %%-%ds\n\n", longestTerm, longestTaggedPath)
		noMatchFormat     = fmt.Sprintf("NO MATCH: %%-%ds\n\n", longestTerm)
	)

	for _, match := range mr {
		if match.IsMatch(0.24) {

			if apply {
				err = os.Rename(match.MatchTerm, match.TaggedPath())
			}
			if err != nil {
				common.Printf(cmd, failFormat, match.MatchTerm)
				common.Println(cmd, match.String())
			} else {
				common.Printf(cmd, format, match.MatchTerm, match.TaggedPath())
			}
		} else {
			common.Printf(cmd, noMatchFormat, match.MatchTerm)
		}
	}

	if !apply {
		return common.Println(cmd, "No files were moved or renamed, please use the --apply or -a flags in order to apply the changes.")
	}

	return nil
}
