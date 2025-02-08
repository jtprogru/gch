package datescalculator

import (
	"testing"
	"time"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name      string
		dateInput string
		daysCount int
		want      *DatesResponse
		wantErr   bool
	}{
		{
			name:      "Default values for count",
			dateInput: "2006-01-02",
			daysCount: 10,
			want: &DatesResponse{
				Today:      time.Now().Format(defaultDateFormat),
				PastDate:   "2005-12-23",
				FutureDate: "2006-01-12",
			},
		},
		{
			name:      "Calculate 1000 days",
			dateInput: "2666-06-06",
			daysCount: 1000,
			want: &DatesResponse{
				Today:      time.Now().Format(defaultDateFormat),
				PastDate:   "2663-09-10",
				FutureDate: "2669-03-02",
			},
		},
		{
			name:      "Calculate Zero days",
			dateInput: "2266-06-06",
			daysCount: 0,
			want: &DatesResponse{
				Today:      time.Now().Format(defaultDateFormat),
				PastDate:   "2266-06-06",
				FutureDate: "2266-06-06",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Calc(tt.dateInput, tt.daysCount)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Calc() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Calc() succeeded unexpectedly")
			}
			if got.FutureDate != tt.want.FutureDate || got.PastDate != tt.want.PastDate {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
