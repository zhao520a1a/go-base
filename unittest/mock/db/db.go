package db

import (
	"context"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xsql/manager"
)

const BPMDBCluster = "BPM"
const WorkflowTable = "bpm_workflow"

type DBManager interface {
	Begin(ctx context.Context) (*manager.Tx, error)
	GetDB(ctx context.Context) (*manager.DB, error)
}
