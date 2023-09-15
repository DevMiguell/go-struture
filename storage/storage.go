package storage

import "v1/types"

type Storage interface {
	Get(int) *types.User
}
