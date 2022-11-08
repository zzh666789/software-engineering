package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func Index1(c *gin.Context) {
	c.HTML(200, "login1.html", nil)
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register1.html", nil)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	a := dao.Mgr.Register(&user)
	if a == 0 {
		c.HTML(200, "register1.html", "用户名重复")
		fmt.Println("用户名重复！")
	} else if a == 2 {
		c.HTML(200, "register1.html", "用户名不合法")
	} else {
		c.HTML(200, "index.html", nil)
	}
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login1.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	u := dao.Mgr.Login(username)
	if u.Username == "" {
		c.HTML(200, "login1.html", "用户名不存在")
		fmt.Println("用户名不存在！")
	} else {
		if password != u.Password {
			fmt.Println("密码错误")
			c.HTML(200, "login1.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/index")
		}
	}
}

func GoIndex(c *gin.Context) {
	dao.Mgr.Deletecount()
	dao.Mgr.Deletecount1()
	dao.Mgr.Deletecount2()
	c.HTML(200, "index.html", nil)
}

func Golevel1(c *gin.Context) {
	var countcode model.Countcode
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(2)
		m := rand.Intn(20)
		n := rand.Intn(20)
		var q int
		var e string
		if a == 0 {
			e = "+"
			q = m + n
		} else {
			e = "-"
			if m < n {
				t := m
				m = n
				n = t
				q = m - n
			} else {
				q = m - n
			}

		}

		countcode = model.Countcode{
			Operate1: m,
			Operate2: n,
			Code:     e,
			Answer:   q,
		}
		dao.Mgr.Save(&countcode)
	}
	Count := dao.Mgr.GetAllCount()
	c.HTML(200, "level1.html", Count)
}

func Golevel1Answer(c *gin.Context) {

	/*s := make([]string, 1000)
	s[0] = c.PostForm("level1answer1")
	s[1] = c.PostForm("level1answer2")
	s[2] = c.PostForm("level1answer3")
	s[3] = c.PostForm("level1answer4")
	s[4] = c.PostForm("level1answer5")
	s[5] = c.PostForm("level1answer6")
	s[6] = c.PostForm("level1answer7")
	s[7] = c.PostForm("level1answer8")
	s[8] = c.PostForm("level1answer9")
	s[9] = c.PostForm("level1answer10")
	Count1 := dao.Mgr.GetAllCount()

	var Score int = 0
	for i, v := range Count1 {
		p, _ := strconv.Atoi(s[i])
		if p == v.Answer {
			fmt.Printf("s: %v\n", s)
			fmt.Printf("v.Answer: %v\n", v.Answer)
			Score += 10
			fmt.Printf("Score: %v\n", Score)
		}
	}*/

	s0 := c.PostForm("level1answer1")
	s1 := c.PostForm("level1answer2")
	s2 := c.PostForm("level1answer3")
	s3 := c.PostForm("level1answer4")
	s4 := c.PostForm("level1answer5")
	s5 := c.PostForm("level1answer6")
	s6 := c.PostForm("level1answer7")
	s7 := c.PostForm("level1answer8")
	s8 := c.PostForm("level1answer9")
	s9 := c.PostForm("level1answer10")
	if s1 == "" {
		s1 = "未输入"
	}
	if s2 == "" {
		s2 = "未输入"
	}
	if s3 == "" {
		s3 = "未输入"
	}
	if s4 == "" {
		s4 = "未输入"
	}
	if s0 == "" {
		s0 = "未输入"
	}
	if s5 == "" {
		s5 = "未输入"
	}
	if s6 == "" {
		s6 = "未输入"
	}
	if s7 == "" {
		s7 = "未输入"
	}
	if s8 == "" {
		s8 = "未输入"
	}
	if s9 == "" {
		s9 = "未输入"
	}
	m := dao.Mgr.Panduan(s0, s1, s2, s3, s4, s5, s6, s7, s8, s9)

	//p := make([]string, 10)

	count := make([]model.Countcode, 100)
	count = dao.Mgr.GetAllCount()
	c.ShouldBindJSON(&count)
	c.HTML(200, "Level1Score.html", gin.H{
		"Score1":         m,
		"Level1answer1":  s0,
		"Level1answer2":  s1,
		"Level1answer3":  s2,
		"Level1answer4":  s3,
		"Level1answer5":  s4,
		"Level1answer6":  s5,
		"Level1answer7":  s6,
		"Level1answer8":  s7,
		"Level1answer9":  s8,
		"Level1answer10": s9,
		"Count1":         gin.H{"operate1": count[0].Operate1, "code": count[0].Code, "operate2": count[0].Operate2, "answer": count[0].Answer},
		"Count2":         gin.H{"operate1": count[1].Operate1, "code": count[1].Code, "operate2": count[1].Operate2, "answer": count[1].Answer},
		"Count3":         gin.H{"operate1": count[2].Operate1, "code": count[2].Code, "operate2": count[2].Operate2, "answer": count[2].Answer},
		"Count4":         gin.H{"operate1": count[3].Operate1, "code": count[3].Code, "operate2": count[3].Operate2, "answer": count[3].Answer},
		"Count5":         gin.H{"operate1": count[4].Operate1, "code": count[4].Code, "operate2": count[4].Operate2, "answer": count[4].Answer},
		"Count6":         gin.H{"operate1": count[5].Operate1, "code": count[5].Code, "operate2": count[5].Operate2, "answer": count[5].Answer},
		"Count7":         gin.H{"operate1": count[6].Operate1, "code": count[6].Code, "operate2": count[6].Operate2, "answer": count[6].Answer},
		"Count8":         gin.H{"operate1": count[7].Operate1, "code": count[7].Code, "operate2": count[7].Operate2, "answer": count[7].Answer},
		"Count9":         gin.H{"operate1": count[8].Operate1, "code": count[8].Code, "operate2": count[8].Operate2, "answer": count[8].Answer},
		"Count10":        gin.H{"operate1": count[9].Operate1, "code": count[9].Code, "operate2": count[9].Operate2, "answer": count[9].Answer},
	})
	dao.Mgr.Deletecount()
}

