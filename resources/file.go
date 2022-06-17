package resources

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
	"os/user"
	"strconv"
	"strings"
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
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	return f.CreateOrOpen()
}

func (f *File) CreateOrOpen() error {
	var file *os.File

	file, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE, os.FileMode(f.Mode))
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)

	_, err = io.Copy(buf, file)
	if err != nil {
		return err
	}

	same, err := CompareSums(buf, strings.NewReader(f.Content))
	if err != nil {
		return err
	}

	if same {
		return nil
	}

	defer file.Close()

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

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

func CompareSums(f, c io.Reader) (bool, error) {
	fileHash := sha256.New()
	contentHash := sha256.New()

	_, err := io.Copy(fileHash, f)
	if err != nil {
		return false, err
	}
	_, err = io.Copy(contentHash, c)
	if err != nil {
		return false, err
	}

	match := bytes.Compare(fileHash.Sum(nil), contentHash.Sum(nil))
	if match == 0 {
		return true, nil
	}

	return false, nil
}
