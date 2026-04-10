
package main

import (
	"fmt"
	"reflect"
)

	type contact struct {
		userID       string
		sendingLimit int32
		//userID       string
		//over here it occupies more space which can be debuged by reflect package
		age          int32
	}

	type perms struct {
		permissionLevel int
		canSend         bool
		canReceive      bool
		canManage       bool
	}
func main()  {
	
	structure := reflect.TypeOf(contact{})
	fmt.Print(structure.Size())
}
