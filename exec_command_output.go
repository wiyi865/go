package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}
	res, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Output() failed with %s\n", err)
	}
	fmt.Println(res)
	fmt.Println(string(res))
}
