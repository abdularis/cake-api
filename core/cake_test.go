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

func TestCreateCakeRequest_ToCake(t *testing.T) {

}

func TestCakeUpdateRequest_ApplyUpdateTo(t *testing.T) {

}
