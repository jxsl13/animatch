package cmd

import (
	"github.com/jxsl13/animatch/anidb"
	"github.com/jxsl13/animatch/common"
	"github.com/spf13/cobra"
)

func NewSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search [text|file|folder|folder/*.mkv]",
		Short: "allows to search for the provided anime name.",
		RunE:  searchCmd,
		Args:  cobra.MinimumNArgs(1),
	}

	// modify public variable of anydb package in case that this flag is set
	cmd.Flags().IntVarP(&anidb.MatchDistanceUpperBound, "distance", "d", 5, "increase this distance in order  to allow a broader range of matches")

	return cmd
}

func searchCmd(cmd *cobra.Command, args []string) error {
	title, anime, err := anidb.BestMatch(args)
	if err != nil {
		return err
	}
	return common.Printf("%s: %s\n", title, anime.AID)
}
