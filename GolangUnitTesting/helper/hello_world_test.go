package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
t.Fail() akan menggagalkan unit test namun tetap melanjutkan eksekusi kode di bawahnya
t.FailNow() akan menggagalkan unit test dan menghentikan eksekusi kode di bawahnya
t.Error() sama seperti t.Fail() namun dapat menambahkan pesan error
t.Fatal() sama seperti t.FailNow() namun dapat menambahkan pesan error
*/
// ini penting dan paling sering digunakan

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ayub",
			request:  "Ayub",
			expected: "Hello Ayub",
		},
		{
			name:     "Subagiya",
			request:  "Subagiya",
			expected: "Hello Subagiya",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Ayub", func(t *testing.T) {
		result := HelloWorld("Ayub")
		require.Equal(t, "Hello Ayub", result, "Result harusnya 'Hello Ayub'")
	})
	t.Run("Subagiya", func(t *testing.T) {
		result := HelloWorld("Subagiya")
		require.Equal(t, "Hello Subagiya", result, "Result harusnya 'Hello Subagiya'")
	})
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ayub")
	if result != "Hello Ayub" {
		// error
		t.Error("Result harusnya 'Hello Ayub'")
	}
	fmt.Println("Test Hello world Ayub selesai dijalankan")
}

func TestHelloWorldSubagiya(t *testing.T) {
	result := HelloWorld("Subagiya")
	if result != "Hello Subagiya" {
		// error
		t.Fatal("Result harusnya 'Hello Subagiya1'")
	}
	fmt.Println("Test Hello world Subagiya selesai dijalankan")
}

func TestHelloWorldKurniawan(t *testing.T) {
	result := HelloWorld("Kurniawan")
	if result != "Hello Kurniawan" {
		// error
		t.Fatal("Result harusnya 'Hello Kurniawan'")
	}
	fmt.Println("Test Hello world Subagiya selesai dijalankan")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ayub")
	require.Equal(t, "Hello Ayub", result, "Result harusnya 'Hello Ayub'")
	fmt.Println("Test Hello world Assert")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ayub")
	assert.Equal(t, "Hello Ayub", result, "Result harusnya 'Hello Ayub'")
	fmt.Println("Test Hello world Assert")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Tidak bisa dijalankan di Mac OS")
	} else if runtime.GOOS == "linux" {
		t.Skip("Tidak bisa dijalankan di Linux OS")
	} else if runtime.GOOS == "windows" {
		fmt.Println("Test berjalan di Windows OS")
	}
	result := HelloWorld("Ayub")
	require.Equal(t, "Hello Ayub", result, "Result harusnya 'Hello Ayub'")
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum Unit Test")

	m.Run()

	// after
	fmt.Println("Setelah Unit Test")
}

// benchmark Table
func BenchMarkSub(b *testing.B) {
	b.Run("Ayub", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ayub")
		}
	})
	b.Run("Subagiya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Subagiya")
		}
	})
}

func BenchmarkHelloWorldAyub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ayub")
	}
}

func BenchmarkHelloWorldSubagiya(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Subagiya")
	}
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Ayub",
			request: "Ayub",
		},
		{
			name:    "Subagiya",
			request: "Subagiya",
		},
		{
			name:    "Subagiya ayub",
			request: "Subagiya ayub",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
