package main

import (
	"fmt"
	"./config"
	"./server"
	"./web"
	"log"
	"time"
)

type logW struct {
}

func (writer logW) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logW))
	config.LoadConf()
	server.Con = []*server.Connects{}
	server.Lan = []*server.Connects{}
	log.Println("[HOST]", config.Host)
	log.Println("[IP]", config.Ip)
	go func() {
		//server.IpsLan()
	}()
	w := web.Http{}
	go func() {
		w.Listen(config.PortWeb)
	}()
	time.Sleep(100)
	s := server.Server{}
	go func() {
		s.Listen(config.PortMaster)
	}()
	time.Sleep(100)

	sum := 1
	for {
		time.Sleep(time.Second)
		if sum%180 == 0 {
			go func() {
				//server.IpsLan()
			}()
		}
		if sum%600 == 10 {
			go func() {
				server.DynamicIP()
			}()
		}
		sum++
	}
}
