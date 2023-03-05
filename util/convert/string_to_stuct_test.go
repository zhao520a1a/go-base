package convert

import (
	"reflect"
	"testing"
)

func TestConnectionStringBuilder(t *testing.T) {
	tests := []struct {
		name             string
		connectionstring string
		want             ConnectionString
	}{
		{
			"good-case-1",
			"Data Source=mysqlserver.local,37001;Initial Catalog=mydatabase;User ID=usertest;Asynchronous Processing=True",
			ConnectionString{
				HOST:                   "mysqlserver.local",
				DATASOURCE:             "mysqlserver.local",
				SERVER:                 "",
				INITIALCATALOG:         "mydatabase",
				USERID:                 "usertest",
				ASYNCHRONOUSPROCESSING: "True",
				PASSWORD:               "",
				DATABASE:               "",
				PORT:                   "37001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectionStringBuilder(tt.connectionstring); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectionStringBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
