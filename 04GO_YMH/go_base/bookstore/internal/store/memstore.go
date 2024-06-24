package store

import (
	"bootstore/store"
	"bootstore/store/factory"
	"sync"
)

/*
*
某个对象，存储的对象
*/
func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*store.Book),
	})
}

type MemStore struct {
	sync.RWMutex // 组合
	books        map[string]*store.Book
}

func (m *MemStore) Create(book *store.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemStore) Update(book *store.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemStore) Get(s string) (store.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemStore) GetAll() ([]store.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemStore) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}
