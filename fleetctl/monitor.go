package fleetctl

import (
	"fmt"
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

func Monitor() {
	all := ListUnits()
	failed := FilterFailedUnits(all)

	if len(failed) > 0 {
		fmt.Println("Failed units:", len(failed))
	}
}

func ListUnits() []Unit {
	listCmd := exec.Command("fleetctl", "list-units", "--no-legend", "--fields=unit,state,load,active,sub")
	cmdOutput, _ := listCmd.Output()

	output := strings.Trim(string(cmdOutput), "\n\n")
	lines := strings.Split(output, "\n")

	units := make([]Unit, len(lines))
	r, _ := regexp.Compile(`\s+`)

	for index, value := range lines {
		unit := r.Split(value, 5)

		units[index] = Unit{
			Id:     strings.TrimSpace(unit[0]),
			State:  strings.TrimSpace(unit[1]),
			Load:   strings.TrimSpace(unit[2]),
			Active: strings.TrimSpace(unit[3]),
			Sub:    strings.TrimSpace(unit[4]),
		}
	}
	return units
}

func FilterFailedUnits(units []Unit) []Unit {
	var result []Unit

	for _, unit := range units {
		if unit.State == "loaded" && unit.Active == "failed" && unit.Sub == "failed" {
			result = append(result, unit)
		}
	}

	return result
}
