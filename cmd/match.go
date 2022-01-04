package cmd

import (
	"os"
	"strings"

	"github.com/jxsl13/animatch/anidb"
	"github.com/jxsl13/animatch/clean"
	"github.com/jxsl13/animatch/common"
	"github.com/jxsl13/animatch/filter"
	"github.com/spf13/cobra"
)

const (
	FlagPathDepth          = "path-depth"
	FlagPathDepthShorthand = "p"
	DefaultPathDepth       = 1

	ErrPatchNotFound = common.Error("path not found")
)

var (
	wrapPathErr = common.WrapErr(ErrPatchNotFound)
)

func NewMatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "match filepath",
		Short: "allows to match a file path to an anime name.",
		RunE:  matchCmd,
		Args:  cobra.MinimumNArgs(1),
	}

	// modify public variable of anydb package in case that this flag is set
	cmd.Flags().IntP(
		FlagDistance,
		FlagDistanceShorthand,
		anidb.DefaultMatchDistanceUpperBound,
		"increase this distance in order to allow a broader range of matches",
	)

	// modify public variable of anydb package in case that this flag is set
	cmd.Flags().IntP(
		FlagPathDepth,
		FlagPathDepthShorthand,
		DefaultPathDepth,
		"allow to add subpath to search query, increasing this value to 2 would add the parent directory to the search",
	)

	return cmd
}

func matchCmd(cmd *cobra.Command, args []string) error {
	depth, _ := common.LookupFlagInt(cmd, FlagPathDepth)

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

	for _, p := range filePaths {

		normalizedTerms := clean.Resolutions(clean.TokenizeAll(clean.SplitPath(clean.RemoveExtension(p), depth)))
		normalizedTerm := strings.Join(normalizedTerms, " ")

		common.Println(cmd, "Path:\n", p, "Search:\n", normalizedTerm)

		// for i, metric := range filter.Metrics {
		// 	distance, title, animeT, err := anidb.Search(normalizedTerm, metric)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	common.Printf(cmd, "%12s distance=%20s id=%5s : %s\n", filter.MetricNames[i], common.FormatFloat64(*distance), animeT.AID, title)
		// }

		distance, title, animeT, err := anidb.Search(normalizedTerm, filter.Metrics)
		if err != nil {
			return err
		}

		common.Printf(cmd, "%s\n%12s distance=%20s id=%6s\n", title, "Summary", common.FormatFloat64(*distance), animeT.AID)
	}

	return nil
}
