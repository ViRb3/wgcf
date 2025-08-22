package trace

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	. "github.com/ViRb3/wgcf/v2/cmd/shared"
	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
)

var shortMsg = "Prints trace information about the current internet connection"

var Cmd = &cobra.Command{
	Use:   "trace",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
Useful for verifying if Warp and Warp+ are working.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := trace(); err != nil {
			log.Fatalf("%+v\n", err)
		}
	},
}

func init() {
}

func trace() error {
	response, err := http.Get("https://cloudflare.com/cdn-cgi/trace")
	if err != nil {
		return errors.WithStack(err)
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.WithStack(err)
	}
	log.Println("Trace result:")
	fmt.Println(strings.TrimSpace(string(bodyBytes)))
	return nil
}
