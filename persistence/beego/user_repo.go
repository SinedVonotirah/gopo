package beego

import (
	"github.com/SinedVonotirah/gopo/persistence/beego/entities"
	"github.com/astaxie/beego/orm"
)

type UserRepo struct {
	connection orm.Ormer
}

func NewUserRepo(connection *Connection) *UserRepo {
	return &UserRepo{connection.beego}
}

func (repo *UserRepo) CreateUser(entity *entities.UserEntity) (int64, error) {
	repo.connection.Begin()
	id, err := repo.connection.Insert(entity)
	if err != nil {
		repo.connection.Rollback()
	} else {
		repo.connection.Commit()
	}
	return id, err
}

func (repo *UserRepo) DeleteUser(userId int64) (int64, error) {
	return repo.connection.Delete(&entities.UserEntity{Id: userId})
}

func (repo *UserRepo) UpdateUser(entity *entities.UserEntity) (int64, error) {
	return repo.connection.Update(entity)
}

func (repo *UserRepo) GetUserById(id int64) (entities.UserEntity, error) {
	user := entities.UserEntity{}
	err := repo.connection.QueryTable("user").Filter("Id", id).RelatedSel().Limit(10000).One(&user)
	return user, err
}

func (repo *UserRepo) GetUserByIdWithOrders(id int64) (entities.UserEntity, error) {
	user := entities.UserEntity{}
	err := repo.connection.QueryTable("user").Filter("Id", id).RelatedSel().Limit(10000).One(&user)
	_, err = repo.connection.LoadRelated(&user, "Orders")

	/*	for _, order := range user.Orders {
		_, err = repo.connection.LoadRelated(order, "Products")

	}*/
	return user, err
}

func (repo *UserRepo) GetUsersByGroupName(groupName string) ([]entities.UserEntity, error) {
	var entities []entities.UserEntity
	qs := repo.connection.QueryTable("user")
	_, err := qs.Filter("UserGroup__Name", groupName).All(&entities)
	return entities, err
}
