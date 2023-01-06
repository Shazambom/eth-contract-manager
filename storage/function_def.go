package storage

type Argument struct {
	Name string `json:"name"`
	Type string `json:"type"`
}


type Function struct {
	Arguments []Argument `json:"arguments"`
}