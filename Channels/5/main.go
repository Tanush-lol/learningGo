package main

func countReports(numSentCh chan int) int {
	var report int
	ok:= true
	for ok != false{
		v,idk:= <-numSentCh
		report += v
		ok = idk
	}	
	return report
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
