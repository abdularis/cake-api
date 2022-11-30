package core

import (
	"context"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("no record found")
)

type CakeRepository interface {
	FindCakeList(ctx context.Context) ([]Cake, error)
	FindCakeByID(ctx context.Context, cakeID uint) (*Cake, error)
	SaveCake(ctx context.Context, cake *Cake) (*Cake, error)
	DeleteCakeByID(ctx context.Context, cakeID uint) error
}
