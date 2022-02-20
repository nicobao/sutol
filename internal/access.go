package internal

import (
	"net/http"

	"github.com/spf13/cobra"
)

func GetHeadersAndAddr(cmd *cobra.Command) (http.Header, string) {
	token, _ := cmd.Flags().GetString("token")
	addr, _ := cmd.Flags().GetString("addr")
	headers := http.Header{"Authorization": []string{"Bearer " + token}}
	return headers, addr
}
