package utils

import "fmt"

type Student struct {
	Username string
	Sex      int
	Grade    string
	Score    float32
}

type StudentMgr struct {
	AllStudent []*Student
}

func NewStudent(username string, sex int, grade string, score float32) (stu *Student) {
	return &Student{
		Username: username,
		Sex:      sex,
		Grade:    grade,
		Score:    score,
	}
}

func (mgr *StudentMgr) AddStudent(stu *Student) (err error) {
	for i, s := range mgr.AllStudent {
		if s.Username == stu.Username {
			mgr.AllStudent[i] = stu
			fmt.Printf("Student %s Updated!\n", stu.Username)
			return
		}
	}
	mgr.AllStudent = append(mgr.AllStudent, stu)
	fmt.Printf("Student %s Added!\n", stu.Username)
	return
}

func (mgr *StudentMgr) ModifyStudent(stu *Student) (err error) {
	found := false
	for i, s := range mgr.AllStudent {
		if s.Username == stu.Username {
			found = true
			mgr.AllStudent[i] = stu
			fmt.Printf("Student %s Updated!\n", stu.Username)
			return
		}
	}
	if !found {
		fmt.Printf("Student %s not found!\n", stu.Username)
		err = fmt.Errorf("Student %s not found!\n", stu.Username)
		return
	}
	return
}

func (mgr *StudentMgr) ShowAllStudent() {
	for _, s := range mgr.AllStudent {
		fmt.Printf("Student %s's info: %#v\n", s.Username, s)
	}
}
