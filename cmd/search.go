package cmd

import (
	"strings"

	"github.com/jxsl13/animatch/anidb"
	"github.com/jxsl13/animatch/clean"
	"github.com/jxsl13/animatch/common"
	"github.com/jxsl13/animatch/filter"
	"github.com/spf13/cobra"
)

const (
	FlagDistance          = "distance"
	FlagDistanceShorthand = "d"
)

func NewSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search [text|file|folder|folder/*.mkv]",
		Short: "allows to search for the provided anime name.",
		RunE:  searchCmd,
		Args:  cobra.MinimumNArgs(1),
	}

	// modify public variable of anydb package in case that this flag is set
	cmd.Flags().IntP(
		FlagDistance,
		FlagDistanceShorthand,
		anidb.DefaultMatchDistanceUpperBound,
		"increase this distance in order to allow a broader range of matches",
	)

	return cmd
}

func searchCmd(cmd *cobra.Command, args []string) error {

	terms := clean.NormalizeAll(clean.TokenizeAll(args))
	normalizedTerm := strings.Join(terms, " ")

	for _, metric := range filter.Metrics {
		distance, title, animeT, err := anidb.Search(normalizedTerm, metric)
		if err != nil {
			return err
		}
		common.Printf(cmd, "%s: %s[distance=%s]\n", title, animeT.AID, common.FormatFloat64(*distance))
	}
	return nil
}
