package cmd

import (
	"errors"
	"os"

	"github.com/happierall/l"
	"github.com/kodeart/warframe-acolytes/internal"
	"github.com/spf13/cobra"
)

var (
	refresh        uint
	silent, notify bool
)

// track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Scans for Warframe acolytes appearance",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := internal.NewTracker(refresh, silent, notify)
		if err != nil {
			l.Errorf("%s", err)
			os.Exit(1)
		}

		t.Run()
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if refresh < 30 {
			return errors.New("minimum refresh rate is 30 seconds")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
	trackCmd.Flags().UintVarP(&refresh, "refresh", "r", 30, "Number of seconds for the world-state refresh\nminimum 30s allowed")
	trackCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Do not beep on alert\n(ignored if notification is off)")
	trackCmd.Flags().BoolVarP(&notify, "notify", "n", false, "Sends a desktop notification (default off)")
}
