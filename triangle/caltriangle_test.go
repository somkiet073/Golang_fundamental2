package triangle_test

import (
	"hellogo/triangle"
	"log"
	"testing"
)

// ในกรณีการเขียน Unit test ชื่อ func ต้องเหมือนกัน
func TestCalTriangle(t *testing.T) {
	v := triangle.Caltriangle(5, 5)
	log.Println(v)
	if 12.5 != v {
		t.Error("test ", v)
	}
}
