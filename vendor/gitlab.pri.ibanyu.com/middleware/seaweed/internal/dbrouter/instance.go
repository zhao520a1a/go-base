package dbrouter

import (
	"context"
	"fmt"
	"sync"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xcontext"
)

//Instancer ...
type Instancer interface {
	GetType() string
	Close() error
	Reload() error
}

//FactoryFunc ...
type factoryFunc func(ctx context.Context, key, group string) (in Instancer, err error)

//InstanceManager 实例管理器
type InstanceManager struct {
	instanceMu sync.RWMutex
	groupMu    sync.RWMutex
	instances  map[string]Instancer
	factory    factoryFunc
	groups     []string
}

//NewInstanceManager ...
func NewInstanceManager(factory factoryFunc, dbChangeChan chan ConfigChange, closeChangeChan, reloadChangeChan chan ChangeIns, groups []string) *InstanceManager {
	instanceManager := &InstanceManager{
		instances: make(map[string]Instancer),
		factory:   factory,
		groups:    groups,
	}
	go instanceManager.dbInsChangeHandler(context.Background(), dbChangeChan)
	go instanceManager.closeChangeHandler(context.Background(), closeChangeChan, groups)
	go instanceManager.reloadChangeHandler(context.Background(), reloadChangeChan, groups)

	return instanceManager
}

//Get 获取实例
func (m *InstanceManager) Get(ctx context.Context, instance string) Instancer {
	fun := "InstanceManager.Get -->"
	var err error
	var in Instancer
	group := xcontext.GetControlRouteGroupWithDefault(ctx, defaultGroup)
	key := m.buildKey(instance, group)
	in, ok := m.getInstance(ctx, key)
	if ok == false {
		in, err = m.buildInstance(ctx, instance, group)
		if err != nil {
			fmt.Printf("%s NewInstance err, instance: %s, err: %s\n", fun, instance, err.Error())
			return nil
		}
	}

	return in
}

//Close ...
func (m *InstanceManager) Close() {
	m.instanceMu.Lock()
	defer m.instanceMu.Unlock()

	for key, in := range m.instances {
		go in.Close()
		delete(m.instances, key)
	}
}

func (m *InstanceManager) buildInstance(ctx context.Context, insName, group string) (Instancer, error) {
	if group != defaultGroup {
		if !m.isInGroup(group) {
			if group == testGroup {
				return nil, fmt.Errorf("db config don't have group: %s", group)
			}
			group = defaultGroup
		}
	}
	key := m.buildKey(insName, group)
	m.instanceMu.Lock()
	defer m.instanceMu.Unlock()

	if in, ok := m.instances[key]; ok {
		return in, nil
	}

	in, err := m.factory(ctx, insName, group)
	if err != nil {
		return nil, err
	}

	m.instances[key] = in
	return in, nil
}

func (m *InstanceManager) closeDbInstance(ctx context.Context, insName, group string) {
	fun := "InstanceManager.closeDbInstance -->"

	key := m.buildKey(insName, group)
	m.instanceMu.Lock()
	defer m.instanceMu.Unlock()

	if in, ok := m.instances[key]; ok {
		delete(m.instances, key)
		go func() {
			if err := in.Close(); err == nil {
				fmt.Printf("%s succeed to close db instance: %s group: %s\n", fun, insName, group)
			} else {
				fmt.Printf("%s close db instance: %s group: %s error: %s\n", fun, insName, group, err.Error())
			}
		}()
	}
}

func (m *InstanceManager) reloadDbInstance(ctx context.Context, insName, group string) {
	fun := "InstanceManager.reloadDbInstance -->"
	key := m.buildKey(insName, group)
	m.instanceMu.Lock()
	defer m.instanceMu.Unlock()
	if in, ok := m.instances[key]; ok {
		go func() {
			if err := in.Reload(); err == nil {
				fmt.Printf("%s succeed to reload db instance: %s group: %s\n", fun, insName, group)
			} else {
				fmt.Printf("%s reload db instance: %s group: %s error: %s\n", fun, insName, group, err.Error())
			}
		}()
	}
}

func (m *InstanceManager) buildKey(insName, group string) string {
	return fmt.Sprintf("%s-%s", group, insName)
}

func (m *InstanceManager) getInstance(ctx context.Context, key string) (Instancer, bool) {
	m.instanceMu.RLock()
	defer m.instanceMu.RUnlock()

	in, ok := m.instances[key]
	return in, ok
}

func (m *InstanceManager) isInGroup(group string) bool {
	m.groupMu.RLock()
	defer m.groupMu.RUnlock()

	for _, configGroup := range m.groups {
		if group == configGroup {
			return true
		}
	}
	return false
}

func (m *InstanceManager) dbInsChangeHandler(ctx context.Context, dbChangeChan chan ConfigChange) {
	for dbInsChange := range dbChangeChan {
		m.handleGroupChange(dbInsChange.dbGroups)
		m.handleDbInsChange(ctx, dbInsChange.dbInstanceChange)
	}
}

func (m *InstanceManager) handleGroupChange(groups []string) {
	m.groupMu.Lock()
	defer m.groupMu.Unlock()

	m.groups = groups
}

func (m *InstanceManager) handleDbInsChange(ctx context.Context, dbInstanceChange map[string][]string) {
	for group, insNames := range dbInstanceChange {
		for _, insName := range insNames {
			m.closeDbInstance(ctx, insName, group)
		}
	}
}

func (m *InstanceManager) closeChangeHandler(ctx context.Context, closeInsChan chan ChangeIns, groups []string) {
	for changeIns := range closeInsChan {
		for _, g := range groups {
			for _, n := range changeIns.insNames {
				m.closeDbInstance(ctx, n, g)
			}
		}
	}
}

func (m *InstanceManager) reloadChangeHandler(ctx context.Context, reloadInsChan chan ChangeIns, groups []string) {
	for changeIns := range reloadInsChan {
		for _, g := range groups {
			for _, n := range changeIns.insNames {
				m.reloadDbInstance(ctx, n, g)
			}
		}
	}
}
