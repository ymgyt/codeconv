package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// base64Encode.goCmd represents the base64Encode.go command
var base64EncodeCmd = &cobra.Command{
	Use:     "encode",
	Aliases: []string{"en"},
	Short:   "encode given data to base64 format",
	Run:     base64EncodeRun,
}

func base64EncodeRun(cmd *cobra.Command, args []string) {
	converter := buildBase64Converter()

	var reader io.Reader
	var writer = os.Stdout
	switch len(args) {
	case 0:
		reader = os.Stdin
	case 1:
		reader = strings.NewReader(args[0])
		if args[0] == "-" {
			reader = os.Stdin
		}
	default:
		fmt.Fprintf(os.Stderr, "multi args not implemented yet")
		os.Exit(1)
	}

	err := converter.Encode(writer, reader)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	base64Cmd.AddCommand(base64EncodeCmd)
}
