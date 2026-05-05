package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	//
	elem, ok := users[name]

	if ok != true {
		return false,errors.New("not found")	
	}
	if elem.scheduledForDeletion == true {
		delete(users,name)
		return true,nil
	}

	return false,nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
