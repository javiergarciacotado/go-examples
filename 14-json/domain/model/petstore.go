package model

type Petstore struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Dogs     []*Pet `json:"dogs,omitempty"`
	Cats     []*Pet `json:"cats,omitempty"`
}
