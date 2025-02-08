package cmd

import "github.com/spf13/cobra"

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "lists all files in system",
	Long:  `queries the p2p network for all available options`,
	Run:   upload,
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringP("path", "p", "~", "path of the file being uploaded")

}

func upload(cmd *cobra.Command, arg []string) {
	filePath, _ := cmd.Flags().GetString("path")

	filePath += "" //lets compile
}
