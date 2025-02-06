package service

import (
	"log"
	"os/exec"
)

// doPing attempts a "ping -c 1 -w 1 <ip>", returning true if successful.
func doPing(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-w", "1", ip)
	err := cmd.Run()
	if err != nil {
		log.Printf("Ping error for %s: %v\n", ip, err)
		return false
	}
	return true
}
