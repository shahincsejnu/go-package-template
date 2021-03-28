package main

import (
	"fmt"
	"os"
	"text/template"
)

type Todo struct {
	Name string
	Description string
}

func main() {
	td := Todo{
		Name:        "Test templates",
		Description: "Let's test a template to see the magic.",
	}

	td2 := Todo{
		Name:        "Go",
		Description: "Contribute to any Go project",
	}

	t, err := template.New("todos").Parse("You have a task named \"{{ .Name }}\" with description: \"{{ .Description }}\"\n")
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, td)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, td2)
	if err != nil {
		fmt.Println(err)
	}
}