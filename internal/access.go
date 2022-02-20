package internal

import (
	"github.com/spf13/cobra"
)

func LoadTokenAndAddr(cmd *cobra.Command) (string, string) {
	// Load data from
	// Returns <token>, <addr>   example: KSEKE9329ejd http:localhost:1234
	token, _ := cmd.Flags().GetString("token")
	addr, _ := cmd.Flags().GetString("addr")
	return token, addr
}
