package deis

import (
	"regexp"
	"strconv"

	"github.com/joiggama/keeper/fleetctl"
)

type Service struct {
	App     string
	Version int
	Name    string
	Id      int
}

func ListServices() []Service {
	var services []Service

	units := fleetctl.ListUnits()

	r, _ := regexp.Compile(`(\w+-\w+)_v(\d+)\.(\w+)\.(\d+)\.(service)`)

	for _, unit := range units {
		submatches := r.FindStringSubmatch(unit.Id)

		if len(submatches) < 6 {
			continue
		}

		version, _ := strconv.Atoi(submatches[2])
		id, _ := strconv.Atoi(submatches[4])

		services = append(services, Service{
			App:     submatches[1],
			Version: version,
			Name:    submatches[3],
			Id:      id,
		})
	}

	return services
}
