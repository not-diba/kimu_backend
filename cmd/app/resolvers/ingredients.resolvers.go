package resolvers

import "kimu_backend/cmd/app/domain/models"

func (r *queryResolver) getIngredientsForRecipe(recipeID string) ([]models.Ingredient, error) {
	rows, err := r.DB.Query(`SELECT name, quantity FROM "Ingredient" WHERE "recipeId" = $1`, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []models.Ingredient
	for rows.Next() {
		var ingredient models.Ingredient
		if err := rows.Scan(&ingredient.Name, &ingredient.Quantity); err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}