package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "animatch [file|folder|folder/*.mkv]",
		Long: `
Use the following command to add auto completions for your terminal:

	source < $(animatch completion $(basename $SHELL))
`,
		Args: cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(NewSearchCmd())
	cmd.AddCommand(NewCacheCmd())

	//cmd.PersistentFlags().Bool("apply", false, "set this flag in order to acually apply changes (renaming of files, etc.)")

	return cmd
}
