package beego

import (
	"fmt"

	"github.com/SinedVonotirah/gopo/persistence/entities"
	"github.com/astaxie/beego/orm"
)

type UserRepo struct {
	connection orm.Ormer
}

func NewUserRepo(connection *Connection) *UserRepo {
	return &UserRepo{connection.beego}
}

func (repo *UserRepo) GetUserById(id int64) (entities.UserEntity, error) {
	user := entities.UserEntity{Id: id}
	err := repo.connection.Read(&user)
	_, err = repo.connection.LoadRelated(&user, "UserGroup")
	_, err = repo.connection.LoadRelated(&user, "Orders")

	for _, order := range user.Orders {

		m2m := repo.connection.QueryM2M(order, "Product")
		fmt.Println(m2m)

		//_, err = repo.connection.LoadRelated(order, "Product")
	}

	return user, err
}
