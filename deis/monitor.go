package deis

import (
	"fmt"
)

func Monitor() {
	apps := ListApps()
	fmt.Println("Apps Count:", len(apps))
	for _, app := range apps {
		fmt.Println("App:", app.Name)
		for _, instance := range app.Instances {
			fmt.Println("- Instance:", instance.Version)
		}
	}

	panic("exit")
}
