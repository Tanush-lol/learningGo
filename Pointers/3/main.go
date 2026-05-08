package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func analyzeMessage(structure *Analytics,message Message)  {
	if message.Success == true {
		structure.MessagesSucceeded++
		structure.MessagesTotal++
		}else{
			structure.MessagesFailed++
			structure.MessagesTotal++
		}	
}
