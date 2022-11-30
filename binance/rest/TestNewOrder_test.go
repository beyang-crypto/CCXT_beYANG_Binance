package rest

import (
	"testing"

	config "github.com/beyang-crypto/CCXT_beYANG_Binance/binance"
)

func TestTestNewOrder(t *testing.T) {

	path := "../../config-prod.yaml"

	conf := config.GetConfig(path)

	cfg := &Configuration{
		Addr:      BaseEndpoint,
		ApiKey:    conf.Api.Key,
		SecretKey: conf.Api.Secret,
		DebugMode: true,
	}

	b := New(cfg)

	t.Log("Тестирование TestNewOrder при разных входных данных")
	{
		timeInForce := TimeInForceGTC
		side := OrderSideBUY
		orderType := OrderTypeLimit
		t.Logf("\tTimeInForce = %s Side = %s", timeInForce, side)
		{
			quantity := 0.01
			price := 9000.0
			newClientOrderId := "my_order_id_1"

			parm := TestNewOrderParam{
				Symbol:           "BTCUSDT",
				Side:             side,
				Type:             orderType,
				TimeInForce:      &timeInForce,
				Quantity:         &quantity,
				NewClientOrderId: &newClientOrderId,
				Price:            &price,
			}

			want := TestNewOrderResp{}

			ans := b.TestNewOrder(parm)

			if ans != want {
				t.Fatalf("TestTestNewOrder got %d, wanted %d", ans, want)
			}
		}

	}

}
