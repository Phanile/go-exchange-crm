package listing

import (
	"context"
	"github.com/Phanile/go-exchange-crm/internal/dto"
	listingv1 "github.com/Phanile/go-exchange-protos/generated/go/crm"
	"google.golang.org/grpc"
)

type Listing interface {
	CreateListingRequest(req *dto.ListingRequestDTO) error
	ValidateListingRequest(req *dto.ValidateListingDTO) error
	ApproveListingRequest(req *dto.ApproveListingDTO) error
	RejectListingRequest(req *dto.RejectListingDTO) error
	DelistCoin(req *dto.DelistRequestDTO) error
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
	err := s.listingProvider.CreateListingRequest(dto.ListingRequestFromProto(req))

	if err != nil {
		return &listingv1.ListingResponse{
			Response: "Возникла ошибка при создании заявки",
		}, err
	}

	return &listingv1.ListingResponse{
		Response: "Монета принята",
	}, nil
}

func (s *ServerAPI) ValidateListingRequest(ctx context.Context, req *listingv1.ValidateRequest) (*listingv1.ValidateResponse, error) {
	err := s.listingProvider.ValidateListingRequest(dto.ValidateRequestFromProto(req))

	if err != nil {
		return &listingv1.ValidateResponse{
			Response: "Ошибка при валидации заявки",
		}, err
	}

	return &listingv1.ValidateResponse{
		Response: "Заявка валидна",
	}, nil
}

func (s *ServerAPI) ApproveListingRequest(ctx context.Context, req *listingv1.ApproveRequest) (*listingv1.ApproveResponse, error) {
	err := s.listingProvider.ApproveListingRequest(dto.ApproveRequestFromProto(req))

	if err != nil {
		return &listingv1.ApproveResponse{
			Response: "Ошибка принятии заявки",
		}, err
	}

	return &listingv1.ApproveResponse{
		Response: "Заявка принята",
	}, nil
}

func (s *ServerAPI) RejectListingRequest(ctx context.Context, req *listingv1.RejectRequest) (*listingv1.RejectResponse, error) {
	err := s.listingProvider.RejectListingRequest(dto.RejectRequestFromProto(req))

	if err != nil {
		return &listingv1.RejectResponse{
			Response: "Ошибка при отказе заявки",
		}, err
	}

	return &listingv1.RejectResponse{
		Response: "В заявке отказано",
	}, nil
}

func (s *ServerAPI) DelistCoin(ctx context.Context, req *listingv1.DelistRequest) (*listingv1.DelistResponse, error) {
	err := s.listingProvider.DelistCoin(dto.DelistRequestFromProto(req))

	if err != nil {
		return &listingv1.DelistResponse{
			Response: "Ошибка при делистинге монеты",
		}, err
	}

	return &listingv1.DelistResponse{
		Response: "Делистинг монеты прошел успешно",
	}, nil
}
