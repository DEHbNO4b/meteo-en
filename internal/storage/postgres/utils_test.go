package postgres

import "testing"

func Test_parseWind(t *testing.T) {

	tests := []struct {
		name    string
		str     string
		avg     int
		max     int
		wantErr bool
	}{
		{
			name:    "pozitive case #1",
			str:     " 9-20{23}",
			avg:     14,
			max:     23,
			wantErr: false,
		},
		{
			name:    "pozitive case #2",
			str:     "9/",
			avg:     9,
			max:     0,
			wantErr: false,
		},
		{
			name:    "pozitive case #3",
			str:     "9{14}",
			avg:     9,
			max:     14,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avg, max, err := parseWind(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseWind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if avg != tt.avg {
				t.Errorf("parseWind() got = %v, want %v", avg, tt.avg)
			}
			if max != tt.max {
				t.Errorf("parseWind() got1 = %v, want %v", max, tt.max)
			}
		})
	}
}
