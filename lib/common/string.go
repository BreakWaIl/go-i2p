package common

import (
	"errors"
)

type String []rune

func (str String) Length() (int, error) {
	if len(str) == 0 {
		// log
		return 0, errors.New("error parsing string: zero length")
	}
	length := Integer([]byte{byte(str[0])})
	inferred_len := length + 1
	str_len := len(str)
	if inferred_len > str_len {
		// log
		return length, errors.New("string parsing warning: string data is shorter than specified by length")
	} else if str_len > inferred_len {
		//log
		return length, errors.New("string parsing warning: string contains data beyond length")
	}
	return length, nil
}

func (str String) Data() ([]rune, error) {
	length, err := str.Length()
	if err != nil {
		switch err.Error() {
		case "error parsing string: zero length":
			return make([]rune, 0), err
		case "string parsing warning: string data is shorter than specified by length":
			return str[1:], err
		case "string parsing warning: string contains data beyond length":
			return str[1 : length+1], err
		}
	}
	return str[1:], nil
}

func (str String) GoString() (string, error) {
	content, err := str.Data()
	return string(content), err
}

func ToI2PString(data []rune) (String, error) {
	data_len := len(data)
	if data_len >= 256 {
		return make([]rune, 0), errors.New("cannot store that much data in I2P string")
	}
	i2p_string := []rune{rune(data_len)}
	i2p_string = append(i2p_string, data...)
	return String(i2p_string), nil
}

func ReadString(data []rune) (String, []rune, error) {
	str := String(data)
	length, err := String(data).Length()
	if err != nil {
		switch err.Error() {
		case "error parsing string: zero length":
			return String{}, make([]rune, 0), err
		case "string parsing warning: string data is shorter than specified by length":
			return str, make([]rune, 0), err
		case "string parsing warning: string contains data beyond length":
			return String(str[:length+1]), str[length+1:], err
		}
	}
	return str, make([]rune, 0), nil
}
