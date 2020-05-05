package errorUtils

import "log"

func HandleErr(err error)  {
	if err != nil{
		log.Fatal("错误:",err)
	}
}
