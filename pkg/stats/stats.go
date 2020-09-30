package stats

import (
	"github.com/khushbakhtmahkamov/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var avg types.Money = 0
	var count int64 = 0
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		count += 1
		avg += payment.Amount
	}
	return types.Money(int64(avg) / count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var avg types.Money = 0
	for _, payment := range payments {

		if payment.Status == types.StatusFail {
			continue
		}
		
		if payment.Category == category {
			avg += payment.Amount
		}
	}
	return types.Money(avg)
}

func CategoriesAvg(payments []types.Payment) ( map[types.Category]types.Money) {

	mp := make(map[types.Category]types.Money)
	for _, v := range payments {

		if _, er := mp[v.Category]; er {
			continue
		}
		filtered := FilterByCategory(payments, v.Category)
		mp[v.Category]=Avg(filtered)
	}


	return mp
}


func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment{

	var filtered []types.Payment

	for _, v := range payments {
		if v.Category == category {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	mp := map[types.Category]types.Money{}
 
	if len(first)>=len(second){
		 for k := range first {
			 mp[k]=second[k]-first[k]
		 }
		 return mp
	}
	 for k := range second {
		 mp[k]=second[k]-first[k]
	 }
 
	return mp
 }
