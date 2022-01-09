package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "animatch [subcommand]",
		Long: `
animatch is an anime video file tagging tool.

It uses the provided file name as reference. File names are cleaned, 
tokenized and normalized in order before comparing them to the 
(also normalized) list of anime titles provided by AniDB.

The main command is:

    animatch tag [file|directory]

which appends an [anidb-12345] tag to the filename.

The Plex Hama.bundle plugin is then able to extract such tags on 
its own and propertly fetch the corresponding meta data.
Another great tagging tool that might help is FileBot 
which can have either the --q flag added in a command line interface
environment or instead the Quary Expression set in the FileBot Node 
user interface.
Any of those two properties may be set to the following value in 
order to extract the anime id and directly query the AniDB database:

    {(fn =~ /\[anidb-(\d+)\]/)[0][1]}

Use the following command to add auto completions for your terminal:

	source < $(animatch completion $(basename $SHELL))
`,
		Args: cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(NewSearchCmd())
	cmd.AddCommand(NewMatchCmd())
	cmd.AddCommand(NewCacheCmd())
	cmd.AddCommand(NewTagCmd())
	cmd.AddCommand(NewMXattrCmd())

	//cmd.PersistentFlags().Bool("apply", false, "set this flag in order to acually apply changes (renaming of files, etc.)")

	return cmd
}
