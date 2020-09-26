package cmd

import (
	"strings"

	"github.com/bubblyworld/logos/ops"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists supported formal logics",
	Long: `Lists the formal logics that logos can manipulate, along with their internal
identifiers.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, l := range ops.SupportedLogics() {
			infoln(l.Name())
			printf("  ID:          ")
			println(l.ID())
			printf("  Description: ")
			println(prefixLines(l.Description(), "               "))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func prefixLines(msg, prefix string) string {
	ll := strings.Split(msg, "\n")
	for i := range ll {
		if i == 0 {
			continue
		}

		ll[i] = prefix + ll[i]
	}

	return strings.Join(ll, "\n")
}
