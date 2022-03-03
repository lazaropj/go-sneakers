package invrepo

import (
	"context"

	"github.com/sksmith/go-micro-example/core"
	"github.com/sksmith/go-micro-example/core/inventory"
	"github.com/sksmith/go-micro-example/db"
	"github.com/sksmith/go-micro-example/test"
)

type MockRepo struct {
	GetProductionEventByRequestIDFunc func(ctx context.Context, requestID string, options ...core.QueryOptions) (pe inventory.ProductionEvent, err error)
	SaveProductionEventFunc           func(ctx context.Context, event *inventory.ProductionEvent, options ...core.UpdateOptions) error

	GetReservationFunc            func(ctx context.Context, ID uint64, options ...core.QueryOptions) (inventory.Reservation, error)
	GetReservationsFunc           func(ctx context.Context, resOptions inventory.GetReservationsOptions, limit, offset int, options ...core.QueryOptions) ([]inventory.Reservation, error)
	GetReservationByRequestIDFunc func(ctx context.Context, requestId string, options ...core.QueryOptions) (inventory.Reservation, error)
	UpdateReservationFunc         func(ctx context.Context, ID uint64, state inventory.ReserveState, qty int64, options ...core.UpdateOptions) error
	SaveReservationFunc           func(ctx context.Context, reservation *inventory.Reservation, options ...core.UpdateOptions) error

	GetProductFunc  func(ctx context.Context, sku string, options ...core.QueryOptions) (inventory.Product, error)
	SaveProductFunc func(ctx context.Context, product inventory.Product, options ...core.UpdateOptions) error

	GetProductInventoryFunc    func(ctx context.Context, sku string, options ...core.QueryOptions) (inventory.ProductInventory, error)
	GetAllProductInventoryFunc func(ctx context.Context, limit int, offset int, options ...core.QueryOptions) ([]inventory.ProductInventory, error)
	SaveProductInventoryFunc   func(ctx context.Context, productInventory inventory.ProductInventory, options ...core.UpdateOptions) error

	BeginTransactionFunc func(ctx context.Context) (core.Transaction, error)

	*test.CallWatcher
}

func (r *MockRepo) SaveProductionEvent(ctx context.Context, event *inventory.ProductionEvent, options ...core.UpdateOptions) error {
	r.AddCall(ctx, event, options)
	return r.SaveProductionEventFunc(ctx, event, options...)
}

func (r *MockRepo) UpdateReservation(ctx context.Context, ID uint64, state inventory.ReserveState, qty int64, options ...core.UpdateOptions) error {
	r.AddCall(ctx, ID, state, options)
	return r.UpdateReservationFunc(ctx, ID, state, qty, options...)
}

func (r *MockRepo) GetProductionEventByRequestID(ctx context.Context, requestID string, options ...core.QueryOptions) (pe inventory.ProductionEvent, err error) {
	r.AddCall(ctx, requestID, options)
	return r.GetProductionEventByRequestIDFunc(ctx, requestID, options...)
}

func (r *MockRepo) SaveReservation(ctx context.Context, reservation *inventory.Reservation, options ...core.UpdateOptions) error {
	r.AddCall(ctx, reservation, options)
	return r.SaveReservationFunc(ctx, reservation, options...)
}

func (r *MockRepo) GetReservation(ctx context.Context, ID uint64, options ...core.QueryOptions) (inventory.Reservation, error) {
	r.AddCall(ctx, ID, options)
	return r.GetReservationFunc(ctx, ID, options...)
}

func (r *MockRepo) GetReservations(ctx context.Context, resOptions inventory.GetReservationsOptions, limit, offset int, options ...core.QueryOptions) ([]inventory.Reservation, error) {
	r.AddCall(ctx, resOptions, limit, offset, options)
	return r.GetReservationsFunc(ctx, resOptions, limit, offset, options...)
}

