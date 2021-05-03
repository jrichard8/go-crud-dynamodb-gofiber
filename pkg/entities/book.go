package entities

// Book Construct your model under entities.
type Book struct {
	Title  string `json:"Title" binding:"required,min=2" bson:"Title"`
	Author string `json:"Author" bson:"Author,omitempty"`
	Rating string `json:"Rating" bson:"Rating,omitempty"`
}

type BookKey struct {
	Title  string `json:"Title" binding:"required,min=2" bson:"Title"`
	Author string `json:"Author" bson:"Author,omitempty"`
}

type Rating struct {
	Rate float64 `dynamodbav:":Rate"`
}
