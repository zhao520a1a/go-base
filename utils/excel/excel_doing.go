package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xlog"
	"gitlab.pri.ibanyu.com/quality/dry.git/errors"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// Tips:
// - Empty row will be skipped.
// - Column larger than len(TitleRow) will be skipped.

const (
	DefaultSheetName = "Sheet1"
	DefaultTagName   = "xlsx"

	cTimeFormat = "2006-01-02 15:04:05"
)

var (
	UserDefinedTagName string
	//ErrTargetNotSettable means the second param of Bind is not settable
	ErrTargetNotSettable = errors.New("[excel-scanner]: target is not settable! a pointer is required")
	//ErrNilRows means the first param can't be a nil
	ErrNilRows     = errors.New("[excel-scanner]: rows can't be nil")
	ErrNilTitleRow = errors.New("[excel-scanner]: title row can't be nil")
	//ErrSliceToString means only []uint8 can be transmuted into string
	ErrSliceToString = errors.New("[excel-scanner]: can't transmute a non-uint8 slice to string")
	//ErrEmptyResult occurs when target of Scan isn't slice and the result of the query is empty
	ErrEmptyResult = errors.New(`[excel-scanner]: empty result`)
)

var _byteUnmarshalerType = reflect.TypeOf(new(ByteUnmarshaler)).Elem()

type ByteUnmarshaler interface {
	UnmarshalByte(data []byte) error
}

type Config struct {
	// The name of the table contained in excel. if it is empty, default is 'Sheet1'.
	SheetNames []string
	// Use the index row as title, every row before title-row will be ignore, default is 0.
	TitleRowIndex int
	// Skip n row after title, default is 0 (not skip), empty row is not counted.
	Skip int
}

//ScanErr will be returned if an underlying type couldn't be AssignableTo type of target field
type ScanErr struct {
	structName, fieldName string
	from, to              reflect.Type
}

func newScanErr(structName, fieldName string, from, to reflect.Type) ScanErr {
	return ScanErr{structName, fieldName, from, to}
}

func (s ScanErr) Error() string {
	return fmt.Sprintf("[excel-scanner]: %s.%s is %s which is not AssignableBy %s", s.structName, s.fieldName, s.to, s.from)
}

func ScanExcelByFilename(ctx context.Context, filename string, target interface{}, config *Config) (err error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	shellDataMap, err := ResolveDataFromFile(ctx, file, config)
	if err != nil {
		return
	}

	for _, rows := range shellDataMap {
		err = ScanExcel(rows, target)
		if err != nil {
			return
		}
	}
	return
}

func ResolveDataFromFile(ctx context.Context, file *excelize.File, config *Config) (shellDataMap map[string][]map[string]string, err error) {
	op := errors.Op("ResolveDataFromFile")
	defer func() {
		if err != nil {
			err = errors.E(op, err, errors.Internal)
		}
	}()

	shellDataMap = make(map[string][]map[string]string)
	for _, sname := range config.SheetNames {
		rows, err := file.GetRows(sname)
		if err != nil {
			return shellDataMap, err
		}
		rowList, err := resolveDataFromRows(ctx, rows, config.TitleRowIndex, config.Skip)
		if err != nil {
			return shellDataMap, err
		}
		if _, ok := shellDataMap[sname]; ok {
			err = errors.E(op, fmt.Errorf("excel have the same shell name  %s ", sname))
			return shellDataMap, err
		}
		shellDataMap[sname] = rowList
	}
	return
}

func resolveDataFromRows(ctx context.Context, rows [][]string, titleRowIndex, skip int) (rowList []map[string]string, err error) {
	op := errors.Op("resolveDataFromRows")
	defer func() {
		if err != nil {
			err = errors.E(op, err)
		}
	}()

	if len(rows) == 0 {
		err = ErrNilRows
		return
	}

	if len(rows) < titleRowIndex {
		err = errors.E(op, "The number of rows %d is less than titleRowIndex", len(rows), titleRowIndex)
		return
	}
	titleRow := rows[titleRowIndex]
	if len(titleRow) == 0 {
		err = ErrNilTitleRow
		return
	}

	for i := titleRowIndex + skip; i < len(rows); i++ {
		row := rows[i]
		// Empty row will be skipped.
		if len(row) == 0 {
			xlog.Warnf(ctx, "%s excel row num %d data is empty", op, i+1)
			continue
		}
		columnMap := make(map[string]string, len(titleRow))
		for i, column := range titleRow {
			columnMap[column] = row[i]
		}
		rowList = append(rowList, columnMap)
	}

	return
}

