package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type CsvUtilMgr struct {
}

var csvUtilMgr *CsvUtilMgr = nil

func GetCsvUtilMgr() *CsvUtilMgr {

	if csvUtilMgr == nil {
		csvUtilMgr = new(CsvUtilMgr)
	}
	return csvUtilMgr
}
func (self *CsvUtilMgr) ReadCsv(filepath string) [][]string {
	csvreadfile, err := os.Open(filepath)
	defer csvreadfile.Close()
	if err != nil {
		//错误记录
		fmt.Println(err.Error())
		return [][]string{}
	}
	r := csv.NewReader(bufio.NewReader(csvreadfile))

	var records [][]string
	var recordIndex int
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if recordIndex == 0 {
			for i := range records {
				if i == 0 && strings.Contains(record[i], "\ufeff") {
					record[i] = strings.Replace(record[i], "\ufeff", "", 1)
					break
				}
			}
			recordIndex = 1
		}
		records = append(records, record)

	}
	return records
}

func (self *CsvUtilMgr) ParseDataSimple(csvData [][]string, dataPtr interface{}, fileName string) (err error) {
	outInnerType := self.GetValueType(dataPtr)
	data := reflect.New(outInnerType.Elem())
	value := reflect.Indirect(data) //返回指针所指的值或者结构
	tagMap := self.GetTagMap(csvData[0])
	fieldMap, trimFlag, keyTag := self.GetFiledMapSimple(value.Interface(), tagMap)
	err = self.genConfig(dataPtr, csvData, tagMap, fieldMap, trimFlag, fileName, keyTag)
	return err
}

func (self *CsvUtilMgr) LoadCsv(filename string, SlicePtr interface{}) {
	csvfile := "C:\\Users\\chang\\OneDrive\\changping的个人文档\\project\\bofyuan\\csv\\" + filename + ".csv" //csv/PlayerLevel.csv
	csvData := self.ReadCsv(csvfile)
	if len(csvData) <= 1 {
		//记录错误
		os.Exit(1) //错误码1 异常退出 使用守护进程进行拉起
		return
	}
	err := self.ParseDataSimple(csvData, SlicePtr, filename)
	if err != nil {
		fmt.Println("error in Loadcsv")
		return
	}
}

func (self *CsvUtilMgr) GetValueType(ptr interface{}) reflect.Type {
	value := reflect.ValueOf(ptr)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	outType := value.Type()
	outInnerType := outType.Elem()
	return outInnerType
}

func (self *CsvUtilMgr) GetTagMap(fields []string) map[int]string {
	tagMap := make(map[int]string, len(fields))
	for i, v := range fields {
		tagMap[i] = v
	}
	return tagMap

}

func (self *CsvUtilMgr) GetFiledMapSimple(config interface{}, tagMap map[int]string) (map[string][]string, map[int]bool, string) {
	t := reflect.TypeOf(config)
	fieldMap := make(map[string][]string, t.NumField())
	trimFlag := make(map[int]bool)
	keyTag := ""
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		trim := field.Tag.Get("trim")
		if trim == "1" || trim == "" {
			tag = self.trimNumber(tag)
		}
		for k, v := range tagMap {
			if (trim == "" || trim == "1") && tag == self.trimNumber(v) {
				tagMap[k] = tag
				break
			}
		}
		data := make([]string, 2, 2)
		data[0] = field.Name
		data[1] = fmt.Sprintf("%v", field.Type)
		fieldMap[tag] = data
		if i == 0 {
			keyTag = tag
		}
	}
	return fieldMap, trimFlag, keyTag
}

func (self *CsvUtilMgr) trimNumber(s string) string {
	subString := strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsNumber(r)
	})

	return subString

}

