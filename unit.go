package roc

import "time"

type Unit struct {
	Key            string
	Data           interface{}
	ExpirationTime time.Time
}

func (u *Unit) Expire() bool {
	return u.ExpirationTime.Before(time.Now().UTC())
}
