package stats

import (
	"fmt"
	"github.com/khushbakhtmahkamov/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "qwerty",
		},
		{
			ID:       2,
			Amount:   200,
			Category: "qwerty",
		},
		{
			ID:       1,
			Amount:   300,
			Category: "qwerty",
		},
	}
	fmt.Println(Avg(payments))
	//Output: 200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "kart",
		},
		{
			ID:       2,
			Amount:   200,
			Category: "kart1",
		},
		{
			ID:       1,
			Amount:   300,
			Category: "kart",
		},
	}
	var category types.Category = "kart"
	fmt.Println(TotalInCategory(payments,category))
	//Output: 400
}