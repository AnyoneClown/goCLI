package cmd

import (
	"time"
	"fmt"
	"github.com/spf13/cobra"
)

func getCurrentTime() string {
	return time.Now().Format(time.RFC1123)
}


var currentTime = &cobra.Command{
	Use:   "time",
	Short: "Gives current time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getCurrentTime())
	},
}


func init() {
	rootCmd.AddCommand(currentTime)
}


