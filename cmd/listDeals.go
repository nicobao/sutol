/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/nicobao/sutol/internal"
	"github.com/spf13/cobra"
)

// listDealsCmd represents the listDeals command
var listDealsCmd = &cobra.Command{
	Use:   "list-deals",
	Short: "List storage market deals (same as Lotus)",
	Long:  `List deals already sent by this lotus instance`,
	Run: func(cmd *cobra.Command, args []string) {
		token, addr = internal.LoadTokenAndAddr(cmd)
		fmt.Println("listDeals called", token, addr)

		headers := http.Header{"Authorization": []string{"Bearer " + token}}
		var api lotusapi.FullNodeStruct
		closer, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin", []interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
		defer closer()

		tipset, err := api.ChainHead(context.Background())
		if err != nil {
			log.Fatalf("calling chain head: %s", err)
		}
		fmt.Printf("Current chain head is: %s", tipset.String())
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
