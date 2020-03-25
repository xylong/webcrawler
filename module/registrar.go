package module

import (
	"fmt"
	"sync"
	"webcrawler/errors"
)

// Registrar 组件注册器的接口
type Registrar interface {
	// Register 注册组件实例
	Register(module Module) (bool, error)
	// Unregister 注销组件实例
	Unregister(mid MID) (bool, error)
	// Get 获取一个指定类型的组件的实例
	// 本函数应该基于负载均衡策略返回实例
	Get(moduleType Type) (Module, error)
	// GetAllByType 获取指定类型的所有组件实例
	GetAllType(moduleType Type) (map[MID]Module, error)
	// GetAll 获取所有组件实例
	GetAll() map[MID]Module
	// Clear 清除所有的组件注册记录
	Clear()
}

// NewRegistrar 创建一个组件注册器的实例
func NewRegistrar() Registrar {
	return myRegistrar{
		moduleTypeMap: map[Type]map[MID]Module{},
	}
}

// myRegistrar 组件注册器的实现类型
type myRegistrar struct {
	// moduleTypeMap 组件类型与对应组件实例的映射
	moduleTypeMap map[Type]map[MID]Module
	// rwlock 组件注册专用读写锁
	rwlock sync.RWMutex
}

func (r *myRegistrar) Register(module Module) (bool, error) {
	if module == nil {
		return false, errors.NewIllegalParameterError("nil module instance")
	}
	mid := module.ID()
	parts, err := SplitMID(mid)
	if err != nil {
		return false, err
	}
	moduleType := legalLetterTypeMap[parts[0]]
	if !CheckType(moduleType, module) {
		errMsg := fmt.Sprintf("incorrect module type: %s", moduleType)
		return false, errors.NewIllegalParameterError(errMsg)
	}
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	modules := r.moduleTypeMap[moduleType]
	if modules == nil {
		modules = map[MID]Module{}
	}
	if _, ok := modules[mid]; ok {
		return false, nil
	}
	modules[mid] = module
	r.moduleTypeMap[moduleType] = modules
	return true, nil
}
