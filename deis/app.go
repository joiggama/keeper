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

func (a *App) LatestVersion() int {
	return a.Versions()[len(a.Versions())-1]
}

func (a *App) Present() bool {
	return (len(a.Instances) > 0)
}

func (a *App) Versions() []int {
	var versions []int

	for _, instance := range a.Instances {
		versions = append(versions, instance.Version)
	}

	sort.Ints(versions)

	return versions
}
