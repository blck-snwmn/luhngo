package luhngo

import "testing"

func Test_luhnn_Verify(t *testing.T) {
	type fields struct {
		sl string
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "sl=number, input=49927398716, valid",
			fields:  fields{sl: "0123456789"},
			args:    args{input: "49927398716"},
			wantErr: false,
		},
		{
			name:    "sl=number(9=>A), input=499273A8716, valid",
			fields:  fields{sl: "012345678A"},
			args:    args{input: "4AA273A8716"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, _ := NewLuhnn(tt.fields.sl)
			if err := l.Verify(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("luhnn.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
