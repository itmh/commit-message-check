package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = CommitMessageCheck(string(buf))
	if err != nil {
		fmt.Println(err)
	}
}
