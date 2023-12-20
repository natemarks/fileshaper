package internal

import (
	"reflect"
	"testing"
)

func TestDeDuplicateLines(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty", args{[]string{}}, []string{}},
		{"single", args{[]string{"a"}}, []string{"a"}},
		{"duplicate", args{[]string{"a", "a"}}, []string{"a"}},
		{"multiple", args{[]string{"apple", "banana", "apple", "orange", "banana", "grape", "apple"}}, []string{"apple", "banana", "orange", "grape"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeDuplicateLines(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeDuplicateLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuplicates(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult DuplicateSet
	}{
		{"empty",
			args{[]string{}},
			DuplicateSet{
				Lines: []DuplicateLineNumbers{}}},
		{"a,b",
			args{[]string{"a", "b"}},
			DuplicateSet{
				Lines: []DuplicateLineNumbers{
					{Line: "a", LineNumbers: []int{0}},
					{Line: "b", LineNumbers: []int{1}},
				},
			},
		},
		{"a,b,b",
			args{[]string{"a", "b", "b"}},
			DuplicateSet{
				Lines: []DuplicateLineNumbers{
					{Line: "a", LineNumbers: []int{0}},
					{Line: "b", LineNumbers: []int{1, 2}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Duplicates(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Duplicates() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestDuplicateSet_String(t *testing.T) {
	type fields struct {
		Lines []DuplicateLineNumbers
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty",
			fields{[]DuplicateLineNumbers{}},
			"",
		},
		{"a,b",
			fields{[]DuplicateLineNumbers{
				{Line: "a", LineNumbers: []int{0}},
				{Line: "b", LineNumbers: []int{1}},
			}},
			"a: [0]\nb: [1]",
		},
		{"a,b,b",
			fields{[]DuplicateLineNumbers{
				{Line: "a", LineNumbers: []int{0}},
				{Line: "b", LineNumbers: []int{1, 2}},
			}},
			"a: [0]\nb: [1 2]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DuplicateSet{
				Lines: tt.fields.Lines,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
