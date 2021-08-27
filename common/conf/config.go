package conf

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	mapOnce   sync.Once
	configMap map[string]string
	envMap    map[string]string
	rwLock    sync.RWMutex
)

func InitConfig(conffile string) {
	mapOnce.Do(func() {
		// 打开文件
		file, _ := os.Open(conffile)
		// 关闭文件
		defer file.Close()

		decoder := json.NewDecoder(file)

		configMap = map[string]string{}
		err := decoder.Decode(&configMap)
		if err != nil {
			panic(err)
		}
	})
}

func ConfigMap(key string) string {
	if envMap == nil {
		rwLock.Lock()
		envMap = map[string]string{}
		rwLock.Unlock()
	}
	if envMap[key] == "" {
		envKey := strings.ToUpper(key)
		reg := regexp.MustCompile(`[-\\.]+`)
		envKey = reg.ReplaceAllString(envKey, `_`)
		env := os.Getenv(envKey)
		if env != "" {
			rwLock.Lock()
			envMap[key] = env
			rwLock.Unlock()
			return env
		}
		rwLock.Lock()
		envMap[key] = configMap[key]
		rwLock.Unlock()
		return configMap[key]
	}
	return envMap[key]
}
