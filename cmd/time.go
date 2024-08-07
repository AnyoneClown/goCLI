package cmd

import (
	"time"
	"fmt"
	"github.com/spf13/cobra"
)

var TimeZone string

func getCurrentTime() string {
	return time.Now().Format(time.RFC1123)
}

func getTimeInZone(timezone string) (string, error) {
    location, err := time.LoadLocation(timezone)
    if err != nil {
        return "", err
    }
    return time.Now().In(location).Format(time.RFC1123), nil
}



var currentTimeCmd = &cobra.Command{
	Use:   "time",
	Short: "Gives current time",
	Long:  `This command provides the current time. If a timezone is specified, it will return the time in that timezone.`,
	Run: func(cmd *cobra.Command, args []string) {
		if TimeZone != "" {
			time, err := getTimeInZone(TimeZone)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("Current time in %s: %s\n", TimeZone, time)
		} else {
			fmt.Printf("Current time: %s\n", getCurrentTime())
		}
	},
}


func init() {
	rootCmd.AddCommand(currentTimeCmd)
	currentTimeCmd.Flags().StringVarP(&TimeZone, "timezone", "z", "", "Specify a timezone (e.g., 'America/New_York')")
}