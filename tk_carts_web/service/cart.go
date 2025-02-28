package service

import (
	"context"

	"github.com/Numsina/tk_carts/tk_carts_web/domain"
	"github.com/Numsina/tk_carts/tk_carts_web/pb/v1"
)

type Cart interface {
	CreateCarts(ctx context.Context, cart domain.Carts, uid int32) error
	DeleteCarts(ctx context.Context, pid, uid int32) error
	UpdateCarts(ctx context.Context, cart domain.Carts, uid int32) error
	ClearCarts(ctx context.Context, uid int32) error
	GetCartsInfo(ctx context.Context, uid int32) ([]domain.Carts, error)
}

type CartRepository struct {
	client pb.CartServiceClient
}

func (c *CartRepository) CreateCarts(ctx context.Context, cart domain.Carts, uid int32) error {
	_, err := c.client.AddItem(ctx, &pb.AddItemReq{
		UserId: uint32(uid),
		Item: &pb.CartItem{
			ProductId: uint32(cart.GoodsID),
			Quantity:  cart.Nums,
			Checked:   cart.Checked,
		},
	})
	return err
}

func (c *CartRepository) DeleteCarts(ctx context.Context, pid, uid int32) error {
	_, err := c.client.DeleteCart(ctx, &pb.DeleteCartReq{
		UserId:    uint32(uid),
		ProductId: uint32(pid),
	})
	return err
}

func (c *CartRepository) UpdateCarts(ctx context.Context, cart domain.Carts, uid int32) error {
	_, err := c.client.UpdateItem(ctx, &pb.UpdateItemReq{
		UserId: uint32(uid),
		Item: &pb.CartItem{
			ProductId: uint32(cart.GoodsID),
			Quantity:  cart.Nums,
			Checked:   cart.Checked,
		},
	})
	return err
}

func (c *CartRepository) ClearCarts(ctx context.Context, uid int32) error {
	_, err := c.client.EmptyCart(ctx, &pb.EmptyCartReq{
		UserId: uint32(uid),
	})
	return err
}

func (c *CartRepository) GetCartsInfo(ctx context.Context, uid int32) ([]domain.Carts, error) {
	cart, err := c.client.GetCart(ctx, &pb.GetCartReq{
		UserId: uint32(uid),
	})

	if err != nil {
		return nil, err
	}

	var data []domain.Carts
	for _, v := range cart.Cart.GetItems() {
		data = append(data, domain.Carts{
			GoodsID: int32(v.ProductId),
			Nums:    v.Quantity,
			Checked: v.Checked,
		})
	}
	return data, nil
}

func NewCartRepository(client pb.CartServiceClient) Cart {
	return &CartRepository{
		client: client,
	}
}
