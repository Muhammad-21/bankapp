package card
import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	card := []types.Card{
		{
			Balance: 10_000,
			Active:  true,
		},
		{
			Balance: 20_000,
			Active:  false,
		},
	}
	//var cart []types.PaymentSource
	for _, carts:=range PaymentSources(card){
	fmt.Println(carts.Number)
	}
//Output:
//5058 xxxx xxxx 8888
}