package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	if d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 100, //一次最大创建100条
		//Logger:          logger.Default.LogMode(logger.Info), //显示执行的sql语句
		AllowGlobalUpdate: false, //禁止无条件更新
	}); err != nil {
		fmt.Println(err)
	} else {
		db = d
	}
}

func main() {
	//创建及同步表
	//db.AutoMigrate(&Student{})

	//add()
	//Search()
	//update()
	//delete()

	// re := db.Create(&Student{Name: "张三", Age: 18, Birthday: time.Now()})
	// fmt.Println(re.RowsAffected, re.Error)
	// re := db.Session(&gorm.Session{SkipHooks: true}).Create([]Student{
	// 	{Name: "张三", Age: 18, Birthday: time.Now()},
	// 	{Name: "李四", Age: 19, Birthday: time.Now()},
	// 	{Name: "王五", Age: 20, Birthday: time.Now()},
	// })
	// fmt.Println(re.RowsAffected, re.Error)

	// var stus []Student
	// re := db.Where(Student{Name: "张三", Age: 18}).Order("id asc").Limit(2).Offset(1).Find(&stus)
	// fmt.Println(re.RowsAffected, re.Error)
	// for _, v := range stus {
	// 	fmt.Println(v)
	// }
	// db.Model(&stus[0]).Updates(Student{Name: "big zhang", Age: 22})

	//db.Model(&Student{}).Where(Student{Name: "big zhang"}).Updates(Student{Name: "big big zhang"})
	//db.Unscoped().Model(&Student{}).Where(Student{Name: "big big zhang"}).Delete(&Student{})

	//原生sql 和构造器
	// type Af struct {
	// 	Name string
	// 	Age  int
	// }
	// var sts []Af
	// db.Raw("select name,age from students").Scan(&sts)
	// for _, v := range sts {
	// 	fmt.Printf("v: %v\n", v)
	// }

	//db.Exec() //执行语句
	//DryRun模式 只生成sql不执行
	// sql := db.Session(&gorm.Session{DryRun: true}).First(&Student{})
	// fmt.Printf("sql.Statement.SQL.String(): %v\n", sql.Statement.SQL.String())

	//事务
	db.Begin()
	db.Rollback()
	db.SavePoint("sp")
	db.RollbackTo("sp")
	db.Commit()
	db.Transaction(func(tx *gorm.DB) (err error) {
		tx.Where("1=1").First(&Student{})
		return
	})
}

func Releation() {
	// //饭卡 老师和同学共有 多对一
	// type Card struct {
	// 	gorm.Model
	// 	No         int
	// 	Expire     time.Time
	// 	OwnnerID   int
	// 	OwnnerType string
	// }

	// type Teacher struct {
	// 	gorm.Model
	// 	Name string
	// 	Sex  string
	// 	Card `gorm:"poly:ownner"`
	// }

	// type Student struct {
	// 	gorm.Model
	// }
}

type Student struct {
	gorm.Model
	Name     string `gorm:"default:'测试'"` //设定默认值
	Age      int    `gorm:"default:18"`
	Birthday time.Time
	Active   bool
}

// 钩子 创建前触发
func (stu *Student) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("创建前触发")
	return
}

// 钩子 创建前触发
func (stu *Student) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("创建前触发")
	return
}

// 钩子 保存前触发
func (stu *Student) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("保存前触发")
	return
}

// 钩子 保存后触发
func (stu *Student) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("保存后触发")
	return
}

