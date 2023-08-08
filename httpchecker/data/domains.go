package data

import "time"

//ssl domains represent domain that is tracked for ssl expiry

type SSLTracking struct {
	ID         int
	Issuer     string
	DomainName string
	Expires    time.Time
	Status     string
}
