package resources

import (
	"bytes"
	"os"
	"text/template"
)

type Template struct {
	Mode   uint32
	Source string
	Dest   string
	Owner  string
	Group  string
	Vars   map[string]interface{}
}

func NewTemplate() *Template {
	return &Template{}
}

func (t *Template) SetMode(m uint32) *Template {
	t.Mode = m
	return t
}

func (t *Template) SetDest(d string) *Template {
	t.Dest = d
	return t
}

func (t *Template) SetOwner(o string) *Template {
	t.Owner = o
	return t
}

func (t *Template) SetGroup(g string) *Template {
	t.Group = g
	return t
}

func (t *Template) SetSource(s string) *Template {
	t.Source = s
	return t
}

func (t *Template) Apply() error {
	_, err := os.Stat(t.Dest)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return t.Create()
}

func (t *Template) Create() error {
	templ, err := template.ParseFiles(t.Source)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := templ.Execute(&buf, t.Vars); err != nil {
		return err
	}

	f := File{
		Mode:    t.Mode,
		Content: string(buf.Bytes()),
		Path:    t.Dest,
		Owner:   t.Owner,
		Group:   t.Group,
	}

	if err := f.Apply(); err != nil {
		return err
	}

	return nil

}

func (t *Template) Remove() error {
	return nil
}
