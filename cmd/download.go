package cmd

import "github.com/spf13/cobra"

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "lists all files in system",
	Long:  `queries the p2p network for all available options`,
	Run:   download,
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringP("path", "p", "~", "path of the file being uploaded")

}

func download(cmd *cobra.Command, arg []string) {
	filePath, _ := cmd.Flags().GetString("path")

	filePath += "" //lets compile
}