func ScanExcel(rows []map[string]string, target interface{}) (err error) {
	op := errors.Op("ScanExcel")
	defer func() {
		if err != nil {
			err = errors.E(op, err, errors.Internal)
		}
	}()

	if nil == target || reflect.ValueOf(target).IsNil() || reflect.TypeOf(target).Kind() != reflect.Ptr {
		return ErrTargetNotSettable
	}

	var rowList []map[string]interface{}
	for _, row := range rows {
		columnMap := make(map[string]interface{})
		for k, v := range row {
			columnMap[k] = v
		}
		rowList = append(rowList, columnMap)
	}
	switch reflect.TypeOf(target).Elem().Kind() {
	case reflect.Slice:
		if nil == rowList {
			return nil
		}
		err = bindSlice(rowList, target)
	default:
		if nil == rowList {
			return ErrEmptyResult
		}
		err = bind(rowList[0], target)
	}

	return err
}

// caller must guarantee to pass a &slice as the second param
func bindSlice(arr []map[string]interface{}, target interface{}) (err error) {
	op := errors.Op("bindSlice")
	defer func() {
		if err != nil {
			err = errors.E(op, err)
		}
	}()

	targetObj := reflect.ValueOf(target)
	if !targetObj.Elem().CanSet() {
		return ErrTargetNotSettable
	}
	length := len(arr)
	valueArrObj := reflect.MakeSlice(targetObj.Elem().Type(), 0, length)
	typeObj := valueArrObj.Type().Elem()
	for i := 0; i < length; i++ {
		newObj := reflect.New(typeObj)
		newObjInterface := newObj.Interface()
		err = bind(arr[i], newObjInterface)
		if nil != err {
			return
		}
		valueArrObj = reflect.Append(valueArrObj, newObj.Elem())
	}
	targetObj.Elem().Set(valueArrObj)
	return
}

func bind(result map[string]interface{}, target interface{}) (err error) {
	op := errors.Op("bind")
	defer func() {
		if r := recover(); nil != r {
			err = errors.E(op, fmt.Errorf("error:[%v], stack:[%s]", r, string(debug.Stack())))
		}
	}()

	valueObj := reflect.ValueOf(target).Elem()
	if !valueObj.CanSet() {
		return ErrTargetNotSettable
	}
	typeObj := valueObj.Type()
	if typeObj.Kind() == reflect.Ptr {
		ptrType := typeObj.Elem()
		newObj := reflect.New(ptrType)
		newObjInterface := newObj.Interface()
		err := bind(result, newObjInterface)
		if nil == err {
			valueObj.Set(newObj)
		}
		return err
	}
	typeObjName := typeObj.Name()

	for i := 0; i < valueObj.NumField(); i++ {
		fieldTypeI := typeObj.Field(i)
		fieldName := fieldTypeI.Name

		valuei := valueObj.Field(i)
		if !valuei.CanSet() {
			continue
		}
		tagName, ok := lookUpTagName(fieldTypeI)
		if !ok || "" == tagName {
			continue
		}
		mapValue, ok := result[tagName]
		if !ok || mapValue == nil {
			continue
		}

		// 指针类型走unmarshal逻辑
		// 如果一个字段是指针类型，则必须先为其分配内存，然后再进行json解组，除非该指针类型实现了ByteUnmarshaler接口
		if fieldTypeI.Type.Kind() == reflect.Ptr && !fieldTypeI.Type.Implements(_byteUnmarshalerType) {
			if fieldTypeI.Type.Elem().Kind() == reflect.Struct {
				err := defaultStructUnmarshal(&valuei, mapValue)
				if err == nil {
					continue
				}
			}
			valuei.Set(reflect.New(fieldTypeI.Type.Elem()))
			valuei = valuei.Elem()
		}
		if fieldTypeI.Type.Kind() == reflect.Slice {
			err := defaultSliceUnmarshal(&valuei, mapValue)
			if err == nil {
				continue
			}
		}

		// 结构体类型走unmarshal逻辑
		if fieldTypeI.Type.Kind() == reflect.Struct && fieldTypeI.Type.String() != "time.Time" {
			vPtr := reflect.New(valuei.Type())
			err := defaultStructUnmarshal(&vPtr, mapValue)
			if err == nil {
				valuei.Set(vPtr.Elem())
				continue
			}
		}

		err := convert(mapValue, valuei, func(from, to reflect.Type) ScanErr {
			return newScanErr(typeObjName, fieldName, from, to)
		})
		if nil != err {
			return err
		}
	}
	return nil
}

