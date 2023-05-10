package interfaces

import (
	"context"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type OrderService interface {
	CheckoutOrder(ctx context.Context, userId uint) (checkOut response.CartResp, err error)
	PlaceOrderByCOD(ctx context.Context, userId uint) (shopOrder response.ShopOrder, err error)
	GetOrderHistory(ctx context.Context, page request.ReqPagination, userId uint) (orderHisory []response.OrderHistory, err error)
}
