
package main

func (e email) cost() int {
	if true==e.isSubscribed {
		return 2*len(e.body)
	}
		return 5*len(e.body)
	
}

func (e email) format() string {
	if true==e.isSubscribed {
		return "'"+e.body+"'"+" | Subscribed"
	}
	return "'"+e.body+"'"+" | Not Subscribed"
}


type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
