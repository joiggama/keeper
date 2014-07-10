package deis

func Monitor() {
	apps := ListApps()

	for _, app := range apps {
		has, instances := app.OldInstances()

		if has == true {
			for _, instance := range instances {
				instance.Stop()
			}
		}

	}

	panic("exit")
}
