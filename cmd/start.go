/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	headerfooter "github.com/golangast/contribute/generateutility/gentemplates/headerfootertemp"
	"github.com/golangast/contribute/generateutility/gentemplates/hometemp"
	"github.com/golangast/contribute/generateutility/gentemplates/router"
	"github.com/golangast/contribute/generateutility/gentemplates/server"
	"github.com/golangast/contribute/generateutility/genutility"
	"github.com/golangast/contribute/generateutility/homeroute"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "to generate the folders and files run ./contribute start",
	Example: `./contribute start`,
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		/*
		   makes folders

		*/
		err := genutility.Makefolder("handler/get/home")
		if err != nil {
			fmt.Print(err)
		}

		err = genutility.Makefolder("router")
		if err != nil {
			fmt.Print(err)
		}

		err = genutility.Makefolder("templates")
		if err != nil {
			fmt.Print(err)
		}
		/*
		   makes files

		*/
		routefile, err := genutility.Makefile("router/router.go")
		if err != nil {
			fmt.Print(err, routefile)
		}
		mainfile, err := genutility.Makefile("main.go")
		if err != nil {
			fmt.Print(err, routefile)
		}
		homefile, err := genutility.Makefile("handler/get/home/home.go")
		if err != nil {
			fmt.Print(err, routefile)
		}
		//make html file
		homehtmlfile, err := genutility.Makefile("templates/home.html")
		if err != nil {
			fmt.Print(err, homehtmlfile)
		}
		//make html file
		footerfile, err := genutility.Makefile("templates/footer.html")
		if err != nil {
			fmt.Print(err, homehtmlfile)
		}
		//make html file
		headerfile, err := genutility.Makefile("templates/header.html")
		if err != nil {
			fmt.Print(err, homehtmlfile)
		}
		/*
		   generate code in those files

		*/
		m := make(map[string]string)
		header := fmt.Sprintf(`{{define "header"}}%s`, "")
		end := fmt.Sprintf(`{{end}}%s`, "")
		m["header"] = header
		m["end"] = end

		//footer map for {{define "footer"}} {{end}}
		mf := make(map[string]string)
		footer := fmt.Sprintf(`{{define "footer"}}%s`, "")
		endf := fmt.Sprintf(`{{end}}%s`, "")
		mf["footer"] = footer
		mf["end"] = endf

		//footer/header map for {{template "footer" .}} {{end}}
		mb := make(map[string]string)
		headerb := fmt.Sprintf(`{{template "header" .}}%s`, "")
		footerb := fmt.Sprintf(`{{template "footer" .}}%s`, "")
		mb["footer"] = footerb
		mb["header"] = headerb
		err = genutility.Writetemplate(router.Routertemp, routefile, nil)
		if err != nil {
			fmt.Print(err)
		}
		err = genutility.Writetemplate(homeroute.Homeroutetemp, homefile, nil)
		if err != nil {
			fmt.Print(err)
		}
		err = genutility.Writetemplate(server.Servertemp, mainfile, nil)
		if err != nil {
			fmt.Print(err)
		}
		//write it to the html file
		err = genutility.Writetemplate(hometemp.Hometemp, homehtmlfile, mb)
		if err != nil {
			fmt.Print(err)
		}
		//write it to the html file
		err = genutility.Writetemplate(headerfooter.Headertemp, headerfile, m)
		if err != nil {
			fmt.Print(err)
		}
		//write it to the html file
		err = genutility.Writetemplate(headerfooter.Footertemp, footerfile, mf)
		if err != nil {
			fmt.Print(err)
		}
		err, out, errout := genutility.Shellout(`go mod init "contribute" && go mod tidy && go mod vendor`)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println(out)
		fmt.Println("--- errs ---")
		fmt.Println(errout)

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
