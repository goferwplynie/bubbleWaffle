package cmd

import (
	"fmt"

	"github.com/goferwplynie/bubbleWaffle/internal/creator"
	"github.com/spf13/cobra"
)

var (
	styleFile  bool
	bubbleZone bool
	keyBinds   bool
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "create bubbletea ui component",
	Long: `creates new bubbletea ui component if it's not present
	by deafult generates model.go, init.go, update.go, view.go
	there are optional style.go, keys.go and bubblezone version of view.go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("creating component '%s'", args[0])
		creator.CreateComponent(".", args[0], &creator.ComponentOptions{
			StyleFile:    styleFile,
			KeybindsFile: keyBinds,
			BubbleZone:   bubbleZone,
		})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolVarP(&keyBinds, "keys", "k", false, "generate keys file")
	addCmd.Flags().BoolVarP(&styleFile, "style", "s", false, "generate styles file")
	addCmd.Flags().BoolVarP(&bubbleZone, "bubblezone", "b", false, "generate view file with bubblezone")
}
