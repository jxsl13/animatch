package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jxsl13/animatch/anidb"
	"github.com/jxsl13/animatch/clean"
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

	// modify public variable of anydb package in case that this flag is set
	cmd.Flags().IntP(
		FlagPathDepth,
		FlagPathDepthShorthand,
		DefaultPathDepth,
		"allows to add subpath to search query, increasing this value to 2 would add the parent directory to the search",
	)

	cmd.Flags().BoolP(
		FlagApply,
		FlagApplyShorthand,
		false,
		"use this flag in order to apply the action to the file system.",
	)

	return cmd
}

func tagCmd(cmd *cobra.Command, args []string) error {
	depth, err := common.LookupFlagInt(cmd, FlagPathDepth)
	if err != nil {
		return err
	}
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

	for _, p := range filePaths {

		pathPrefix := clean.RemoveExtension(p)
		ext := filepath.Ext(p)

		normalizedTerms := clean.LanguageTags(
			clean.ScreenResolutions(
				clean.TokenizeAll(
					clean.SplitPath(
						clean.Domains(
							clean.Tags(pathPrefix),
						), depth))))

		normalizedTerm := strings.Join(normalizedTerms, " ")

		common.Println(cmd, "Path    : ", p)
		common.Println(cmd, "Search  : ", normalizedTerm)
		distance, title, animeT, err := anidb.Search(normalizedTerm, filter.Metrics)
		if err != nil {
			return err
		}
		common.Println(cmd, "Found   : ", *title)

		newPath := fmt.Sprintf("%s [anidb-%d]%s", pathPrefix, animeT.AID, ext)
		common.Println(cmd, "Result  : ", newPath)
		common.Println(cmd, "Distance: ", distance, "\n")

		if apply {
			err = os.Rename(p, newPath)
			if err != nil {
				return err
			}
		}
	}

	if !apply {
		common.Println(cmd, "No files were moved or renamed, please use the --apply or -a flags in order to apply the changes.")
	}

	return nil
}
