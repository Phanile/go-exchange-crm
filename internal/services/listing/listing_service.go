package listing

import (
	"context"
	"fmt"
	"github.com/Phanile/go-exchange-crm/internal/dto"
	"log/slog"
	"runtime"
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

func (s *ListingService) CreateListingRequest(req *dto.ListingRequestDTO) error {
	const op = "ListingService.CreateListingRequest"

	s.log.With(
		slog.String("op", op),
	)

	fmt.Printf("goroutines count before worker pool: %d", runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inputCh := rpcGenerator(ctx, req.RPCServerAddresses)
	resultCh := make(chan *Result)

	startWorkerPool(ctx, workersCount, inputCh, resultCh)

	for res := range resultCh {
		if res.err != nil {
			cancel()
			s.log.Error("rpc check err:", res.err)
			return res.err
		}
	}

	s.log.Info("all RPCs are valid")
	fmt.Printf("goroutines count after worker pool: %d", runtime.NumGoroutine())

	return nil
}

func (s *ListingService) ValidateListingRequest(req *dto.ValidateListingDTO) error {
	const op = "ListingService.ValidateListingRequest"

	return nil
}

func (s *ListingService) ApproveListingRequest(req *dto.ApproveListingDTO) error {
	const op = "ListingService.ApproveListingRequest"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return s.publisher.PublishListingCoin(ctx, req.ListingId)
}

func (s *ListingService) RejectListingRequest(req *dto.RejectListingDTO) error {
	const op = "ListingService.RejectListingRequest"

	return nil
}

func (s *ListingService) DelistCoin(req *dto.DelistRequestDTO) error {
	const op = "ListingService.DelistCoin"

	return nil
}
