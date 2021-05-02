package entities

// Book Construct your model under entities.
type Book struct {
	Title  string `json:"title" binding:"required,min=2" bson:"title"`
	Author string `json:"author" bson:"author,omitempty"`
}
