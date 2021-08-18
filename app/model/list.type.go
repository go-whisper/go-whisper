package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

type StringList []string
type UintList []uint
type StringMapList []StringMap
type InterfaceMapList []InterfaceMap

func NewStringList(s string) StringList {
	toStringList := func(str, separator string) StringList {
		strList := make(StringList, 0)
		tmp := strings.Split(s, separator)
		if len(tmp) > 1 {
			for _, v := range tmp {
				strList = append(strList, v)
			}
		}
		return strList
	}
	// 尝试用逗号分割
	strList := toStringList(s, ",")
	if len(strList) > 1 {
		return strList
	}

	// 尝试用空格分割
	strList = toStringList(s, " ")
	if len(strList) > 1 {
		return strList
	}

	return strList
}

func (v StringList) Value() (driver.Value, error) {
	return json.Marshal(v)
}
func (v *StringList) Scan(data interface{}) error {
	switch d := data.(type) {
	case string:
		return json.Unmarshal([]byte(d), &v)
	case []byte:
		return json.Unmarshal(data.([]byte), &v)
	}
	return errors.New("未知类型")
}

func (v UintList) Value() (driver.Value, error) {
	return json.Marshal(v)
}
func (v *UintList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &v)
}

func (v StringMapList) Value() (driver.Value, error) {
	return json.Marshal(v)
}
func (v *StringMapList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &v)
}

func (v InterfaceMapList) ToStringMapList() StringMapList {
	res := make(StringMapList, len(v))
	for idx, mapItem := range v {
		res[idx] = mapItem.ToStringMap()
	}
	return res
}
