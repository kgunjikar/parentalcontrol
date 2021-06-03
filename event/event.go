package event

import (
	"net"
	"time"
	"fmt"
	"sync"
	"github.com/jasonlvhit/gocron"
	"log"
	"os"
	"io/ioutil"
)

type SitesVisited struct {
	DestIP net.IP
	TimeStamp time.Time
}

var flows map[string]SitesVisited
var eventMutex sync.RWMutex

func startCron() {

	s := gocron.NewScheduler()
	s.Every(2).Days().Do(CleanupStale)
	<- s.Start()
}

func Init() {
	flows = make(map[string]SitesVisited, 0)
	startCron()
}

func logToFile() {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "event-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}
	log.Printf("Writing logfile:%s to disk", tmpFile)

	eventMutex.RLock()
	text := fmt.Sprintf("%#v", flows)
	eventMutex.RUnlock()
	if _, err = tmpFile.Write([]byte(text)); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func LogEvent(dst net.IP, now time.Time) {
	eventMutex.RLock()
	_, ok := flows[dst.String()]
	eventMutex.RUnlock()
	if !ok {
		newSiteVisit := new(SitesVisited)
		newSiteVisit.TimeStamp = now
		newSiteVisit.DestIP = dst
		log.Printf("NewSite: %#v", newSiteVisit)
		eventMutex.Lock()
		flows[dst.String()] = *newSiteVisit
		eventMutex.Unlock()
	}
}

func CleanupStale() {
	logToFile()
	eventMutex.Lock()
	defer eventMutex.Unlock()
	for _, val := range flows {
		currentTime := time.Now()
		duration := currentTime.Sub(val.TimeStamp)
		if duration.Hours() > 48 {
			delete(flows,val.DestIP.String())
		}
	}
}
