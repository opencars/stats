package store

type Store interface {
	Authorization() AuthRepository
}