// 增加数据
func add() {
	var stu = Student{
		Name:     "赵六",
		Age:      26,
		Birthday: time.Now(),
	}
	//默认创建
	// d := db.Create(&stu)
	// fmt.Printf("d.RowsAffected: %v\n", d.RowsAffected)

	//db.Select("Name", "CreateAt", "Birthday").Create(&stu) //指定字段创建
	//db.Omit("name").Create(&stu) //忽略字段创建

	//批量插入
	// var stus = []Student{
	// 	{Name: "张三", Age: 18, Birthday: time.Now()},
	// 	{Name: "李四", Age: 19, Birthday: time.Now()},
	// 	{Name: "王五", Age: 20, Birthday: time.Now()},
	// 	{Name: "赵六", Age: 21, Birthday: time.Now()},
	// 	{Name: "钱七", Age: 22, Birthday: time.Now()},
	// 	{Name: "？8", Age: 23, Birthday: time.Now()},
	// }
	//db.Create(&stus) //一次创建
	// db.CreateInBatches(&stus, 3) //分批创建 每次创建2个

	//跳过钩子创建
	// d := db.Session(&gorm.Session{SkipHooks: true}).Create(&stu)
	// fmt.Printf("d.RowsAffected: %v\n", d.RowsAffected)

	//使用MAP创建 会忽略model以及对应的钩子
	// d := db.Model(&stu).Create(map[string]interface{}{"Name": "张三", "Age": 18})
	// fmt.Printf("d.RowsAffected: %v\n", d.RowsAffected)

	//创建或更新
	d := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&stu)
	fmt.Printf("d.RowsAffected: %v\n", d.RowsAffected)
}

// 查询数据
func Search() {
	//var tar Student
	//db.First(&tar, 3) //查询id为3的数据
	//db.First(&tar, "name = ?", "张三") //条件查询
	//db.First(&tar, "name = '张三'")

	//db.First(&tar, "name = ?", "张三") // 第一条数据
	//db.Last(&tar, "name = ?", "张三") //最后一条数据
	//d := db.Take(&tar, "name = ? order by birthday", "张三")
	// d := db.Take(&tar, []int{7, 8, 9})
	// fmt.Printf("tar: %v\n", tar)
	// fmt.Println(d.RowsAffected, d.Error)

	var stus []Student
	//d2 := db.Where("name=?", "张三").Find(&stus, []int{7, 8, 9, 10, 11, 23, 24, 1})
	//d2 := db.Where(&Student{Name: "张三", Age: 18}).Find(&stus)
	//d2 := db.Where(map[string]interface{}{"Name": "张三", "Age": 18}).Find(&stus)
	d2 := db.Not([]int{1, 2, 3, 4, 5}).Or([]int{7, 8, 9}).Order("age asc").Limit(10).Offset(10).Distinct("name").Find(&stus)
	fmt.Println(d2.RowsAffected, d2.Error)
	for k, v := range stus {
		fmt.Println(k, v)
	}

	// d3, _ := db.Table("students").Where(&Student{Name: "张三", Age: 18}).Rows()
	// for d3.Next() {
	// 	fmt.Println(d3)
	// 	d3.Scan(&Student{})
	// }
}

// 更新数据
func update() {

	// db.Create(&Student{
	// 	Name:     "张三",
	// 	Age:      18,
	// 	Birthday: time.Now(),
	// 	Active:   true,
	// })

	//指定对象更新
	// var stu Student
	// db.First(&stu, 1)
	// // db.Model(&stu).Update("name", "李四") //单值更新
	// db.Model(&stu).Updates(Student{Name: "张三", Age: 20}) //多值更新

	//未指定对象更新
	//db.Model(&Student{}).Where("name=?", "张三").Update("Age", 22) //更新单值
	//db.Model(&Student{}).Where("name=?", "张三").Updates(Student{Age: 25})

	// stu = Student{}
	// db.First(&stu, 2)
	// db.Model(&stu).Updates(Student{Name: "李四", Age: 19})

	// var stu = Student{}
	// db.First(&stu, 1)
	//db.Model(&stu).Updates(map[string]interface{}{"name": "王五", "age": 20})
	//指定字段更新 不会更新时间和触发钩子
	db.Model(Student{}).Where("1=1").UpdateColumns(Student{Name: "132", Age: 28})
}

// 删除数据 软删除 记录还在数据库
func delete() {
	var stu Student
	//db.First(&stu, 4)
	db.Delete(&stu, 4)            //软删除
	db.Unscoped().Delete(&stu, 1) //硬删除
}
