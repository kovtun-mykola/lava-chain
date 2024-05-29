package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lavanet/lava/x/rewards/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group rewards queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQueryPools())
	cmd.AddCommand(CmdQueryBlockReward())
	cmd.AddCommand(CmdQueryShowIprpcData())
	cmd.AddCommand(CmdQueryIprpcProviderRewardEstimation())
	cmd.AddCommand(CmdQueryIprpcSpecReward())
	cmd.AddCommand(CmdQueryProviderReward())
	cmd.AddCommand(CmdQueryGenerateIbcIprpcTx())
	// this line is used by starport scaffolding # 1

	return cmd
}
