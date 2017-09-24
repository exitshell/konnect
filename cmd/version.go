package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCmd - Show detailed version information.
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show detailed version information",
	Long:  "Show detailed version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Konnect - version %v\n", getVersion())
	},
}
