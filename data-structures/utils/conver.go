/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package utils

import (
	"errors"
	"math"
	"strconv"
)

var (
	ErrConvertToString = errors.New("can not convert to string")
	ErrConvertToInt    = errors.New("can not convert to integer")
	ErrConvertToFloat  = errors.New("can not convert to float")
)

func ToString(rawVal interface{}) (ret string, err error) {
	if rawVal == nil {
		return
	}
	switch rawVal.(type) {
	case string:
		ret = rawVal.(string)
	case int:
		ret = strconv.Itoa(rawVal.(int))
	case int32:
		ret = strconv.Itoa(int(rawVal.(int32)))
	case int64:
		ret = strconv.FormatInt(rawVal.(int64), 10)
	case int8:
		ret = strconv.Itoa(int(rawVal.(int8)))
	case int16:
		ret = strconv.Itoa(int(rawVal.(int16)))
	case uint8:
		ret = strconv.FormatUint(uint64(rawVal.(uint8)), 10)
	case uint16:
		ret = strconv.FormatUint(uint64(rawVal.(uint16)), 10)
	case uint32:
		ret = strconv.FormatUint(uint64(rawVal.(uint32)), 10)
	case uint64:
		ret = strconv.FormatUint(uint64(rawVal.(uint64)), 10)
	case uint:
		ret = strconv.FormatUint(uint64(rawVal.(uint)), 10)
	case bool:
		b := rawVal.(bool)
		ret = strconv.FormatBool(b)
	case float32:
		f32 := rawVal.(float32)
		ret = strconv.FormatFloat(float64(f32), 'f', -1, 32)
	case float64:
		f64 := rawVal.(float64)
		ret = strconv.FormatFloat(f64, 'f', -1, 64)
		// TODO: 增加更多类型
	default:
		err = ErrConvertToString
	}
	return
}
func ToInt(rawVal interface{}) (ret int64, err error) {
	if rawVal == nil {
		return
	}
	switch rawVal.(type) {
	case int8:
		ret = int64(rawVal.(int8))
	case int16:
		ret = int64(rawVal.(int16))
	case int32:
		ret = int64(rawVal.(int32))
	case int64:
		ret = int64(rawVal.(int64))
	case int:
		ret = int64(rawVal.(int))
	case uint8:
		ret = int64(rawVal.(uint8))
	case uint16:
		ret = int64(rawVal.(uint16))
	case uint32:
		ret = int64(rawVal.(uint32))
	case uint64:
		ret = int64(rawVal.(uint64))
	case uint:
		ret = int64(rawVal.(uint))
	case string:
		ret, err = strconv.ParseInt(rawVal.(string), 10, 64)
	case float32:
		ret = int64(rawVal.(float32))
	case float64:
		ret = int64(rawVal.(float64))
	}
	return
}
func ToBool(rawVal interface{}) (ret bool, err error) {
	ret = false
	if rawVal == nil {
		return
	}
	switch rawVal.(type) {
	case int8:
		ret = rawVal.(int8) != 0
	case int16:
		ret = rawVal.(int16) != 0
	case int32:
		ret = rawVal.(int32) != 0
	case int64:
		ret = rawVal.(int64) != 0
	case int:
		ret = rawVal.(int) != 0
	case uint8:
		ret = rawVal.(uint8) != 0
	case uint16:
		ret = rawVal.(uint16) != 0
	case uint32:
		ret = rawVal.(uint32) != 0
	case uint64:
		ret = rawVal.(uint64) != 0
	case uint:
		ret = rawVal.(uint) != 0
	case float32:
		ret = math.IsNaN(float64(rawVal.(float32)))
	case float64:
		f64 := rawVal.(float64)
		ret = math.IsNaN(f64)
	case string:
		s := rawVal.(string)
		ret = s != ""
	case bool:
		ret = rawVal.(bool)
	default:
		// 其他非空的值都认为是true
		ret = true
	}
	return
}
func ToFloat(rawVal interface{}) (ret float64, err error) {
	if rawVal == nil {
		return
	}
	switch rawVal.(type) {
	case int8:
		ret = float64(rawVal.(int8))
	case int16:
		ret = float64(rawVal.(int16))
	case int32:
		ret = float64(rawVal.(int32))
	case int64:
		ret = float64(rawVal.(int64))
	case int:
		ret = float64(rawVal.(int))
	case uint8:
		ret = float64(rawVal.(uint8))
	case uint16:
		ret = float64(rawVal.(uint16))
	case uint32:
		ret = float64(rawVal.(uint32))
	case uint64:
		ret = float64(rawVal.(uint64))
	case float32:
		ret = float64(rawVal.(float32))
	case float64:
		ret = rawVal.(float64)
	case string:
		ret, err = strconv.ParseFloat(rawVal.(string), 64)
	default:
		err = ErrConvertToFloat
	}
	return
}
