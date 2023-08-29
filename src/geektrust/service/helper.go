package service

import (
	"geektrust/config"
	"geektrust/model"
)

func getFareAndDiscount(cardsArr []model.Card, fare int, station, cardName string) (int, int) {
	discount := 0
	var emptyString string
	for i := range cardsArr {
		if cardsArr[i].Name != cardName {
			continue
		}
		switch cardsArr[i].PreviosStation {
		case emptyString:
			switch cardsArr[i].CurrentStation {
			case emptyString:
				break
			default:
				fare /= 2
				discount = fare
				cardsArr[i].PreviosStation = cardsArr[i].CurrentStation
			}
		case station:
			cardsArr[i].PreviosStation = emptyString
		default:
			break
		}
		cardsArr[i].CurrentStation = station

		bal := cardsArr[i].Balance - fare
		cardsArr[i].Balance -= fare
		if bal < 0 {
			cardsArr[i].Balance = 0
			fare += (-bal) * config.InsufficientBalFeePerc / 100
		}
		break
	}
	return fare, discount
}
