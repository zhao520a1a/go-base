package dbrouter

import (
	"context"
	"fmt"
	"sync"

	"github.com/coreos/etcd/client"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xetcd"
)

// Config ...
type Config struct {
	DBName   string
	DBType   string
	DBAddr   []string
	UserName string
	PassWord string
}

// ChangeIns ...
type ChangeIns struct {
	insNames []string
}

// ConfigChange 配置变更
type ConfigChange struct {
	dbInstanceChange map[string][]string
	dbGroups         []string
}

//Configer 配置接口抽象
type Configer interface {
	GetInstanceName(ctx context.Context, cluster, table string) string
	GetInstanceConfig(ctx context.Context, instance, group string) *Config
	GetAllGroups(ctx context.Context) []string
}

//etcdConfig etcd配置
type etcdConfig struct {
	etcdAddr []string
	parser   *Parser
	parserMu sync.RWMutex
}

//NewEtcdConfiger 新建配置实例
func NewEtcdConfiger(ctx context.Context, dbChangeChan chan ConfigChange) (Configer, error) {
	fun := "NewEtcdConfiger -->"
	etcdConfig := &etcdConfig{
		// TODO etcd address如何获取
		etcdAddr: []string{"http://infra0.etcd.ibanyu.com:20002", "http://infra1.etcd.ibanyu.com:20002", "http://infra2.etcd.ibanyu.com:20002", "http://infra3.etcd.ibanyu.com:20002", "http://infra4.etcd.ibanyu.com:20002", "http://old0.etcd.ibanyu.com:20002", "http://old1.etcd.ibanyu.com:20002", "http://old2.etcd.ibanyu.com:20002"},
	}
	err := etcdConfig.init(ctx, dbChangeChan)
	if err != nil {
		fmt.Printf("%s init etcd configer err: %s\n", fun, err.Error())
		return nil, err
	}
	return etcdConfig, nil
}

func (e *etcdConfig) init(ctx context.Context, dbChangeChan chan ConfigChange) error {
	fun := "EtcdConfig.init -->"
	etcdInstance, err := xetcd.NewEtcdInstance(e.etcdAddr)
	if err != nil {
		return err
	}

	initCh := make(chan error)
	var initOnce sync.Once
	etcdInstance.Watch(ctx, "/roc/db/route", func(response *client.Response) {
		parser, err := NewParser([]byte(response.Node.Value))
		if err != nil {
			fmt.Printf("%s init db parser err: %+v\n", fun, err.Error())
		} else {
			fmt.Printf("succeed to init new parser\n")
			if oldParser := e.GetParser(ctx); oldParser != nil {
				dbConfigChange := compareParsers(*oldParser, *parser)
				fmt.Printf("parser changes: %+v\n", dbConfigChange)
				e.SetParser(ctx, parser)
				dbChangeChan <- dbConfigChange
			} else {
				e.SetParser(ctx, parser)
			}
		}

		initOnce.Do(func() {
			initCh <- err
		})
	})
	// 做一次同步，等parser初始化完成
	err = <-initCh
	close(initCh)
	return err
}

func (e *etcdConfig) GetParser(ctx context.Context) *Parser {
	e.parserMu.RLock()
	defer e.parserMu.RUnlock()

	return e.parser
}

func (e *etcdConfig) SetParser(ctx context.Context, parser *Parser) {
	e.parserMu.Lock()
	defer e.parserMu.Unlock()

	e.parser = parser
}

//GetInstanceConfig 获取实例配置
func (e *etcdConfig) GetInstanceConfig(ctx context.Context, insName, group string) *Config {
	parser := e.GetParser(ctx)
	info := parser.getInstanceConfig(insName, group)
	return &Config{
		DBType:   info.DBType,
		DBAddr:   info.DBAddr,
		DBName:   info.DBName,
		UserName: info.UserName,
		PassWord: info.PassWord,
	}
}

//GetInstanceName 根据cluster以及table 获取实例名称
func (e *etcdConfig) GetInstanceName(ctx context.Context, cluster, table string) string {
	parser := e.GetParser(ctx)
	return parser.getInstanceName(cluster, table)
}

//GetAllGroups 获取所有的渠道（default or testing or ...）
func (e *etcdConfig) GetAllGroups(ctx context.Context) []string {
	var groups []string
	parser := e.GetParser(ctx)

	for group := range parser.dbIns {
		groups = append(groups, group)
	}
	return groups
}

// SendInsChange ...
func SendInsChange(insMap map[string]bool, ch chan ChangeIns) {
	var insNames []string
	for k := range insMap {
		insNames = append(insNames, k)
	}
	if len(insNames) > 0 {
		ch <- ChangeIns{insNames}
	}
}
