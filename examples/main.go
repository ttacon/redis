package main

import (
	"flag"
	"fmt"

	"github.com/kr/pretty"
	"github.com/ttacon/redis"
)

var (
	keyName = flag.String("k", "", "the key you want to see the type of")
	cmd     = flag.String("c", "", "the command to call")
)

func main() {
	flag.Parse()

	if len(*keyName) == 0 {
		fmt.Println("please provide a key to inspect")
		return
	}

	c, err := redis.NewClient(":6379")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	switch *cmd {
	case "type":
		resp, err := c.Type(*keyName)
		fmt.Println("err: ", err)
		fmt.Println("resp: ", resp)
	case "smembers":
		resp, err := c.SMembers(*keyName)
		fmt.Println("err: ", err)
		pretty.Print(resp)
		fmt.Println()
	case "select":
		ok, err := c.Select(1)
		fmt.Println("err: ", err)
		fmt.Println("ok: ", ok)
	case "quit":
		ok, err := c.Quit()
		fmt.Println("ok: ", ok)
		fmt.Println("err: ", err)
	}
}
