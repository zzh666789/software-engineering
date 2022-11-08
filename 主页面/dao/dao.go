package dao

import (
	"blog/model"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	Register(user *model.User) int
	Login(username string) model.User
	Save(code *model.Countcode)
	Save1(code *model.Countcode1)
	Save2(code *model.Countcode2)
	GetAllCount() []model.Countcode
	GetAllCount1() []model.Countcode1
	GetAllCount2() []model.Countcode2
	Deletecount()
	Deletecount1()
	Deletecount2()
	Panduan(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 string) int
	Panduan1(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 string) int
	Panduan2(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 string) int
	Level3count()
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

var EditId int

func init() {
	dsn := "root:cxykxr2002062909@tcp(127.0.0.1:3306)/count?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Countcode{})
	db.AutoMigrate(&model.Countcode1{})
	db.AutoMigrate(&model.Countcode2{})
}

func (mgr *manager) Register(user *model.User) int {
	var a model.User
	if user.Username != "" {
		mgr.db.Where("username = ?", user.Username).Find(&a)
		if a.Username == "" {
			mgr.db.Create(user)
			return 1
		} else {
			return 0
		}
	} else {
		return 2
	}
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) Save(code *model.Countcode) {
	mgr.db.Create(code)
}

func (mgr *manager) Save1(code *model.Countcode1) {
	mgr.db.Create(code)
}

func (mgr *manager) Save2(code *model.Countcode2) {
	mgr.db.Create(code)
}

func (mgr *manager) GetAllCount() []model.Countcode {
	var counts = make([]model.Countcode, 10)
	mgr.db.Find(&counts)
	return counts
}

func (mgr *manager) GetAllCount1() []model.Countcode1 {
	var counts = make([]model.Countcode1, 10)
	mgr.db.Find(&counts)
	return counts
}

func (mgr *manager) GetAllCount2() []model.Countcode2 {
	var counts = make([]model.Countcode2, 10)
	mgr.db.Find(&counts)
	fmt.Println("333")
	return counts
}

func (mgr *manager) Deletecount() {
	mgr.db.Exec("truncate table countcodes;")
}

func (mgr *manager) Deletecount1() {
	mgr.db.Exec("truncate table countcode1;")
}

func (mgr *manager) Deletecount2() {
	mgr.db.Exec("truncate table countcode2;")
}

func (mgr *manager) Panduan(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 string) int {
	var ans1, ans2, ans3, ans4, ans5, ans6, ans7, ans8, ans9, ans10 model.Countcode
	var score int = 0
	mgr.db.Where("id = ?", 1).Find(&ans1)
	mgr.db.Where("id = ?", 2).Find(&ans2)
	mgr.db.Where("id = ?", 3).Find(&ans3)
	mgr.db.Where("id = ?", 4).Find(&ans4)
	mgr.db.Where("id = ?", 5).Find(&ans5)
	mgr.db.Where("id = ?", 6).Find(&ans6)
	mgr.db.Where("id = ?", 7).Find(&ans7)
	mgr.db.Where("id = ?", 8).Find(&ans8)
	mgr.db.Where("id = ?", 9).Find(&ans9)
	mgr.db.Where("id = ?", 10).Find(&ans10)
	if a1 == "未输入" {
		a1 = "-9999"
	}
	if a2 == "未输入" {
		a2 = "-9999"
	}
	if a3 == "未输入" {
		a3 = "-9999"
	}
	if a4 == "未输入" {
		a4 = "-9999"
	}
	if a5 == "未输入" {
		a5 = "-9999"
	}
	if a6 == "未输入" {
		a6 = "-9999"
	}
	if a7 == "未输入" {
		a7 = "-9999"
	}
	if a8 == "未输入" {
		a8 = "-9999"
	}
	if a9 == "未输入" {
		a9 = "-9999"
	}
	if a10 == "未输入" {
		a10 = "-9999"
	}
	b1, _ := strconv.Atoi(a1)
	fmt.Printf("ans1: %v\n", ans1)
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("ans1.Answer: %v\n", ans1.Answer)
	b2, _ := strconv.Atoi(a2)
	b3, _ := strconv.Atoi(a3)
	b4, _ := strconv.Atoi(a4)
	b5, _ := strconv.Atoi(a5)
	b6, _ := strconv.Atoi(a6)
	b7, _ := strconv.Atoi(a7)
	b8, _ := strconv.Atoi(a8)
	b9, _ := strconv.Atoi(a9)
	b10, _ := strconv.Atoi(a10)
	if b1 == ans1.Answer {
		score += 10
	}
	if b2 == ans2.Answer {
		score += 10
	}
	if b3 == ans3.Answer {
		score += 10
	}
	if b4 == ans4.Answer {
		score += 10
	}
	if b5 == ans5.Answer {
		score += 10
	}
	if b6 == ans6.Answer {
		score += 10
	}
	if b7 == ans7.Answer {
		score += 10
	}
	if b8 == ans8.Answer {
		score += 10
	}
	if b9 == ans9.Answer {
		score += 10
	}
	if b10 == ans10.Answer {
		score += 10
	}
	fmt.Printf("score: %v\n", score)

	return score
}

func (mgr *manager) Panduan1(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 string) int {
	var ans1, ans2, ans3, ans4, ans5, ans6, ans7, ans8, ans9, ans10 model.Countcode1
	var score int = 0
	mgr.db.Where("id = ?", 1).Find(&ans1)
	mgr.db.Where("id = ?", 2).Find(&ans2)
	mgr.db.Where("id = ?", 3).Find(&ans3)
	mgr.db.Where("id = ?", 4).Find(&ans4)
	mgr.db.Where("id = ?", 5).Find(&ans5)
	mgr.db.Where("id = ?", 6).Find(&ans6)
	mgr.db.Where("id = ?", 7).Find(&ans7)
	mgr.db.Where("id = ?", 8).Find(&ans8)
	mgr.db.Where("id = ?", 9).Find(&ans9)
	mgr.db.Where("id = ?", 10).Find(&ans10)
	if a1 == "未输入" {
		a1 = "-9999"
	}
	if a2 == "未输入" {
		a2 = "-9999"
	}
	if a3 == "未输入" {
		a3 = "-9999"
	}
	if a4 == "未输入" {
		a4 = "-9999"
	}
	if a5 == "未输入" {
		a5 = "-9999"
	}
	if a6 == "未输入" {
		a6 = "-9999"
	}
	if a7 == "未输入" {
		a7 = "-9999"
	}
	if a8 == "未输入" {
		a8 = "-9999"
	}
	if a9 == "未输入" {
		a9 = "-9999"
	}
	if a10 == "未输入" {
		a10 = "-9999"
	}
	b1, _ := strconv.Atoi(a1)
	fmt.Printf("ans1: %v\n", ans1)
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("ans1.Answer: %v\n", ans1.Answer)
	b2, _ := strconv.Atoi(a2)
	b3, _ := strconv.Atoi(a3)
	b4, _ := strconv.Atoi(a4)
	b5, _ := strconv.Atoi(a5)
	b6, _ := strconv.Atoi(a6)
	b7, _ := strconv.Atoi(a7)
	b8, _ := strconv.Atoi(a8)
	b9, _ := strconv.Atoi(a9)
	b10, _ := strconv.Atoi(a10)
	if b1 == ans1.Answer {
		score += 10
	}
	if b2 == ans2.Answer {
		score += 10
	}
	if b3 == ans3.Answer {
		score += 10
	}
	if b4 == ans4.Answer {
		score += 10
	}
	if b5 == ans5.Answer {
		score += 10
	}
	if b6 == ans6.Answer {
		score += 10
	}
	if b7 == ans7.Answer {
		score += 10
	}
	if b8 == ans8.Answer {
		score += 10
	}
	if b9 == ans9.Answer {
		score += 10
	}
	if b10 == ans10.Answer {
		score += 10
	}
	fmt.Printf("score: %v\n", score)
	return score
}

func (mgr *manager) Panduan2(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 string) int {
	var ans1, ans2, ans3, ans4, ans5, ans6, ans7, ans8, ans9, ans10 model.Countcode2
	var score int = 0
	mgr.db.Where("id = ?", 1).Find(&ans1)
	mgr.db.Where("id = ?", 2).Find(&ans2)
	mgr.db.Where("id = ?", 3).Find(&ans3)
	mgr.db.Where("id = ?", 4).Find(&ans4)
	mgr.db.Where("id = ?", 5).Find(&ans5)
	mgr.db.Where("id = ?", 6).Find(&ans6)
	mgr.db.Where("id = ?", 7).Find(&ans7)
	mgr.db.Where("id = ?", 8).Find(&ans8)
	mgr.db.Where("id = ?", 9).Find(&ans9)
	mgr.db.Where("id = ?", 10).Find(&ans10)
	if a1 == "未输入" {
		a1 = "-9999"
	}
	if a2 == "未输入" {
		a2 = "-9999"
	}
	if a3 == "未输入" {
		a3 = "-9999"
	}
	if a4 == "未输入" {
		a4 = "-9999"
	}
	if a5 == "未输入" {
		a5 = "-9999"
	}
	if a6 == "未输入" {
		a6 = "-9999"
	}
	if a7 == "未输入" {
		a7 = "-9999"
	}
	if a8 == "未输入" {
		a8 = "-9999"
	}
	if a9 == "未输入" {
		a9 = "-9999"
	}
	if a10 == "未输入" {
		a10 = "-9999"
	}
	b1, _ := strconv.Atoi(a1)
	fmt.Printf("ans1: %v\n", ans1)
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("ans1.Answer: %v\n", ans1.Answer)
	b2, _ := strconv.Atoi(a2)
	b3, _ := strconv.Atoi(a3)
	b4, _ := strconv.Atoi(a4)
	b5, _ := strconv.Atoi(a5)
	b6, _ := strconv.Atoi(a6)
	b7, _ := strconv.Atoi(a7)
	b8, _ := strconv.Atoi(a8)
	b9, _ := strconv.Atoi(a9)
	b10, _ := strconv.Atoi(a10)
	if b1 == ans1.Answer {
		score += 10
	}
	if b2 == ans2.Answer {
		score += 10
	}
	if b3 == ans3.Answer {
		score += 10
	}
	if b4 == ans4.Answer {
		score += 10
	}
	if b5 == ans5.Answer {
		score += 10
	}
	if b6 == ans6.Answer {
		score += 10
	}
	if b7 == ans7.Answer {
		score += 10
	}
	if b8 == ans8.Answer {
		score += 10
	}
	if b9 == ans9.Answer {
		score += 10
	}
	if b10 == ans10.Answer {
		score += 10
	}
	fmt.Printf("score: %v\n", score)
	return score
}

func (mgr *manager) Level3count() {
	var countcode2 model.Countcode2
	for i := 0; i < 10; i++ {
		s1 := make([]int, 10000)
		s2 := make([]int, 10000)
		var pos int = 999
		var div int = 1
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Microsecond * 200)
		temp := rand.Intn(20)
		for i := 2; i >= 0; i-- { // 0+   1-   2*   3/
			if div == 1 {
				rand.Seed(time.Now().UnixNano())
				time.Sleep(time.Microsecond * 200)
				s1[i] = rand.Intn(4)
			} else {
				rand.Seed(time.Now().UnixNano())
				time.Sleep(time.Microsecond * 200)
				s1[i] = rand.Intn(3)
			}
			if s1[i] == 3 {
				div = 0
				pos = i
			}
		}

		if div == 0 {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Microsecond * 200)
			s2[pos+1] = rand.Intn(50)
			s2[pos] = s2[pos+1] * temp
		}

		for i := 0; i < 4; i++ {
			if div == 0 && (i == pos || i == pos+1) {
			} else {
				rand.Seed(time.Now().UnixNano())
				time.Sleep(time.Microsecond * 200)
				s2[i] = rand.Intn(100)
			}
		}

		stack1 := make([]int, 1000)
		stack2 := make([]int, 1000)
		top1 := 0
		top2 := 0
		stack1[0] = s2[0]
		top1++

		for i := 1; i < 4; i++ {
			stack1[top1] = s2[i]
			stack2[top2] = s1[i-1]
			if stack2[top2] == 2 {
				stack1[top1-1] = stack1[top1] * stack1[top1-1]
				top1--
				top2--
			} else if stack2[top2] == 3 {
				stack1[top1-1] = stack1[top1-1] / stack1[top1]
				top1--
				top2--
			}
			top1++
			top2++
		}
		top1--
		top2--
		for i := 0; i <= top2; i++ {
			if stack2[i] == 0 {
				stack1[i+1] = stack1[i] + stack1[i+1]
			}
			if stack2[i] == 1 {
				stack1[i+1] = stack1[i] - stack1[i+1]
			}
		}
		var count int = stack1[top1]

		var q string
		if s1[0] == 0 {
			q = "+"
		} else if s1[0] == 1 {
			q = "-"
		} else if s1[0] == 3 {
			q = "/"
		} else {
			q = "*"
		}
		var p string
		if s1[1] == 0 {
			p = "+"
		} else if s1[1] == 1 {
			p = "-"
		} else if s1[1] == 3 {
			p = "/"
		} else {
			p = "*"
		}
		var r string
		if s1[2] == 0 {
			r = "+"
		} else if s1[2] == 1 {
			r = "-"
		} else if s1[2] == 3 {
			r = "/"
		} else {
			r = "*"
		}
		countcode2 = model.Countcode2{
			Operate1: s2[0],
			Operate2: s2[1],
			Operate3: s2[2],
			Operate4: s2[3],
			Code1:    q,
			Code2:    p,
			Code3:    r,
			Answer:   count}
		fmt.Printf("countcode2: %v\n", countcode2)
		Mgr.Save2(&countcode2)
	}
}
