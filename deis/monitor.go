package deis

import (
	"fmt"
)

func Monitor() {
	apps := ListApps()
  fmt.Println(apps)

	panic("exit")
}
