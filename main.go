package main

import (
	_ "bytes"
	"fmt"
	"log"
	"net"
	"sniffer/devices"
	"sniffer/event"
	"sniffer/logger"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"go.uber.org/ratelimit"
	"sync/atomic"
	"time"
)

var (
	iface       = "wlp4s0mon"
	ifaceNormal = "wlp4s0"
	buffer      = int32(1600)
	filter      = "tcp and port 80"
)
var normal bool = true
var rl = ratelimit.New(1)
var prev = time.Now()

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

	if normal == true {
		iface = ifaceNormal
	}

	handler, err := pcap.OpenLive(iface, buffer, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handler.Close()
	if normal == true {
		if err := handler.SetBPFFilter(filter); err != nil {
			log.Fatal(err)
		}
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range source.Packets() {
		if normal == true {
			go harvestHTTP(packet)
		} else {
			dot11LayerParse(packet)
		}
	}
}

func lookupIP(addr string) []string {
	out, err := net.LookupAddr(addr)
	if err != nil {
		return nil
	}
	return out
}

func dot11LayerParse(p gopacket.Packet) []byte {
	atomic.AddUint64(&countdot11, 1)
	dot11 := p.Layer(layers.LayerTypeDot11)
	if dot11 != nil {
		dot11, _ := dot11.(*layers.Dot11)
		if dot11 != nil {
			if dot11.DataLayer != nil {
				fmt.Printf("LayerType:%#v\n", dot11)
				/*ethPacket := gopacket.NewPacket(eth, layers.LayerTypeEthernet, gopacket.Lazy)
				harvestHTTP(ethPacket)*/
			}
		}
	}
	return nil
}

var countip, countdot11 uint64

func harvestHTTP(p gopacket.Packet) {
	atomic.AddUint64(&countip, 1)
	ipv4 := p.Layer(layers.LayerTypeIPv4)
	if ipv4 != nil {
		ipv4, _ := ipv4.(*layers.IPv4)
		// the flags are empty in many of the packets of this example capture file
		fmt.Printf("Destination IP: %v %s\n", ipv4.DstIP, lookupIP(ipv4.DstIP.String()))
		now := rl.Take()
		event.LogEvent(ipv4.DstIP, time.Now())
		prev = now
	}
	fmt.Printf("\n")
}
