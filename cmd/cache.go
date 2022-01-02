package cmd

import (
	"github.com/jxsl13/animatch/anidb"
	"github.com/jxsl13/animatch/common"
	"github.com/spf13/cobra"
)

const (
	ErrAPI = common.Error("api error")
)

var (
	wrapErrApi = common.WrapErr(ErrAPI)
)

func NewCacheCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cache",
		Short: "interacting with the title cache",
		Args:  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(NewCacheUpdateCmd())
	return cmd
}

func NewCacheUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update the anime title cache",
		Args:  cobra.NoArgs,
		RunE:  updateCacheFromAPI,
	}

	return cmd
}

func updateCacheFromAPI(cmd *cobra.Command, args []string) error {
	c, err := anidb.DefaultTitlesCache()
	if err != nil {
		return wrapErrApi(err)
	}
	if _, err := c.GetFreshTitles(); err != nil {
		return wrapErrApi(err)
	}
	if err := c.Save(); err != nil {
		return wrapErrApi(err)
	}
	return common.Println(cmd, "anime title cache was updated successfully")
}
