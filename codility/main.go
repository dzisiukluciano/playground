package main

import "fmt"

type Node struct {
	next *Node
	value string
}

func (n *Node) isEmpty() bool {
	return n.value == ""
}

var pushChars map[string]bool = map[string]bool{
	"(": true,
	"{": true,
	"[": true,
}

var pairs map[string]string = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

func isPush(s string) bool {
	return pushChars[s]
}

func isPop(s string) bool {
	return !isPush(s)
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	head := &Node{}
	for _, c := range s {
		if isPush(string(c)) {
			newNode := &Node{
				value: string(c),
				next: head,
			}
			head = newNode
		} else {
			if head.isEmpty() {
				return false
			}
			popedNode := head
			head = head.next
			expected := pairs[string(c)]
			if expected != popedNode.value {
				return false
			}
		}
	}

	return head.isEmpty()
}


func main() {
	fmt.Println(isValid("()") == true)
	fmt.Println(isValid("()[]{}") == true)
	fmt.Println(isValid("(]") == false)
	fmt.Println(isValid(")(") == false)
	fmt.Println(isValid("([)]") == false)
	fmt.Println(isValid("{[]}") == true)
	fmt.Println(isValid(")") == false)
	fmt.Println(isValid("([]") == false)
	fmt.Println(isValid("{") == false)
	fmt.Println(isValid("([)])") == false)
	fmt.Println(isValid("(((((((())))))))") == true)
	fmt.Println(isValid("") == true)
}