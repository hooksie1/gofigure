package resources

import (
	"os/exec"
)

type DNFPkg struct {
	Name string
}

func NewDNFPkg(name string) *DNFPkg {
	return &DNFPkg{
		Name: name,
	}
}

func (p *DNFPkg) Apply() error {
	cmd := exec.Command("dnf", "install", "-y", p.Name)

	return cmd.Run()
}

func (p *DNFPkg) Remove() error {
	cmd := exec.Command("dnf", "remove", "-y", p.Name)

	return cmd.Run()

}
