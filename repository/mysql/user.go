package mysql

import (
	"database/sql"
	"fmt"
	"questionGame/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt []uint8

	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("can't scan row: %w", err)
	}
	return false, nil

}

func (d *MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number) values(?, ?)`, u.Name, u.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("cannot execute command: %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}