func Golevel2(c *gin.Context) {
	var countcode1 model.Countcode1
	for i := 0; i < 10; i++ {
		s1 := make([]int, 1000)
		for i := 0; i < 2; i++ { // 0 +   1-   2*
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Microsecond * 200)
			s1[i] = rand.Intn(3)
		}
		s2 := make([]int, 1000)
		for i := 0; i < 3; i++ {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Microsecond * 200)
			s2[i] = rand.Intn(50)
		}
		stack1 := make([]int, 1000)
		stack2 := make([]int, 1000)
		top1 := 0
		top2 := 0
		stack1[0] = s2[0]
		top1++
		for i := 1; i < 3; i++ {
			stack1[top1] = s2[i]
			stack2[top2] = s1[i-1]
			if stack2[top2] == 2 {
				stack1[top1-1] = stack1[top1] * stack1[top1-1]
				top1--
				top2--
			}
			top1++
			top2++
		}
		top1--
		top2--
		/*for top2 >= 0 {
			if stack2[top2] == 0 {
				stack1[top1-1] = stack1[top1] + stack1[top1-1]
			}
			if stack2[top2] == 1 {
				stack1[top1-1] = stack1[top1-1] - stack1[top1]
			}
			top2--
			top1--
		}
		var count int = stack1[0]*/
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
		} else {
			q = "*"
		}
		var p string
		if s1[1] == 0 {
			p = "+"
		} else if s1[1] == 1 {
			p = "-"
		} else {
			p = "*"
		}
		countcode1 = model.Countcode1{
			Operate1: s2[0],
			Operate2: s2[1],
			Operate3: s2[2],
			Code1:    q,
			Code2:    p,
			Answer:   count,
		}
		dao.Mgr.Save1(&countcode1)
	}
	Count := dao.Mgr.GetAllCount1()
	c.HTML(200, "level2.html", Count)
}

