package dbrouter

import (
	"encoding/json"
	"fmt"
)

const (
	defaultGroup = ""
	testGroup    = "test"
)

// Parser ETCD配置抽象
type Parser struct {
	//逻辑集群配置
	dbCls *dbCluster
	//实例连接配置
	dbIns map[string]map[string]*dbInsInfo
}

// DB ETCD配置映射
type routeConfig struct {
	Cluster   map[string][]*dbLookupCfg `json:"cluster"`
	Instances map[string]*dbInsCfg      `json:"instances"`
}

type dbLookupCfg struct {
	Instance string `json:"instance"`
	Match    string `json:"match"`
	Express  string `json:"express"`
}

type dbInsCfg struct {
	Dbtype string          `json:"dbtype"`
	Dbname string          `json:"dbname"`
	Dbcfg  json.RawMessage `json:"dbcfg"`
	Level  int32           `json:"level"`
	Ins    []itemDbInsCfg  `json:"ins"`
}

// 配置用于压测、泳道场景
type itemDbInsCfg struct {
	Dbcfg json.RawMessage `json:"dbcfg"`
	Group string          `json:"group"`
}

type dbInsInfo struct {
	Instance string
	DBType   string
	DBName   string
	Level    int32
	DBAddr   []string `json:"addrs"`
	UserName string   `json:"user"`
	PassWord string   `json:"passwd"`
}

func (m *dbLookupCfg) String() string {
	return fmt.Sprintf("ins:%s exp:%s match:%s", m.Instance, m.Express, m.Match)
}

func (m *Parser) String() string {
	return fmt.Sprintf("%v", m.dbCls.clusters)
}

//getInstanceName 根据逻辑集群名称、表名获取实例名称
func (m *Parser) getInstanceName(cluster, table string) string {
	insName := m.dbCls.getInstance(cluster, table)
	return insName
}

//getInstanceConfig 根据实例名称、group渠道获取实例配置
func (m *Parser) getInstanceConfig(insName, group string) *dbInsInfo {
	return m.getConfig(insName, group)
}

func (m *Parser) getConfig(insName, group string) *dbInsInfo {
	if infoMap, ok := m.dbIns[group]; ok {
		if info, ok := infoMap[insName]; ok {
			return info
		}
	} else if group != testGroup {
		if info, ok := infoMap[defaultGroup]; ok {
			return info
		}
	}
	return &dbInsInfo{}
}

// 检查用户输入的合法性
// 1. 只能是字母或者下划线
// 2. 首字母不能为数字，或者下划线
func checkVarname(varname string) error {
	if len(varname) == 0 {
		return fmt.Errorf("is empty")
	}

	f := varname[0]
	if !((f >= 'a' && f <= 'z') || (f >= 'A' && f <= 'Z')) {
		return fmt.Errorf("first char is not alpha")
	}

	for _, c := range varname {

		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			continue
		} else if c >= '0' && c <= '9' {
			continue
		} else if c == '_' {
			continue
		} else {
			return fmt.Errorf("is contain not [a-z] or [A-Z] or [0-9] or _")
		}
	}

	return nil
}

