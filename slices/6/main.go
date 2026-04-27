package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word1 := range msg{
		for _,word2 := range badWords{
			if word1 == word2 {
				return i 
			}
		}	
	}
	return -1
}

