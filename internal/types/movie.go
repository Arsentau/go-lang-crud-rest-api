package types

type Movie struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Bought      bool   `json:"bought"`
	Description string `json:"description"`
}
