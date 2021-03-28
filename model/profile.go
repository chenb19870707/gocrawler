package model

import "strconv"

type Profile struct {
	Name string				//名字
	Age int					//年龄
	Height int				//身高
	Weight int				//体重
	Constellation string	//星座
	Salary string			//薪水
	Marry string			//婚姻状况
}

func (p Profile) String() string {
	return p.Name + " " + p.Marry + " " + strconv.Itoa(p.Age) + " odls " + strconv.Itoa(p.Height) + " cm " + strconv.Itoa(p.Weight) + " kg "
}