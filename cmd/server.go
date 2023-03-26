/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/golangast/contribute/generateutility/genutility"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "run",
	Short: "To run the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server is running now......")
		fmt.Println("Please visit http://localhost:5002/home")
		err, out, errout := genutility.Shellout(`go run main.go`)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println(out)
		fmt.Println("--- errs ---")
		fmt.Println(errout)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

}
