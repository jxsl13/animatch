package cmd

import (
	"strings"

	"github.com/jxsl13/animatch/anidb"
	"github.com/jxsl13/animatch/clean"
	"github.com/jxsl13/animatch/common"
	"github.com/spf13/cobra"
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

	return cmd
}

func matchCmd(cmd *cobra.Command, args []string) error {
	terms := clean.Overlap(clean.NormalizeAll(clean.SplitPath(clean.RemoveExtension(strings.Join(args, " ")), 2)))

	distance, err := common.LookupFlagInt(cmd, FlagDistance)
	if err != nil {
		return err
	}
	title, animeT, err := anidb.BestMatch(terms, distance)
	if err != nil {
		return err
	}

	return common.Printf(cmd, "%s: %s\n", title, animeT.AID)
}
