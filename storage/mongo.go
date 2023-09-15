package storage

import "v1/types"

type MogoStorage struct{}

func (s *MogoStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
