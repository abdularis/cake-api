package core

import (
	"context"
)

type CakeUseCase interface {
	GetListCakes(ctx context.Context) ([]Cake, error)
	GetCakeDetail(ctx context.Context, cakeID uint) (*Cake, error)
	CreateNewCake(ctx context.Context, request CreateCakeRequest) (*Cake, error)
	UpdateCake(ctx context.Context, cakeID uint, request CakeUpdateRequest) (*Cake, error)
	DeleteCakeByID(ctx context.Context, cakeID uint) error
}

func NewCakeUseCase(repository CakeRepository) CakeUseCase {
	return &cakeUseCaseImpl{repo: repository}
}

type cakeUseCaseImpl struct {
	repo CakeRepository
}

func (u *cakeUseCaseImpl) GetListCakes(ctx context.Context) ([]Cake, error) {
	return u.repo.FindCakeList(ctx)
}

func (u *cakeUseCaseImpl) GetCakeDetail(ctx context.Context, cakeID uint) (*Cake, error) {
	return u.repo.FindCakeByID(ctx, cakeID)
}

func (u *cakeUseCaseImpl) CreateNewCake(ctx context.Context, request CreateCakeRequest) (*Cake, error) {
	cake, err := request.ToCake()
	if err != nil {
		return nil, err
	}
	return u.repo.SaveCake(ctx, cake)
}

func (u *cakeUseCaseImpl) UpdateCake(ctx context.Context, cakeID uint, request CakeUpdateRequest) (*Cake, error) {
	cake, err := u.repo.FindCakeByID(ctx, cakeID)
	if err != nil {
		return nil, err
	}
	err = request.ApplyUpdateTo(cake)
	if err != nil {
		return nil, err
	}
	return u.repo.SaveCake(ctx, cake)
}

func (u *cakeUseCaseImpl) DeleteCakeByID(ctx context.Context, cakeID uint) error {
	return u.repo.DeleteCakeByID(ctx, cakeID)
}
