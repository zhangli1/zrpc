package main

import "zrpc/server"

func main() {
	server.NewZrpc("127.0.0.1", "1234").Startup()
}
