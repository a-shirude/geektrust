package service

import (
	"testing"

	"geektrust/config"
	"geektrust/model"
)

func Test_getFareAndDiscount(t *testing.T) {
	type args struct {
		cardsArr []model.Card
		fare     int
		station  string
		cardName string
	}
	tests := []struct {
		name     string
		args     args
		wantFare int
		wantDisc int
	}{
		{
			name: "start from central ",
			args: args{
				cardsArr: []model.Card{
					{
						Name:    "MC1",
						Balance: 500,
					},
				},
				fare:     200,
				station:  config.Central,
				cardName: "MC1",
			},
			wantFare: 200,
			wantDisc: 0,
		},
		{
			name: "return journey with discount",
			args: args{
				cardsArr: []model.Card{
					{
						Name:           "MC1",
						Balance:        300,
						CurrentStation: config.Central,
					},
				},
				fare:     200,
				station:  config.Airport,
				cardName: "MC1",
			},
			wantFare: 100,
			wantDisc: 100,
		},
		{
			name: "insufficent balance",
			args: args{
				cardsArr: []model.Card{
					{
						Name:    "MC1",
						Balance: 100,
					},
				},
				fare:     200,
				station:  config.Central,
				cardName: "MC1",
			},
			wantFare: 202,
			wantDisc: 0,
		},
		{
			name: "3rd time journey",
			args: args{
				cardsArr: []model.Card{
					{
						Name:           "MC1",
						Balance:        300,
						CurrentStation: config.Airport,
						PreviosStation: config.Central,
					},
				},
				fare:     200,
				station:  config.Central,
				cardName: "MC1",
			},
			wantFare: 200,
			wantDisc: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getFareAndDiscount(tt.args.cardsArr, tt.args.fare, tt.args.station, tt.args.cardName)
			if got != tt.wantFare {
				t.Errorf("getFareAndDiscount() got = %v, want %v", got, tt.wantFare)
			}
			if got1 != tt.wantDisc {
				t.Errorf("getFareAndDiscount() got1 = %v, want %v", got1, tt.wantDisc)
			}
		})
	}
}
