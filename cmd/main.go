package main

import (
	"log"

	"github.com/hooksie1/gofigure"
	"github.com/hooksie1/gofigure/resources"
)

func main() {
	file := &resources.File{
		Path:    "./test.txt",
		Mode:    0644,
		Content: "This is a test\n",
		Owner:   "johnhooks",
		Group:   "staff",
	}

	if err := gofigure.Exists(file); err != nil {
		log.Println(err)
	}

	t := resources.NewTemplate().SetSource("./thing.tmpl").SetDest("./thing.txt").SetMode(0644).SetOwner("johnhooks").SetGroup("staff")

	t.Vars = map[string]interface{}{
		"name": "john",
	}

	if err := gofigure.Exists(t); err != nil {
		log.Println(err)
	}

}
