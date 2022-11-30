package core

import (
	"testing"
)

type testData struct {
	Cake  Cake
	Valid bool
}

func TestCake_Validate(t *testing.T) {
	testingData := []testData{
		{
			Cake:  Cake{},
			Valid: false,
		},
		{
			Cake: Cake{
				Title:       "title here",
				Description: "",
			},
			Valid: false,
		},
		{
			Cake: Cake{
				Title:       "title here",
				Description: "description",
				Rating:      -1,
			},
			Valid: false,
		},
		{
			Cake: Cake{
				Title:       "title here",
				Description: "description",
				Rating:      10.2,
			},
			Valid: false,
		},
		{
			Cake: Cake{
				Title:       "title here",
				Description: "description",
				Rating:      2,
				Image:       "",
			},
			Valid: false,
		},
		{
			Cake: Cake{
				Title:       "title here",
				Description: "description",
				Rating:      0,
				Image:       "image url",
			},
			Valid: true,
		},
	}

	for _, td := range testingData {
		err := td.Cake.Validate()
		if td.Valid {
			if err != nil {
				t.Errorf("data should be valid but error returned: %s", err)
			}
		} else {
			if err == nil {
				t.Errorf("data should error but no error returned")
			}
		}
	}
}

func TestCakeUpdateRequest_ApplyUpdateTo(t *testing.T) {
	update := CakeUpdateRequest{}
	cake := Cake{
		Title:       "title here",
		Description: "description",
		Rating:      0,
		Image:       "image url",
	}
	oriCake := Cake{
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
	}

	err := update.ApplyUpdateTo(&cake)
	if err != nil {
		t.Errorf(err.Error())
	}

	if cake != oriCake {
		t.Errorf("cake object should have not change anything")
	}

	title := "title updated"
	description := "desc updated"
	rating := float32(2.3)
	image := "image updated"

	update.Title = &title
	update.Description = &description
	update.Rating = &rating
	update.Image = &image
	err = update.ApplyUpdateTo(&cake)
	if err != nil {
		t.Errorf(err.Error())
	}

	if cake.Title != title {
		t.Errorf("cake title should be updated to %s", title)
	}
	if cake.Description != description {
		t.Errorf("cake description should be updated to %s", title)
	}
	if cake.Rating != rating {
		t.Errorf("cake rating should be updated to %s", title)
	}
	if cake.Image != image {
		t.Errorf("cake image should be updated to %s", title)
	}
}
