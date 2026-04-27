
package main

func getMessageCosts(messages []string) []float64 {
	// ?
	length:= len(messages)
	mySlice:= make([]float64,length)
	for i := 0; i < length; i++ {
		mySlice[i]=float64(len(messages[i]))*0.01;
	}
	return mySlice
}
