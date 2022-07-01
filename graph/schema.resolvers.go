package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/someday-94/gqlgen-todos/graph/generated"
	"github.com/someday-94/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	book := &model.Book{
		ID:      fmt.Sprintf("T%d", rand.Int()),
		Title:   input.Title,
		Content: input.Content,
	}
	r.books = append(r.books, book)
	return book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
