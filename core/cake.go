package core

import (
	"errors"
	"fmt"
	"time"
)

var ErrDataValidation = errors.New("data validation err")

type Cake struct {
	ID          uint
	Title       string
	Description string
	Rating      float32
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c *Cake) Validate() error {
	switch {
	case c.Title == "":
		return fmt.Errorf("%w: title cannot be empty", ErrDataValidation)
	case c.Description == "":
		return fmt.Errorf("%w: description cannot be empty", ErrDataValidation)
	case c.Rating < 0 || c.Rating > 10:
		return fmt.Errorf("%w: rating should be in range 0 - 10", ErrDataValidation)
	case c.Image == "":
		return fmt.Errorf("%w: image url cannot be empty", ErrDataValidation)
	default:
		return nil
	}
}

type CreateCakeRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
}

func (create *CreateCakeRequest) ToCake() (*Cake, error) {
	cake := &Cake{
		Title:       create.Title,
		Description: create.Description,
		Rating:      create.Rating,
		Image:       create.Image,
	}
	if err := cake.Validate(); err != nil {
		return nil, err
	}
	return cake, nil
}

type CakeUpdateRequest struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Rating      *float32 `json:"rating"`
	Image       *string  `json:"image"`
}

func (update *CakeUpdateRequest) ApplyUpdateTo(cake *Cake) error {
	if update.Title != nil {
		cake.Title = *update.Title
	}
	if update.Description != nil {
		cake.Description = *update.Description
	}
	if update.Rating != nil {
		cake.Rating = *update.Rating
	}
	if update.Image != nil {
		cake.Image = *update.Image
	}
	return cake.Validate()
}
