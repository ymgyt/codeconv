package cmd

import (
	stdbase64 "encoding/base64"

	"github.com/spf13/cobra"

	"github.com/ymgyt/codeconv/internal/base64"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:     "base64",
	Aliases: []string{"b64"},
	Short:   "base64 convert base64 format",
	Long: `base64 convert base64 format

USAGE
	codeconv base64 <encode|decode> [ARGS] [OPTIONS]`,
}

func buildBase64Converter() *base64.Converter {
	return &base64.Converter{Opt: &base64.Options{
		Encoding: stdbase64.StdEncoding,
	}}
}

func init() {
	rootCmd.AddCommand(base64Cmd)
}
