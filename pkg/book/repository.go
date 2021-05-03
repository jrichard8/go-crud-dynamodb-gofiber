package book

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go-crud-dynamodb-gofiber/pkg/entities"
	"strconv"
)

type Repository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	ReadBooks() (*[]entities.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(book *entities.BookKey) (*entities.Book, error)
}

func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {

	b, err := attributevalue.MarshalMap(book)
	if err != nil {
		return nil, err
	}
	param := &dynamodb.PutItemInput{
		Item:         b,
		TableName:    aws.String(r.table),
		ReturnValues: types.ReturnValueNone,
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
	param := &dynamodb.ScanInput{
		TableName: aws.String(r.table),
	}
	scan, err := r.Dyn.Scan(context.TODO(), param)
	if err != nil {
		return nil, err
	}

	books := make([]entities.Book, len(scan.Items))

	for i, item := range scan.Items {
		book := entities.Book{}
		err := attributevalue.UnmarshalMap(item, &book)
		if err != nil {
			return nil, err
		}
		books[i] = book
	}
	return &books, nil
}

func (r *repository) UpdateBook(book *entities.Book) (*entities.Book, error) {
	bKey := entities.BookKey{
		Title:  book.Title,
		Author: book.Author,
	}
	key, err := attributevalue.MarshalMap(&bKey)
	if err != nil {
		return nil, err
	}
	irate, _ := strconv.ParseFloat(book.Rating, 64)
	rating := entities.Rating{
		Rate: irate,
	}

	rate, err := attributevalue.MarshalMap(&rating)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		ExpressionAttributeValues: rate,
		UpdateExpression:          aws.String("SET Rating = :Rate"),
		TableName:                 aws.String(r.table),
		ReturnValues:              types.ReturnValueUpdatedNew,
	}

	item, err := r.Dyn.UpdateItem(context.TODO(), input)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	res := entities.Book{}
	err = attributevalue.UnmarshalMap(item.Attributes, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *repository) DeleteBook(bKey *entities.BookKey) (*entities.Book, error) {

	key, err := attributevalue.MarshalMap(bKey)
	if err != nil {
		return nil, err
	}

	param := &dynamodb.DeleteItemInput{
		Key:          key,
		ReturnValues: types.ReturnValueAllOld,
	}

	item, err := r.Dyn.DeleteItem(context.TODO(), param)
	if err != nil {
		return nil, err
	}
	deletedBook := entities.Book{}
	err = attributevalue.UnmarshalMap(item.Attributes, &deletedBook)
	if err != nil {
		return nil, err
	}
	return &deletedBook, nil
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
