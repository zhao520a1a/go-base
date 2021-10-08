package convert

import (
	"reflect"
	"strings"
)

func ConnectionStringBuilder(connectionstring string) ConnectionString {
	splittedcs := strings.Split(connectionstring, ";")

	csstruct := ConnectionString{}

	for i := 0; i < len(splittedcs); i++ {

		actualitem := splittedcs[i]
		splitteditem := strings.Split(actualitem, "=")

		fieldname := strings.ToUpper(strings.Replace(splitteditem[0], " ", "", -1))
		value := splitteditem[1]

		if fieldname == "DATASOURCE" || fieldname == "SERVER" {
			splittedport := strings.Split(value, ",")

			val := reflect.ValueOf(&csstruct)
			(val.Elem()).FieldByName(fieldname).SetString(splittedport[0])

			if len(splittedport) > 1 {
				(val.Elem()).FieldByName("PORT").SetString(splittedport[1])
			}

		} else {

			val := reflect.ValueOf(&csstruct)
			(val.Elem()).FieldByName(fieldname).SetString(value)
		}

		if csstruct.DATASOURCE != "" {
			csstruct.HOST = csstruct.DATASOURCE
		}

		if csstruct.SERVER != "" {
			csstruct.HOST = csstruct.SERVER
		}
	}

	return csstruct
}

type ConnectionString struct {
	HOST                   string `json:"Host,omitempty"`
	DATASOURCE             string `json:"DataSource,omitempty"`
	SERVER                 string `json:"Server,omitempty"`
	INITIALCATALOG         string `json:"InitialCatalog,omitempty"`
	USERID                 string `json:"UserID,omitempty"`
	ASYNCHRONOUSPROCESSING string `json:"AsynchronousProcessing,omitempty"`
	PASSWORD               string `json:"Password,omitempty"`
	DATABASE               string `json:"Database,omitempty"`
	PORT                   string `json:"Port,omitempty"`
}
