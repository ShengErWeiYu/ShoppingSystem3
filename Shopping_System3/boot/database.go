package boot

//连接数据库
import (
	g "Shopping_System/global"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func DatabaseInit() {
	MysqlInit()
	RedisInit()
}
func MysqlInit() {
	var err error
	dsn := "root:haoting39872256@tcp(127.0.0.1:3306)/shopping?charset=utf8mb4&parseTime=True&loc=Local"
	g.Xdb, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	g.Xdb.SetMaxOpenConns(20)
	g.Xdb.SetMaxIdleConns(10)
	log.Println("MYSQL connect successfully")
}
func RedisInit() {
	g.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	log.Println("Redis Connect successfully")
}