func (self *CsvUtilMgr) genConfig(dataPtr interface{}, csvData [][]string, tagMap map[int]string,
	fieldMap map[string][]string, trimFlag map[int]bool, fileName string, keyTag string) (err error) {
	dataVal := reflect.Indirect(reflect.ValueOf(dataPtr))
	outInnerType := self.GetValueType(dataPtr)

	//if fileName == "expspeedup" {
	//litter.Dump(tagMap)
	//litter.Dump(fieldMap)
	//}

	for r := 1; r < len(csvData); r++ {
		data := reflect.New(outInnerType.Elem())

		key := 0
		for c := 0; c < len(csvData[r]); c++ {
			tag := tagMap[c]
			if _, ok := trimFlag[c]; ok {
				tag = self.trimNumber(tag)
			}

			fieldInfo, ok := fieldMap[tag]
			if !ok {
				continue
			}

			if len(fieldInfo) != 2 {
				continue
			}
			fieldName := fieldInfo[0]
			filedType := fieldInfo[1]
			cellValue := csvData[r][c]
			switch filedType {
			case "int":
				v, err := strconv.Atoi(cellValue)
				if err != nil {
					fmt.Println(err.Error(), ", fileName:"+fileName, ", fieldName:", fieldName, ", index:", r)
					break
				}
				reflect.Indirect(data).FieldByName(fieldName).SetInt(int64(v))
				if tag == keyTag && key == 0 {
					key = v
					//fmt.Println("fileName:", fileName, ", tag:", tag)
				}
			case "int64":
				v, err := strconv.ParseInt(cellValue, 10, 64)
				if err != nil {
					fmt.Println(err.Error(), ", fileName:"+fileName, ", fieldName:", fieldName)
					break
				}
				reflect.Indirect(data).FieldByName(fieldName).SetInt(v)
			case "string":
				reflect.Indirect(data).FieldByName(fieldName).SetString(cellValue)
			case "[]int":
				v, err := strconv.Atoi(cellValue)
				if err != nil {
					fv, err := strconv.ParseFloat(cellValue, 64)
					if err != nil {
						fmt.Println(err.Error(), ", fileName:"+fileName, ", fieldName:", fieldName, ", row:", r, ", col:", c, ", ", strings.Join(csvData[r], ","))
					} else {
						c := reflect.Indirect(data).FieldByName(fieldName)
						newSlice := reflect.Append(c, reflect.ValueOf(int(fv*10)))
						reflect.Indirect(data).FieldByName(fieldName).Set(newSlice)
					}
					break
				} else {
					c := reflect.Indirect(data).FieldByName(fieldName)
					newSlice := reflect.Append(c, reflect.ValueOf(v))
					reflect.Indirect(data).FieldByName(fieldName).Set(newSlice)
				}
			case "[]int64":
				v, err := strconv.ParseInt(cellValue, 10, 0)
				if err != nil {
					fmt.Println(err.Error(), ", fileName:"+fileName)
					break
				}
				c := reflect.Indirect(data).FieldByName(fieldName)
				newSlice := reflect.Append(c, reflect.ValueOf(v))
				reflect.Indirect(data).FieldByName(fieldName).Set(newSlice)
			case "[]string":
				c := reflect.Indirect(data).FieldByName(fieldName)
				newSlice := reflect.Append(c, reflect.ValueOf(cellValue))
				reflect.Indirect(data).FieldByName(fieldName).Set(newSlice)
			case "float32":
				fv, err := strconv.ParseFloat(cellValue, 32)
				if err != nil {
					fmt.Println(err.Error(), ", fileName:"+fileName, ", fieldName:", fieldName, ", row:", r, ", col:", c, ", ", strings.Join(csvData[r], ","))
				} else {
					reflect.Indirect(data).FieldByName(fieldName).SetFloat(fv)
				}
			case "float64":
				fv, err := strconv.ParseFloat(cellValue, 64)
				if err != nil {
					fmt.Println(err.Error(), ", fileName:"+fileName, ", fieldName:", fieldName, ", row:", r, ", col:", c, ", ", strings.Join(csvData[r], ","))
				} else {
					reflect.Indirect(data).FieldByName(fieldName).SetFloat(fv)
				}
			case "[]float32":
				fv, err := strconv.ParseFloat(cellValue, 32)
				if err != nil {
					fmt.Println(err.Error(), ", fileName:"+fileName, ", fieldName:", fieldName, ", row:", r, ", col:", c, ", ", strings.Join(csvData[r], ","))
				} else {
					c := reflect.Indirect(data).FieldByName(fieldName)
					newSlice := reflect.Append(c, reflect.ValueOf(float32(fv)))
					reflect.Indirect(data).FieldByName(fieldName).Set(newSlice)
				}
			case "[]float64":
				fv, err := strconv.ParseFloat(cellValue, 64)
				if err != nil {
					fmt.Println(err.Error(), ", fileName:"+fileName, ", fieldName:", fieldName, ", row:", r, ", col:", c, ", ", strings.Join(csvData[r], ","))
				} else {
					c := reflect.Indirect(data).FieldByName(fieldName)
					newSlice := reflect.Append(c, reflect.ValueOf(fv))
					reflect.Indirect(data).FieldByName(fieldName).Set(newSlice)
				}
			}
		}

		kind := reflect.TypeOf(dataVal.Interface()).Kind()
		if kind == reflect.Slice {
			dataVal.Set(reflect.Append(dataVal, data))
		} else if kind == reflect.Map {
			dataVal.SetMapIndex(reflect.ValueOf(key), data)
		}
	}

	return nil
}
