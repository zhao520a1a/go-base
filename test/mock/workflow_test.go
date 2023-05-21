package mock

import (
	"context"
	"database/sql"
	"testing"

	"github.com/zhao520a1a/go-base/test/mock/bpm"
	"github.com/zhao520a1a/go-base/test/mock/db"
	"github.com/zhao520a1a/go-base/test/mock/test"

	"github.com/stretchr/testify/suite"
)

type WorkflowControllerTestSuite struct {
	suite.Suite
	testingDBName string
	testingDB     *sql.DB
}

func (m *WorkflowControllerTestSuite) SetupTest() {
	m.testingDBName = "testing_bpm"
	m.testingDB = test.SetupTestingMySQL(m.testingDBName)

	_ = &ConfigService{
		GetNotifierEmailFlagFn: func(ctx context.Context) (bool, error) {
			return false, nil
		},
		GetNotNotifyConfigForEmailFn: func(ctx context.Context) ([]string, error) {
			return []string{}, nil
		},
	}

	_ = NewDBManagerMock(db.BPMDBCluster, db.WorkflowTable, m.testingDB)

	_ = &Notifier{
		NotifyFn: func(ctx context.Context, notif *bpm.Notification) (err error) {
			return
		},
		NotifyInvoked: false,
	}

}

func (m *WorkflowControllerTestSuite) TearDownTest() {
	m.testingDB.Close()
}

func TestWorkflowControllerTestSuite(t *testing.T) {
	suite.Run(t, new(WorkflowControllerTestSuite))
}

func (m *WorkflowControllerTestSuite) TestListWorkflow() {
	//ast := assert.New(m.T())
	//ctx := context.TODO()

	//res, err := m.controller.ListWorkflow(ctx, &bpm.ListWorkflowReq{
	//	Id:       100000,
	//	Offset:   0,
	//	Limit:    2,
	//	Operator: "zhenghe3119",
	//})

	//ast.NoError(err)
	//ast.Nil(res.Errinfo)
	//xlog.Info(ctx, res)
}
