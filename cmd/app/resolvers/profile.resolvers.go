package resolvers

import (
	"context"
	"database/sql"
	"kimu_backend/cmd/app/domain/models"
)

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context) (*models.Profile, error) {
	row := r.DB.QueryRow(`SELECT id, "firstName", "lastName", "profileImg", "email", "phoneNumber", "county", "country", "createdAt" FROM "Profile" LIMIT 1`)

	profile := &models.Profile{}

	err := row.Scan(
		&profile.ID,
		&profile.FirstName,
		&profile.LastName,
		&profile.ProfileImg,
		&profile.Email,
		&profile.PhoneNumber,
		&profile.County,
		&profile.Country,
		&profile.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return profile, nil
}
