package cmd

import (
	"os"
	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "goCLI",
	Short: "Upplication that allows you to make fast copy of SCP from windows to remote server",
	Long: `This application allows you to make fast and efficient copies of files
from a Windows machine to a remote server using SCP (Secure Copy Protocol).
It simplifies the process of transferring files securely over the network.`,
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


