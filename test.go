package main

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
)

func main() {
	_, _ := parser.ParseOne("SELECT 1")
	fmt.Println("vim-go")
}