func (r *MockRepo) SaveProduct(ctx context.Context, product inventory.Product, options ...core.UpdateOptions) error {
	r.AddCall(ctx, product, options)
	return r.SaveProductFunc(ctx, product, options...)
}

func (r *MockRepo) GetProduct(ctx context.Context, sku string, options ...core.QueryOptions) (inventory.Product, error) {
	r.AddCall(ctx, sku, options)
	return r.GetProductFunc(ctx, sku, options...)
}

func (r *MockRepo) GetProductInventory(ctx context.Context, sku string, options ...core.QueryOptions) (inventory.ProductInventory, error) {
	r.AddCall(ctx, sku, options)
	return r.GetProductInventoryFunc(ctx, sku, options...)
}

func (r *MockRepo) SaveProductInventory(ctx context.Context, productInventory inventory.ProductInventory, options ...core.UpdateOptions) error {
	r.AddCall(ctx, productInventory, options)
	return r.SaveProductInventoryFunc(ctx, productInventory, options...)
}

func (r *MockRepo) GetAllProductInventory(ctx context.Context, limit int, offset int, options ...core.QueryOptions) ([]inventory.ProductInventory, error) {
	r.AddCall(ctx, limit, offset, options)
	return r.GetAllProductInventoryFunc(ctx, limit, offset, options...)
}

func (r *MockRepo) BeginTransaction(ctx context.Context) (core.Transaction, error) {
	r.AddCall(ctx)
	return r.BeginTransactionFunc(ctx)
}

func (r *MockRepo) GetReservationByRequestID(ctx context.Context, requestId string, options ...core.QueryOptions) (inventory.Reservation, error) {
	r.AddCall(ctx, requestId, options)
	return r.GetReservationByRequestIDFunc(ctx, requestId, options...)
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		SaveProductionEventFunc: func(ctx context.Context, event *inventory.ProductionEvent, options ...core.UpdateOptions) error {
			return nil
		},
		GetProductionEventByRequestIDFunc: func(ctx context.Context, requestID string, options ...core.QueryOptions) (pe inventory.ProductionEvent, err error) {
			return inventory.ProductionEvent{}, nil
		},
		SaveReservationFunc: func(ctx context.Context, reservation *inventory.Reservation, options ...core.UpdateOptions) error {
			return nil
		},
		GetReservationFunc: func(ctx context.Context, ID uint64, options ...core.QueryOptions) (inventory.Reservation, error) {
			return inventory.Reservation{}, nil
		},
		GetReservationsFunc: func(ctx context.Context, resOptions inventory.GetReservationsOptions, limit, offset int, options ...core.QueryOptions) ([]inventory.Reservation, error) {
			return nil, nil
		},
		SaveProductFunc: func(ctx context.Context, product inventory.Product, options ...core.UpdateOptions) error { return nil },
		GetProductFunc: func(ctx context.Context, sku string, options ...core.QueryOptions) (inventory.Product, error) {
			return inventory.Product{}, nil
		},
		GetAllProductInventoryFunc: func(ctx context.Context, limit int, offset int, options ...core.QueryOptions) ([]inventory.ProductInventory, error) {
			return nil, nil
		},
		BeginTransactionFunc: func(ctx context.Context) (core.Transaction, error) { return db.NewMockTransaction(), nil },
		GetReservationByRequestIDFunc: func(ctx context.Context, requestId string, options ...core.QueryOptions) (inventory.Reservation, error) {
			return inventory.Reservation{}, nil
		},
		UpdateReservationFunc: func(ctx context.Context, ID uint64, state inventory.ReserveState, qty int64, options ...core.UpdateOptions) error {
			return nil
		},
		GetProductInventoryFunc: func(ctx context.Context, sku string, options ...core.QueryOptions) (inventory.ProductInventory, error) {
			return inventory.ProductInventory{}, nil
		},
		SaveProductInventoryFunc: func(ctx context.Context, productInventory inventory.ProductInventory, options ...core.UpdateOptions) error {
			return nil
		},
		CallWatcher: test.NewCallWatcher(),
	}
}
