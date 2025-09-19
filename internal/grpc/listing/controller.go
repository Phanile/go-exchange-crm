package listing

import (
	"context"
	"github.com/Phanile/go-exchange-crm/internal/dto"
	listingv1 "github.com/Phanile/go-exchange-protos/generated/go/crm"
	"google.golang.org/grpc"
)

type Listing interface {
	CreateListingRequest(ctx context.Context, req *dto.ListingRequestDTO) error
	//ValidateListingRequest(ctx context.Context, req *dto.ValidateListingDTO) error
	//ApproveListingRequest(ctx context.Context, req *dto.ApproveListingDTO) error
	//RejectListingRequest(ctx context.Context, req *dto.RejectListingDTO) error
	//DelistCoin(ctx context.Context, req *dto.DelistRequestDTO) error
}

type ServerAPI struct {
	listingv1.UnimplementedListingServer
	listingProvider Listing
}

func Register(server *grpc.Server, listingProvider Listing) {
	listingv1.RegisterListingServer(server, &ServerAPI{
		listingProvider: listingProvider,
	})
}

func (s *ServerAPI) CreateListingRequest(ctx context.Context, req *listingv1.ListingRequest) (*listingv1.ListingResponse, error) {
	err := s.listingProvider.CreateListingRequest(ctx, dto.ListingRequestFromProto(req))

	if err != nil {
		return &listingv1.ListingResponse{
			Response: "Возникла ошибка при создании заявки",
		}, err
	}

	return &listingv1.ListingResponse{
		Response: "Монета принята",
	}, nil
}
