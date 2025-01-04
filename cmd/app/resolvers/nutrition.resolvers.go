package resolvers

import "kimu_backend/cmd/app/domain/models"

func (r *queryResolver) getNutritionForRecipe(recipeID string) ([]models.Nutrition, error) {
	rows, err := r.DB.Query(`SELECT "nutritionItem", quantity FROM "Nutrition" WHERE "recipeId" = $1`, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nutrition []models.Nutrition
	for rows.Next() {
		var nutrion models.Nutrition
		if err := rows.Scan(&nutrion.NutritionItem, &nutrion.Quantity); err != nil {
			return nil, err
		}
		nutrition = append(nutrition, nutrion)
	}
	return nutrition, nil
}