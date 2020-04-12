package main
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)
//定义与表结构一致的结构体
type User struct {
	Id int64
	Name string
	Salt string
	Age int
	Passwd string `xorm:"varchar(32)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
func main() {
	//连接MySQL数据库
	engine, err := xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	if err != nil{
		fmt.Println(err)
		return
	}
	engine.ShowSQL(true)//设置在控制台输出SQL语句，默认为false
	var users []User//声明结构体切片
	err = engine.Find(&users)//查询一条数据
	fmt.Println(users)
}


