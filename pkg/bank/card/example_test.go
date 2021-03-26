package card
import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	card := []types.Card{
		{
			PAN: "5058 1111 1111 8888",
			Balance: 10_000,
			Active:  true,
		},
	}
	for _, carts:=range PaymentSources(card){
		fmt.Println(carts.Number)
	}
//Output:
//5058 1111 1111 8888
}