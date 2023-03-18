package resources

import (
	"fmt"
	"os"
	"strings"

	"github.com/coreos/go-systemd/v22/dbus"
)

type SystemdType string

const (
	Timer   SystemdType = "Timer"
	Service SystemdType = "Service"
)

func (s SystemdType) String() string {
	return string(s)
}

type SystemdUnit struct {
	Type        SystemdType
	Description string
	Name        string
	Schedule    string
}

var timerTemplate string = `
[Unit]
Description={{ .Description }}
Requires={{ .Name }}.service

[{{ .Type }}]
Unit={{ .Name }}.service
OnCalendar={{ .Schedule }}

[Install]
WantedBy=timers.target

`

func NewSystemdUnit() *SystemdUnit {
	return &SystemdUnit{}
}

func (s *SystemdUnit) SetType(t SystemdType) *SystemdUnit {
	s.Type = t
	return s
}

func (s *SystemdUnit) SetDescription(d string) *SystemdUnit {
	s.Description = d
	return s
}

func (s *SystemdUnit) SetName(n string) *SystemdUnit {
	s.Name = n
	return s
}

func (s *SystemdUnit) SetSchedule(c string) *SystemdUnit {
	s.Schedule = c
	return s
}

func (s *SystemdUnit) Apply() error {
	path := fmt.Sprintf("/etc/systemd/system/%s.%s", s.Name, strings.ToLower(s.Type.String()))
	t := NewTemplate().SetSource(strings.NewReader(timerTemplate)).SetDest(path).SetMode(0644).SetOwner("root").SetGroup("root").SetVars(map[string]any{
		"Name":        s.Name,
		"Description": s.Description,
		"Type":        s.Type.String(),
		"Schedule":    s.Schedule,
	})

	if err := t.Apply(); err != nil {
		return err
	}

	conn, err := dbus.New()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, _, err = conn.EnableUnitFiles([]string{path}, false, false)
	if err != nil {
		return err
	}

	return nil

}

func (s *SystemdUnit) Remove() error {
	path := fmt.Sprintf("/etc/systemd/system/%s.%s", s.Name, strings.ToLower(s.Type.String()))
	if err := os.Remove(path); err != nil {
		return err
	}

	conn, err := dbus.New()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.DisableUnitFiles([]string{path}, false)
	if err != nil {
		return err
	}

	return nil
}
