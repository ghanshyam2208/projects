package iterations

import "testing"

func TestRepeat(t *testing.T) {
	t.Run(" test repeated 5 times", func(t *testing.T) {
		actual := Repeat("a")
		expected := "aaaaa"

		if actual != expected {
			t.Errorf("expected '%s actual '%s' ", expected, actual)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
