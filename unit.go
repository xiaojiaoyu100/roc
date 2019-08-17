package roc

import "time"

// Unit represents a store unit
type Unit struct {
	Key            string
	Data           interface{}
	ExpirationTime time.Time
}

// Expire returns true when unit has expired.
func (u *Unit) Expire() bool {
	return u.ExpirationTime.Before(time.Now().UTC())
}
