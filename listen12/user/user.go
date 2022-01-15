package user

type User struct {
	Username string
	Sex      int
	Age      int
	AvataUrl string
}

// 构造函数
func NewUser(username, avataurl string, age, sex int) *User {
	return &User{
		Username: username,
		Sex:      sex,
		Age:      age,
		AvataUrl: avataurl,
	}
}
