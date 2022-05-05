package main

import (
	"github.com/ralugr/filter-service/pkg/controller"
)

func main() {
	controller.FilterMessage(`{"id": "123","body": "likes to perch on rocks"}`)
}
