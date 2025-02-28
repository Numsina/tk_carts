package service

import (
	"context"

	"github.com/Numsina/tk_carts/tk_carts_srv/dao"
	"github.com/Numsina/tk_carts/tk_carts_srv/domain"
	"github.com/Numsina/tk_carts/tk_carts_srv/logger"
)

type Cart interface {
	CreateCarts(ctx context.Context, cart domain.Carts, uid int32) error
	DeleteCarts(ctx context.Context, pid, uid int32) error
	UpdateCarts(ctx context.Context, cart domain.Carts, uid int32) error
	ClearCarts(ctx context.Context, uid int32) error
	GetCartsInfo(ctx context.Context, uid int32) ([]domain.Carts, error)
}

type CartService struct {
	dao dao.Cart
	log *logger.Logger
}

func (c *CartService) CreateCarts(ctx context.Context, cart domain.Carts, uid int32) error {
	d := c.toDao(cart)
	d.UserID = uid
	return c.dao.InsertCarts(ctx, d)
}

func (c *CartService) DeleteCarts(ctx context.Context, pid, uid int32) error {
	return c.dao.DeleteCarts(ctx, pid, uid)
}

func (c *CartService) UpdateCarts(ctx context.Context, cart domain.Carts, uid int32) error {
	d := c.toDao(cart)
	d.UserID = uid
	return c.dao.UpdateCarts(ctx, d)
}

func (c *CartService) ClearCarts(ctx context.Context, uid int32) error {
	return c.dao.ClearCarts(ctx, uid)
}

func (c *CartService) GetCartsInfo(ctx context.Context, uid int32) ([]domain.Carts, error) {
	data, err := c.dao.QueryCartsInfo(ctx, uid)
	if err != nil {
		return nil, err
	}

	var ds []domain.Carts
	for _, v := range data {
		ds = append(ds, c.toDomain(v))
	}
	return ds, nil
}

func NewCartRepository(dao dao.Cart, log *logger.Logger) Cart {
	return &CartService{
		dao: dao,
		log: log,
	}
}

func (c *CartService) toDao(data domain.Carts) dao.Carts {
	return dao.Carts{
		Id:      data.Id,
		GoodsID: data.GoodsID,
		Nums:    data.Nums,
		Checked: data.Checked,
	}
}

func (c *CartService) toDomain(data dao.Carts) domain.Carts {
	return domain.Carts{
		Id:      data.Id,
		GoodsID: data.GoodsID,
		Nums:    data.Nums,
		Checked: data.Checked,
	}
}
