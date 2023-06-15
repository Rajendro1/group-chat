package db

import "log"

func CreateUsersToDatabase(username string) (bool, string, error) {
	var id string
	sqlQuery := `INSERT INTO users(username) VALUES ($1) RETURNING id;`
	if err := DB.QueryRow(sqlQuery, username).Scan(&id); err != nil {
		log.Println("CreateUsersToDatabase QueryRow ", err.Error())
		return false, id, err
	}
	return true, id, nil
}
