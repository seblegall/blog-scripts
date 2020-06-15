package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatalf("please, provide a git repository as command argument")
	}

	out, err := GetShortStatus(args[0])
	if err != nil {
		log.Fatalf("unable to read git repository status : %s", err.Error())
	}

	status := ParseShort(out)

	fmt.Println("Branch is ", status.Branch)
	for _, fs := range status.FilesStatus {
		fmt.Printf("    %s\n", fs)
	}

}
