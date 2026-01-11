package main

import (
	"fmt"
	"strings"
)

func main(){
	AgentName :="Vikram"
	role := "Squad Commander"
	fmt.Printf("Hi Agent %s\n",strings.ToUpper(AgentName))
	fmt.Printf("Role: %s\n",strings.ToUpper(role))
}
