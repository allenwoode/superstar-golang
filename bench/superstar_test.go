package main

import (
	"awesome/service"
	"testing"
)

var superService service.SuperstarService

func init() {
	superService = service.NewSuperstarService()
}

func TestGetAll(t *testing.T) {
	all := superService.GetAll()
	t.Log(all)
}

func BenchmarkGetAll(b *testing.B) {

	for i := 0; i < b.N; i++ {
		superService.GetAll()
	}
}
