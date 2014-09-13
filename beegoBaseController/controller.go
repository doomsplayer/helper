package beegoBaseController

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type more []string

func (m *more) Add(s string) string {
	*m = append(*m, s)
	return ``
}

type Base struct {
	beego.Controller
}

type ErrMap map[error]string

func InfoPrepend(s string) func(err error) string {
	return func(err error) string { return fmt.Sprint(s, err) }
}
func (this *Base) CheckHtml(condition interface{}, code int, i interface{}, jmp ...string) {

	switch condition.(type) {
	case bool:
		{
			if !condition.(bool) {
				this.Ctx.Output.Status = code
				switch i.(type) {
				case func() string:
					{
						this.Data["error"] = i.(func() string)()
					}
				case error:
					{
						this.Data["error"] = i.(error).Error()
					}
				case string:
					{
						this.Data["error"] = i.(string)
					}
				case func():
					{
						i.(func())()
						this.Data["error"] = "equation condition not meet"
					}
				default:
					{
						this.Data["error"] = "equation condition not meet"
					}
				}
				if len(jmp) != 0 {
					this.Data[`jump`] = jmp[0]
				}
				this.TplNames = "error.html"
				this.Render()
				this.StopRun()

			}

		}
	case error:
		{
			err := condition.(error)
			if err != nil {
				this.Ctx.Output.Status = code
				switch i.(type) {
				case func() string:
					{
						this.Data["error"] = i.(func() string)()
					}
				case func(error) string:
					{
						this.Data["error"] = i.(func(error) string)(err)
					}
				case func(error):
					{
						i.(func(error))(err)
						this.Data["error"] = err.Error()
					}
				case func():
					{
						i.(func())()
						this.Data["error"] = err.Error()
					}
				case error:
					{
						this.Data["error"] = i.(error).Error()
					}
				case string:
					{
						this.Data["error"] = i.(string)
					}
				case map[error]string:
					{
						m := i.(map[error]string)
						es, ok := m[err]
						if ok {
							this.Data[`error`] = es
						} else {
							this.Data[`error`] = err.Error()
						}

					}
				case ErrMap:
					{
						m := i.(ErrMap)
						es, ok := m[err]
						if ok {
							this.Data[`error`] = es
						} else {
							this.Data[`error`] = err.Error()
						}
					}
				default:
					{
						this.Data["error"] = err.Error()
					}
				}
				if len(jmp) != 0 {
					this.Data[`jump`] = jmp[0]
				}
				this.TplNames = "error.html"
				this.Render()
				this.StopRun()

			}
		}
	default:
		{

			if condition != nil {
				this.Ctx.Output.Status = code
				panic(fmt.Errorf("controller.Check: input is not bool or error"))
			}

		}
	}

}
func (this *Base) OkHtml(i interface{}, jmp ...string) {
	this.Data[`info`] = i
	if len(jmp) != 0 {
		this.Data[`jump`] = jmp[0]
	}
	this.TplNames = "error.html"
	this.Render()
}

