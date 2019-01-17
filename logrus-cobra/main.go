package main

import (
	"github.com/seblegall/blog-scripts/logrus-cobra/cmd"
)

func main() {
	if err := cmd.NewMyCmd().Execute(); err != nil {
		panic(err)
	}
}
