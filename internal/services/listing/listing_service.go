package listing

import (
	"context"
	"github.com/Phanile/go-exchange-crm/internal/dto"
	"log/slog"
)

type ListingService struct {
	log *slog.Logger
}

func NewListingService(log *slog.Logger) *ListingService {
	return &ListingService{
		log: log,
	}
}

func (s *ListingService) CreateListingRequest(ctx context.Context, req *dto.ListingRequestDTO) error {
	const op = "ListingService.CreateListingRequest"

	return nil
}
