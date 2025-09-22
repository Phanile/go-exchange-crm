package listing

import (
	"context"
	"github.com/Phanile/go-exchange-crm/internal/dto"
	"log/slog"
)

type ListingService struct {
	log       *slog.Logger
	publisher ListingPublisher
}

func NewListingService(log *slog.Logger, publisher ListingPublisher) *ListingService {
	return &ListingService{
		log:       log,
		publisher: publisher,
	}
}

func (s *ListingService) CreateListingRequest(ctx context.Context, req *dto.ListingRequestDTO) error {
	const op = "ListingService.CreateListingRequest"

	return nil
}

func (s *ListingService) ValidateListingRequest(ctx context.Context, req *dto.ValidateListingDTO) error {
	const op = "ListingService.ValidateListingRequest"

	return nil
}

func (s *ListingService) ApproveListingRequest(ctx context.Context, req *dto.ApproveListingDTO) error {
	const op = "ListingService.ApproveListingRequest"

	s.publisher.PublishListingCoin(ctx, req.ListingId)

	return nil
}

func (s *ListingService) RejectListingRequest(ctx context.Context, req *dto.RejectListingDTO) error {
	const op = "ListingService.RejectListingRequest"

	return nil
}

func (s *ListingService) DelistCoin(ctx context.Context, req *dto.DelistRequestDTO) error {
	const op = "ListingService.DelistCoin"

	return nil
}
