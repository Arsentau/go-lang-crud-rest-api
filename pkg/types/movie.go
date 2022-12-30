package types

type Movie struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Bought      bool   `json:"director"`
	Description string `json:"description"`
}
