package passwd_test

import (
	"fmt"
	"testing"

	"github.com/jtprogru/gch/internal/passwd"
	"github.com/jtprogru/gch/internal/utils"
	"github.com/magiconair/properties/assert"
)

func TestGetPasswd(t *testing.T) {
	testCases := []struct {
		conf passwd.Config
	}{
		{
			conf: passwd.Config{
				Length:         64,
				IncludeDigits:  false,
				IncludeSymbols: false,
			},
		},
		{
			conf: passwd.Config{
				Length:         32,
				IncludeDigits:  true,
				IncludeSymbols: false,
			},
		},
		{
			conf: passwd.Config{
				Length:         10,
				IncludeDigits:  false,
				IncludeSymbols: true,
			},
		},
		{
			conf: passwd.Config{
				Length:         64,
				IncludeDigits:  true,
				IncludeSymbols: false,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Length(%d)_IncludeDigits(%t)_IncludeSymbols(%t)", tt.conf.Length, tt.conf.IncludeDigits, tt.conf.IncludeSymbols), func(t *testing.T) {
			gotPasswd := passwd.GetPasswd(tt.conf)
			assert.Equal(t, len(gotPasswd), tt.conf.Length, fmt.Sprintf("expected password length %d, got %d", tt.conf.Length, len(gotPasswd)))
			assert.Equal(t, utils.CheckDigit(gotPasswd), tt.conf.IncludeDigits, fmt.Sprintf("expected password to contain digits, got "+gotPasswd))
			assert.Equal(t, utils.CheckPunct(gotPasswd), tt.conf.IncludeSymbols, fmt.Sprintf("expected password to contain punct, got "+gotPasswd))
			assert.Equal(t, utils.CheckDigitAndPunt(gotPasswd), tt.conf.IncludeSymbols && tt.conf.IncludeDigits, fmt.Sprintf("expected password to contain only letters, got "+gotPasswd))
		})
	}
}

func BenchmarkGetPasswd(b *testing.B) {
	benchmarks := []passwd.Config{
		{Length: 16, IncludeDigits: true, IncludeSymbols: true},
		{Length: 64, IncludeDigits: true, IncludeSymbols: true},
		{Length: 1_024, IncludeDigits: true, IncludeSymbols: true},
		{Length: 1_048_576, IncludeDigits: true, IncludeSymbols: true},
		{Length: 3_145_728, IncludeDigits: true, IncludeSymbols: true},
	}

	for _, cfg := range benchmarks {
		cfg := cfg // capture range variable
		b.Run(fmt.Sprintf("len(%d)_digits(%t)_symbols(%t)", cfg.Length, cfg.IncludeDigits, cfg.IncludeSymbols), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = passwd.GetPasswd(cfg)
			}
		})
	}
}
