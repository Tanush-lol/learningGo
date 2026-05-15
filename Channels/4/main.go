package main

func addEmailsToQueue(emails []string) chan string {
	length:= len(emails)

	ch:= make(chan string,length)	
	for i := 0; i < length; i++ {
		ch <- emails[i]	
	}
	return ch
}
