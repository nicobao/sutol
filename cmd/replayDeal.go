package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"

	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/ipfs/go-cid"
	"github.com/nicobao/sutol/internal"

	"github.com/spf13/cobra"
)

// replayDealCmd represents the replayDeal command
var replayDealCmd = &cobra.Command{
	Use:   "replay-deal [proposal-cid]",
	Short: "Replay the given deal",
	Long: `Replay (send again) the given deal. 

	Use 'list-deals' command to find the 'proposal-cid' positional argument.
	This command will use the default lotus wallet. 
	'DealStartEpoch' is set to -1.
	'FastRetrieval' is set to false.
	'ProviderCollateral is set to 0.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a proposal-cid argument")
		}
		_, err := cid.Decode(args[0])
		if err != nil {
			return errors.New("positional argument must be a valid CID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cidProposalToReplay, _ := cid.Decode(args[0])
		var api lotusapi.FullNodeStruct
		ctx := context.Background()
		headers, addr := internal.GetHeadersAndAddr(cmd)
		closer, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin", []interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
		defer closer()
		if err != nil {
			log.Fatalf("creating new merge client: %s", err)
		}

		dealInfo, err := api.ClientGetDealInfo(ctx, cidProposalToReplay)
		if err != nil {
			log.Fatalf("calling client get deal info: %s", err)
		}

		walletAddr, err := api.WalletDefaultAddress(ctx)
		if err != nil {
			log.Fatalf("error while fetching default wallet addr: %s", err)
		}

		// Actually replay the given deal
		sdParams := lotusapi.StartDealParams{
			Data:               dealInfo.DataRef,
			Wallet:             walletAddr, // cannot figure out where to find this info in previous deal - should be made configurable at least
			Miner:              dealInfo.Provider,
			EpochPrice:         dealInfo.PricePerEpoch,
			MinBlocksDuration:  dealInfo.Duration,  // not sure whether it's the minimum duration or the actual one?
			DealStartEpoch:     abi.ChainEpoch(-1), // cannot figure out where to find this info in previous deal - should be made configurable at least
			FastRetrieval:      false,              // cannot figure out where to find this info in previous deal - should be made configurable at least
			VerifiedDeal:       dealInfo.Verified,
			ProviderCollateral: big.NewInt(0), // cannot figure out where to find this info in previous deal - should be made configurable at least,
		}
		newDealInfo, err := api.ClientStartDeal(ctx, &sdParams)
		fmt.Println("Deal replayed", newDealInfo)

	},
}

func init() {
	rootCmd.AddCommand(replayDealCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replayDealCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replayDealCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
