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

func (v InterfaceMap) ToStringMap() StringMap {
	to := make(StringMap)
	for mapK, mapV := range v {
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