func Golevel2Answer(c *gin.Context) {

	s0 := c.PostForm("level2answer1")
	s1 := c.PostForm("level2answer2")
	s2 := c.PostForm("level2answer3")
	s3 := c.PostForm("level2answer4")
	s4 := c.PostForm("level2answer5")
	s5 := c.PostForm("level2answer6")
	s6 := c.PostForm("level2answer7")
	s7 := c.PostForm("level2answer8")
	s8 := c.PostForm("level2answer9")
	s9 := c.PostForm("level2answer10")
	if s1 == "" {
		s1 = "未输入"
	}
	if s2 == "" {
		s2 = "未输入"
	}
	if s3 == "" {
		s3 = "未输入"
	}
	if s4 == "" {
		s4 = "未输入"
	}
	if s0 == "" {
		s0 = "未输入"
	}
	if s5 == "" {
		s5 = "未输入"
	}
	if s6 == "" {
		s6 = "未输入"
	}
	if s7 == "" {
		s7 = "未输入"
	}
	if s8 == "" {
		s8 = "未输入"
	}
	if s9 == "" {
		s9 = "未输入"
	}
	m := dao.Mgr.Panduan1(s0, s1, s2, s3, s4, s5, s6, s7, s8, s9)

	time.Sleep(time.Nanosecond * 100)
	count := make([]model.Countcode1, 100)
	count = dao.Mgr.GetAllCount1()
	c.ShouldBindJSON(&count)

	dao.Mgr.Deletecount1()
	c.HTML(200, "Level2Score.html", gin.H{
		"Score1":         m,
		"Level1answer1":  s0,
		"Level1answer2":  s1,
		"Level1answer3":  s2,
		"Level1answer4":  s3,
		"Level1answer5":  s4,
		"Level1answer6":  s5,
		"Level1answer7":  s6,
		"Level1answer8":  s7,
		"Level1answer9":  s8,
		"Level1answer10": s9,
		"Count1":         gin.H{"operate1": count[0].Operate1, "code1": count[0].Code1, "operate2": count[0].Operate2, "code2": count[0].Code2, "operate3": count[0].Operate3, "answer": count[0].Answer},
		"Count2":         gin.H{"operate1": count[1].Operate1, "code1": count[1].Code1, "operate2": count[1].Operate2, "code2": count[1].Code2, "operate3": count[1].Operate3, "answer": count[1].Answer},
		"Count3":         gin.H{"operate1": count[2].Operate1, "code1": count[2].Code1, "operate2": count[2].Operate2, "code2": count[2].Code2, "operate3": count[2].Operate3, "answer": count[2].Answer},
		"Count4":         gin.H{"operate1": count[3].Operate1, "code1": count[3].Code1, "operate2": count[3].Operate2, "code2": count[3].Code2, "operate3": count[3].Operate3, "answer": count[3].Answer},
		"Count5":         gin.H{"operate1": count[4].Operate1, "code1": count[4].Code1, "operate2": count[4].Operate2, "code2": count[4].Code2, "operate3": count[4].Operate3, "answer": count[4].Answer},
		"Count6":         gin.H{"operate1": count[5].Operate1, "code1": count[5].Code1, "operate2": count[5].Operate2, "code2": count[5].Code2, "operate3": count[5].Operate3, "answer": count[5].Answer},
		"Count7":         gin.H{"operate1": count[6].Operate1, "code1": count[6].Code1, "operate2": count[6].Operate2, "code2": count[6].Code2, "operate3": count[6].Operate3, "answer": count[6].Answer},
		"Count8":         gin.H{"operate1": count[7].Operate1, "code1": count[7].Code1, "operate2": count[7].Operate2, "code2": count[7].Code2, "operate3": count[7].Operate3, "answer": count[7].Answer},
		"Count9":         gin.H{"operate1": count[8].Operate1, "code1": count[8].Code1, "operate2": count[8].Operate2, "code2": count[8].Code2, "operate3": count[8].Operate3, "answer": count[8].Answer},
		"Count10":        gin.H{"operate1": count[9].Operate1, "code1": count[9].Code1, "operate2": count[9].Operate2, "code2": count[9].Code2, "operate3": count[9].Operate3, "answer": count[9].Answer},
	})
}

func Golevel3(c *gin.Context) {
	dao.Mgr.Level3count()
	/*var countcode2 model.Countcode2
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
				continue
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
			}
			if stack2[top2] == 3 {
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
		dao.Mgr.Save2(&countcode2)
	}*/

	Count := dao.Mgr.GetAllCount2()
	c.HTML(200, "level3.html", Count)

}

