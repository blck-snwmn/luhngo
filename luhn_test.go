package luhngo

import (
	"testing"
)

func Test_rtoi(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "r=/ -> /", args: args{r: '/'}, want: 0, wantErr: true},
		{name: "r=0 -> 0", args: args{r: '0'}, want: 0, wantErr: false},
		{name: "r=1 -> 1", args: args{r: '1'}, want: 1, wantErr: false},
		{name: "r=8 -> 8", args: args{r: '8'}, want: 8, wantErr: false},
		{name: "r=9 -> 9", args: args{r: '9'}, want: 9, wantErr: false},
		{name: "r=: -> :", args: args{r: ':'}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rtoi(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("rtoi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("rtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "input=49927398716, valid",
			args:    args{input: "49927398716"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Verify(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
