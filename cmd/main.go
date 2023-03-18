package main

import (
	"log"
	"os"

	"github.com/hooksie1/gofigure"
	"github.com/hooksie1/gofigure/resources"
)

func main() {
	file := &resources.File{
		Path:    "./test.txt",
		Mode:    0644,
		Content: []byte("This is a test"),
		Owner:   "jhooks",
		Group:   "jhooks",
	}

	if err := gofigure.Exists(file); err != nil {
		log.Println(err)
	}

	tmpl, err := os.Open("thing.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	t := resources.NewTemplate().SetSource(tmpl).SetDest("./thing.sh").SetMode(0755).SetOwner("jhooks").
		SetGroup("jhooks").SetVars(map[string]any{
		"name": "john",
	})

	if err := gofigure.Exists(t); err != nil {
		log.Println(err)
	}

	script := resources.NewCommand("./thing.sh").SetOutput(os.Stdout)
	cmd := resources.NewCommand("echo").SetArgs("script", "ran", "successfully").SetOutput(os.Stdout)

	if err := gofigure.Exists(script); err != nil {
		log.Println(err)
	}

	if err := gofigure.Exists(cmd); err != nil {
		log.Println(err)
	}

	timer := resources.NewSystemdUnit().SetDescription("this is a test").SetName("test").SetType(resources.Timer).SetSchedule("*-*-*-*")

	if err := gofigure.Exists(timer); err != nil {
		log.Println(err)
	}

}
