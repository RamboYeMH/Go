package store

type Book struct {
	Id      string   // 图书ISBN ID
	Name    string   // 图书名称
	Authors []string // 图书作者
	Press   string   // 出版社
}

func NewBook(id string, name string, authors []string, press string) *Book {
	return &Book{Id: id, Name: name, Authors: authors, Press: press}
}

// Store book 存取的接口类型 Store
type Store interface {
	Create(*Book) error       // 创建一个新图书条目
	Update(*Book) error       // 更新某图书条目
	Get(string) (Book, error) // 获取某图书信息
	GetAll() ([]Book, error)  // 获取所有图书信息
	Delete(string) error      //删除某图书条目
}
