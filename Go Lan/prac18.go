// Develop a real-time chat application with multiple chat rooms represented by channels. Use the 
// select statement to multiplex incoming messages from different channels. Allow users to join and 
// leave chat rooms concurrently. 




package main

import (
	"fmt"
	"sync"
)

type ChatRoom struct {
	name     string
	messages chan string
}

func NewChatRoom(name string) *ChatRoom {
	return &ChatRoom{
		name:     name,
		messages: make(chan string),
	}
}

func (cr *ChatRoom) Join(user string) {
	fmt.Printf("%s joined %s\n", user, cr.name)
}

func (cr *ChatRoom) Leave(user string) {
	fmt.Printf("%s left %s\n", user, cr.name)
}

func (cr *ChatRoom) Broadcast() {
	for msg := range cr.messages {
		fmt.Printf("[%s] %s\n", cr.name, msg)
	}
}

func main() {
	room1 := NewChatRoom("Room 1")
	room2 := NewChatRoom("Room 2")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		room1.Join("MrX")
		room1.Broadcast()
	}()

	go func() {
		defer wg.Done()
		room2.Join("MrY")
		room2.Broadcast()
	}()

	// Simulate messages being sent
	room1.messages <- "Hello from MrX!"
	room2.messages <- "Hi from MrY!"

	close(room1.messages)
	close(room2.messages)

	wg.Wait()
}
