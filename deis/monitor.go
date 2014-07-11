package deis

import (
	"log"
)

func Monitor() {
	apps := ListApps()

	for _, app := range apps {

		for _, service := range app.Services {

			has, instances := service.OldInstances()

			if has == true {
        log.Println("Old instances detected for Service:", service.Name, "from App:", service.AppName)

				for _, instance := range instances {
					switch instance.Status {
					case "-":
						log.Println(instance.UnitName())
					default:
						instance.Stop()
					}
				}
			}
		}

	}

}
