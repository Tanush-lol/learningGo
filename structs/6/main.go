
package main

type User struct {
	Name string
	Membership
}
type Membership struct{
	Type string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	member:= User{}

	member.Name= name
	member.Membership.Type = membershipType

	if "premium" == membershipType{
		member.Membership.MessageCharLimit = 1000
			}else {
		member.Membership.MessageCharLimit= 100
	}
	return member
}
