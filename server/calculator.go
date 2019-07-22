

package server

import (
	"strconv"
	"strings"
)

func Calculate(s string)  float64 {
	if s[len(s)-1]=='\n'{
		s=s[:len(s)-1]
	}
	var result float64
	if strings.Contains(s,"+"){
		nums:=strings.Split(s,"+")
		firstNum,_:=strconv.ParseFloat( nums[0],64)
		secpndNum,_:= strconv.ParseFloat(nums[1],64)
		result=firstNum+secpndNum
	} else if strings.Contains(s,"-"){
		nums:=strings.Split(s,"-")
		firstNum,_:=strconv.ParseFloat( nums[0],64)
		secpndNum,_:= strconv.ParseFloat(nums[1],64)
		result=firstNum-secpndNum
	}else if strings.Contains(s,"*") {

		nums := strings.Split(s, "*")
		firstNum, _ := strconv.ParseFloat(nums[0], 64)
		secpndNum, _ := strconv.ParseFloat(nums[1], 64)
		result = firstNum * secpndNum
	} else {
		nums:=strings.Split(s,"/")
		firstNum,_:=strconv.ParseFloat( nums[0],64)
		secpndNum,_:= strconv.ParseFloat(nums[1],64)
		result=firstNum/secpndNum
	}

	return result
}

