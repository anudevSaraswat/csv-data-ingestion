package schema

import (
	"csv-app/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func FetchUsers(db *sql.DB, filter models.User) ([]models.User, error) {

	queryStr := `SELECT user_id, first_name, last_name, sex, email, phone, dob, job_title FROM user_data `

	values := []interface{}{}
	where := []string{}

	if filter.UserID != "" {
		where = append(where, ` user_id = $1 `)
		values = append(values, filter.UserID)
	}

	if filter.FirstName != "" {
		where = append(where, ` first_name LIKE `+fmt.Sprintf("$%d", len(where)+1))
		values = append(values, filter.FirstName)
	}

	if filter.LastName != "" {
		where = append(where, ` last_name LIKE `+fmt.Sprintf("$%d", len(where)+1))
		values = append(values, filter.LastName)
	}

	if filter.Sex != "" {
		where = append(where, ` sex = `+fmt.Sprintf("$%d", len(where)+1))
		values = append(values, filter.Sex)
	}

	if filter.Email != "" {
		where = append(where, ` email = `+fmt.Sprintf("$%d", len(where)+1))
		values = append(values, filter.Email)
	}
	if filter.Phone != "" {
		where = append(where, ` phone = `+fmt.Sprintf("$%d", len(where)+1))
		values = append(values, filter.Phone)
	}

	if filter.DOB != "" {
		where = append(where, ` dob = `+fmt.Sprintf("$%d", len(where)+1))
		values = append(values, filter.DOB)
	}

	if filter.JobTitle != "" {
		where = append(where, ` job_title LIKE `+fmt.Sprintf("$%d", len(where)+1))
		values = append(values, filter.JobTitle)
	}

	if len(where) > 0 {
		queryStr += "WHERE " + strings.Join(where, " AND ")
	}

	rows, err := db.Query(queryStr, values...)
	if err != nil {
		log.Println("(FetchUsers) err in db.Query:", err)
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Sex,
			&user.Email, &user.Phone, &user.DOB, &user.JobTitle)
		if err != nil {
			log.Println("(FetchUsers) err in rows.Scan:", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}
