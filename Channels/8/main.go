package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	var take,save = true,true
	for take && save {
		select {
			case _,ok1:=<-snapshotTicker:
			if ok1 {
				takeSnapshot(logChan)
			}else{
				take = false
			}

		case _,ok2:= <-saveAfter:
			if ok2 {
				saveSnapshot(logChan)
				return 	
			}
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
