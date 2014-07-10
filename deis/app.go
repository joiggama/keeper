package deis

import (
	"sort"
)

type App struct {
	Name      string
	Instances []Service
}

func ListApps() []App {
	var list []App

	apps := map[string]App{}

	services := ListServices()

	for _, service := range services {
		if service.Status != "running" {
			continue
		}

		app := apps[service.App]
		if app.Present() != true {
			apps[service.App] = App{Name: service.App, Instances: []Service{service}}
		} else {
			app.Instances = append(app.Instances, service)
			apps[service.App] = app
		}
	}

	for _, app := range apps {
		list = append(list, app)
	}

	return list
}

func (self *App) OldInstances() (bool, []Service) {
	var old_instances []Service

	if len(self.Instances) < 2 {
		return false, old_instances
	}

	current_version := self.LatestVersion()

	for _, instance := range self.Instances {
		if instance.Version != current_version {
			old_instances = append(old_instances, instance)
		}
	}

	return len(old_instances) > 0, old_instances
}

func (self *App) LatestVersion() int {
	return self.Versions()[len(self.Versions())-1]
}

func (self *App) Present() bool {
	return (len(self.Instances) > 0)
}

func (self *App) Versions() []int {
	var versions []int

	for _, instance := range self.Instances {
		versions = append(versions, instance.Version)
	}

	sort.Ints(versions)

	return versions
}
