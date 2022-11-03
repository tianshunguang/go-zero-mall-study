package random

import (
	"fmt"
	"github.com/o1egl/govatar"
)

func GenAvatar(name, gender string) string {

	filename := "avatar_" + name + ".jpg"
	avatarSavePath := "D:\\workspace\\src\\test\\picture\\" + filename
	fmt.Println(avatarSavePath)
	//avatarGetUrl:="192.168.1.1/"+filename
	if gender == string(1) {
		err := govatar.GenerateFileForUsername(govatar.MALE, name, avatarSavePath)
		if err != nil {
			fmt.Println(err)
			return ""
		} else {
			err := govatar.GenerateFileForUsername(govatar.FEMALE, name, avatarSavePath)
			if err != nil {
				fmt.Println(err)
				return ""
			}
		}
	}
	return avatarSavePath
}

