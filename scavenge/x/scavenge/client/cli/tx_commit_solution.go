package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"scavenge/x/scavenge/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCommitSolution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-solution [solution-hash] [solution-scavenger-hash]",
		Short: "Broadcast message commit-solution",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			solution := args[0]

			solutionHash := sha256.Sum256([]byte(solution))

			solutionHashStr := hex.EncodeToString(solutionHash[:])

			scavenger := clientCtx.GetFromAddress().String()

			solutionScavengerHash := sha256.Sum256([]byte(solutionHashStr + scavenger))

			solutionScavengerHashStr := hex.EncodeToString(solutionScavengerHash[:])

			msg := types.NewMsgCommitSolution(
				clientCtx.GetFromAddress().String(),
				string(solutionHashStr),
				string(solutionScavengerHashStr),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
