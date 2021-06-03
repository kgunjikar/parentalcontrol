module sniffer

go 1.14

require (
	event v0.0.0-00010101000000-000000000000
	github.com/google/gopacket v1.1.19
	github.com/jasonlvhit/gocron v0.0.1 // indirect
	github.com/takama/daemon v1.0.0
	go.uber.org/ratelimit v0.2.0
)

replace event => ./event
