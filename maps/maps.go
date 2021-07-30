package main

import "errors"

type Dictionary map[string]string

var NotFoundErr = errors.New("word is not found")
var KeyExistsErr = errors.New("word you trying to add already exists")

func (dictionary Dictionary) Search(word string) (string, error) {

	w, ok := dictionary[word]

	if !ok {
		return w, NotFoundErr
	}

	return dictionary[word], nil
}

func (dictionary Dictionary) Add(key, val string) error {

	_, err := dictionary.Search(key)

	switch err {
	case NotFoundErr:
		dictionary[key] = val
	case nil:
		return KeyExistsErr
	default:
		return err
	}

	return nil
}
