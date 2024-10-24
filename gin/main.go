package main

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
)

type Student struct {
	gorm.Model
	Name     string
	Card     Card
	Teachers []Teacher `gorm:"many2many:student_teachers;"`
}

type Teacher struct {
	gorm.Model
	Name     string
	Students []Student `gorm:"many2many:student_teachers;"`
}

type Card struct {
	gorm.Model
	StudentID uint
	Num       uint
}

func (s *Student) BeforeCreate() {
	s.Name = "default"
	return
}

func main() {
	s := &Student{}
	s.BeforeCreate()
	r := gin.Default()
	// 获取 Gin 默认的验证器引擎
	// MySQL 连接字符串
	dsn := "host=localhost user=postgres password=123456 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// 打开数据库连接
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// 自动迁移
	db.AutoMigrate(&Student{}, &Teacher{}, &Card{})


	r.POST("/addStu", func(c *gin.Context) {
		var stu Student
		// 读取请求体内容到字节数组
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Failed to read request body",
			})
			return
		}
		// 关闭请求体
		c.Request.Body.Close()
		if err := sonic.Unmarshal(body, &stu); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Create(&stu)
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.POST("/delStu/:id", func(c *gin.Context) {
		c.Param("id")
		studentID := c.Param("id")
		.Where

		// 查找学生
		var student Student
		if err := db.Preload("Teachers").First(&student, studentID).Error; err != nil {
			c.JSON(404, gin.H{
				"message": "Student not found",
			})
			return
		}

		// 删除学生的选课记录
		if err := db.Model(&student).Association("Enrollments").Clear(); err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to clear enrollments",
			})
			return
		}

		db.Delete(&student).Where("id = ?", studentID)
        c.JSON(200, gin.H{
            "message": "success",
        })
    })

    r.GET("/getStu", func(c *gin.Context) {
        var stu []Student
        db.Preload("Teachers").Preload("Card").Find(&)
        c.JSON(200, gin.H{
            "message": "success",
        })
	})

	r.GET("/getStu/:id", func(c *gin.Context) {
		id := c.Param("id")
		var stu Student
		db.Preload("Teachers").Preload("Card").First(&stu, id)
		c.JSON(200, gin.H{
			"message": "success",
			"data":    stu,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
