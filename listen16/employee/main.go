package main

import "fmt"

type Employer interface {
	CalcSalary() float64
	Name() string
}

type Programmer struct {
	name string
	base float64
	bonus float64
}

func NewProgrammer(name string, base, bonus float64) Programmer {
	return Programmer{
		name: name,
		base: base,
		bonus: bonus,
	}
}

func (p Programmer) CalcSalary() float64 {
	return p.base
}

func (p Programmer) Name() string {
	return p.name
}

type Seller struct {
	name string
	base float64
	bonus float64 // 提成系数
	sales float64  // 销售额
}

func NewSeller(name string, base, bonus, sales float64) Seller {
	return Seller{
		name: name,
		base: base,
		bonus: bonus,
		sales: sales,
	}
}

func (s Seller) CalcSalary() float64 {
	return s.base + s.sales * s.bonus
}

func (s Seller) Name() string {
	return s.name
}

func CalAll(emps []Employer) float64 {
	var cost float64

	for _, emp := range emps {
		fmt.Printf("%s这个月工资为: %.2f\n", emp.Name(), emp.CalcSalary())
		cost += emp.CalcSalary()
	}
	return cost
}

func main() {
	p1 := NewProgrammer("码农1", 13000.0, 0)
	p2 := NewProgrammer("码农2", 12500.0, 0)
	p3 := NewProgrammer("码农3", 15000.0, 0)

	s1 := NewSeller("销售1", 6000, 0.08, 78563.68)
	s2 := NewSeller("销售2", 5500, 0.06, 128974.5)
	s3 := NewSeller("销售3", 7000, 0.1, 125789.79)

	var employList []Employer
	employList = append(employList, p1)
	employList = append(employList, p2)
	employList = append(employList, p3)

	employList = append(employList, s1)
	employList = append(employList, s2)
	employList = append(employList, s3)

	cost := CalAll(employList)
	fmt.Printf("这个月人力总成本为%.2f", cost)
}