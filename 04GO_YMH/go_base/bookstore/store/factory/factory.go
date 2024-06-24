package factory

import (
	"bootstore/store"
	"fmt"
	"sync"
)

/**
对工厂可以“生产”的、满足 Store 接口的实例类型进行管理。
factory 包还提供了 Register 函数，让各个实现 Store 接口的类型可以把自己“注册”到工厂中来
*/
// 初始化常量
var (
	providerMu sync.RWMutex
	providers  = make(map[string]store.Store)
)

// Register 注册这种name类型的存储器
func Register(name string, p store.Store) {
	providerMu.Lock()
	defer providerMu.Unlock()
	if p == nil {
		panic("store: Register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("store Register called twice for provider " + name)
	}
	providers[name] = p
}

func New(providerName string) (store store.Store, err error) {
	providerMu.Lock()
	// map 获取值
	p, ok := providers[providerName]
	providerMu.Unlock()
	if !ok {
		return nil, fmt.Errorf("store: unknow provider %s", providerName)
	}
	return p, nil
}
