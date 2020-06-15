package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

//GetShortStatus read the git status of the repository located at path
func GetShortStatus(path string) (io.Reader, error) {
	return execOutput(fmt.Sprintf("git -C %s status -s -b", path))
}

//It is useful to declare a var instead of a function for testing purpose
var execOutput = func(c string) (io.Reader, error) {
	out, err := exec.Command("/bin/sh", "-c", c).Output()

	return bytes.NewReader(out), err
}
