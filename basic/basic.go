package basic

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

// A lo hecho, pecho. Esto es probablemente más farragoso
// de lo estrictamente necesario. Una vez leida la linea entera
// se podría usar Sscanf/Sscanln, que ya sabe muy bien tratar
// con diferentes tipos de datos ... Pero ahi queda, hasta que
// se cambie.
// También se podria usar un Scanner en vez de el buffer reader

const MaxUint = ^uint(0)

var (
	input          = bufio.NewReader(os.Stdin)
	output         = bufio.NewWriter(os.Stdout)
	intSize        int
	MustBeAPointer = errors.New("basic.Input: arg must be a pointer")
)

func init() {
	if MaxUint == math.MaxInt64 {
		intSize = 64
	} else {
		intSize = 32
	}
}
func getString() (string, error) {
	str, err := input.ReadString('\n')
	if err != nil {
		return "", err
	}
	return str[:len(str)-1], err
}
func getIntSize(size int) (int64, error) {
	str, err := getString()
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(str, 0, size)
	if err != nil {
		return 0, err
	}
	return i, err
}
func getInt32() (int32, error) {
	i64, err := getIntSize(32)
	return int32(i64), err
}
func getInt64() (int64, error) {
	return getIntSize(64)
}
func getInt() (int, error) {
	i64, err := getIntSize(intSize)
	return int(i64), err
}
func getUintSize(size int) (uint64, error) {
	str, err := getString()
	if err != nil {
		return 0, err
	}
	u, err := strconv.ParseUint(str, 0, size)
	if err != nil {
		return 0, err
	}
	return u, err
}
func getUint32() (uint32, error) {
	i64, err := getUintSize(32)
	return uint32(i64), err
}
func getUint64() (uint64, error) {
	return getUintSize(64)
}
func getUint() (uint, error) {
	i64, err := getUintSize(intSize)
	return uint(i64), err
}
func getFloatSize(size int) (float64, error) {
	str, err := getString()
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(str, size)
	if err != nil {
		return 0, err
	}
	return f, err
}
func getFloat32() (float32, error) {
	f64, err := getFloatSize(32)
	return float32(f64), err
}
func getFloat64() (float64, error) {
	return getFloatSize(64)
}
func Input(msg string, val interface{}) (err error) {
	if msg != "" {
		_, err = output.WriteString(msg)
		if err != nil {
			return err
		}
		output.Flush()
	}
	switch val := val.(type) {
	case *int:
		*val, err = getInt()
	case *int64:
		*val, err = getInt64()
	case *int32:
		*val, err = getInt32()
	case *uint:
		*val, err = getUint()
	case *uint64:
		*val, err = getUint64()
	case *uint32:
		*val, err = getUint32()
	case *string:
		*val, err = getString()
	case *float64:
		*val, err = getFloat64()
	case *float32:
		*val, err = getFloat32()
	case int, int64, int32, uint, uint32, uint64, string, float32, float64:
		panic(MustBeAPointer)
	default:
		panic(fmt.Errorf("basic.Input: unsupported type %T", val))
	}
	return err
}
