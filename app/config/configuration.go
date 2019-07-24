package conf

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var (
	// GlobalConfig 全局配置集合
	GlobalConfig map[string]string
)

// InitConfig 用于读取指定路径下的配置文件
/**
import conf "tom_club/config" // 导入的是文件夹名，包名和文件夹名不一致。此包用于测试该写法
//导入配置文件
configMap := conf.InitConfig("config.ini")
//获取配置里host属性的value
fmt.Println(configMap["host"])
//查看配置文件里所有键值对
fmt.Println(configMap)
*/
func InitConfig(path string) map[string]string {
	//初始化
	myMap := make(map[string]string)

	//打开文件指定目录，返回一个文件f和错误信息
	f, err := os.Open(path)
	defer f.Close()

	//异常处理 以及确保函数结尾关闭文件流
	if err != nil {
		panic(err)
	}

	//创建一个输出流向该文件的缓冲流*Reader
	r := bufio.NewReader(f)
	for {
		//读取，返回[]byte 单行切片给b
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		//去除单行属性两端的空格
		s := strings.TrimSpace(string(b))
		//fmt.Println(s)

		//判断等号=在该行的位置
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		//取得等号左边的key值，判断是否为空
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}

		//取得等号右边的value值，判断是否为空
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		//这样就成功吧配置文件里的属性key=value对，成功载入到内存中c对象里
		myMap[key] = value
	}
	GlobalConfig = myMap
	return myMap
}

// GetConfig 获取初始化之后的配置文件
func GetConfig(path string) (map[string]string, error) {
	if GlobalConfig != nil {
		return GlobalConfig, nil
	}

	// re init
	return InitConfig(path), nil
}

// SimpleGetConfig 获取初始化之后的配置文件
func SimpleGetConfig() map[string]string {
	if GlobalConfig != nil {
		return GlobalConfig
	}

	// re init
	return InitConfig("config.ini")
}
