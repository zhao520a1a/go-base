package db

import (
	"context"
	"database/sql"
)

const BPMDBCluster = "BPM"
const WorkflowTable = "bpm_workflow"

type DBManager interface {
	Begin(ctx context.Context) (*sql.Tx, error)
	GetDB(ctx context.Context) (*sql.DB, error)
}
