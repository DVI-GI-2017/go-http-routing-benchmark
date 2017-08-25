package validator

import (
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	emailPatern       = regexp.MustCompile(".+@.+\\..+")
	dateiso8601Patern = regexp.MustCompile("^(\\d{4})-(\\d{2})-(\\d{2})T(\\d{2}):(\\d{2}):(\\d{2})(Z|(\\+|-)\\d{2}(:?\\d{2})?)$")
)

type ValidateError struct {
	Name string
	Err  error
}

type Validator struct {
	errs []ValidateError
}

func New() *Validator {
	return &Validator{
		errs: []ValidateError{},
	}
}

func (v *Validator) HasError() bool {
	if len(v.errs) > 0 {
		return true
	}

	return false
}

func (v *Validator) Messages() map[string]string {
	msgs := make(map[string]string)
	for i := 0; i < len(v.errs); i++ {
		if v.errs[i].Name != "" {
			msgs[v.errs[i].Name] = v.errs[i].Err.Error()
		}
	}

	return msgs
}

func (v *Validator) GetError() error {
	if len(v.errs) > 0 {
		return v.errs[0].Err
	}

	return nil
}

func (v *Validator) addError(err error, name []string) {
	n := ""
	if len(name) > 0 {
		n = name[0]
	}

	v.errs = append(v.errs, ValidateError{n, err})
}

func (v *Validator) RequiredString(val string, err error, name ...string) {
	if len(strings.TrimSpace(val)) == 0 {
		v.addError(err, name)
	}
}

func (v *Validator) RequiredInt(val int, err error, name ...string) {
	if val == 0 {
		v.addError(err, name)
	}
}

func (v *Validator) RequiredFloat64(val float64, err error, name ...string) {
	if val == 0 {
		v.addError(err, name)
	}
}

func (v *Validator) RequiredBool(val bool, err error, name ...string) {
	if !val {
		v.addError(err, name)
	}
}

func (v *Validator) RequiredEmail(val string, err error, name ...string) {
	if val == "" {
		v.addError(err, name)
	}

	v.Email(val, err, name...)
}

func (v *Validator) NotNil(val interface{}, err error, name ...string) {
	if val == nil {
		v.addError(err, name)
	}
}

func (v *Validator) RequiredTime(val time.Time, err error, name ...string) {
	if val.IsZero() {
		v.addError(err, name)
	}
}

func (v *Validator) MinInt(val int, n int, err error, name ...string) {
	if val > n {
		return
	}

	v.addError(err, name)
}

func (v *Validator) MaxInt(val int, n int, err error, name ...string) {
	if val > n {
		v.addError(err, name)
	}
}

func (v *Validator) MinFloat64(val float64, n float64, err error, name ...string) {
	if val < n {
		return
	}

	v.addError(err, name)
}

func (v *Validator) MaxFloat64(val float64, n float64, err error, name ...string) {
	if val > n {
		return
	}
	v.addError(err, name)
}

func (v *Validator) MinChar(val string, n int, err error, name ...string) {
	if utf8.RuneCountInString(val) < n {
		v.addError(err, name)
	}
}

func (v *Validator) MaxChar(val string, n int, err error, name ...string) {
	if utf8.RuneCountInString(val) > n {
		v.addError(err, name)
	}
}

func (v *Validator) Email(val string, err error, name ...string) {
	if val == "" {
		return
	}
	if !emailPatern.MatchString(val) {
		v.addError(err, name)
	}
}

func (v *Validator) Gender(val string, err error, name ...string) {
	if val != `male` && val != `female` {
		v.addError(err, name)
	}
}

func (v *Validator) Confirm(val, confirm string, err error, name ...string) {
	if val != confirm {
		v.addError(err, name)
	}
}

func (v *Validator) ISO8601DataTime(val string, err error, name ...string) {
	if val == "" {
		return
	}
	if !dateiso8601Patern.MatchString(val) {
		v.addError(err, name)
	}
}

func (v *Validator) Length(val int, atleast int, err error, name ...string) {
	if val <= atleast {
		v.addError(err, name)
	}
}

func (v *Validator) InString(val string, in []string, err error, name ...string) {
	for _, k := range in {
		if k == val {
			return
		}
	}

	v.addError(err, name)
}
