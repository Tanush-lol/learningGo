package main

import "strings"

func countDistinctWords(messages []string) int {
	map1:= make(map[string]bool)
	var totalDistinctWords int

for i := 0; i < len(messages); i++ {
		temp := strings.Fields(messages[i])

		if len(temp)>1 {
			messages = messages[i+1:]
			messages= append(temp,messages...)
		}

		if _, ok := map1[messages[i]]; !ok {
    	map1[messages[i]] = true
			totalDistinctWords ++
		} 

	}
	return totalDistinctWords
}
