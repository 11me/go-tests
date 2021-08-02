package main

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",

			struct {
				Name string
			}{"Jo"},

			[]string{"Jo"},
		},

		{
			"struct with two string fileds",

			struct {
				Name string
				City string
			}{"Test", "Moscow"},

			[]string{"Test", "Moscow"},
		},

		{
			"struct with integer field",

			struct {
				Name string
				Age  int
			}{"Jo", 25},

			[]string{"Jo"},
		},

		{
			"struct with nested fields",

			struct {
				Name    string
				Profile Profile
			}{
				"Jo",
				Profile{33, "Moscow"},
			},

			[]string{"Jo", "Moscow"},
		},

		{
			"pointer to things",

			&struct {
				Name    string
				Profile Profile
			}{
				"Jo",
				Profile{33, "Moscow"},
			},
			[]string{"Jo", "Moscow"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {

				got = append(got, input)

			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
