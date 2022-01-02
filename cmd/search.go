package cmd

import (
	"github.com/jxsl13/animatch/anidb"
	"github.com/jxsl13/animatch/common"
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

	distance, err := common.LookupFlagInt(cmd, FlagDistance)
	if err != nil {
		return err
	}
	title, animeT, err := anidb.BestMatch(args, distance)
	if err != nil {
		return err
	}

	anime, err := anidb.MetaData(animeT.AID)
	if err != nil {
		return err
	}

	return common.Printf("%s: %s\n", title, anime)
}
