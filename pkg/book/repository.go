package book

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go-crud-dynamodb-gofiber/pkg/entities"
)

type Repository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	ReadBooks() (*[]entities.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(book *entities.Book) error
}

func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {

	b, err := attributevalue.MarshalMap(book)
	if err != nil {
		return nil, err
	}
	param := &dynamodb.PutItemInput{
		Item:         b,
		TableName:    aws.String(r.table),
		ReturnValues: types.ReturnValueAllNew,
	}
	item, err := r.Dyn.PutItem(context.TODO(), param)
	if err != nil {
		return nil, err
	}
	result := entities.Book{}
	err = attributevalue.UnmarshalMap(item.Attributes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) ReadBooks() (*[]entities.Book, error) {
	return nil, nil
}

func (r *repository) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return nil, nil
}

func (r *repository) DeleteBook(book *entities.Book) error {
	return nil
}

type repository struct {
	Dyn   *dynamodb.Client
	table string
}

//NewRepo is the single instance repo that is being created.
func NewRepo(dyn *dynamodb.Client, table string) Repository {
	return &repository{
		Dyn:   dyn,
		table: table,
	}
}
