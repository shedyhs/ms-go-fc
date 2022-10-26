package gateway

import "github.com.br/shedyhs/ms-fullcycle/wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
