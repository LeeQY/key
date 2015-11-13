package key

import "testing"

func getNilString() *string {
	return nil
}

func BenchmarkMakeKey(b *testing.B) {
	key := "key"
	prefix := "prefix:"

	for i := 0; i < b.N; i++ {
		MakeKey(&key, &prefix)
	}
}

func TestMakeKey(t *testing.T) {
	key := "key"
	prefix := "prefix:"

	a, err := MakeKey(&key, getNilString())
	if err == nil {
		t.Fatal("Should have error while prefix is empty.")
	}

	a, err = MakeKey(getNilString(), &prefix)
	if err == nil {
		t.Fatal("Should have error while key is empty.")
	}

	a, err = MakeKey(&key, &prefix)
	if err != nil {
		t.Fatal(err)
	}

	if *a != "prefix:key" {
		t.Fatal("Result is not corrected.")
	}
}
