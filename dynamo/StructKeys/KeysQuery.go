package StructKeys

type KeysQuery[T any] struct {
	haveTwoKeys bool
	*structKeys[T]
	checkKeys *CheckKeys[T]
}

func NewKeysQuery[T any](haveTwoKeys bool) *KeysQuery[T] {
	k := &KeysQuery[T]{haveTwoKeys: haveTwoKeys, structKeys: newStructKeys[T](haveTwoKeys)}
	k.checkKeys = NewCheckKeys[T](k)
	return k
}

func (s *KeysQuery[T]) BuildTypesAttribute(values []interface{}, needAllKeys bool) (*TypesAttribute[T], error) {
	if needAllKeys {
		err := s.checkKeys.checkIsCorrectAllKeysRequired(values)
		if err != nil {
			return nil, err
		}
	} else {
		err := s.checkKeys.checkIfHaveOneKeyOrMore(values)
		if err != nil {
			return nil, err
		}
	}

	return NewTypesAttribute[T](values, s), nil
}
