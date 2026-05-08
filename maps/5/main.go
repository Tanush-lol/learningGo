package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	map1:= make(map[string]bool)
	var totalDistinctWords int

	i:=0
for i < len(messages){
		temp := strings.Fields(messages[i])
	
		if strings.TrimSpace(messages[i]) == "" {
			i++
		continue
		}

		if len(temp)>1 {

			messages = messages[i+1:]
			messages= append(temp,messages...)
			i=0
			
		}

		if _, ok := map1[strings.ToLower(messages[i])]; !ok {
    	map1[strings.ToLower(messages[i])] = true
			totalDistinctWords ++
		} 
		i++
	}
	return totalDistinctWords
}
