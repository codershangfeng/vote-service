package persistence

import (
	"fmt"
	"sort"
)

type (
	Repository interface {
		GetVoteEntity(id int64) *VoteEntity
		GetVoteEntities() []VoteEntity
		SaveVoteEntity(v VoteEntity)
		DeleteVoteEntity(id int64) error
	}

	RepositoryImpl struct {
		database map[int64]VoteEntity
	}

	VoteEntity struct {
		ID      int64
		Options []string
		Topic   string
	}
)

// Should NOT refering to this variable in each public method
var db map[int64]VoteEntity

func init() {
	db = make(map[int64]VoteEntity)
}

func NewRepository() Repository {
	return RepositoryImpl{database: db}
}

// FIXME: Refactor the interface with 2 return values: (VaultEntity, ok)
func (r RepositoryImpl) GetVoteEntity(id int64) *VoteEntity {
	if v, ok := r.database[id]; ok {
		return &v
	}
	return nil
}

func (r RepositoryImpl) GetVoteEntities() []VoteEntity {
	if len(r.database) == 0 {
		return nil
	}

	votes := make([]VoteEntity, 0, len(r.database))
	for _, v := range r.database {
		votes = append(votes, v)
	}
	sort.Slice(votes[:], func(i, j int) bool {
		return votes[i].ID < votes[j].ID
	})
	return votes
}

func (r RepositoryImpl) SaveVoteEntity(v VoteEntity) {
	r.database[v.ID] = v
}

func (r RepositoryImpl) DeleteVoteEntity(id int64) error {
	if _, ok := r.database[id]; ok {
		delete(r.database, id)
		return nil
	}
	return fmt.Errorf("vote with ID[%d] does not exist", id)
}
