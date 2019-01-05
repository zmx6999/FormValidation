package FormValidation

import (
	"reflect"
	"errors"
	"strings"
	"regexp"
	"time"
	"strconv"
)

type FieldValidation struct {
	FieldName string
	FieldValue interface{}
	ValidMethodName string
	ValidMethodArgs []interface{}
	ErrMsg string
	Trim bool
	ValidEmpty bool
}

func IsEmpty(x interface{},trim bool) bool {
	if x==nil {
		return true
	}
	str,isStr:=GetStr(x,trim)
	if isStr && str=="" {
		return true
	}
	return false
}

func GetStr(x interface{},trim bool) (string,bool) {
	t:=reflect.TypeOf(x).String()
	if t=="string" {
		str := x.(string)
		if trim {
			str=strings.Trim(str," ")
		}
		return str,true
	}
	return "",false
}

func ValidStartEndDate(startStr string,endStr string) bool {
	if startStr=="" || endStr=="" {
		return true
	}
	start,err:=time.Parse("2006-01-02 15:04:05",startStr+" 00:00:00")
	if err!=nil {
		return false
	}
	end,err:=time.Parse("2006-01-02 15:04:05",endStr+" 00:00:00")
	if err!=nil {
		return false
	}
	if end.Before(start) {
		return false
	}
	return true
}

func ValidStartEndTime(startStr string,endStr string) bool {
	if startStr=="" || endStr=="" {
		return true
	}
	start,err:=time.Parse("2006-01-02 15:04:05",startStr+":00")
	if err!=nil {
		return false
	}
	end,err:=time.Parse("2006-01-02 15:04:05",endStr+":00")
	if err!=nil {
		return false
	}
	if end.Before(start) {
		return false
	}
	return true
}

func (fv *FieldValidation) ValidateField() (bool,error) {
	_,found:=reflect.TypeOf(fv).MethodByName(fv.ValidMethodName)
	if !found {
		return false,errors.New("VALIDATION METHOD NOT FOUND")
	}
	method:=reflect.ValueOf(fv).MethodByName(fv.ValidMethodName)
	numIn:=method.Type().NumIn()
	in:=make([]reflect.Value,numIn)
	for i:=0; i<numIn; i++ {
		in[i]=reflect.ValueOf(fv.ValidMethodArgs[i])
	}
	r:=method.Call(in)[0].Bool()
	if r {
		return true,nil
	} else {
		return false,errors.New(fv.ErrMsg)
	}
}

func (fv *FieldValidation) Require() bool {
	return !IsEmpty(fv.FieldValue,fv.Trim)
}

func (fv *FieldValidation) Length(minLength int,maxLength int) bool {
	str,isStr:=GetStr(fv.FieldValue,fv.Trim)
	if isStr {
		n:=len([]byte(str))
		if n<minLength {
			return false
		}
		if maxLength>0 && n>maxLength {
			return false
		}
		return true
	}
	return false
}

func (fv *FieldValidation) Format(format string) bool {
	str,isStr:=GetStr(fv.FieldValue,fv.Trim)
	if isStr {
		re:=regexp.MustCompile(format)
		return re.MatchString(str)
	}
	return false
}

func (fv *FieldValidation) Int() bool {
	return fv.Format(`^(\+|-)?\d+$`)
}

func (fv *FieldValidation) Float() bool {
	return fv.Format(`^[-+]?\d*\.?\d*$`)
}

func (fv *FieldValidation) Email() bool {
	return fv.Format(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
}

func (fv *FieldValidation) URL() bool {
	return fv.Format(`^https?:\/\/(([a-zA-Z0-9_-])+(\.)?)*(:\d+)?(\/((\.)?(\?)?=?&?[a-zA-Z0-9_-](\?)?)*)*$`)
}

func (fv *FieldValidation) ChineseMobile() bool {
	return fv.Format(`^(0|86|17951)?(13[0-9]|15[0-9]|166|17[0-9]|18[0-9]|14[57])[0-9]{8}$`)
}

func (fv *FieldValidation) ChineseIdCard() bool {
	return fv.Format(`^(^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$)|(^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])((\d{4})|\d{3}[Xx])$)$`)
}

func (fv *FieldValidation) Date() bool {
	return fv.Format(`^[1-2][0-9][0-9][0-9]-[0-1]{0,1}[0-9]-[0-3]{0,1}[0-9]$`)
}

func (fv *FieldValidation) StartDate(endDate interface{}) bool {
	startStr,startIsStr:=GetStr(fv.FieldValue,fv.Trim)
	if !startIsStr {
		return false
	}
	endStr,endIsStr:=GetStr(endDate,true)
	if !endIsStr {
		return false
	}
	return ValidStartEndDate(startStr,endStr)
}

func (fv *FieldValidation) EndDate(startDate interface{}) bool {
	startStr,startIsStr:=GetStr(startDate,true)
	if !startIsStr {
		return false
	}
	endStr,endIsStr:=GetStr(fv.FieldValue,fv.Trim)
	if !endIsStr {
		return false
	}
	return ValidStartEndDate(startStr,endStr)
}

func (fv *FieldValidation) StartTime(endTime interface{}) bool {
	startStr,startIsStr:=GetStr(fv.FieldValue,fv.Trim)
	if !startIsStr {
		return false
	}
	endStr,endIsStr:=GetStr(endTime,true)
	if !endIsStr {
		return false
	}
	return ValidStartEndTime(startStr,endStr)
}

func (fv *FieldValidation) EndTime(startTime interface{}) bool {
	startStr,startIsStr:=GetStr(startTime,true)
	if !startIsStr {
		return false
	}
	endStr,endIsStr:=GetStr(fv.FieldValue,fv.Trim)
	if !endIsStr {
		return false
	}
	return ValidStartEndTime(startStr,endStr)
}

func (fv *FieldValidation) BeforeToday() bool {
	endDate:=time.Now().Format("2006-01-02")
	return fv.StartDate(endDate)
}

func (fv *FieldValidation) AfterToday() bool {
	startDate:=time.Now().Format("2006-01-02")
	return fv.EndDate(startDate)
}

func (fv *FieldValidation) BeforeNow() bool {
	endTime:=time.Now().Format("2006-01-02 15:04")
	return fv.StartTime(endTime)
}

func (fv *FieldValidation) AfterNow() bool {
	startTime:=time.Now().Format("2006-01-02 15:04")
	return fv.EndTime(startTime)
}

func (fv *FieldValidation) Unsigned() bool {
	if fv.Int() || fv.Float() {
		str,isStr:=GetStr(fv.FieldValue,fv.Trim)
		if isStr {
			n,err:=strconv.ParseFloat(str,64)
			if err==nil && n>=0 {
				return true
			}
		}
	}
	return false
}

func (fv *FieldValidation) Positive() bool {
	if fv.Int() || fv.Float() {
		str,isStr:=GetStr(fv.FieldValue,fv.Trim)
		if isStr {
			n,err:=strconv.ParseFloat(str,64)
			if err==nil && n>0 {
				return true
			}
		}
	}
	return false
}
