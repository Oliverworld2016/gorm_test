package gorm

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
    ID uint
    Name string
    Gender string
    Hobby string
}

// 使用`AnimalID`作为主键
type Animal struct {
    AnimalID int64 `gorm:"primary_key"`
    Name string     `gorm:"column:animal_name"`
    Age int64
}

func main()  {
    db,err := gorm.Open("mysql","root:123456@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    // db.AutoMigrate(&UserInfo{})
    //
    // // u1 := UserInfo{1,"abin","男","篮球"}
    // // u2 := UserInfo{2,"小花","女","足球"}
    // // db.Create(&u1)
    // // db.Create(&u2)
    //
    // var u = new(UserInfo)
    // db.First(u)
    // fmt.Printf("%#v\n",u)
    //
    // var uu = new(UserInfo)
    // db.Find(&uu,"hobby = ?","足球")
    // fmt.Printf("%#v\n",uu)
    //
    // db.Model(&u).Update("hobby","羽毛球")
    //
    // db.Delete(&u)
    db.AutoMigrate(&Animal{})

}