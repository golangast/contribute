package genutility

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func FindText(p, str string) (bool, error) {
	// Open file for reading.
	var file, err = os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)
		if err != nil {
			return false, err
		}
		if strings.Contains(string(text), str) {
			return true, nil

		}
		// Break if finally arrived at end of file
		if err == io.EOF {
			if err != nil {
				return false, err
			}
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			if err != nil {
				return false, err
			}
			fmt.Println(err)

		}
	}

	file.Close()
	return false, nil
}

//f is for file, o is for old text, n is for new text
func UpdateText(f string, o string, n string) error {
	input, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)

	if err = ioutil.WriteFile(f, output, 0666); err != nil {
		fmt.Println(err)
	}

	return nil
}

// make any folder
func Makefolder(p string) error {
	if err := os.MkdirAll(p, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("~~~~could not create"+p, err)
	} else {
		fmt.Println("Directory " + p + " successfully created with sticky bits and full permissions")
	}
	return nil
}

// make any file
func Makefile(p string) (*os.File, error) {
	file, err := os.Create(p)
	if err != nil {
		return file, err
	}
	return file, nil
}

// error checker
func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// write any template to file
func Writetemplate(temp string, f *os.File, d map[string]string) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("queue").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}
func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
