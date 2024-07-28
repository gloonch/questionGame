package mysql

import (
	"database/sql"
	"fmt"
	"questionGame/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {

	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("can't scan row: %w", err)
	}
	return false, nil

}

func (d *MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number, password) values(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("cannot execute command: %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}

func (d *MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {

	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		return entity.User{}, false, err
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, fmt.Errorf("can't scan row: %w", err)
	}
	return user, true, nil

}

func (d *MySQLDB) GetUserByID(userID uint) (entity.User, error) {

	row := d.db.QueryRow(`select * from users where id = ?`, userID)

	user, err := scanUser(row)
	if err != nil {
		return entity.User{}, err
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("record not found")
		}

		return entity.User{}, fmt.Errorf("can't scan row: %w", err)
	}
	return user, nil

}

func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

	return user, err

}
