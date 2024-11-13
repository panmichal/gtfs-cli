/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/panmichal/gtfs-cli/gtfs"
	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Shows basic information about GTFS data",
	Long:  `Shows basic information about GTFS data`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires GTFS path or URL")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gtfsPath := args[0]
		files := gtfs.Parse(gtfsPath)
		for files.RouteFile.Scan() {
			fmt.Println(files.RouteFile.Text())
		}
		gtfs.CreateFeed(files)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
