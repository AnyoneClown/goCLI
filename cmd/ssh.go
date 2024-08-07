package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	keyFile string
	hostAddress string
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Connect to EC2 instance via SSH",
	Long:  `This command connects to the specified EC2 instance using SSH.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connecting to EC2 instance...")

		sshCommand := exec.Command("ssh", "-i", keyFile, hostAddress)
		
		sshCommand.Stdin = os.Stdin
		sshCommand.Stdout = os.Stdout
		sshCommand.Stderr = os.Stderr

		err := sshCommand.Run()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)

	sshCmd.Flags().StringVarP(&keyFile, "key", "k", "/c/Users/dumde/TBandKF.pem", "Path to the SSH key file")
	sshCmd.Flags().StringVarP(&hostAddress, "hostadress", "a", "ubuntu@ec2-52-21-129-119.compute-1.amazonaws.com", "Host to connect to")
}