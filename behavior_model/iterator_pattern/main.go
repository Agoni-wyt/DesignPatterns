package main

import (
	"fmt"
	"iterator_pattern/iterator"
)

func main() {

	user1 := &iterator.User{
		Name: "a",
		Age:  30,
	}
	user2 := &iterator.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}

	createIterator := userCollection.CreateIterator()

	for createIterator.HasNext() {
		user := createIterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
