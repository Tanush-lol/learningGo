package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil,errors.New("invalid sizes")
	}
	userMaps:= make(map[string]user)

	for i := 0; i < len(names); i++ {
		userMaps[names[i]] = user{
			name: names[i],
			phoneNumber: phoneNumbers[i],
		}	
	}
	return userMaps,nil
}

type user struct {
	name        string
	phoneNumber int
}
