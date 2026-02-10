package cmd

import (
	"errors"
	"fmt"

	dockscanner "github.com/arasdenizhan/dockenv/internal/scanner"

	"github.com/spf13/cobra"
)

var (
	imageFlagValue     string
	containerFlagValue string
)


var scanCmd = &cobra.Command{
    Use:   "scan",
    Short: "Scan docker target",
    PreRunE: func(cmd *cobra.Command, args []string) error {
		if imageFlagValue == "" && containerFlagValue == "" {
			return errors.New("you must provide either -i <image> or -c <container>")
		}

		if imageFlagValue != "" && containerFlagValue != "" {
			return errors.New("cannot use both -i and -c together")
		}

		return nil
	},
    Run: func(cmd *cobra.Command, args []string) {
		if imageFlagValue != "" {
			fmt.Println("Scanning IMAGE:", imageFlagValue)
			dockscanner.ScanTarget(imageFlagValue, true)
			return
		}

		if containerFlagValue != "" {
			fmt.Println("Scanning CONTAINER:", containerFlagValue)
			dockscanner.ScanTarget(containerFlagValue, false)
			return
		}
    },
}

func init() {
    rootCmd.AddCommand(scanCmd)
    scanCmd.Flags().StringVarP(&imageFlagValue, "image", "i", "", "scan docker image")
	scanCmd.Flags().StringVarP(&containerFlagValue, "container", "c", "", "scan docker container")
}
