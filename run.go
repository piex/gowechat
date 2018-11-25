package gowechat

import (
	"fmt"
	"os"
	"time"
)

func run(desc string, f func() error) {
	start := time.Now()
	fmt.Println(desc)
	if err := f(); err != nil {
		fmt.Println("FAIL, exit now", err)
		os.Exit(1)
	}

	fmt.Printf(" SUCCESS, use time %f s\n", time.Now().Sub(start).Seconds())
}
