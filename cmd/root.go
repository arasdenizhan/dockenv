package cmd

import (
	"os"

	"github.com/arasdenizhan/dockenv/internal/output"

	"github.com/spf13/cobra"
)

var version = "dev"
var noColor bool

var rootCmd = &cobra.Command{
    Use:   "dockenv",
    Short: "Docker Environment Security Scanner",
    Long:  "dockenv scans docker containers and images for risky environment variables",
    Version: version,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable colored output")
}

func initConfig() {
    output.NoColor = noColor
}
