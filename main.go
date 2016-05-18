package main

import (
	"fmt"

	"os/exec"
)

func main() {
	fmt.Println("Hi")

	// todo
	// spawn nginx, pass reference over to scheduler (reference has Reload() method on the type)
	// spawn ecs poller, pass reference to scheduler (has Refresh() method on the type)
	// spawn scheduler, is the brain of linking the two together


	// just seeing if nginx is there
	raw, err := exec.Command("cat", "/etc/nginx/nginx.conf").Output()
	if err != nil {
		fmt.Println("nginx not found")
		return
	}

	s := string(raw)
	fmt.Println(s)
}
