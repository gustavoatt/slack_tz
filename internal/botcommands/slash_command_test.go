package botcommands

import (
	"reflect"
	"testing"
)

func TestParseSlackTzCommand(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []UserInfo
	}{
		{
			name: "Single user",
			args: args{"<@UGXFXFG90|gustavoatt>"},
			want: []UserInfo{
				{
					UserID:   "UGXFXFG90",
					UserName: "gustavoatt",
				},
			},
		},
		{
			name: "Multiple users",
			args: args{"<@UGXFXFG90|gustavoatt> and this other user <@UGAFZFG70|pepito>"},
			want: []UserInfo{
				{
					UserID:   "UGXFXFG90",
					UserName: "gustavoatt",
				},
				{
					UserID:   "UGAFZFG70",
					UserName: "pepito",
				},
			},
		},
		{
			name: "No match",
			args: args{"This is a long message with no user in it"},
			want: []UserInfo{},
		},
		{
			name: "Only User ID",
			args: args{"This doesn't have a | <@UGXFXFG90>"},
			want: []UserInfo{
				{
					UserID: "UGXFXFG90",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseSlackTzCommand(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSlackTzCommand(%v) = %v, want %v", tt.args.text, got, tt.want)
			}
		})
	}
}
