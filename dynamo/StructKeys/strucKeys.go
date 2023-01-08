package StructKeys

import (
	"fmt"
	"reflect"
	"strings"
)

// structKeys take the T type and extract the primary and secondary keys
type structKeys[T any] struct {
	t           T
	haveTwoKeys bool
}

func newStructKeys[T any](haveTwoKeys bool) *structKeys[T] {
	return &structKeys[T]{haveTwoKeys: haveTwoKeys}
}

func (s *structKeys[T]) getPrimaryKey() string {
	return strings.Split(reflect.TypeOf(s.t).Field(0).Tag.Get("dynamodbav"), ",")[0]
}

func (s *structKeys[T]) getSecondaryKey() string {
	return strings.Split(reflect.TypeOf(s.t).Field(1).Tag.Get("dynamodbav"), ",")[0]
}

func (s *structKeys[T]) getPrimaryKeyType() reflect.Type {
	return reflect.TypeOf(s.t).Field(0).Type
}

func (s *structKeys[T]) getSecondaryKeyType() reflect.Type {
	return reflect.TypeOf(s.t).Field(1).Type
}
func (s *structKeys[T]) verifyIfPrimaryKeyIsSameType(key interface{}) error {
	if reflect.TypeOf(key) != s.getPrimaryKeyType() {
		return fmt.Errorf("the type of the primary key is incorrect, expected %s, got %s", s.getPrimaryKeyType(), reflect.TypeOf(key))
	}
	return nil
}

func (s *structKeys[T]) verifyIfSecondaryKeyIsSameType(key interface{}) error {
	if reflect.TypeOf(key) != s.getSecondaryKeyType() {
		return fmt.Errorf("the type of the secondary key is incorrect, expected %s, got %s", s.getSecondaryKeyType(), reflect.TypeOf(key))
	}
	return nil
}