func (this *Base) CheckJson(condition interface{}, code int, i interface{}) {

	switch condition.(type) {
	case bool:
		{
			if !condition.(bool) {
				this.Ctx.Output.Status = code

				switch i.(type) {
				case func() string:
					{
						this.Data["json"] = i.(func() string)()
					}
				case error:
					{
						this.Data["json"] = i.(error).Error()
					}
				case string:
					{
						this.Data["json"] = i.(string)
					}
				case func():
					{
						i.(func())()
						this.Data["json"] = "equation condition not meet"
					}
				default:
					{
						this.Data["json"] = "equation condition not meet"
					}
				}
				this.ServeJson()
				this.StopRun()

			}

		}
	case error:
		{
			err := condition.(error)
			if err != nil {
				this.Ctx.Output.Status = code
				switch i.(type) {
				case func() string:
					{
						this.Data["json"] = i.(func() string)()
					}
				case func(error) string:
					{
						this.Data["json"] = i.(func(error) string)(err)
					}
				case func(error):
					{
						i.(func(error))(err)
						this.Data["json"] = err.Error()
					}
				case func():
					{
						i.(func())()
						this.Data["json"] = err.Error()
					}
				case error:
					{
						this.Data["json"] = i.(error).Error()
					}
				case string:
					{
						this.Data["json"] = i.(string)
					}
				case map[error]string:
					{
						m := i.(map[error]string)
						es, ok := m[err]
						if ok {
							this.Data[`json`] = es
						} else {
							this.Data[`json`] = err.Error()
						}

					}
				case ErrMap:
					{
						m := i.(ErrMap)
						es, ok := m[err]
						if ok {
							this.Data[`json`] = es
						} else {
							this.Data[`json`] = err.Error()
						}
					}
				default:
					{
						this.Data["json"] = err.Error()
					}
				}
				this.ServeJson()
				this.StopRun()

			}
		}
	default:
		{
			if condition != nil {
				this.Ctx.Output.Status = code
				panic(fmt.Errorf("controller.Check: input is not bool or error"))
			}

		}
	}

}
func (this *Base) OkJson(i interface{}) {
	this.Data[`json`] = i
	this.ServeJson()
}
func (this *Base) CheckFlash(condition interface{}, i interface{}, to string) {
	flash := beego.NewFlash()

	switch condition.(type) {
	case bool:
		{
			if !condition.(bool) {

				switch i.(type) {
				case func() string:
					{
						flash.Error(i.(func() string)())
					}
				case error:
					{
						flash.Error(i.(error).Error())
					}
				case string:
					{
						flash.Error(i.(string))
					}
				case func():
					{
						i.(func())()
						flash.Error("equation condition not meet")
					}
				default:
					{
						flash.Error("equation condition not meet")
					}
				}
				flash.Store(&this.Controller)
				this.Redirect(to, 302)
				this.StopRun()

			}

		}
	case error:
		{
			err := condition.(error)

			if err != nil {

				switch i.(type) {
				case func() string:
					{
						flash.Error(i.(func() string)())
					}
				case func(error) string:
					{
						flash.Error(i.(func(error) string)(err))
					}
				case func():
					{
						i.(func())()
						flash.Error(err.Error())
					}
				case func(error):
					{
						i.(func(error))(err)
						flash.Error(err.Error())
					}
				case error:
					{
						flash.Error(i.(error).Error())
					}
				case string:
					{
						flash.Error(i.(string))
					}
				case map[error]string:
					{
						m := i.(map[error]string)
						es, ok := m[err]
						if ok {
							flash.Error(es)
						} else {
							flash.Error(err.Error())
						}

					}
				case ErrMap:
					{
						m := i.(ErrMap)
						es, ok := m[err]
						if ok {
							flash.Error(es)
						} else {
							flash.Error(err.Error())
						}
					}
				default:
					{
						flash.Error(err.Error())
					}
				}
				flash.Store(&this.Controller)
				this.Redirect(to, 302)
				this.StopRun()

			}
		}
	default:
		{
			if condition != nil {
				panic(fmt.Errorf("controller.Check: input is not bool or error"))
			}

		}
	}

}
func (this *Base) OkFlash(i interface{}, to string) {
	flash := beego.NewFlash()
	flash.Notice(fmt.Sprint(i))
	flash.Store(&this.Controller)
	this.Redirect(to, 302)
}

func (this *Base) ParseFormAndValidCheckJson(form interface{}, f interface{}) {
	err := this.ParseFormAndValid(form)
	this.CheckJson(err, 400, f)
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

func (this *Base) NewPaginator(per, nums int) *Paginator {
	paginator := NewPaginator(this.Ctx.Input.Request, per, nums)
	this.Data[`paginator`] = paginator
	return paginator
}
func (b *Base) Prepare() {
	b.Data[`moreStyles`] = &more{}
	b.Data[`beforeScripts`] = &more{}
	b.Data[`laterScripts`] = &more{}
	b.Data["position"] = ``
	b.Data[`title`] = ""
	beego.ReadFromRequest(&b.Controller)

}
