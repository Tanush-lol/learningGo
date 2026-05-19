package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	
	errUpdate :=	em.recipient.updateStatus(em.status)
	if errUpdate !=nil {
		return	fmt.Errorf("error updating user status: %w",errUpdate)	
	}

	errTrack := a.track(em.status)
	if errTrack!=nil {
		return	fmt.Errorf("error tracking user bounce: %w",errTrack)	
	}
	return nil
}
