package ds

import (
	"errors"
	"fmt"
	"strata/pkg/engine"
	"strconv"
)



// The meta key stores the metadata for the list and the head and tail pointers
// In our app its store
// meta:list:list:head -> head index
func getListMetaKey(list, side string) string {
	return fmt.Sprintf("meta:list:%s:%s", list, side)
}

// The list key stores the value for the list
// In our app its store
// list:[list]:[list_index] -> value
func getListKey(listName string, index int64) string {
	return fmt.Sprintf("list:%s:%d", listName, index)
}

func initializeList(e engine.StorageEngine, list string) (int64, int64) {
	headStr, _, _ := e.Get(getListMetaKey(list, "head"))
	tailStr, _, _ := e.Get(getListMetaKey(list, "tail"))

	if headStr == "" && tailStr == "" {
		e.Set(getListMetaKey(list, "head"), "0")
		e.Set(getListMetaKey(list, "tail"), "-1")
	}

	head, _ := strconv.ParseInt(headStr, 10, 64)
	tail, _ := strconv.ParseInt(tailStr, 10, 64)

	if head > tail {
		e.Set(getListMetaKey(list, "head"), "0")
		e.Set(getListMetaKey(list, "tail"), "-1")
	}
	return head, tail
}

func LPush(e engine.StorageEngine, list, value string) (int64, error) {
	head, tail := initializeList(e, list)
	head--

	err := e.Set(getListKey(list, head), value)
	if err != nil {
		return 0, err
	}

	e.Set(getListMetaKey(list, "head"), strconv.FormatInt(head, 10))
	return tail - head + 1, nil
}

func LPop(e engine.StorageEngine, list string) (string, error) {
	headStr, _, _ := e.Get(getListMetaKey(list, "head"))
	tailStr, _, _ := e.Get(getListMetaKey(list, "tail"))

	if headStr == "" && tailStr == "" {
		return "", errors.New("list is empty")
	}

	head, _ := strconv.ParseInt(headStr, 10, 64)
	tail, _ := strconv.ParseInt(tailStr, 10, 64)

	if head > tail {
		return "", errors.New("list is empty")
	}

	value, _, err := e.Get(getListKey(list, head))
	if err != nil {
		return "", err
	}

	_, err = e.Del(getListKey(list, head))
	if err != nil {
		return "", err
	}

	e.Set(getListMetaKey(list, "head"), strconv.FormatInt(head+1, 10))
	return value, nil

}

func RPush(e engine.StorageEngine, list, value string) (int64, error) {
	head, tail := initializeList(e, list)
	tail++

	err := e.Set(getListKey(list, tail), value)
	if err != nil {
		return 0, err
	}

	e.Set(getListMetaKey(list, "tail"), strconv.FormatInt(tail, 10))
	return tail - head + 1, nil
}

func RPop(e engine.StorageEngine, list string) (string, error) {
	headStr, _, _ := e.Get(getListMetaKey(list, "head"))
	tailStr, _, _ := e.Get(getListMetaKey(list, "tail"))

	if headStr == "" && tailStr == "" {
		return "", errors.New("list is empty")
	}

	head, _ := strconv.ParseInt(headStr, 10, 64)
	tail, _ := strconv.ParseInt(tailStr, 10, 64)

	if head > tail {
		return "", errors.New("list is empty")
	}

	value, _, err := e.Get(getListKey(list, tail))
	if err != nil {
		return "", err
	}

	_, err = e.Del(getListKey(list, tail))
	if err != nil {
		return "", err
	}

	e.Set(getListMetaKey(list, "tail"), strconv.FormatInt(tail-1, 10))
	return value, nil
}
