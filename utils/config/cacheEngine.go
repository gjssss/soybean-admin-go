package config

import (
	"encoding/json"
	"errors"
)

type CacheEngine int

const (
	CacheUnknown CacheEngine = iota // 默认未知类型
	CacheRedis                      // Redis
	CacheMemory                     // 内存
)

// 定义字符串映射
var storageTypeToString = map[CacheEngine]string{
	CacheRedis:  "redis",
	CacheMemory: "memory",
}

var stringToStorageType = map[string]CacheEngine{
	"redis":  CacheRedis,
	"memory": CacheMemory,
}

// String 方法实现 fmt.Stringer 接口
func (s CacheEngine) String() string {
	return storageTypeToString[s]
}

// UnmarshalJSON 实现 JSON 解析
func (s *CacheEngine) UnmarshalJSON(data []byte) error {
	// 去掉 JSON 解析的双引号
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// 查找映射
	if val, ok := stringToStorageType[str]; ok {
		*s = val
		return nil
	}
	return errors.New("未知的 storage_type 值: " + str)
}

// MarshalJSON 实现 JSON 序列化
func (s CacheEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
