package simple_bench

import "testing"

func BenchmarkEmptyEqualEmpty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EqualEmpty("")
	}
}

func BenchmarkEmptyLenZero(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LenZero("")
	}
}

func BenchmarkEmptyNotEmpty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NotEmpty("")
	}
}

func BenchmarkEmptyLenSupZero(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LenSupZero("")
	}
}

func BenchmarkNotEmptyEqualEmpty(b *testing.B) {
        for n := 0; n < b.N; n++ {
                EqualEmpty("not_empty")
        }
}

func BenchmarkNotEmptyLenZero(b *testing.B) {
        for n := 0; n < b.N; n++ {
                LenZero("not_empty")
        }
}

func BenchmarkNotEmptyNotEmpty(b *testing.B) {
        for n := 0; n < b.N; n++ {
                NotEmpty("not_empty")
        }
}

func BenchmarkNotEmptyLenSupZero(b *testing.B) {
        for n := 0; n < b.N; n++ {
                LenSupZero("not_empty")
        }
}
