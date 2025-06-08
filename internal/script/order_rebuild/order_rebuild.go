package order_rebuild

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/HeRedBo/pkg/es"
	"github.com/HeRedBo/pkg/httpclient"
	"github.com/HeRedBo/pkg/sign"
	"github.com/HeRedBo/pkg/strutil"
	"net/http"
	"net/url"
	"shop-schedule/config"
	"shop-schedule/global"
	"shop-schedule/internal/po"
	"time"
)

type Response struct {
	Code int              `json:"status"`
	Msg  string           `json:"msg"`
	Data CursorResultList `json:"data"`
}
type CursorResultList struct {
	Content    []po.StoreOrder `json:"content"`
	NextID     int64           `json:"next_id"`
	ExtendData interface{}     `json:"extendData"`
}

const GetOrdersUri = "/dev/v1/orders/user"

func Rebuild(userID string) {
	t := time.Now()
	orderSuffixLen := 4
	nextID := "0"
	for {
		global.LOG.Warn("start rebuild")
		params := url.Values{}
		params.Add("next_id", nextID)
		uri := GetOrdersUri + "/" + userID
		authorization, date, err := sign.New(global.CONFIG.Api.OrderAK, global.CONFIG.Api.OrderSK, time.Minute*3).Generate(uri, http.MethodGet, params)
		if err != nil {
			global.LOG.Error("sign.New error", err, params)
			return
		}
		headerAuth := httpclient.WithHeader(config.HeaderAuthField, authorization)
		headerAuthDate := httpclient.WithHeader(config.HeaderAuthDateField, date)
		code, body, err := httpclient.Get(global.CONFIG.Api.ShopHost+uri, params,
			headerAuth, headerAuthDate, httpclient.WithTTL(time.Second*5))
		if err != nil || code != http.StatusOK {
			global.LOG.Errorf("get user order error %v ; params %+v ; "+
				"code %d ; body: %s", err, params, code, body)
			return
		}
		fmt.Println("body", string(body))
		res := Response{}
		json.Unmarshal(body, &res)
		if res.Code != 200 {
			global.LOG.Errorf("get user order error %v ; params %+v ; "+
				"code %d ; body: %s", err, params, code, body)
			return
		}

		for _, order := range res.Data.Content {
			index := po.OrderIndex{
				OrderId:       order.OrderId,
				OrderIdSuffix: order.OrderId,
				Names:         []string{},
				ProductIds:    []int64{},
				Uid:           order.Uid,
				PayTime:       order.PayTime,
				PayType:       order.PayType,
				RefundStatus:  order.RefundStatus,
				ShippingType:  order.ShippingType,
				OrderStatus:   order.OrderStatus,
				CreateTime:    order.CreateTime,
				UpdateTime:    order.UpdateTime,
			}
			if len(order.OrderId) >= orderSuffixLen {
				index.OrderIdSuffix = order.OrderId[len(order.OrderId)-4:]
			}
			for _, cart := range order.CartInfo {
				index.Names = append(index.Names, cart.ProductInfo.StoreName)
				index.ProductIds = append(index.ProductIds, cart.ProductInfo.Id)
			}
			fmt.Println("index", index)
			esClient := es.GetClient(es.DefaultClient)
			esClient.Create(context.Background(), global.IndexName, index.OrderId, strutil.Int64ToString(index.Uid), index)
		}

		if res.Data.NextID == 0 {
			break
		} else {
			nextID = strutil.Int64ToString(res.Data.NextID)
		}
	}
	global.LOG.Warn("rebuild over cost", time.Since(t).Seconds())
}
