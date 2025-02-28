package handler

import (
	"context"

	"github.com/Numsina/tk_carts/tk_carts_srv/domain"
	"github.com/Numsina/tk_carts/tk_carts_srv/pb/v1"
	"github.com/Numsina/tk_carts/tk_carts_srv/service"
)

type CartHandler struct {
	svc service.Cart
	pb.UnimplementedCartServiceServer
}

func NewCartHandler(svc service.Cart) *CartHandler {
	return &CartHandler{svc: svc}

}

func (c *CartHandler) AddItem(ctx context.Context, req *pb.AddItemReq) (*pb.AddItemResp, error) {
	err := c.svc.CreateCarts(ctx, domain.Carts{
		GoodsID: int32(req.Item.ProductId),
		Nums:    req.Item.Quantity,
		Checked: req.Item.Checked,
	}, int32(req.UserId))
	if err != nil {
		return nil, err
	}

	return &pb.AddItemResp{}, nil
}

func (c *CartHandler) UpdateItem(ctx context.Context, req *pb.UpdateItemReq) (*pb.UpdateItemResp, error) {
	err := c.svc.UpdateCarts(ctx, domain.Carts{
		GoodsID: int32(req.Item.ProductId),
		Nums:    req.Item.Quantity,
		Checked: req.Item.Checked,
	}, int32(req.UserId))
	if err != nil {
		return nil, err
	}

	return &pb.UpdateItemResp{}, nil
}

func (c *CartHandler) GetCart(ctx context.Context, req *pb.GetCartReq) (*pb.GetCartResp, error) {
	info, err := c.svc.GetCartsInfo(ctx, int32(req.UserId))
	if err != nil {
		return nil, err
	}
	var data []*pb.CartItem
	for _, v := range info {
		data = append(data, &pb.CartItem{
			ProductId: uint32(v.GoodsID),
			Quantity:  v.Nums,
			Checked:   v.Checked,
		})
	}
	return &pb.GetCartResp{Cart: &pb.Cart{
		Items:  data,
		UserId: req.UserId,
	}}, nil
}

func (c *CartHandler) EmptyCart(ctx context.Context, req *pb.EmptyCartReq) (*pb.EmptyCartResp, error) {
	err := c.svc.ClearCarts(ctx, int32(req.UserId))
	return nil, err
}

func (c *CartHandler) DeleteCart(ctx context.Context, req *pb.DeleteCartReq) (*pb.DeleteCartResp, error) {
	err := c.svc.DeleteCarts(ctx, int32(req.ProductId), int32(req.UserId))
	return nil, err
}
