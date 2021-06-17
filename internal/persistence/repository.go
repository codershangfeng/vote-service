package persistence

import (
	"fmt"
	"sort"
	"sync"
)

type (
	Repository interface {
		GetVoteEntity(id int64) *VoteEntity
		GetVoteEntities() []VoteEntity
		SaveVoteEntity(v VoteEntity) VoteEntity
		DeleteVoteEntity(id int64) error
		UpdateVoteEntity(v VoteEntity) error
	}

	RepositoryImpl struct {
		database    map[int64]VoteEntity
		idGenerator IDGenerator
	}

	VoteEntity struct {
		ID      int64
		Options []string
		Topic   string
	}

	IDGenerator interface {
		IncAndGet() int64
	}

	IDGeneratorImpl struct {
		counter int64
		rwMutex sync.RWMutex
	}
)

// Should NOT refering to this variable in each public method
var db map[int64]VoteEntity
var ig IDGenerator

func init() {
	db = make(map[int64]VoteEntity)
	ig = &IDGeneratorImpl{}
}

func NewRepository() Repository {
	return RepositoryImpl{database: db, idGenerator: ig}
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

func (r RepositoryImpl) SaveVoteEntity(v VoteEntity) VoteEntity {
	v.ID = r.idGenerator.IncAndGet()
	r.database[v.ID] = v
	return v
}

func (r RepositoryImpl) DeleteVoteEntity(id int64) error {
	if _, ok := r.database[id]; ok {
		delete(r.database, id)
		return nil
	}
	return fmt.Errorf("vote with ID[%d] does not exist", id)
}

func (r RepositoryImpl) UpdateVoteEntity(v VoteEntity) error {
	r.database[v.ID] = v
	return nil
}

func (igi *IDGeneratorImpl) IncAndGet() int64 {
	igi.rwMutex.Lock()
	igi.counter++
	igi.rwMutex.Unlock()
	return igi.counter
}
