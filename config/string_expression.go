package config

import (
	"encoding/json"
	"fmt"
	"github.com/expr-lang/expr"
	"html/template"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// StringExpression replace the python String like "%(var)s" to string
type StringExpression struct {
	env map[string]string // the environment variable used to replace the var in the python expression
}

// NewStringExpression create a new StringExpression with the environment variables
func NewStringExpression(envs ...string) *StringExpression {
	se := &StringExpression{env: make(map[string]string)}

	for _, env := range os.Environ() {
		t := strings.SplitN(env, "=", 2)
		se.env["ENV_"+t[0]] = t[1]
	}
	n := len(envs)
	for i := 0; i+1 < n; i += 2 {
		se.env[envs[i]] = envs[i+1]
	}

	hostname, err := os.Hostname()
	if err == nil {
		se.env["host_node_name"] = hostname
	}

	return se

}

// Add add the environment variable (key,value)
func (se *StringExpression) Add(key string, value string) *StringExpression {
	se.env[key] = value
	return se
}

// Eval evaluate the expression include "%(var)s"  and return the string after replacing the var
func (se *StringExpression) Eval(s string) (string, error) {

	for {
		// find variable start indicator
		start := strings.Index(s, "%(")

		if start == -1 {
			return s, nil
		}

		end := start + 1
		n := len(s)

		// find variable end indicator
		for end < n && s[end] != ')' {
			end++
		}

		// find the type of the variable
		typ := end + 1
		for typ < n && !((s[typ] >= 'a' && s[typ] <= 'z') || (s[typ] >= 'A' && s[typ] <= 'Z')) {
			typ++
		}

		// evaluate the variable
		if typ < n {
			varName := s[start+2 : end]

			if s[typ] == 'd' {
				for k, v := range se.env {
					varName = strings.Replace(varName, k, v, -1)
				}
				program, err := expr.Compile(varName, expr.Env(se.env))
				if err != nil {
					return "", fmt.Errorf("fail to find the environment variable %s", varName)
				}
				varValue, err := expr.Run(program, se.env)
				//进行计算
				s = s[0:start] + fmt.Sprintf("%d", varValue) + s[typ+1:]
			} else if s[typ] == 's' {
				program, err := expr.Compile(varName, expr.Env(se.env))
				if err != nil {
					return "", fmt.Errorf("fail to find the environment variable %s", varName)
				}
				varValue, err := expr.Run(program, se.env)
				result, _ := ToStringE(varValue)
				s = s[0:start] + result + s[typ+1:]
			} else {
				return "", fmt.Errorf("not implement type:%v", s[typ])
			}
			fmt.Println(s)
		} else {
			return "", fmt.Errorf("invalid string expression format")
		}
	}

}

var (
	errorType       = reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

// Copied from html/template/content.go.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a any) any {
	if a == nil {
		return nil
	}
	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ToStringE converts any type to a string type.
func ToStringE(i any) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case json.Number:
		return s.String(), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}
