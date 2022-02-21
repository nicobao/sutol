package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/ipfs/go-cid"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// replayDealCmd represents the replayDeal command

var (
	fastRetrieval      bool
	providerCollateral int64
	dealStartEpoch     int64
	replayDealCmd      = &cobra.Command{
		Use:   "replay-deal [proposal-cid]",
		Short: "Replay the given deal",
		Long: `Replay (send again) the given deal. 

		Use 'list-deals' command to find the 'proposal-cid' positional argument.
		This command will use the default lotus wallet.`,
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
			// fastRetrieval, _ := cmd.Flags().GetBool("fast-retrieval")
			// providerCollateral, _ := cmd.Flags().GetInt64("provider-collateral")
			// dealStartEpoch, _ := cmd.Flags().GetInt64("deal-start-epoch")

			cidProposalToReplay, _ := cid.Decode(args[0])

			fmt.Println("ProviderCollateral", providerCollateral, "FastRetrieval", fastRetrieval, "dealStartEpoch", dealStartEpoch)
			var api lotusapi.FullNodeStruct
			ctx := context.Background()
			headers := http.Header{"Authorization": []string{"Bearer " + token}}
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
				MinBlocksDuration:  dealInfo.Duration, // not sure whether it's the minimum duration or the actual one?
				DealStartEpoch:     abi.ChainEpoch(dealStartEpoch),
				FastRetrieval:      fastRetrieval,
				VerifiedDeal:       dealInfo.Verified,
				ProviderCollateral: big.NewInt(providerCollateral),
			}
			newDealInfo, err := api.ClientStartDeal(ctx, &sdParams)
			fmt.Println("Deal replayed", newDealInfo)

		},
	}
)

func init() {
	rootCmd.AddCommand(replayDealCmd)
	initConfigAndPrint(false) // TODO fix that,  see https://github.com/spf13/cobra/issues/1176
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replayDealCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	fastRetrievalStr := "fast-retrieval"
	replayDealCmd.Flags().BoolVarP(&fastRetrieval, fastRetrievalStr, "f", false, "fast retrieval (default to SUTOL_FAST_RETRIEVAL else config file else false)")
	viper.BindPFlag(fastRetrievalStr, replayDealCmd.Flags().Lookup("fast-retrieval"))
	viper.SetDefault(fastRetrievalStr, false)
	fastRetrieval = viper.GetBool(fastRetrievalStr)

	providerCollateralStr := "provider-collateral"
	replayDealCmd.Flags().Int64VarP(&providerCollateral, providerCollateralStr, "p", 0, "provider collateral (default to SUTOL_PROVIDER_COLLATERAL else config file else 0)")
	viper.BindPFlag(providerCollateralStr, replayDealCmd.Flags().Lookup(providerCollateralStr))
	viper.SetDefault(providerCollateralStr, 0)
	providerCollateral = viper.GetInt64(providerCollateralStr)

	dealStartEpochStr := "deal-start-epoch"
	replayDealCmd.Flags().Int64VarP(&dealStartEpoch, dealStartEpochStr, "d", -1, "deal start epoch (default to SUTOL_DEAL_START_EPOCH else config file else -1)")
	viper.BindPFlag(dealStartEpochStr, replayDealCmd.Flags().Lookup(dealStartEpochStr))
	viper.SetDefault(dealStartEpochStr, -1)
	dealStartEpoch = viper.GetInt64(dealStartEpochStr)

}
