package cmd

import (
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade CodeChoki to the latest version",
	Long:  "Pull the latest changes from GitHub and upgrade CodeChoki to the latest available version.",
	Run: func(cmd *cobra.Command, args []string) {
		//upgrade.UpgradeCodeChoki()
	},
}
