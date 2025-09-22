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

func ValidateRequestFromProto(req *listingv1.ValidateRequest) *ValidateListingDTO {
	return &ValidateListingDTO{
		ListingId: int(req.Id),
	}
}

type ApproveListingDTO struct {
	ListingId int
}

func ApproveRequestFromProto(req *listingv1.ApproveRequest) *ApproveListingDTO {
	return &ApproveListingDTO{
		ListingId: int(req.Id),
	}
}

type RejectListingDTO struct {
	ListingId int
}

func RejectRequestFromProto(req *listingv1.RejectRequest) *RejectListingDTO {
	return &RejectListingDTO{
		ListingId: int(req.Id),
	}
}

type DelistRequestDTO struct {
	CoinId int
}

func DelistRequestFromProto(req *listingv1.DelistRequest) *DelistRequestDTO {
	return &DelistRequestDTO{
		CoinId: int(req.Id),
	}
}
