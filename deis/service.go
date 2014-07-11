package deis

import (
	"sort"
)

type Service struct {
	AppName   string
	Name      string
	Instances []Instance
}

func ListServices() []Service {
	var list []Service

	services := map[string]Service{}

	for _, instance := range ListInstances() {

		service_id := instance.AppName + "." + instance.ServiceName
		service := services[service_id]

		if service.Present() != true {
			services[service_id] = Service{
				AppName:   instance.AppName,
				Name:      instance.ServiceName,
				Instances: []Instance{instance},
			}
		} else {
			service.Instances = append(service.Instances, instance)
			services[service_id] = service
		}
	}

	return list
}

func (self *Service) LatestVersion() int {
	return self.Versions()[len(self.Versions())-1]
}

func (self *Service) OldInstances() (bool, []Instance) {
	var old []Instance

	if len(self.Instances) < 2 {
		return false, old
	}

	current_version := self.LatestVersion()

	for _, instance := range self.Instances {
		if instance.Version < (current_version - 2) {
			old = append(old, instance)
		}
	}

	return len(old) > 0, old
}

func (self *Service) Present() bool {
	return len(self.Instances) > 0
}

func (self *Service) Versions() []int {
	var versions []int

	for _, instance := range self.Instances {
		versions = append(versions, instance.Version)
	}

	sort.Ints(versions)

	return versions
}
