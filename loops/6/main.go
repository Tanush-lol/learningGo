
package main

func countConnections(groupSize int) int {
	// ?
	var connections int
	for i:=groupSize;i>0;i--{
		connections =connections+i-1
	}
	return connections }
