package convert

import (
	"fmt"
	"testing"
)

type MsgSearchFilter struct {
	From          int32  `json:"from"`
	Size          int32  `json:"size"`
	SortDirection string `json:"sort_direction"`

	CreatedAtStart int64          `json:"created_at_start"`
	CreatedAtEnd   int64          `json:"created_at_end"`
	Content        string         `json:"content"`
	MsgIdFirst     int64          `json:"msg_id_first"`
	MsgIdStart     int64          `json:"msg_id_start"`
	MsgIdEnd       int64          `json:"msg_id_end"`
	StaffOpName    string         `json:"staff_op_name"`
	DialogIds      []int64        `json:"dialog_id"`
	AccountPair    []*AccountPair `json:"account_pairs"`
	RegexModelIds  []int64        `json:"regex_model_ids"`

	NeedGroupByDialogId bool `json:"need_group_by_dialog_id"`

	SearchAfter []interface{} `json:"search_after"`
}

type AccountPair struct {
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
}

func TestToMap(t *testing.T) {
	filter := &MsgSearchFilter{}
	toMap, err := ToMap(filter, "json")
	if err != nil {
		return
	}
	fmt.Println(toMap)
}
