// SPDX-License-Identifier: BSD-3-Clause

package strcase_test

import (
	"testing"

	"github.com/zchee/strcase"
)

func BenchmarkToCamelCase(b *testing.B) {
	tests := []string{
		"test",
		"test_case",
		" test  case ",
		"",
		"many_many_words",
		"AnyKind of_string",
		"odd-fix",
		"numbers2And55with000",
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, s := range tests {
				_ = strcase.ToCamelCase(s)
			}
		}
	})
}

func BenchmarkToLowerCamelCase(b *testing.B) {
	tests := []string{
		"foo-bar",
		"TestCase",
		"",
		"AnyKind of_string",
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, s := range tests {
				_ = strcase.ToLowerCamelCase(s)
			}
		}
	})
}

func BenchmarkToSnakeCase(b *testing.B) {
	tests := []string{
		"testCase",
		"TestCase",
		"Test Case",
		" Test Case",
		"test",
		"test",
		"test_case",
		"",
		"ManyManyWords",
		"manyManyWords",
		"AnyKind of_string",
		"numbers2and55with000",
		"JSONData",
		"userID",
		"AAAbbb",
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, s := range tests {
				_ = strcase.ToSnakeCase(s)
			}
		}
	})
}

func BenchmarkToDelimited(b *testing.B) {
	tests := []string{
		"testCase",
		"TestCase",
		"Test Case",
		" Test Case",
		"test",
		"test_case",
		"",
		"ManyManyWords",
		"manyManyWords",
		"AnyKind of_string",
		"numbers2and55with000",
		"JSONData",
		"userID",
		"AAAbbb",
		"test-case",
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, s := range tests {
				_ = strcase.ToDelimited(s, '@')
			}
		}
	})
}