func Golevel3Answer(c *gin.Context) {
	s0 := c.PostForm("level3answer1")
	s1 := c.PostForm("level3answer2")
	s2 := c.PostForm("level3answer3")
	s3 := c.PostForm("level3answer4")
	s4 := c.PostForm("level3answer5")
	s5 := c.PostForm("level3answer6")
	s6 := c.PostForm("level3answer7")
	s7 := c.PostForm("level3answer8")
	s8 := c.PostForm("level3answer9")
	s9 := c.PostForm("level3answer10")
	if s1 == "" {
		s1 = "未输入"
	}
	if s2 == "" {
		s2 = "未输入"
	}
	if s3 == "" {
		s3 = "未输入"
	}
	if s4 == "" {
		s4 = "未输入"
	}
	if s0 == "" {
		s0 = "未输入"
	}
	if s5 == "" {
		s5 = "未输入"
	}
	if s6 == "" {
		s6 = "未输入"
	}
	if s7 == "" {
		s7 = "未输入"
	}
	if s8 == "" {
		s8 = "未输入"
	}
	if s9 == "" {
		s9 = "未输入"
	}
	m := dao.Mgr.Panduan2(s0, s1, s2, s3, s4, s5, s6, s7, s8, s9)
	time.Sleep(time.Nanosecond * 100)

	count := make([]model.Countcode2, 100)
	count = dao.Mgr.GetAllCount2()
	c.ShouldBindJSON(&count)

	c.HTML(200, "Level3Score.html", gin.H{
		"Score1":         m,
		"Level1answer1":  s0,
		"Level1answer2":  s1,
		"Level1answer3":  s2,
		"Level1answer4":  s3,
		"Level1answer5":  s4,
		"Level1answer6":  s5,
		"Level1answer7":  s6,
		"Level1answer8":  s7,
		"Level1answer9":  s8,
		"Level1answer10": s9,
		"Count1":         gin.H{"operate1": count[0].Operate1, "code1": count[0].Code1, "operate2": count[0].Operate2, "code2": count[0].Code2, "operate3": count[0].Operate3, "code3": count[0].Code3, "operate4": count[0].Operate4, "answer": count[0].Answer},
		"Count2":         gin.H{"operate1": count[1].Operate1, "code1": count[1].Code1, "operate2": count[1].Operate2, "code2": count[1].Code2, "operate3": count[1].Operate3, "code3": count[1].Code3, "operate4": count[1].Operate4, "answer": count[1].Answer},
		"Count3":         gin.H{"operate1": count[2].Operate1, "code1": count[2].Code1, "operate2": count[2].Operate2, "code2": count[2].Code2, "operate3": count[2].Operate3, "code3": count[2].Code3, "operate4": count[2].Operate4, "answer": count[2].Answer},
		"Count4":         gin.H{"operate1": count[3].Operate1, "code1": count[3].Code1, "operate2": count[3].Operate2, "code2": count[3].Code2, "operate3": count[3].Operate3, "code3": count[3].Code3, "operate4": count[3].Operate4, "answer": count[3].Answer},
		"Count5":         gin.H{"operate1": count[4].Operate1, "code1": count[4].Code1, "operate2": count[4].Operate2, "code2": count[4].Code2, "operate3": count[4].Operate3, "code3": count[4].Code3, "operate4": count[4].Operate4, "answer": count[4].Answer},
		"Count6":         gin.H{"operate1": count[5].Operate1, "code1": count[5].Code1, "operate2": count[5].Operate2, "code2": count[5].Code2, "operate3": count[5].Operate3, "code3": count[5].Code3, "operate4": count[5].Operate4, "answer": count[5].Answer},
		"Count7":         gin.H{"operate1": count[6].Operate1, "code1": count[6].Code1, "operate2": count[6].Operate2, "code2": count[6].Code2, "operate3": count[6].Operate3, "code3": count[6].Code3, "operate4": count[6].Operate4, "answer": count[6].Answer},
		"Count8":         gin.H{"operate1": count[7].Operate1, "code1": count[7].Code1, "operate2": count[7].Operate2, "code2": count[7].Code2, "operate3": count[7].Operate3, "code3": count[7].Code3, "operate4": count[7].Operate4, "answer": count[7].Answer},
		"Count9":         gin.H{"operate1": count[8].Operate1, "code1": count[8].Code1, "operate2": count[8].Operate2, "code2": count[8].Code2, "operate3": count[8].Operate3, "code3": count[8].Code3, "operate4": count[8].Operate4, "answer": count[8].Answer},
		"Count10":        gin.H{"operate1": count[9].Operate1, "code1": count[9].Code1, "operate2": count[9].Operate2, "code2": count[9].Code2, "operate3": count[9].Operate3, "code3": count[9].Code3, "operate4": count[9].Operate4, "answer": count[9].Answer},
	})
	dao.Mgr.Deletecount2()
}
