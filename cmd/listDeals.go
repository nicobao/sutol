/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/nicobao/sutol/internal"
	"github.com/spf13/cobra"
)

// listDealsCmd represents the listDeals command
var listDealsCmd = &cobra.Command{
	Use:   "list-deals",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, url = internal.LoadTokenAndUrl(cmd)
		fmt.Println("listDeals called", token, url)
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
