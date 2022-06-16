package resources

import (
	"io"
	"os"
	"os/user"
	"strconv"
)

type File struct {
	Mode    uint32
	Content string
	Path    string
	Owner   string
	Group   string
}

func NewFile() *File {
	return &File{}
}

func (f *File) Apply() error {
	_, err := os.Stat(f.Path)
	if err == nil {
		return nil
	}

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return f.Create()
}

func (f *File) Create() error {
	file, err := os.Create(f.Path)
	if err != nil {
		return err
	}

	defer file.Close()

	if err := f.Write(file); err != nil {
		return err
	}

	return f.Chown()
}

func (f *File) Chown() error {
	usr, err := user.Lookup(f.Owner)
	if err != nil {
		return err
	}

	group, err := user.LookupGroup(f.Group)
	if err != nil {
		return err
	}

	uid, err := strconv.Atoi(usr.Uid)
	if err != nil {
		return err
	}

	gid, err := strconv.Atoi(group.Gid)
	if err != nil {
		return err
	}

	return os.Chown(f.Path, uid, gid)
}

func (f *File) Write(w io.Writer) error {
	if _, err := w.Write([]byte(f.Content)); err != nil {
		return err
	}
	return nil
}

func (f *File) Remove() error {
	return os.Remove(f.Path)
}
