package core

import (
	"context"
	"errors"
)

var (
	ErrGeneralDB      = errors.New("general db error")
	ErrRecordNotFound = errors.New("record not found")
)

type CakeRepository interface {
	FindCakeList(ctx context.Context) ([]Cake, error)
	FindCakeByID(ctx context.Context, cakeID uint) (*Cake, error)
	SaveCake(ctx context.Context, cake *Cake) (*Cake, error)
	DeleteCakeByID(ctx context.Context, cakeID uint) error
}
