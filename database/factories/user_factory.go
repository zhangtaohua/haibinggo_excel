// Package factories 存放工厂方法
package factories

import (
	"haibinggo/app/models/user"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	model1 := user.User{
		Name:     "yanghaibing",
		NickName: "杨海冰",
		Email:    "yanghaibing@163.com",
		Phone:    "15986670126",
		Password: "SeCrEt123456",
	}

	objs = append(objs, model1)

	model1 = user.User{
		Name:     "admin",
		NickName: "管理员",
		Email:    "admin@163.com",
		Phone:    "19954347299",
		Password: "SeCrEt123456",
	}

	objs = append(objs, model1)

	model1 = user.User{
		Name:     "finance",
		NickName: "财务",
		Email:    "finance@163.com",
		Phone:    "13077912310",
		Password: "SeCrEt",
	}

	objs = append(objs, model1)

	// for i := 0; i < times; i++ {
	// 	model := user.User{
	// 		Name:     "user_" + cast.ToString(i+4),
	// 		NickName: cast.ToString(i+4) + "号", // faker.Name()
	// 		Email:    faker.Email(),
	// 		Phone:    helpers.RandomNumber(11),
	// 		Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
	// 	}
	// 	objs = append(objs, model)
	// }

	return objs
}
