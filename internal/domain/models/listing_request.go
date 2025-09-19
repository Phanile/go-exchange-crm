package models

import "time"

type ListingRequest struct {
	Id                 int
	CreatedAt          time.Time
	StatusId           int
	ChainID            int
	FullName           string
	ShortName          string
	Currency           string
	RPCServerAddresses []string
	WhitePaperLink     string
}
