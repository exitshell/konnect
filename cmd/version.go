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
		fmt.Println("Konnect:")
		fmt.Printf("  version %v\n", AppVersion)
		fmt.Printf("  build   %v\n", AppBuild)
		fmt.Printf("  date    %v\n", AppDate)
	},
}
