
package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	// ?
	messages:= [3]string{primary,secondary,tertiary}
	
	cost1:=len(primary)
	cost2:=len(secondary)
	cost3:=len(tertiary)

	cost := [3]int{cost1,cost1+cost2,cost1+cost2+cost3}
	return messages,cost
}
