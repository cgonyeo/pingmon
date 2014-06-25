package main

import (
	"net"
	"time"

	golog "github.com/op/go-logging"
	"github.com/tatsushid/go-fastping"
)

var log = golog.MustGetLogger("main")

func main() {
	log.Debug("Started")

	p := fastping.NewPinger()
	addr8, err := net.ResolveIPAddr("ip4:icmp", "8.8.8.8")
	if err != nil {
		log.Error("Adding 8.8.8.8: %v", err)
	}
	p.AddIPAddr(addr8)

	addr10, err := net.ResolveIPAddr("ip4:icmp", "10.10.0.1")
	if err != nil {
		log.Error("Adding 10.10.0.1: %v", err)
	}
	p.AddIPAddr(addr10)

	pings := make(chan string, 2)

	p.AddHandler("receive", func(addr *net.IPAddr, rtt time.Duration) {
		log.Debug("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		pings <- addr.String()
	})
	p.AddHandler("idle", func() {
		var ips []string
		for i := 0; i < 2; i++ {
			select {
			case ip, ok := <-pings:
				if ok {
					ips = append(ips, ip)
				} else {
					//This shouldn't happen
					log.Fatal("pings channel closed")
				}
			default:
				//No ip received
			}
		}
		if len(ips) == 1 {
			log.Error("Error: only pinged " + ips[0] + " successfully")
		}
		if len(ips) == 0 {
			log.Error("Error: couldn't ping anything")
		}
	})

	_, errch := p.RunLoop()
	for {
		log.Error("Ping failed: %v", <-errch)
	}
}
