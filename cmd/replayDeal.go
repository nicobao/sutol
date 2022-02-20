package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// replayDealCmd represents the replayDeal command
var replayDealCmd = &cobra.Command{
	Use:   "replay-deal",
	Short: "Replay the given deal",
	Long:  `Replay (send again) the given deal`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("replayDeal called")
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
