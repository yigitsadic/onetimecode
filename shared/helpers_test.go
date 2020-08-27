package shared

import "testing"

func TestCreateRandomValue(t *testing.T) {
	t.Run("should generate random string", func(t *testing.T) {
		if got := CreateRandomValue(); len(got) != 7 {
			t.Errorf("Generated code should be 7 characters")
		}
	})
}
