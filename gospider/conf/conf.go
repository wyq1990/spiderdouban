package conf

import (
	"douban/model"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	model.Database(os.Getenv("MYSQL_DSN"))
}