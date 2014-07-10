package deis

import (
	"log"
)

func Monitor() {
	apps := ListApps()

	for _, app := range apps {
		has, instances := app.OldInstances()

		if has == true {
			log.Println("Old instances detected for App:", app.Name)
			for _, instance := range instances {
				instance.Stop()
			}
		}

	}
}
