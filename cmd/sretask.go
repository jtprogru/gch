/*
Copyright © 2023 Michael <jtprogru@gmail> Savin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

var (
	// sretaskCmd represents the sretask command
	sretaskCmd = &cobra.Command{
		Use:   "sretask",
		Short: "Create template with SRE task",
		Long:  `Create template with SRE task for interview.`,
		Run: func(cmd *cobra.Command, args []string) {
			if title == "" {
				panic("Title is not defined")
			}
			fileName := fmt.Sprintf("%v/Tasker - %v.md", tasksPath, title)
			_, err := os.Stat(fileName)
			if os.IsNotExist(err) {
				rendertask(title, fileName)
			} else {
				fmt.Printf("File %v exist\nPlease use another name of SRE Task\n", fileName)
			}

		},
	}
	title        string
	tasksPath    string
	taskTemplate = `---
id: {{ .Id }}
creation_date: {{ .CreationDate }}
modification_date: {{ .ModificationDate }}
type: simple_note
tags:
- tasker
---

# Tasker - {{ .Title }}

## Questions

<- write your question on this place ->

## Answer

<- write your answer on this place ->

***`
)

func init() {
	rootCmd.AddCommand(sretaskCmd)
	sretaskCmd.Flags().StringVarP(&title, "title", "T", "", "Title for SRE Task")
	sretaskCmd.Flags().StringVarP(&tasksPath, "path", "P", "./", "Path to repo with all tasks")
}

type taskTmpl struct {
	Id               string
	CreationDate     string
	ModificationDate string
	Title            string
}

func rendertask(t string, fileName string) {
	var timeFormat = "2006-01-02T15:04:05"
	var f *os.File
	var err error
	tmpl, err := template.New("").Parse(taskTemplate)
	if err != nil {
		fmt.Printf("Open template error: %v", err)
		panic(err)
	}
	task := taskTmpl{}

	task.Id = GetUUID()
	task.CreationDate = time.Now().Format(timeFormat)
	task.ModificationDate = time.Now().Format(timeFormat)
	task.Title = t

	f, err = os.Create(fileName)
	if err != nil {
		fmt.Printf("Create file error: %v\n", err)
		panic(err)
	}

	err = tmpl.Execute(f, task)
	if err != nil {
		fmt.Printf("Execute template error: %v\n", err)
		panic(err)
	}

	err = f.Close()
	if err != nil {
		fmt.Printf("Close file error: %v\n", err)
		panic(err)
	}

	fmt.Println("Please open file for edit task:")
	fmt.Printf("%v\n", fileName)

}
