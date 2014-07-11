package deis

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/joiggama/keeper/fleetctl"
)

type Instance struct {
	AppName     string
	Version     int
	ServiceName string
	Id          int
	Status      string
}

func ListInstances() []Instance {
	var instances []Instance

	r, _ := regexp.Compile(`(\w+-\w+)_v(\d+)\.(\w+)\.(\d+)\.(service)`)

	for _, unit := range fleetctl.ListUnits() {
		submatches := r.FindStringSubmatch(unit.Id)

		if len(submatches) < 6 {
			continue
		}

		id, _ := strconv.Atoi(submatches[4])
		version, _ := strconv.Atoi(submatches[2])

		instances = append(instances, Instance{
			AppName: submatches[1],
			Id:      id,
			Version: version,
			Status:  unit.Sub,
		})
	}

	return instances
}

func (self *Instance) Destroy() {
	unit := self.UnitName()
	log.Println("Destroying:", unit)

	destroyCmd := exec.Command("fleetctl", "destroy", unit+".service")
	destroyLogCmd := exec.Command("fleetctl", "destroy", unit+".log.service")
	destroyAnnCmd := exec.Command("fleetctl", "destroy", unit+".announce.service")

	destroyCmd.Start()
	destroyLogCmd.Start()
	destroyAnnCmd.Start()

	err := destroyCmd.Wait()

	if err != nil {
		log.Println(err)
	}

}

func (self *Instance) Stop() {
	unit := self.UnitName()
	log.Println("Stopping:", unit)

	stopCmd := exec.Command("fleetctl", "stop", unit+".service")
	stopLogCmd := exec.Command("fleetctl", "stop", unit+".log.service")
	stopAnnCmd := exec.Command("fleetctl", "stop", unit+".announce.service")

	stopCmd.Start()
	stopLogCmd.Start()
	stopAnnCmd.Start()

	err := stopCmd.Wait()

	if err != nil {
		log.Println(err)
	}

}

func (self *Instance) UnitName() string {
	return fmt.Sprintf("%s_v%d.%s.%d.",
		self.AppName,
		self.Version,
		self.ServiceName,
		self.Id,
	)
}
