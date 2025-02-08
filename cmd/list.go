package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all files in system",
	Long:  `queries the p2p network for all available options`,
	Run:   list,
}

func list(cmd *cobra.Command, arg []string) {

}
