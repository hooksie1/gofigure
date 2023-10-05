# gofigure

Go figure is a configuration package for Golang. I've always wanted to have something like Ansible but in a single binary that I could 
ship to servers. 


Gofigure has prebuilt resources for configuration. Currently this is limited to: `file`, `template`, `command`, `systemd`, and very limited support for `package`.

The Resource iterface is very simple, so creating your own Resources should be easy.

## Examples

The cmd directory has some examples, but here's an exmaple of creating a file:

```go
	file := &resources.File{
		Path:    "./test.txt",
		Mode:    0644,
		Content: []byte("This is a test"),
		Owner:   "someUser",
		Group:   "someGroup",
	}

	if err := gofigure.Exists(file); err != nil {
		log.Println(err)
	}
```

And here's an example of a template:

```go
	tmpl, err := os.Open("thing.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	t := resources.NewTemplate().SetSource(tmpl).SetDest("./thing.sh").SetMode(0755).SetOwner("someUser").
		SetGroup("someGroup").SetVars(map[string]any{
		"name": "John",
	})

	if err := gofigure.Exists(t); err != nil {
		log.Println(err)
	}
```
