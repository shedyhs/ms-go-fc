package gateway

import "github.com.br/shedyhs/ms-fullcycle/wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
