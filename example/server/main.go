package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"zrpc/server"

	"gopkg.in/gcfg.v1"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pwd, _ := os.Getwd()
	executeDir := *flag.String("d", pwd, "execute directory") + "/"

	configFilePath := flag.String("c", "base.cfg", "config file")
	flag.Parse()
	fmt.Println("Current execute directory:", executeDir)

	config := executeDir + *configFilePath
	fmt.Println(config)
	var server_cfg Config
	gcfg.ReadFileInto(&server_cfg, config)

	server.NewZrpc(server_cfg.Base.Host, server_cfg.Base.Port).Startup()
}
