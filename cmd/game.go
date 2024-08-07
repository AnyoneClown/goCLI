package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var pathForGame = map[string]string{
    "dota2": "C:/Program Files (x86)/Steam/steamapps/common/dota 2 beta/game/bin/win64/dota2.exe",
}


var gameCmd = &cobra.Command{
	Use:   "game [game to play]",
	Short: "Start game from steam",
	Long:  `This command starts game from steam`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gameName := args[0]
		gamePath, exists := pathForGame[gameName]
		if !exists {
			fmt.Printf("Error: Unknown game '%s'\n", gameName)
			return
		}

		fmt.Printf("Starting %s...\n", gameName)

		gameCommand := exec.Command(gamePath)
		
		err := gameCommand.Start()
		if err != nil {
			fmt.Printf("Error starting game: %s\n", err)
			return
		}

		fmt.Printf("%s has been launched. This window will close in 5 seconds.\n", gameName)
		time.Sleep(5 * time.Second)
	},
}

func init() {
	rootCmd.AddCommand(gameCmd)

}
