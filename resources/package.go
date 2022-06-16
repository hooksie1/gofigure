package resources

import (
	"os/exec"
)

type Pkg struct {
	Name string
}

func NewPkg() *Pkg {
	return &Pkg{}
}

func (p *Pkg) Apply() error {
	cmd := exec.Command("dnf", "install", "-y", p.Name)

	return cmd.Run()
}

func (p *Pkg) Remove() error {
	cmd := exec.Command("apt", "remove", "-y", p.Name)

	return cmd.Run()

}
