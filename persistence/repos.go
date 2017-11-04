package persistence

import "github.com/SinedVonotirah/gopo/persistence/entities"

type UserRepo interface {
	Insert(*entities.UserEntity) error
	GetUserById(id int64) (entities.UserEntity, error)
}

type UserGroupRepo interface {
}

type OrderRepo interface {
}

type ProductRepo interface {
}

type ProductDetailsRepo interface {
}
