package model

// Car represents car entity
type Car struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Brand string `json:"brand"`
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}
