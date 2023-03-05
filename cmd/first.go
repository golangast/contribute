/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// firstCmd represents the first command
var firstCmd = &cobra.Command{
	Use:   "first",
	Short: "A brief description of your command",
	Long:  `print stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("first called")
		fmt.Println("stuff")

	},
}

func init() {
	rootCmd.AddCommand(firstCmd)
}
