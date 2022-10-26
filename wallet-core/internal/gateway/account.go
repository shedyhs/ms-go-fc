package gateway

import "github.com.br/shedyhs/ms-fullcycle/wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
