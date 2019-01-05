package FormValidation

type GroupValidation struct {
	Form map[string]interface{}
	FieldValidations []*FieldValidation
}

func (gv *GroupValidation) Validate() (bool,error) {
	n:=len(gv.FieldValidations)
	for i:=0; i<n; i++ {
		fv:=gv.FieldValidations[i]
		fv.FieldValue=gv.Form[fv.FieldName]
		if IsEmpty(fv.FieldValue,fv.Trim) && !fv.ValidEmpty {
			continue
		}
		v,err:=fv.ValidateField()
		if !v {
			return false,err
		}
	}
	return true,nil
}

func (gv *GroupValidation) PatchValidate() (bool,map[string]error) {
	valid:=true
	errMap:=map[string]error{}
	n:=len(gv.FieldValidations)
	for i:=0; i<n; i++ {
		fv:=gv.FieldValidations[i]
		if errMap[fv.FieldName]!=nil {
			continue
		}
		fv.FieldValue=gv.Form[fv.FieldName]
		if IsEmpty(fv.FieldValue,fv.Trim) && !fv.ValidEmpty {
			continue
		}
		v,err:=fv.ValidateField()
		if !v {
			valid=false
			errMap[fv.FieldName]=err
		}
	}
	return valid,errMap
}
