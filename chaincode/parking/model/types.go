package parkingservice

import (
	"time"
)

type ParkingTime struct {
	Id                string           `json:"id"`
	ParkingStart      time.Time        `json:"parkingStart"`
	ParkingEnd        time.Time        `json:"parkingEnd"`
	ParkingType       string           `json:"parkingType"` //RESERVATION,PARKING
	CostPerMinute     int              `json:"costPerMinute"`
	Cost              int              `json:"cost"`
	Parkingspot       Parkingspot      `json:"parkingspot"`
	Renter            User             `json:"renter"`
	CurrentTimestamps CurrentTimestamp `json:"currentTimestamps"`
}

type Parkingspot struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	CostPerMinute CurrencyAmount `json:"costPerMinute"`
	Owner         User           `json:"owner"`
}

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Balance Balance `json:"balance"`
}

type CurrencyAmount struct {
	Amount       int    `json:"amount"`
	CurrencyName string `json:"currencyName"`
}

type Balance struct {
	Id           string `json:"id"`
	CurrencyName string `json:"currencyName"`
	Amount       int    `json:"amount"`
}

type CurrentTimestamp struct {
	TimeWindow            time.Duration `json:"timeWindow"`
	TransactionTime       time.Time     `json:"transactionTime"`
	TimeWindowCurrentTime time.Time     `json:"timeWindowCurrentTime"`
	TimeServerCurrentTime time.Time     `json:"timeServerCurrentTime"`
	ChaincodeCurrentTime  time.Time     `json:"chaincodeCurrentTime"`
	Errors                []string      `json:"errors"`
}

type HyperledgerFabricTimestamp struct {
	CurrentTime time.Time `json:"currentTime"`
}
