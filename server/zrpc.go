package server

type Zrpc struct {
	Host string
	Port string
}

func NewZrpc(host string, port string) *Zrpc {
	return &Zrpc{Host: host, Port: port}
}

func (self *Zrpc) Startup() {
	startup(self.Host, self.Port)
}
