// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type BookTripInput struct {
	Pickup      *LocationInput `json:"pickup"`
	Destination *LocationInput `json:"destination"`
	CabType     CabType        `json:"cabType"`
}

type CancelTrip struct {
	Cancel bool   `json:"cancel"`
	Reason string `json:"reason"`
}

type Cash struct {
	Currency Currency `json:"currency"`
	Amount   float64  `json:"amount"`
}

type Location struct {
	Lat float64 `json:"Lat"`
	Lon float64 `json:"Lon"`
}

type LocationInput struct {
	Lat float64 `json:"Lat"`
	Lon float64 `json:"Lon"`
}

type NearbyCab struct {
	Event    NearbyCabEvent `json:"event"`
	Location *Location      `json:"location"`
	CabID    string         `json:"cabID"`
}

type NearbyCabInput struct {
	CurLocation *LocationInput `json:"curLocation"`
	Radius      float64        `json:"radius"`
	Type        CabType        `json:"type"`
}

type RegisterInput struct {
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirmPassword"`
	FirstName       string  `json:"firstName"`
	LastName        *string `json:"lastName"`
}

type CabType string

const (
	CabTypeSuv    CabType = "SUV"
	CabTypeMicro  CabType = "Micro"
	CabTypeMini   CabType = "Mini"
	CabTypeLuxury CabType = "Luxury"
)

var AllCabType = []CabType{
	CabTypeSuv,
	CabTypeMicro,
	CabTypeMini,
	CabTypeLuxury,
}

func (e CabType) IsValid() bool {
	switch e {
	case CabTypeSuv, CabTypeMicro, CabTypeMini, CabTypeLuxury:
		return true
	}
	return false
}

func (e CabType) String() string {
	return string(e)
}

func (e *CabType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CabType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CabType", str)
	}
	return nil
}

func (e CabType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Currency string

const (
	CurrencyUsd Currency = "USD"
	CurrencyInr Currency = "INR"
	CurrencyEur Currency = "EUR"
)

var AllCurrency = []Currency{
	CurrencyUsd,
	CurrencyInr,
	CurrencyEur,
}

func (e Currency) IsValid() bool {
	switch e {
	case CurrencyUsd, CurrencyInr, CurrencyEur:
		return true
	}
	return false
}

func (e Currency) String() string {
	return string(e)
}

func (e *Currency) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Currency(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Currency", str)
	}
	return nil
}

func (e Currency) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Gender string

const (
	GenderM Gender = "M"
	GenderF Gender = "F"
	GenderO Gender = "O"
)

var AllGender = []Gender{
	GenderM,
	GenderF,
	GenderO,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderM, GenderF, GenderO:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NearbyCabEvent string

const (
	NearbyCabEventEnter NearbyCabEvent = "ENTER"
	NearbyCabEventLeave NearbyCabEvent = "LEAVE"
	NearbyCabEventMove  NearbyCabEvent = "MOVE"
)

var AllNearbyCabEvent = []NearbyCabEvent{
	NearbyCabEventEnter,
	NearbyCabEventLeave,
	NearbyCabEventMove,
}

func (e NearbyCabEvent) IsValid() bool {
	switch e {
	case NearbyCabEventEnter, NearbyCabEventLeave, NearbyCabEventMove:
		return true
	}
	return false
}

func (e NearbyCabEvent) String() string {
	return string(e)
}

func (e *NearbyCabEvent) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NearbyCabEvent(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NearbyCabEvent", str)
	}
	return nil
}

func (e NearbyCabEvent) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TripsInput string

const (
	TripsInputAll       TripsInput = "ALL"
	TripsInputCanceled  TripsInput = "CANCELED"
	TripsInputCompleted TripsInput = "COMPLETED"
	TripsInputActive    TripsInput = "ACTIVE"
)

var AllTripsInput = []TripsInput{
	TripsInputAll,
	TripsInputCanceled,
	TripsInputCompleted,
	TripsInputActive,
}

func (e TripsInput) IsValid() bool {
	switch e {
	case TripsInputAll, TripsInputCanceled, TripsInputCompleted, TripsInputActive:
		return true
	}
	return false
}

func (e TripsInput) String() string {
	return string(e)
}

func (e *TripsInput) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TripsInput(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TripsInput", str)
	}
	return nil
}

func (e TripsInput) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
