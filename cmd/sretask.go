package cmd

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/spf13/cobra"

	"github.com/jtprogru/gch/internal/uuids"
)

var (
	// sretaskCmd represents the sretask command.
	sretaskCmd = &cobra.Command{ //nolint:gochecknoglobals,nolintlint // This is normal.
		Use:   "sretask",
		Short: "Create template with SRE task",
		Long:  `Create template with SRE task for interview.`,
		Run: func(_ *cobra.Command, _ []string) {
			if title == "" {
				panic("Title is not defined")
			}
			fileName := fmt.Sprintf("%v/Tasker - %v.md", tasksPath, title)
			_, err := os.Stat(fileName)
			if os.IsNotExist(err) {
				rendertask(title, fileName)
			} else {
				_, _ = fmt.Printf("File %v exist\nPlease use another name of SRE Task\n", fileName) //nolint:errcheck,nolintlint // Ignore errors for test.
			}
		},
	}
	title     string //nolint:gochecknoglobals,nolintlint // This is normal.
	tasksPath string //nolint:gochecknoglobals,nolintlint // This is normal.
)

func init() { //nolint:gochecknoinits,nolintlint // Init func is needed for cobra.
	rootCmd.AddCommand(sretaskCmd)
	sretaskCmd.Flags().StringVarP(&title, "title", "T", "", "Title for SRE Task")
	sretaskCmd.Flags().StringVarP(&tasksPath, "path", "P", "./", "Path to repo with all tasks")
}

type taskTmpl struct {
	ID               string
	CreationDate     string
	ModificationDate string
	Title            string
}

func rendertask(title, fileName string) {
	taskTemplate := `---
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

	timeFormat := "2006-01-02T15:04:05"
	var f *os.File
	var err error
	tmpl, err := template.New("").Parse(taskTemplate)
	if err != nil {
		_, _ = fmt.Printf("Open template error: %v", err) //nolint:errcheck,nolintlint // Ignore errors for test.
		panic(err)
	}
	task := taskTmpl{}

	task.ID = uuids.GetUUID()
	task.CreationDate = time.Now().Format(timeFormat)
	task.ModificationDate = time.Now().Format(timeFormat)
	task.Title = title

	f, err = os.Create(fileName)
	if err != nil {
		_, _ = fmt.Printf("Create file error: %v\n", err) //nolint:errcheck,nolintlint // Ignore errors for test.
		panic(err)
	}

	err = tmpl.Execute(f, task)
	if err != nil {
		_, _ = fmt.Printf("Execute template error: %v\n", err) //nolint:errcheck,nolintlint // Ignore errors for test.
		panic(err)
	}

	err = f.Close()
	if err != nil {
		_, _ = fmt.Printf("Close file error: %v\n", err) //nolint:errcheck,nolintlint // Ignore errors for test.
		panic(err)
	}

	_, _ = fmt.Println("Please open file for edit task:") //nolint:errcheck,nolintlint // Ignore errors for test.
	_, _ = fmt.Printf("%v\n", fileName)                   //nolint:errcheck,nolintlint // Ignore errors for test.
}
