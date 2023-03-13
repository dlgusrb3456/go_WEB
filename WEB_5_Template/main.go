package main

import (
	"html/template" // text/template 패키지도 존재함. 해당 패키지 사용하면 template에서 사용하는 특수문자가 그냥 나옴
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool { //템플릿파일 내부에서 이 함수를 불러서 사용할 수 있음.
	return u.Age > 30
}

func main() {
	user1 := User{Name: "Lee", Email: "dlgusrb@naver.com", Age: 24}
	user2 := User{Name: "Lee2", Email: "dlgusrb2@naver.com", Age: 40}
	users := []User{user1, user2}

	//tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail:{{.Email}}\nAge: {{.Age}}\n")
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl") //파일 템플릿 사용
	if err != nil {
		panic(err)
	}

	// tmpl.Execute(os.Stdout, user1) //user의 data가 tmpl에 채워짐
	// tmpl.Execute(os.Stdout, user2) //직접 template 내부를 적은 경우에 Execute 사용
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user1) //template을 파일화 한 이후 사용시 ExecuteTemplate 사용 + 어떤 파일 사용할건지 명시 + data
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user1) // inner template으로 template 안에서 template을 사용함. 구분지어서 사용 가능함
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)

	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users) // 리스트를 data로 넘김. + template 파일에서 {{range .}} & {{end}}를 이용해 리스트를 순회함
}
