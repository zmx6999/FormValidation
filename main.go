package main

import (
	"FormValidation"
	"fmt"
)

func main()  {
	request:=map[string]interface{}{}
	request["name"]="  Mary    "
	request["password"]="  123456  "
	request["mobile"]="    "
	request["email"]="  "
	request["home_page"]="    https://www.google.com  "
	request["id_card"]="    "
	request["start_date"]="  2019-01-01  "
	request["end_date"]="  2019-01-05  "
	request["start_time"]="  2019-01-01 15:35  "
	request["end_time"]="  2019-02-05 15:35  "
	request["age"]="  28  "
	request["year_income"]="  123456.78  "

	fvs:=[]*FormValidation.FieldValidation{
		&FormValidation.FieldValidation{
			FieldName:"name",
			ValidMethodName:"Require",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"name cannot be empty",
			Trim:true,
			ValidEmpty:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"password",
			ValidMethodName:"Require",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"password cannot be empty",
			ValidEmpty:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"password",
			ValidMethodName:"Length",
			ValidMethodArgs:[]interface{}{6,16},
			ErrMsg:"password length 6-16",
		},
		&FormValidation.FieldValidation{
			FieldName:"mobile",
			ValidMethodName:"ChineseMobile",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"invalid mobile",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"email",
			ValidMethodName:"Email",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"invalid email",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"home_page",
			ValidMethodName:"URL",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"invalid url",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"id_card",
			ValidMethodName:"ChineseIdCard",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"invalid ID card",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"start_date",
			ValidMethodName:"Date",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"start date invalid",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"start_date",
			ValidMethodName:"BeforeToday",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"start date should be earlier than today",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"end_date",
			ValidMethodName:"Date",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"end date invalid",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"end_date",
			ValidMethodName:"AfterToday",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"end date should be later than today",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"end_date",
			ValidMethodName:"EndDate",
			ValidMethodArgs:[]interface{}{request["start_date"]},
			ErrMsg:"end date should be later than start date",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"start_time",
			ValidMethodName:"BeforeNow",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"start time should be earlier than now",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"end_time",
			ValidMethodName:"AfterNow",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"end time should be later than now",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"end_time",
			ValidMethodName:"EndTime",
			ValidMethodArgs:[]interface{}{request["start_time"]},
			ErrMsg:"end time should be later than start time",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"age",
			ValidMethodName:"Int",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"age should be an integer",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"age",
			ValidMethodName:"Unsigned",
			ValidMethodArgs:[]interface{}{request["start_time"]},
			ErrMsg:"age>=0",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"age",
			ValidMethodName:"Positive",
			ValidMethodArgs:[]interface{}{request["start_time"]},
			ErrMsg:"age>0",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"year_income",
			ValidMethodName:"Float",
			ValidMethodArgs:[]interface{}{},
			ErrMsg:"year_income should be a float",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"year_income",
			ValidMethodName:"Unsigned",
			ValidMethodArgs:[]interface{}{request["start_time"]},
			ErrMsg:"year_income>=0",
			Trim:true,
		},
		&FormValidation.FieldValidation{
			FieldName:"year_income",
			ValidMethodName:"Positive",
			ValidMethodArgs:[]interface{}{request["start_time"]},
			ErrMsg:"year_income>0",
			Trim:true,
		},
	}

	gv:=&FormValidation.GroupValidation{
		request,
		fvs,
	}
	v,err:=gv.Validate()
	fmt.Println(v)
	if err!=nil {
		fmt.Println(err)
	}
	v2,err2:=gv.PatchValidate()
	fmt.Println(v2)
	if len(err2)>0 {
		fmt.Println(err2)
	}
}
