package models

type Recipe struct {
	ID           string       `json:"id"`
	RecipeName   string       `json:"recipeName"`
	CategoryName string       `json:"categoryName"`
	Duration     string       `json:"duration"`
	People       int          `json:"people"`
	Description  string       `json:"description"`
	Amount       int          `json:"amount"`
	ImageUrl     string       `json:"imageUrl"`
	Instructions []string     `json:"instructions"`
	Ingredients  []Ingredient `json:"ingredients"`
	Nutrition    []Nutrition  `json:"nutrition"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Nutrition struct {
	NutritionItem string `json:"nutritionItem"`
	Quantity      string `json:"quantity"`
}