func lookUpTagName(typeObj reflect.StructField) (string, bool) {
	var tName string
	if "" != UserDefinedTagName {
		tName = UserDefinedTagName
	} else {
		tName = DefaultTagName
	}
	name, ok := typeObj.Tag.Lookup(tName)
	if !ok {
		return "", false
	}
	name = resolveTagName(name)
	return name, ok
}

// TODO support split(|) etc
func resolveTagName(tag string) string {
	idx := strings.IndexByte(tag, ',')
	if -1 == idx {
		return tag
	}
	return tag[:idx]
}

func defaultStructUnmarshal(valuei *reflect.Value, mapValue interface{}) error {
	var pt reflect.Value
	initFlag := false
	// init pointer
	if valuei.IsNil() {
		pt = reflect.New(valuei.Type().Elem())
		initFlag = true
	} else {
		pt = *valuei
	}
	err := json.Unmarshal(mapValue.([]byte), pt.Interface())
	if nil != err {
		structName := pt.Elem().Type().Name()
		return fmt.Errorf("[excel-scanner]: %s.Unmarshal fail to unmarshal the bytes, err: %s", structName, err)
	}
	if initFlag {
		valuei.Set(pt)
	}
	return nil
}

func defaultSliceUnmarshal(valuei *reflect.Value, mapValue interface{}) error {
	var pt reflect.Value
	initFlag := false
	// init pointer
	if valuei.IsNil() {
		// 创建slice
		itemslice := reflect.MakeSlice(valuei.Type(), 0, 0)
		// 指针赋值
		pt = reflect.New(itemslice.Type())
		// 指针指向slice
		pt.Elem().Set(itemslice)
		initFlag = true
	} else {
		pt = *valuei
	}
	err := json.Unmarshal(mapValue.([]byte), pt.Interface())
	if nil != err {
		structName := pt.Elem().Type().Name()
		return fmt.Errorf("[excel-scanner]: %s.Unmarshal fail to unmarshal the bytes, err: %s", structName, err)
	}
	if initFlag {
		valuei.Set(pt.Elem())
	}
	return nil
}

type convertErrWrapper func(from, to reflect.Type) ScanErr

func convert(mapValue interface{}, valuei reflect.Value, wrapErr convertErrWrapper) error {
	vit := valuei.Type()
	mvt := reflect.TypeOf(mapValue)
	if nil == mvt {
		return nil
	}
	//[]byte tp []byte && time.Time to time.Time
	if mvt.AssignableTo(vit) {
		valuei.Set(reflect.ValueOf(mapValue))
		return nil
	}
	//time.Time to string
	switch assertT := mapValue.(type) {
	case time.Time:
		return handleConvertTime(assertT, mvt, vit, &valuei, wrapErr)
	}

	//according to go-mysql-driver/mysql, driver.Value type can only be:
	//int64 or []byte(> maxInt64)
	//float32/float64
	//[]byte
	//time.Time if parseTime=true or DATE type will be converted into []byte
	switch mvt.Kind() {
	case reflect.Int64:
		if isIntSeriesType(vit.Kind()) {
			valuei.SetInt(mapValue.(int64))
		} else if isUintSeriesType(vit.Kind()) {
			valuei.SetUint(uint64(mapValue.(int64)))
		} else if vit.Kind() == reflect.Bool {
			v := mapValue.(int64)
			if v > 0 {
				valuei.SetBool(true)
			} else {
				valuei.SetBool(false)
			}
		} else if vit.Kind() == reflect.String {
			valuei.SetString(strconv.FormatInt(mapValue.(int64), 10))
		} else {
			return wrapErr(mvt, vit)
		}
	case reflect.Float32:
		if isFloatSeriesType(vit.Kind()) {
			valuei.SetFloat(float64(mapValue.(float32)))
		} else {
			return wrapErr(mvt, vit)
		}
	case reflect.Float64:
		if isFloatSeriesType(vit.Kind()) {
			valuei.SetFloat(mapValue.(float64))
		} else {
			return wrapErr(mvt, vit)
		}
	case reflect.Slice:
		return handleConvertSlice(mapValue, mvt, vit, &valuei, wrapErr)
	default:
		return wrapErr(mvt, vit)
	}
	return nil
}

