package dto

import listingv1 "github.com/Phanile/go-exchange-protos/generated/go/crm"

type ListingRequestDTO struct {
	ChainID            int
	FullName           string
	ShortName          string
	Currency           string
	RPCServerAddresses []string
	WhitePaperLink     string
}

func ListingRequestFromProto(req *listingv1.ListingRequest) *ListingRequestDTO {
	return &ListingRequestDTO{
		ChainID:            int(req.ChainId),
		FullName:           req.FullName,
		ShortName:          req.ShortName,
		Currency:           req.Currency,
		RPCServerAddresses: req.RpcServerAddresses,
		WhitePaperLink:     req.WhitePaperLink,
	}
}

type ValidateListingDTO struct {
	ListingId int
}

type ApproveListingDTO struct {
	ListingId int
}

type RejectListingDTO struct {
	ListingId int
}

type DelistRequestDTO struct {
	CoinId int
}
