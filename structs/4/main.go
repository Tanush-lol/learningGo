
package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (auth authenticationInfo) getBasicAuth() string   {
	fmt.Println(auth.username)
	return "Authorization: Basic "+auth.username+":"+auth.password
}


// authenticationInfo struct is passed into func which actually isn't a func but a method that is named getBasicAuth
// hence called that to pass upon the variables
