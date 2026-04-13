
package main

type formatter interface{
	format() string
}
type plainText struct{
	message string
}
type bold struct{
	message string
}
type code struct{
	message string
}

func (mes plainText) format()  string{
	return mes.message 
}

func (mes bold) format()  string{
	return "**"+mes.message+"**"
}
func (mes code) format()  string{
	return "`"+mes.message+"`"	
}
// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}
