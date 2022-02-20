package internal

import (
	"github.com/spf13/cobra"
)

func LoadTokenAndUrl(cmd *cobra.Command) (string, string) {
	// Load data from
	// Returns <token>, <url>   example: KSEKE9329ejd http:localhost:1234
	token, _ := cmd.Flags().GetString("token")
	url, _ := cmd.Flags().GetString("url")
	return token, url
}
