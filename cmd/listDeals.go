package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"

	"github.com/spf13/cobra"
)

// listDealsCmd represents the listDeals command
var listDealsCmd = &cobra.Command{
	Use:   "list-deals",
	Short: "List storage market deals (same as Lotus)",
	Long:  `List deals already sent by this lotus instance`,
	Run: func(cmd *cobra.Command, args []string) {
		var api lotusapi.FullNodeStruct
		headers := http.Header{"Authorization": []string{"Bearer " + token}}
		closer, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin", []interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
		if err != nil {
			log.Fatalf("creating new merge client: %s", err)
		}
		defer closer()

		dealInfo, err := api.ClientListDeals(context.Background())
		if err != nil {
			log.Fatalf("calling client list deals: %s", err)
		}
		// TODO: refine pretty-printing of these values
		fmt.Println("List of deals already sent using this instance of lotus:")
		for i := 0; i < len(dealInfo); i++ {
			fmt.Println(dealInfo[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(listDealsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listDealsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listDealsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
