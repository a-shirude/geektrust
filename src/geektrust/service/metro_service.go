package service

import (
	"fmt"
	"sort"
	"strconv"

	"geektrust/config"
	"geektrust/model"
)

type MetroCardServiceInterface interface {
	MetroCard(input [][]string)
}

type MetroCardServiceImpl struct{}

func NewMetroCardServiceImpl() MetroCardServiceInterface {
	return &MetroCardServiceImpl{}
}

func (s *MetroCardServiceImpl) MetroCard(input [][]string) {
	var cardsArr []model.Card
	stationArr := initStationArr()

	for _, inp := range input {
		switch inp[0] {
		case config.Balance:
			bal, _ := strconv.Atoi(inp[2])
			name := inp[1]
			card := initCard(bal, name)
			cardsArr = append(cardsArr, card)

		case config.CheckIn:
			passType := inp[2]
			station := inp[3]
			cardName := inp[1]
			checkIn(stationArr, cardsArr, station, cardName, passType)

		case config.PrintSummary:
			print(stationArr)
		}
	}
}

func initCard(bal int, name string) model.Card {
	card := model.Card{
		Name:    name,
		Balance: bal,
	}
	return card
}

func initStationArr() []model.Station {
	var stationArr []model.Station
	var passArr []model.Passenger
	stationList := [...]string{config.Central, config.Airport}
	for i := range stationList {
		passArr := append(passArr,
			model.Passenger{Type: config.Adult, Fare: config.AdultFare},
			model.Passenger{Type: config.Kid, Fare: config.KidFare},
			model.Passenger{Type: config.SenioCitizen, Fare: config.SenioCitizenFare},
		)
		stationArr = append(stationArr,
			model.Station{Name: stationList[i], Passenger: passArr},
		)
	}
	return stationArr
}

func checkIn(stationArr []model.Station, cardsArr []model.Card, station, cardName, passType string) {
	var fare, discount int
	for i := range stationArr {
		if stationArr[i].Name != station {
			continue
		}
		for j := range stationArr[i].Passenger {
			if stationArr[i].Passenger[j].Type != passType {
				continue
			}
			fare = stationArr[i].Passenger[j].Fare
			stationArr[i].Passenger[j].Count += 1
		}
		fare, discount = getFareAndDiscount(cardsArr, fare, station, cardName)
		stationArr[i].Collection += fare
		stationArr[i].Discount += discount
	}
}

func print(stationArr []model.Station) {
	for _, st := range stationArr {
		sort.Slice(st.Passenger[:], func(i, j int) bool {
			return st.Passenger[i].Count > st.Passenger[j].Count
		})
		fmt.Println(config.TotalCollection, st.Name, st.Collection, st.Discount)
		fmt.Println(config.PassTypeSummary)
		for i := range st.Passenger {
			if st.Passenger[i].Count != 0 {
				fmt.Println(st.Passenger[i].Type, st.Passenger[i].Count)
			}
		}
	}
}
