package queries

import (
	"SOCIAL_MEDIA_APP/platform/database/postgress"
	"SOCIAL_MEDIA_APP/platform/database/postgress/models"
	"context"
	"database/sql"
	"fmt"
)

var (
	ctx context.Context
)

// CheckIfUserNameExists checks if a userName exists in the users table and returns the userid if it does.
func CheckIfUserNameExists(userName string) (bool, string, error) {
	DBInstance, err := postgress.GetPostgressDBInstance()
	if err != nil {
		return false, "", nil
	}

	var result models.IsUserNameExists
	query := `SELECT COUNT(*) AS count, MAX(userid) AS userid FROM users WHERE userName = $1 `

	err = DBInstance.GetContext(ctx, &result, query, userName)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		return false, "", fmt.Errorf("query failed: %w", err)
	}

	return result.Count > 0, result.UserID.String, nil
}
