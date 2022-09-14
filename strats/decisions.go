package strats

type Decision string

// Actions
const (
	Hit     Decision = "H"
	Stand   Decision = "S"
	Double  Decision = "D"
	Split   Decision = "Y"
	noSplit Decision = "N"
)
