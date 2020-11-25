package main

import (

    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
    ID int64
    // Name *string `gorm:"default:'小王子'"`
    // Name sql.NullString `gorm:"default:'小王子'"`
    Name string `gorm:"default:'小王子'"`
    Age int64
}

type Email struct {
    ID int64
    Address string
    UserID int64
}

func main() {
    db, err := gorm.Open("mysql","root:123456@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    db.AutoMigrate(User{})
    db.AutoMigrate(Email{})

    //     user := User{Name:"abin",Age:18}
//     fmt.Println(db.NewRecord(user))
// db.Create(&user)
//     fmt.Println(db.NewRecord(user))


// var user = User{Name:"",Age:99}
// db.Debug().Create(&user)

// user := User{Name:new(string),Age:18}
// db.Create(&user)


// user := User{Name:sql.NullString{"",true},Age:66}
// db.Create(&user)
// var user User
// db.First(&user,3)
//     fmt.Printf("%#v\n",user)

// var user User
    // db.FirstOrInit(&user,User{Age:88})
    // fmt.Printf("%#v\n",user)
    // db.FirstOrInit(&user,User{Age:99})
    //     fmt.Printf("%#v\n",user)
    // db.Where(User{Age:88}).FirstOrInit(&user)
    // fmt.Printf("%#v\n",user)
// db.FirstOrInit(&user,map[string]interface{}{"age":18})
//     fmt.Printf("%#v\n",user)


// db.Where(User{Age:18}).Attrs("name","莎莎").FirstOrInit(&user)
//     fmt.Printf("%#v\n",user)
//     db.Where(User{Age:55}).Attrs("name","莎莎").FirstOrInit(&user)
//     fmt.Printf("%#v\n",user)

    // db.Where(User{Age:55}).Assign("name","莎莎").FirstOrInit(&user)
    // fmt.Printf("%#v\n",user)
    // db.Where(User{Age:18}).Assign("name","莎莎").FirstOrInit(&user)
    // fmt.Printf("%#v\n",user)

    // db.Where(User{Age:18}).Attrs("name","朱朱").FirstOrCreate(&user)
    // fmt.Printf("%#v\n",user)

    // db.Where(User{Age:18}).Assign("name","朱朱").FirstOrCreate(&user)
    // fmt.Printf("%#v\n",user)

    // var users []User
    // db.Table("users").Select("name,age").Find(&users)
    // fmt.Printf("%#v\n",users)

    // db.Select([]string{"name","age"}).Find(&users)
    // fmt.Printf("%#v\n",users)

//     var user1 []User
//     var user2 []User
// db.Order("age desc").Find(&user1).Order("age",true).Find(&user2)
//     fmt.Printf("%#v\n",user1)
//     fmt.Printf("%#v\n",user2)

// var users []User
// db.Offset(2).Find(&users)
//         fmt.Printf("%#v\n",users)

//     type Result struct {
//         Name string
//         Address string
//     }
// var results []Result
// db.Table("users").Select("users.name,emails.address").Joins("join emails on users.id=emails.user_id").Scan(&results)
// fmt.Printf("%#v\n",results)

// var users []User
// var names []string
// db.Find(&users).Pluck("name",&names)
//     fmt.Printf("%#v\n",names)

    // var names []string
    // db.Model(User{}).Pluck("name",&names)
    // fmt.Printf("%#v\n",names)

    // var users1 []User
    // var users2 []User

    // type Result struct {
    //     Name string
    //     Age int64
    // }
    // var result []Result

    // db.Select("name,age").Find(&users1)
    // fmt.Printf("%#v\n",users1)
    // db.Table("users").Select("name,age").Scan(&users2)
    // fmt.Printf("%#v\n",users2)
    // db.Table("users").Select("name,age").Scan(&result)
    // fmt.Printf("%#v\n",result)

// var user User
// db.First(&user)
// user.Name = "七米"
// user.Age = 199
// db.Debug().Save(&user)
//     fmt.Printf("%#v\n",user)

// var user User
// db.Model(&user).Where("id = ?", 7).Updates(map[string]interface{}{"name":"小花花","age":0,})
// fmt.Printf("%#v\n",user)


    // var user User
    // db.Model(&user).Where("id = ?", 6).Updates(User{Name:"大花花",Age:0})
    // fmt.Printf("%#v\n",user)

    // var user User
    //    // db.Model(&user).Where("id = ?", 7).Omit("name").Updates(map[string]interface{}{"name":"大沙沙","age":21,})
    //    // fmt.Printf("%#v\n",user)


var user User
db.First(&user)

db.Model(&user).Update("age",gorm.Expr("age * ? + ?",2,1))
fmt.Printf("%#v\n",user)

}