package models

type Issue struct {
	Description string
	Severity    string
	File        string
	Line        int
	Column      int
}
