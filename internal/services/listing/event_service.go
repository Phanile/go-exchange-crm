package listing

import (
	"context"
	"log/slog"
)

type ListingPublisher interface {
	PublishListingCoin(ctx context.Context, id int) error
}

func (s *ListingService) PublishListingCoin(ctx context.Context, id int) error {
	const op = "ListingService.PublishListingCoin"

	s.log.With(
		slog.String("op", op),
		slog.Int("coinId", id),
	)

	return s.publisher.PublishListingCoin(ctx, id)
}
