module sniffer

go 1.14

require (
	github.com/google/gopacket v1.1.19
	github.com/jasonlvhit/gocron v0.0.1
	github.com/spf13/viper v1.15.0
	go.uber.org/ratelimit v0.2.0
	go.uber.org/zap v1.24.0
)

replace event => ./event
