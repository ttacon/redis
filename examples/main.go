package main

import (
	"flag"
	"fmt"

	"github.com/ttacon/pretty"
	"github.com/ttacon/redis"
)

var (
	keyName  = flag.String("k", "", "the key you want to see the type of")
	keyName2 = flag.String("k2", "", "second key")
	cmd      = flag.String("c", "", "the command to call")
	port     = flag.String("p", "6379", "the port to connect to redis on")
)

func main() {
	flag.Parse()

	if len(*keyName) == 0 {
		fmt.Println("please provide a key to inspect")
		return
	}

	c, err := redis.NewClient(fmt.Sprintf(":%s", *port))
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
	case "ping":
		pong, err := c.Ping()
		fmt.Println("pong: ", pong)
		fmt.Println("err: ", err)
	case "echo":
		echo, err := c.Echo("a cool message")
		fmt.Println("echo: ", echo)
		fmt.Println("err: ", err)
	case "auth":
		ok, err := c.Auth("bananaschipsandpoop")
		fmt.Println("err: ", err)
		fmt.Println("ok: ", ok)
	case "ttl":
		ttl, err := c.TTL(*keyName)
		fmt.Println("err: ", err)
		fmt.Println("ttl: ", ttl)
	case "randomkey":
		key, err := c.RandomKey()
		fmt.Println("err: ", err)
		fmt.Println("key: ", key)
	case "sadd":
		numAdded, err := c.SAdd(*keyName, "yolo")
		fmt.Println("numAdded: ", numAdded)
		fmt.Println("err: ", err)
	case "scard":
		card, err := c.SCard(*keyName)
		fmt.Println("err: ", err)
		fmt.Println("card: ", card)
	case "sdiff":
		dif, err := c.SDiff(*keyName)
		fmt.Println("err: ", err)
		fmt.Println("dif: ", dif)
	case "sismember":
		isMember, err := c.SIsMember(*keyName, "yolo2")
		fmt.Println("err: ", err)
		fmt.Println("isMember: ", isMember)
	case "smove":
		wasMoved, err := c.SMove(*keyName, *keyName2, "yolo")
		fmt.Println("err: ", err)
		fmt.Println("wasMoved: ", wasMoved)
	case "scan":
		strs, next, err := c.Scan(*keyName, 0)
		fmt.Println("strs: ", strs)
		fmt.Println("next: ", next)
		fmt.Println("err: ", err)
	}
}
