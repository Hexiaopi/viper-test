package main

import (
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"viper-test/config"
)

// 配置优先级：命令行参数 > 环境变量 > 配置文件 > 默认值
func main() {
	configFilePath := pflag.String("config", "", "配置文件路径")

	pflag.String("http.host", "localhost", "服务主机地址")
	pflag.Int("http.port", 8080, "服务端口")
	pflag.Bool("http.debug", false, "是否启用调试模式")
	pflag.String("db.user", "root", "数据库用户名")
	pflag.String("db.password", "password", "数据库密码")
	pflag.String("db.host", "localhost", "数据库主机地址")
	pflag.Int("db.port", 3306, "数据库端口")
	pflag.String("db.database", "test", "数据库名称")

	// 解析命令行参数
	pflag.Parse()
	// 绑定命令行参数到viper
	viper.BindPFlags(pflag.CommandLine)

	// 设置配置文件信息
	if *configFilePath != "" {
		viper.SetConfigFile(*configFilePath) // 指定配置文件路径
	} else {
		viper.SetConfigName("config") // 配置文件名称（不带扩展名）
		viper.SetConfigType("yaml")   // 配置文件类型
		viper.AddConfigPath(".")      // 当前目录
	}

	// 从环境变量读取配置
	viper.SetEnvPrefix("APP")                              // 环境变量前缀
	viper.AutomaticEnv()                                   // 自动读取匹配的环境变量
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // 替换点为下划线

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("配置文件未找到，将使用命令行参数、环境变量和程序内置默认值（如未指定则使用默认值）")
		} else {
			log.Fatalf("无法读取配置文件: %v", err)
		}
	} else {
		log.Println("使用配置文件:", viper.ConfigFileUsed())
	}

	// 1.直接获取配置值
	log.Printf("主机: %s\n", viper.GetString("http.host"))
	log.Printf("端口: %d\n", viper.GetInt("http.port"))
	log.Printf("调试模式: %t\n", viper.GetBool("http.debug"))
	log.Printf("db 用户名: %s\n", viper.GetString("db.user"))
	log.Printf("db 密码: %s\n", viper.GetString("db.password"))
	log.Printf("db 主机: %s\n", viper.GetString("db.host"))
	log.Printf("db 端口: %d\n", viper.GetInt("db.port"))
	log.Printf("db 数据库名: %s\n", viper.GetString("db.database"))
	// 2.使用结构体获取配置值
	var c config.App
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("配置解析失败: %v", err)
	}
	log.Println("解析后的配置:", c)
}
