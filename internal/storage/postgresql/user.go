package postgresql

import "internal/model"

func (s *PostgresqlDB) CreateUser(user *model.User) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresqlDB) ReadCredential(name string) (hashedPassword string, err error) {
	//TODO implement me
	panic("implement me")
}
