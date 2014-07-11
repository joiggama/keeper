package deis

type App struct {
	Name      string
	Services  []Service
}

func ListApps() []App {
	var list []App

	apps := map[string]App{}

	for _, service := range ListServices() {
		app := apps[service.AppName]
		if app.Present() != true {
			apps[service.AppName] = App{
        Name: service.AppName,
        Services: []Service{ service },
      }
		} else {
			app.Services = append(app.Services, service)
			apps[service.AppName] = app
		}
	}

	for _, app := range apps {
		list = append(list, app)
	}

	return list
}

func (self *App) Present() bool {
	return (len(self.Services) > 0)
}


