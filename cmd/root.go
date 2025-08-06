package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	version = "v1.0.0"
)

var rootCmd = &cobra.Command{
	Use:   "choki",
	Short: "Advanced Code Snippet Manager write in  Go",
	Long: `CodeChoki V1.0.0 - Professional Code Snippet Manager
	Created by: MostafaSensei106
	GitHub: https://github.com/MostafaSensei106/CodeChoki`,
}

func init() {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("CodeChoki {{.Version}}\n")
	rootCmd.AddCommand(upgradeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("❌ Error: %v", err)
		os.Exit(1)
	}
}
