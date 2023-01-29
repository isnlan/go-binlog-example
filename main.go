package main

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/canal"
)

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%s:%d", "127.0.0.1", 3306)
	cfg.User = "root"
	cfg.Password = "root"
	cfg.Flavor = "mysql"
	cfg.Dump.ExecutionPath = ""

	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	coords, err := c.GetMasterPos()
	if err != nil {
		panic(err)
	}

	fmt.Printf("postion: %v\n", coords)

	c.SetEventHandler(&binlogHandler{})
	err = c.RunFrom(coords)
	if err != nil {
		panic(err)
	}
}
