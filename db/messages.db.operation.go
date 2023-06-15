package db

import "log"

func CreateMessageToDatabase(user_id, content string) (bool, string, error) {
	var id string
	sqlQuery := `INSERT INTO message(user_id, content, created_at) VALUES ($1, $2, NOW()) RETURNING id;`
	if err := DB.QueryRow(sqlQuery, user_id, content).Scan(&id); err != nil {
		log.Println("CreateMessageToDatabase QueryRow ", err.Error())
		return false, id, err
	}
	return true, id, nil
}
