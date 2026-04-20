package cast_test

import (
	"testing"

	"github.com/bdlm/cast/v2"
)

func BenchmarkTo_chan_int_fromInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := cast.To[chan int](42)
		<-ch
	}
}

func BenchmarkTo_chan_int_fromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := cast.To[chan int]("42")
		<-ch
	}
}

func BenchmarkTo_chan_string_fromInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := cast.To[chan string](42)
		<-ch
	}
}

func BenchmarkTo_chan_int_withLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch, _ := cast.ToE[chan int](42, cast.Op{Flag: cast.LENGTH, Val: 10})
		<-ch
	}
}