//NewParser ...
func NewParser(jscfg []byte) (*Parser, error) {
	fun := "NewParser -->"

	r := &Parser{
		dbCls: &dbCluster{
			clusters: make(map[string]*clsEntry),
		},
		dbIns: make(map[string]map[string]*dbInsInfo),
	}

	var cfg routeConfig
	err := json.Unmarshal(jscfg, &cfg)
	if err != nil {
		return nil, fmt.Errorf("dbrouter config unmarshal error: %s", err.Error())
	}

	cls := cfg.Cluster
	for c, ins := range cls {
		if er := checkVarname(c); er != nil {
			fmt.Printf("%s cluster config name err:%s\n", fun, err)
			continue
		}

		if len(ins) == 0 {
			fmt.Printf("%s empty instance in cluster:%s\n", fun, c)
			continue
		}

		for _, v := range ins {
			if len(v.Express) == 0 {
				fmt.Printf("%s empty express in cluster:%s instance:%s\n", fun, c, v.Instance)
				continue
			}

			if er := checkVarname(v.Match); er != nil {
				fmt.Printf("%s match in cluster:%s instance:%s err:%s\n", fun, c, v.Instance, err)
				continue
			}

			if er := checkVarname(v.Instance); er != nil {
				fmt.Printf("%s instance name in cluster:%s instance:%s err:%s\n", fun, c, v.Instance, err)
				continue
			}

			if err := r.dbCls.addInstance(c, v); err != nil {
				return nil, fmt.Errorf("load instance lookup rule err:%s", err.Error())
			}
		}
	}

	inss := cfg.Instances
	for ins, db := range inss {
		if er := checkVarname(ins); er != nil {
			fmt.Printf("%s instances name config err:%s\n", fun, er.Error())
			continue
		}

		dbtype := db.Dbtype
		dbname := db.Dbname
		cfg := db.Dbcfg
		level := db.Level
		// 外层配置为默认的group
		group := defaultGroup
		info, err := parseDbIns(dbtype, dbname, ins, cfg, level)
		if err != nil {
			fmt.Printf("%s parse default dbInsInfo error, %s\n", fun, err.Error())
		} else {
			if _, ok := r.dbIns[group]; ok {
				r.dbIns[group][ins] = info
			} else {
				r.dbIns[group] = make(map[string]*dbInsInfo)
				r.dbIns[group][ins] = info
			}
		}

		for _, itemIns := range db.Ins {
			cfg = itemIns.Dbcfg
			group = itemIns.Group
			info, err := parseDbIns(dbtype, dbname, ins, cfg, level)
			if err != nil {
				fmt.Printf("%s parse %s dbInsInfo error, %s\n", fun, group, err.Error())
			} else {
				if _, ok := r.dbIns[group]; ok {
					r.dbIns[group][ins] = info
				} else {
					r.dbIns[group] = make(map[string]*dbInsInfo)
					r.dbIns[group][ins] = info
				}
			}
		}

		dbInsLen := len(r.dbIns[defaultGroup])
		for group, insMap := range r.dbIns {
			if len(insMap) > dbInsLen {
				return nil, fmt.Errorf("db instance %s is lack of default group", ins)
			} else if len(insMap) < dbInsLen {
				return nil, fmt.Errorf("db instance %s is lack of group %s", ins, group)
			}
		}
	}

	return r, nil
}

func parseDbIns(dbtype, dbname, ins string, dbcfg json.RawMessage, level int32) (*dbInsInfo, error) {
	if er := checkVarname(dbtype); er != nil {
		return nil, fmt.Errorf("dbtype instance:%s err:%s", ins, er.Error())
	}

	if er := checkVarname(dbname); er != nil {
		return nil, fmt.Errorf("dbname instance:%s err:%s", ins, er.Error())
	}

	if len(dbcfg) == 0 {
		return nil, fmt.Errorf("empty dbcfg instance:%s", ins)
	}

	var info = new(dbInsInfo)
	err := json.Unmarshal(dbcfg, info)
	if err != nil {
		return nil, fmt.Errorf("unmarshal err, cfg:%s", string(dbcfg))
	}
	info.DBType = dbtype
	info.DBName = dbname
	info.Instance = ins
	info.Level = level

	return info, nil
}

func compareParsers(originParser Parser, newParser Parser) ConfigChange {
	// 原来实例中修改的、删除的，要通知数据库连接池关闭掉实例的数据库连接
	var dbInsChangeMap = make(map[string][]string)
	var groups []string
	for group, originIns := range originParser.dbIns {
		var dbInsChanges []string
		if newIns, ok := newParser.dbIns[group]; ok {
			dbInsChanges = compareDbInstances(originIns, newIns)
			if len(dbInsChanges) == 0 {

			}
		} else {
			dbInsChanges = compareDbInstances(originIns, make(map[string]*dbInsInfo))
		}
		if len(dbInsChanges) > 0 {
			dbInsChangeMap[group] = dbInsChanges
		}
	}

	for group := range newParser.dbIns {
		groups = append(groups, group)
	}

	return ConfigChange{
		dbInstanceChange: dbInsChangeMap,
		dbGroups:         groups,
	}
}

func compareDbInstances(originDbInstances map[string]*dbInsInfo, newDbInstances map[string]*dbInsInfo) []string {
	var dbInstanceChanges []string
	for insName, originDbInsInfo := range originDbInstances {
		if newDbInsInfo, ok := newDbInstances[insName]; ok {
			if !compareDbInfo(originDbInsInfo, newDbInsInfo) {
				dbInstanceChanges = append(dbInstanceChanges, insName)
			}
		} else {
			dbInstanceChanges = append(dbInstanceChanges, insName)
		}
	}
	return dbInstanceChanges
}

func compareDbInfo(dbInsInfo1 *dbInsInfo, dbInsInfo2 *dbInsInfo) bool {
	return dbInsInfo1.DBName == dbInsInfo2.DBName && dbInsInfo1.UserName == dbInsInfo2.UserName &&
		dbInsInfo1.PassWord == dbInsInfo2.PassWord && compareStringList(dbInsInfo1.DBAddr, dbInsInfo2.DBAddr)
}

func compareStringList(stringList1 []string, stringList2 []string) bool {
	if len(stringList1) != len(stringList2) {
		return false
	}

	for index := range stringList1 {
		if stringList1[index] != stringList2[index] {
			return false
		}
	}

	return true
}
