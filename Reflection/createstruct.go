package main

import (
	"reflect"
	"fmt"
	"strconv"
	"encoding/binary"
	"math"
)

type Person struct {
	Name 	string
	Age 	int64
	Salary  float64
}

// CloneToByteType takes a reflect value of struct and returns a new interface with same exported field as passed struct and all data converted to bytes.
func CloneToByteType(value reflect.Value) interface{} {
	val := value.Elem()
	sti := CreateByteType(value.Type())
	// populate struct with bytes
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		reflect.ValueOf(sti).Elem().Field(i).SetBytes(CastTypeToByte(valueField))
	}
	return sti
}

// CreateByteType takes a reflect value of struct and returns a new interface with same exported field as passed struct
// All values are instantiated to zero values.
func CreateByteType(t reflect.Type) interface{} {
	val := t.Elem()
	sf := make([]reflect.StructField, val.NumField())
	// make struct fields
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Field(i)

		s := reflect.StructField{
			Name: typeField.Name,
			Type: reflect.TypeOf([]byte{}),
		}
		sf[i] = s
	}
	// create new struct
	st := reflect.StructOf(sf)
	sti := reflect.New(st).Interface()
	return sti
}

// CloneFromByteType takes 'p' reflect.Type of intended type and byte type interface with values.
// It create new instance of 'p' type and populates the fields from byte types.
func CloneFromByteType(p reflect.Type, in interface{}) interface{} {
	returnValue := reflect.New(p).Interface()
	val := reflect.ValueOf(returnValue).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := reflect.ValueOf(in).Elem().FieldByName(val.Type().Field(i).Name)
		val.Field(i).Set(CastByteToType(val.Field(i).Kind(), valueField))
	}
	return returnValue
}

func CastTypeToByte(value reflect.Value) []byte {
	switch value.Kind() {
	case reflect.Int64:
		return []byte(strconv.Itoa(int(value.Int())))
	case reflect.Float64:
		return float64ToByte(value.Float())
	case reflect.String:
		return []byte(value.String())
	default:
		return []byte{}
	}
}

func CastByteToType(kind reflect.Kind, valueField reflect.Value) reflect.Value {
	switch kind {
	case reflect.Int64:
		i,_  := strconv.Atoi(string(valueField.Bytes()))
		return reflect.ValueOf(int64(i))
	case reflect.Float64:
		return reflect.ValueOf(math.Float64frombits(binary.BigEndian.Uint64(valueField.Bytes())))
	case reflect.String:
		return reflect.ValueOf(string(valueField.Bytes()))
	default:
		return valueField
	}
}


func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}


func main() {
	p := Person{Name:"Prashant", Age:30, Salary:1.11}
	conP := CloneToByteType(reflect.ValueOf(&p))
	fmt.Printf("%+v\n",conP)
	pp := CloneFromByteType(reflect.TypeOf(p), conP).(*Person)
	fmt.Printf("%+v",pp)
}