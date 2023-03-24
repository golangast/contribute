/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/golangast/contribute/generateutility/gentemplates/body"
	"github.com/golangast/contribute/generateutility/gentemplates/routes"
	"github.com/golangast/contribute/generateutility/genutility"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:           "gen",
	Short:         "generate the file",
	Long:          ``,
	Example:       `go run main.go gen --route "newtestt"`,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		var route string
		route, _ = cmd.Flags().GetString("route")

		//update router with route
		found, err := genutility.FindText("router/router.go", `e.GET("/`+route+`", `+route+`.`+cases.Title(language.Und, cases.NoLower).String(route)+`)`)
		if err != nil {
			fmt.Print(err)
		}
		if !found {
			err := genutility.UpdateText("router/router.go", "//#routes", `e.GET("/`+route+`", `+route+`.`+cases.Title(language.Und, cases.NoLower).String(route)+`)`+"\n"+`//#routes`)
			if err != nil {
				fmt.Print(err)
			}
		}
		//update router with route
		found, err = genutility.FindText("router/router.go", `github.com/golangast/contribute/handler/get/`+route)
		if err != nil {
			fmt.Print(err)
		}
		if !found {
			err := genutility.UpdateText("router/router.go", "//#import", `"github.com/golangast/contribute/handler/get/`+route+`"`+"\n"+`//#import`)
			if err != nil {
				fmt.Print(err)
			}
		}
		//make folder for handler
		err = genutility.Makefolder("handler/get/" + route)
		if err != nil {
			fmt.Print(err)
		}
		err = genutility.Makefolder("handler/get/" + route)
		if err != nil {
			fmt.Print(err)
		}
		//make route file
		routefile, err := genutility.Makefile("handler/get/" + route + "/" + route + ".go")
		if err != nil {
			fmt.Print(err, routefile)
		}
		//make html file
		htmlfile, err := genutility.Makefile("server/templates/" + route + ".html")
		if err != nil {
			fmt.Print(err, htmlfile)
		}
		//write it to the handler file
		r := make(map[string]string)
		r["routes"] = route
		err = genutility.Writetemplate(routes.Routetemp, routefile, r)
		if err != nil {
			fmt.Print(err)
		}
		//write it to the html file
		err = genutility.Writetemplate(body.Bodytemp, htmlfile, r)
		if err != nil {
			fmt.Print(err)
		}
		err, out, errout := genutility.Shellout(`go mod tidy && go mod vendor && go build`)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println(out)
		fmt.Println("--- errs ---")
		fmt.Println(errout)

	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringP("route", "r", "", "Set your route")

}
