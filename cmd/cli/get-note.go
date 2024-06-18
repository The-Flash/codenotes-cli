package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single note",
	Run:   func(cmd *cobra.Command, args []string) {},
}
