package cal


import "testing"
func TestAdd(t *testing.T) {
	result := Add(20, 38)
	if result != 8.0 {
		t.Errorf("add fail, got %f, wanted %f", result, 58.0)
	}
}
