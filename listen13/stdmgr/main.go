package main

import (
	"fmt"
	"os"
	"stdmgr/utils"
)

func enterStudentInfo() (stu *utils.Student) {
	var (
		username string
		score    float32
		grade    string
		sex      int
	)
	fmt.Println("Please enter username: ")
	fmt.Scanf("%s\n", &username)

	fmt.Println("Please enter sex[0|1]: ")
	fmt.Scanf("%d\n", &sex)

	fmt.Println("Please enter grade[1-6]: ")
	fmt.Scanf("%s\n", &grade)

	fmt.Println("Please enter score: ")
	fmt.Scanf("%f\n", &score)

	stu = utils.NewStudent(username, sex, grade, score)
	return stu
}

func ShowMenu() {
	fmt.Println("1. Add Student")
	fmt.Println("2. Modify Student")
	fmt.Println("3. Show All Student")
	fmt.Println("4. Exit")
	fmt.Print("Please enter the corresponding serial number: ")
}

func main() {
	mgr := utils.StudentMgr{}
	for {
		ShowMenu()
		var sel int
		fmt.Scanf("%d", &sel)
		switch sel {
		case 1:
			std := enterStudentInfo()
			if err := mgr.AddStudent(std); err != nil {
				fmt.Println(err)
			}
		case 2:
			std := enterStudentInfo()
			if err := mgr.ModifyStudent(std); err != nil {
				fmt.Println(err)
			}
		case 3:
			mgr.ShowAllStudent()
		case 4:
			os.Exit(0)
		}
	}
}
