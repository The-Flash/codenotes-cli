package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a note",
	Run:   func(cmd *cobra.Command, args []string) {},
}
