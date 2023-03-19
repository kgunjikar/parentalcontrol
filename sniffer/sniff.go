package controlsniffer

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"go.uber.org/ratelimit"
	"net"
	"sniffer/event"
	"sniffer/logger"
	"sync/atomic"
	"time"
)

func lookupIP(addr string) []string {
	out, err := net.LookupAddr(addr)
	if err != nil {
		return nil
	}
	return out
}

func Dot11LayerParse(p gopacket.Packet) []byte {
	atomic.AddUint64(&countdot11, 1)
	dot11 := p.Layer(layers.LayerTypeDot11)
	if dot11 != nil {
		dot11, _ := dot11.(*layers.Dot11)
		if dot11 != nil {
			if dot11.DataLayer != nil {
				logger.Log.Infof("LayerType:%#v\n", dot11)
				/*ethPacket := gopacket.NewPacket(eth, layers.LayerTypeEthernet, gopacket.Lazy)
				harvestHTTP(ethPacket)*/
			}
		}
	}
	return nil
}

var rl = ratelimit.New(1)
var prev = time.Now()
var countip, countdot11 uint64

func HarvestHTTP(p gopacket.Packet) {
	atomic.AddUint64(&countip, 1)
	ipv4 := p.Layer(layers.LayerTypeIPv4)
	if ipv4 != nil {
		ipv4, _ := ipv4.(*layers.IPv4)
		// the flags are empty in many of the packets of this example capture file
		logger.Log.Infof("Destination IP: %v %s\n", ipv4.DstIP, lookupIP(ipv4.DstIP.String()))
		now := rl.Take()
		event.LogEvent(ipv4.DstIP, time.Now())
		prev = now
	}
	logger.Log.Infof("\n")
}
