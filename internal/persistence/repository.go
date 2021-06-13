package persistence

// Should NOT refering to this variable in each public method
var db map[int64]VoteEntity

func init() {
	db = make(map[int64]VoteEntity)
}

type Repository struct {
	database map[int64]VoteEntity
}

func NewRepository() *Repository {
	return &Repository{database: db}
}

func (r Repository) GetVote(id int64) *VoteEntity {
	if v, ok := r.database[id]; ok {
		return &v
	}
	return nil
}

func (r Repository) GetVotes() []VoteEntity {
	if len(r.database) == 0 {
		return nil
	}

	votes := make([]VoteEntity, 0, len(r.database))
	for _, v := range r.database {
		votes = append(votes, v)
	}
	return votes
}

func (r Repository) SaveVote(v VoteEntity) {
	r.database[v.ID] = v
}

type VoteEntity struct {
	ID      int64
	Options []string
	Topic   string
}
