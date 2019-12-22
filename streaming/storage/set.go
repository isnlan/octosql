package storage

//TODO: this code is of questionable quality. It might be useful to change it later

import (
	"github.com/cube2222/octosql"
	"github.com/mitchellh/hashstructure"
	"github.com/pkg/errors"
)

type Set struct {
	tx StateTransaction
}

type SetIterator struct {
	it Iterator
}

func NewSet(tx StateTransaction) *Set {
	return &Set{
		tx: tx,
	}
}

func (set *Set) Insert(value octosql.Value) (bool, error) {
	hash, bytes, err := getHashAndBytes(value)
	if err != nil {
		return false, errors.Wrap(err, "couldn't parse given value in insert")
	}

	hashedTxn := set.tx.WithPrefix(hash)
	state := NewValueState(hashedTxn)

	var tupleValue octosql.Value
	tupleBytes, err := state.getBytes()

	if err == ErrKeyNotFound {
		tupleValue = octosql.ZeroTuple()
		tupleBytes = tupleValue.MonotonicMarshal()
	} else if err == nil {
		err := tupleValue.MonotonicUnmarshal(tupleBytes)
		if err != nil {
			return false, errors.Wrap(err, "failed to parse set element")
		}
	} else if err != nil {
		return false, errors.Wrap(err, "failed to read set elements")
	}

	if isMemberOfTuple(tupleValue.AsSlice(), value) {
		return false, nil
	}

	newTupleBytes := addValueBytesToTupleBytes(tupleBytes, bytes)

	err = state.setBytes(newTupleBytes)
	if err != nil {
		return false, errors.Wrap(err, "couldn't actualize new value for set")
	}

	return true, nil
}

func (set *Set) Contains(value octosql.Value) (bool, error) {
	hash, _, err := getHashAndBytes(value)
	if err != nil {
		return false, errors.Wrap(err, "couldn't parse given value in contains")
	}

	hashedTxn := set.tx.WithPrefix(hash)
	state := NewValueState(hashedTxn)

	var tuple octosql.Value
	tupleBytes, err := state.getBytes()

	if err == ErrKeyNotFound {
		return false, nil
	} else if err != nil {
		return false, errors.Wrap(err, "failed to read set")
	}

	err = tuple.MonotonicUnmarshal(tupleBytes)
	if err != nil {
		return false, errors.Wrap(err, "failed to parse set element")
	}

	values := tuple.AsSlice()
	return isMemberOfTuple(values, value), nil
}

func NewSetIterator(it Iterator) *SetIterator {
	return &SetIterator{
		it: it,
	}
}

func getHashAndBytes(value octosql.Value) ([]byte, []byte, error) {
	valueBytes := value.MonotonicMarshal()

	hashValue, err := hashstructure.Hash(valueBytes, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't hash value")
	}

	hashKey := octosql.MonotonicMarshalUint64(hashValue, true)

	return hashKey, valueBytes, nil
}

func addValueBytesToTupleBytes(tuple, newBytes []byte) []byte {
	newTuple := tuple[:len(tuple)-1]
	newTuple = append(newTuple, newBytes...)
	newTuple = append(newTuple, octosql.TupleDelimiter)

	return newTuple
}

func isMemberOfTuple(values []octosql.Value, value octosql.Value) bool {
	for _, v := range values {
		if octosql.AreEqual(v, value) {
			return true
		}
	}

	return false
}
