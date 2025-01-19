package schema

import (
	"csv-app/models"
	"database/sql"
	"strings"
)

func FetchUsers(db *sql.DB, filter models.User) ([]models.User, error) {

	queryStr := `SELECT user_id, name, email, dob, city FROM user_data `

	values := []interface{}{}
	where := []string{}

	if filter.UserID != "" {
		where = append(where, ` user_id = $1 `)
		values = append(values, filter.UserID)
	}

	if filter.Name != "" {
		where = append(where, ` name LIKE CONCAT('%%', $2, '%%') `)
		values = append(values, filter.Name)
	}

	if filter.Email != "" {
		where = append(where, ` email = $3 `)
		values = append(values, filter.Email)
	}

	if filter.DOB != "" {
		where = append(where, ` dob = $4 `)
		values = append(values, filter.DOB)
	}

	if filter.City != "" {
		where = append(where, ` city LIKE CONCAT('%%', $5, '%%') `)
		values = append(values, filter.City)
	}

	if len(where) > 0 {
		queryStr += "WHERE " + strings.Join(where, " AND ")
	}

	rows, err := db.Query(queryStr, values...)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.DOB, &user.City)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}
