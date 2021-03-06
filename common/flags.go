package common

import (
	"strconv"

	"github.com/spf13/cobra"
)

func LookupFlagInt(cmd *cobra.Command, flagName string) (int, error) {
	return strconv.Atoi(cmd.Flags().Lookup(flagName).Value.String())
}

func LookupFlagBool(cmd *cobra.Command, flagName string) (bool, error) {
	return strconv.ParseBool(cmd.Flags().Lookup(flagName).Value.String())
}
