package storage

type Argument struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

//TODO May need to add some sort of function to calculate the Value? This may not be needed as this layer has
// no coupling to the contracts themselves. It may be valid to just have the layer with the knowledge of a particular
// type of contract to know how to calculate the Value.

type Function struct {
	Arguments []Argument `json:"arguments"`
}