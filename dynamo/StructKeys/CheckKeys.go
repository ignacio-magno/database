package StructKeys

import "fmt"

type CheckKeys[T any] struct {
	k *KeysQuery[T]
}

func NewCheckKeys[T any](keysQuery *KeysQuery[T]) *CheckKeys[T] {
	return &CheckKeys[T]{
		k: keysQuery,
	}
}

// compare if the values contain the required keys to this model
func (s *CheckKeys[T]) checkIsCorrectAllKeysRequired(values []interface{}) error {
	if (s.k.haveTwoKeys && len(values) != 2) || (!s.k.haveTwoKeys && len(values) != 1) {
		return fmt.Errorf("the number of keys is incorrect by the model, expected %d, got %d", len(values), len(values))
	}

	if err := s.k.verifyIfPrimaryKeyIsSameType(values[0]); err != nil {
		return err
	}

	if s.k.haveTwoKeys {
		if err := s.k.verifyIfSecondaryKeyIsSameType(values[1]); err != nil {
			return err
		}
	}
	return nil
}

func (s *CheckKeys[T]) checkIfHaveOneKeyOrMore(values []interface{}) error {
	if len(values) == 0 && len(values) > 2 {
		return fmt.Errorf("the number of keys is incorrect by the model, expected at least 1 or 2, got %d", len(values))
	}

	if err := s.k.verifyIfPrimaryKeyIsSameType(values[0]); err != nil {
		return err
	}

	if s.k.haveTwoKeys && len(values) == 2 {
		if err := s.k.verifyIfSecondaryKeyIsSameType(values[1]); err != nil {
			return err
		}
	}

	return nil
}
