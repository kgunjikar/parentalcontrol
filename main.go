package main

import (
	_ "bytes"
	"fmt"
	"log"
	"sniffer/devices"
	controlsniffer "sniffer/sniffer"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"sniffer/config"
	"sniffer/event"
	"sniffer/logger"
)

func main() {
	fmt.Println("--= GoSniff =--")
	fmt.Println("A simple packet sniffer in golang")

	if err := devices.DeviceInit(); err != nil {
		panic(err)
	}
	if err := logger.LoggerInit(); err != nil {
		panic(err)
	}
	go event.Init()

	handler, err := pcap.OpenLive(config.Config.IfName, config.Config.BufferSize, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()
	if config.Config.HTTPOnly == true {
		if err := handler.SetBPFFilter(config.Config.Filter); err != nil {
			log.Fatal(err)
		}
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range source.Packets() {
		if config.Config.HTTPOnly == true {
			go controlsniffer.HarvestHTTP(packet)
		} else {
			controlsniffer.Dot11LayerParse(packet)
		}
	}
}
