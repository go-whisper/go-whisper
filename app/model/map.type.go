package model

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/go-whisper/go-whisper/app/instance"
	"go.uber.org/zap"
)

type StringMap map[string]string
type InterfaceMap map[string]interface{}

func (strMap StringMap) Value() (driver.Value, error) {
	return json.Marshal(strMap)
}
func (strMap *StringMap) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &strMap)
}

// GetInt 返回 key 对应值的 int 值
// 如果 key 不存在或者其值无法转换成 int 形式，则返回的 ok 值为 false
func (strMap StringMap) GetInt(key string) (val int, ok bool) {
	if v, _has := strMap[key]; _has {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			instance.Logger().Error("StringMap.GetInt() 获取key的值无法转换到 int 类型", zap.String("key", key), zap.Error(err))
			return 0, false
		}
		return int(num), true
	}
	return 0, false
}

func (iMap InterfaceMap) Value() (driver.Value, error) {
	if len(iMap) == 0 {
		return []byte("{}"), nil
	}
	return json.Marshal(iMap)
}
func (iMap *InterfaceMap) Scan(data interface{}) error {
	var d []byte
	switch dType := data.(type) {
	case string:
		d = []byte(dType)
	case []byte:
		d = dType
	}
	return json.Unmarshal(d, &iMap)
}

func (iMap InterfaceMap) GetUint(key string, defVal ...uint) (val uint, has bool) {
	def := uint(0)
	if len(defVal) == 1 {
		def = defVal[0]
	}
	iV, h := iMap[key]
	if !h {
		return def, false
	}
	switch v := iV.(type) {
	case int:
		return uint(v), true
	case uint:
		return v, true
	case int64:
		return uint(v), true
	case uint64:
		return uint(v), true
	case string:
		num, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, false
		}
		return uint(num), true
	case float64:
		return uint(v), true
	default:
		// 异常的类型
	}
	return def, false
}

func (iMap InterfaceMap) ToStringMap() StringMap {
	to := make(StringMap)
	for mapK, mapV := range iMap {
		switch mapVT := mapV.(type) {
		case string:
			to[mapK] = mapVT
		case int:
			to[mapK] = strconv.Itoa(mapVT)
		case int64:
			to[mapK] = strconv.FormatInt(mapVT, 10)
		case float64:
			tmpStr := strconv.FormatFloat(mapVT, 'f', 2, 64)
			if strings.HasSuffix(tmpStr, ".00") {
				tmpStr = tmpStr[:len(tmpStr)-3]
			}
			to[mapK] = tmpStr
		default:
			instance.Logger().Error("InterfaceMapList.ToStringMapList() mapV invalid type")
		}
	}
	return to
}
