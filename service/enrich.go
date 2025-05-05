package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AgeResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GenderResponse struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type NationalityResponse struct {
	Name        string `json:"name"`
	Nationality []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"nationality"`
}

// Получение возраста через API
func GetAge(name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result AgeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Age, nil
}

// Получение пола через API
func GetGender(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result GenderResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Gender, nil
}

// Получение национальности через API
func GetNationality(name string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result NationalityResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var nationalities []string
	for _, nationality := range result.Nationality {
		nationalities = append(nationalities, nationality.CountryID)
	}

	return nationalities, nil
}
