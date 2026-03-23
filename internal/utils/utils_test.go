package utils_test

import (
	"fmt"
	"testing"

	"github.com/jtprogru/gch/internal/utils"
	"github.com/magiconair/properties/assert"
)

func TestCheckDigit(t *testing.T) {
	testCases := []struct {
		s string
		a bool
	}{
		{
			s: "1q2w3e4r",
			a: true,
		},
		{
			s: "1234567890",
			a: true,
		},
		{
			s: "qwertyuiop1",
			a: true,
		},
		{
			s: "qwertyuiop",
			a: false,
		},
		{
			s: "!@#$%^&*()",
			a: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.s, func(t *testing.T) {
			got := utils.CheckDigit(tt.s)
			assert.Equal(t, got, tt.a, fmt.Sprintf("want %t, got %t", tt.a, got))
		})
	}
}

func TestCheckPunct(t *testing.T) {
	testCases := []struct {
		s string
		a bool
	}{
		{
			s: "1q2w3e4r",
			a: false,
		},
		{
			s: "qwertyuiop1",
			a: false,
		},
		{
			s: "qwertyuiop",
			a: false,
		},
		{
			s: "1234567890",
			a: false,
		},
		{
			s: "!@#$%^&*()",
			a: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.s, func(t *testing.T) {
			got := utils.CheckPunct(tt.s)
			assert.Equal(t, got, tt.a, fmt.Sprintf("want %t, got %t", tt.a, got))
		})
	}
}

func TestCheckDigitAndPunt(t *testing.T) {
	testCases := []struct {
		s string
		a bool
	}{
		{
			s: "1q2w3e4r",
			a: false,
		},
		{
			s: "qwertyuiop1",
			a: false,
		},
		{
			s: "qwertyuiop",
			a: false,
		},
		{
			s: "1234567890",
			a: false,
		},
		{
			s: "!@#$%^&*()",
			a: false,
		},
		{
			s: "!1@2#3$4%5^6&7*8(9)0",
			a: true,
		},
		{
			s: "q!w1e@r2t#y3u$i4o%p5^6&7*8(9)0",
			a: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.s, func(t *testing.T) {
			got := utils.CheckDigitAndPunt(tt.s)
			assert.Equal(t, got, tt.a, fmt.Sprintf("want %t, got %t", tt.a, got))
		})
	}
}
