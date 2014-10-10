package beegoBaseController

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strconv"
)

func (this *Base) ParseFormAndValidCheckJson(form interface{}, mapping interface{}) {
	err := this.ParseFormAndValid(form)
	this.CheckJson(err, 400, mapping)
}

func (this *Base) ParseFormAndValid(form interface{}) (err error) {
	err = this.ParseForm(form)
	if err != nil {
		return err
	}

	valid := validation.Validation{}
	b, err := valid.Valid(form)
	if err != nil {
		return err
	}
	if !b {
		errmsg := ``
		for _, err := range valid.Errors {
			errmsg += fmt.Sprint(err.Key, " ", err.Message, "\n")
		}
		return fmt.Errorf(errmsg)
	}
	return nil
}
func (this *Base) ParseJson(obj interface{}) (err error) {
	dec := json.NewDecoder(this.Ctx.Input.Request.Body)
	err = dec.Decode(obj)
	return
}
func (this *Base) GetParamInt(field string) int {
	param := this.Ctx.Input.Param(field)
	ret, _ := strconv.Atoi(param)
	return ret
}
func (this *Base) GetInts(field string) []int {
	ss := this.GetStrings(field)
	retv := []int{}
	for _, v := range ss {
		r, _ := strconv.Atoi(v)
		retv = append(retv, r)
	}
	return retv
}
