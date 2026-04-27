package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {

	var costSlice []float64

	for i:=0;i<len(costs);i++{
		if costs[i].day == day{
			costSlice=append(costSlice,costs[i].value)
		}
						
	}
	if costSlice== nil {
		return []float64{} 
	}
	return costSlice
}