func handleConvertTime(assertT time.Time, mvt, vit reflect.Type, valuei *reflect.Value, wrapErr convertErrWrapper) error {
	if vit.Kind() == reflect.String {
		sTime := assertT.Format(cTimeFormat)
		valuei.SetString(sTime)
		return nil
	}
	return wrapErr(mvt, vit)
}

func isIntSeriesType(k reflect.Kind) bool {
	return k >= reflect.Int && k <= reflect.Int64
}

func isUintSeriesType(k reflect.Kind) bool {
	return k >= reflect.Uint && k <= reflect.Uint64
}

func isFloatSeriesType(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

func handleConvertSlice(mapValue interface{}, mvt, vit reflect.Type, valuei *reflect.Value, wrapErr convertErrWrapper) error {
	mapValueSlice, ok := mapValue.([]byte)
	if !ok {
		return ErrSliceToString
	}
	mapValueStr := string(mapValueSlice)
	vitKind := vit.Kind()
	switch {
	case vitKind == reflect.String:
		valuei.SetString(mapValueStr)
	case isIntSeriesType(vitKind):
		intVal, err := strconv.ParseInt(mapValueStr, 10, 64)
		if nil != err {
			return wrapErr(mvt, vit)
		}
		valuei.SetInt(intVal)
	case isUintSeriesType(vitKind):
		uintVal, err := strconv.ParseUint(mapValueStr, 10, 64)
		if nil != err {
			return wrapErr(mvt, vit)
		}
		valuei.SetUint(uintVal)
	case isFloatSeriesType(vitKind):
		floatVal, err := strconv.ParseFloat(mapValueStr, 64)
		if nil != err {
			return wrapErr(mvt, vit)
		}
		valuei.SetFloat(floatVal)
	case vitKind == reflect.Bool:
		intVal, err := strconv.ParseInt(mapValueStr, 10, 64)
		if nil != err {
			return wrapErr(mvt, vit)
		}
		if intVal > 0 {
			valuei.SetBool(true)
		} else {
			valuei.SetBool(false)
		}
	default:
		if _, ok := valuei.Interface().(ByteUnmarshaler); ok {
			return byteUnmarshal(mapValueSlice, valuei, wrapErr)
		}
		return wrapErr(mvt, vit)
	}
	return nil
}

// valuei Here is the type of ByteUnmarshaler
func byteUnmarshal(mapValueSlice []byte, valuei *reflect.Value, wrapErr convertErrWrapper) error {
	var pt reflect.Value
	initFlag := false
	// init pointer
	if valuei.IsNil() {
		pt = reflect.New(valuei.Type().Elem())
		initFlag = true
	} else {
		pt = *valuei
	}
	err := pt.Interface().(ByteUnmarshaler).UnmarshalByte(mapValueSlice)
	if nil != err {
		structName := pt.Elem().Type().Name()
		return fmt.Errorf("[excle-scanner]: %s.UnmarshalByte fail to unmarshal the bytes, err: %s", structName, err)
	}
	if initFlag {
		valuei.Set(pt)
	}
	return nil
}
