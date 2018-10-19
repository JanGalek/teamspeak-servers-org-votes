package main

type Server struct {
	Name string `json:"name"`
	Month string `json:"month"`
	Voters []Voter `json:"voters"`
}