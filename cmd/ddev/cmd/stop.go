package cmd

import (
	"github.com/drud/ddev/pkg/ddevapp"
	"github.com/drud/ddev/pkg/util"
	"github.com/spf13/cobra"
)

// DdevStopCmd represents the stop command
var DdevStopCmd = &cobra.Command{
	Use:   "stop [sitename]",
	Short: "Stop the development environment for a site.",
	Long: `Stop the development environment for a site. You can run 'ddev stop'
from a site directory to stop that site, or you can specify a running site
to stop by running 'ddev stop <sitename>.`,
	Run: func(cmd *cobra.Command, args []string) {
		var siteName string

		if len(args) > 1 {
			util.Failed("Too many arguments provided. Please use 'ddev stop' or 'ddev stop [sitename]'")
		}

		if len(args) == 1 {
			siteName = args[0]
		}

		app, err := ddevapp.GetActiveApp(siteName)
		if err != nil {
			util.Failed("Failed to stop: %v", err)
		}

		err = app.Stop()
		if err != nil {
			util.Failed("Failed to stop containers for %s: %v", app.GetName(), err)
		}

		util.Success("Application has been stopped.")
	},
}

func init() {
	RootCmd.AddCommand(DdevStopCmd)
}
