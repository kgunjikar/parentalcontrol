package devices

import (
	"github.com/google/gopacket/pcap"
	"log"
)

func deviceExists(name string) bool {
	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Panic(err)
	}

	for _, device := range devices {
		if device.Name == name {
			return true
		}
	}
	return false
}

func DeviceInit() error {
	if !deviceExists(config.IfName) {
		log.Fatal("Unable to open device ", iface)
	}
	return nil
}
