package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "curl -sSL $V2RAY_CONF_URL | sed -e '/\"listen\": \"127.0.0.1\",/d' | sed -e 's|\"port\": 1755|\"port\": 8080|' > v2ray/config.json")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	v2ray := exec.Cmd{
		Dir:    "./v2ray",
		Path:   "./v2ray",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = v2ray.Run()
	if err != nil {
		panic(err)
	}
}
