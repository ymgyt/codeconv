package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	loggingLevel int
)

var rootCmd = &cobra.Command{
	Use:   "codeconv",
	Short: "codeconv encode/decode supported code",
	Long: `codeconv convert(encode/decode) supported code.

USAGE:
	codeconv <CODE> <encode|decode> [ARGS] [OPTIONS]

CODES:
  - base64
  - url
  - jwt`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.codeconv.yaml)")
	rootCmd.PersistentFlags().IntVarP(&loggingLevel, "log", "v", 0, "logging level")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
