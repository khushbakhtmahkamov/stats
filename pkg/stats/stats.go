package stats

import (
	"github.com/khushbakhtmahkamov/bank/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var avg types.Money = 0
	var count int64 = 0
	for _, payment := range payments {
		count += 1
		avg += payment.Amount
	}
	return types.Money(int64(avg) / count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var avg types.Money = 0
	for _, payment := range payments {
		if payment.Category == category {
			avg += payment.Amount
		}
	}
	return types.Money(avg)
}