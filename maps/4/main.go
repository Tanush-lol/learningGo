package main
 

func getNameCounts(names []string) map[rune]map[string]int {


		map1:= make(map[rune]map[string]int)

	for _,m := range names{
		firstRune:= []rune(m)[0]
		
		if _, ok := map1[firstRune]; !ok{
			map1[firstRune]= make(map[string]int)
		}
		map1[firstRune][m]++

	}

	return map1
}
