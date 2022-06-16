package gofigure

type Resource interface {
	Apply() error
	Remove() error
}

func Exists(r Resource) error {
	return r.Apply()
}

func Absent(r Resource) error {
	return r.Remove()
}
