package models

type Card struct {
	Face  bool
	Name  string
	Value int
}

var cards = []Card{
	{Face: false, Name: "2", Value: 2},
	{Face: false, Name: "3", Value: 3},
	{Face: false, Name: "4", Value: 4},
	{Face: false, Name: "5", Value: 5},
	{Face: false, Name: "6", Value: 6},
	{Face: false, Name: "7", Value: 7},
	{Face: false, Name: "8", Value: 8},
	{Face: false, Name: "9", Value: 9},
	{Face: true, Name: "10", Value: 10},
	{Face: true, Name: "J", Value: 10},
	{Face: true, Name: "Q", Value: 10},
	{Face: true, Name: "K", Value: 10},
	{Face: true, Name: "A", Value: 1},
}
