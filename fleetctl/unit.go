package fleetctl

import (
	"os/exec"
	"regexp"
	"strings"
)

type Unit struct {
	Id     string
	State  string
	Load   string
	Active string
	Sub    string
}

func ListUnits() []Unit {
	var units []Unit

	listCmd := exec.Command("fleetctl",
		"list-units",
		"--no-legend",
		"--fields=unit,state,load,active,sub",
	)

	cmdOut, _ := listCmd.Output()

	r, _ := regexp.Compile(`\s+`)

	for _, value := range strings.Split(strings.Trim(string(cmdOut), "\n\n"), "\n") {
		unit := r.Split(value, 5)

		units = append(units, Unit{
			Id:     strings.TrimSpace(unit[0]),
			State:  strings.TrimSpace(unit[1]),
			Load:   strings.TrimSpace(unit[2]),
			Active: strings.TrimSpace(unit[3]),
			Sub:    strings.TrimSpace(unit[4]),
		})
	}

	return units
}
