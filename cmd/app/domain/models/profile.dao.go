package models

type Profile struct {
	ID          string  `json:"id"`
	FirstName   *string `json:"firstName"`   // Nullable in GraphQL
	LastName    *string `json:"lastName"`    // Nullable in GraphQL
	ProfileImg  *string `json:"profileImg"`  // Nullable in GraphQL
	Email       *string `json:"email"`       // Nullable in GraphQL
	PhoneNumber *string `json:"phoneNumber"` // Nullable in GraphQL
	County      *string `json:"county"`      // Nullable in GraphQL
	Country     *string `json:"country"`     // Nullable in GraphQL
}