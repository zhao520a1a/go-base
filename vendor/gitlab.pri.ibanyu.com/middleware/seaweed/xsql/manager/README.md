## Manager

Manager用于管理所有的DB连接池，每个连接池实例是标准库中的sql.DB 
Mangager解析ETCD中的DB配置，并且监听配置的变化，管理DB连接池的建立与关闭

### Example
以下为简单示例，项目中无需手写，参照[生成sql](http://confluence.pri.ibanyu.com/pages/viewpage.action?pageId=20381903)进行代码生成

```go
import (
	"context"

	_ "github.com/go-sql-driver/mysql"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xsql/manager"
)

var err error
var db manager.XDB
var err error
// ACCOUNT 路由名 ; user 数据库表名
// 每次通过manager.GetDB获取db进行使用，不要把db存为全局变量
db = manager.GetDB(context.Todo(), "ACCOUT", "user")
db.QueryContext(ctx, query, args...)
// 开启事务
var tx *manager.Tx
tx, err = db.Begin(ctx)
tx.ExecContext(ctx, query, args...)
tx.Commit(ctx)
```
