/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/spf13/cobra"
)

// secondCmd represents the second command
var secondCmd = &cobra.Command{
	Use:   "second",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("second called")

	},
}

func init() {
	rootCmd.AddCommand(secondCmd)

}
