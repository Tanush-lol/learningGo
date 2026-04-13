
package main

func maxMessages(thresh int) int {
	var cost int
	var maxMessages int
	for ; cost < thresh; maxMessages++ {
		cost +=100+maxMessages
		if cost> thresh {
				break
			}	
	}
	return maxMessages
}
