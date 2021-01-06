package stack

import "testing"

func Test_isValid591(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"</A></A></A></A>"}, false},
		{"", args{"<A></A><B></B>"}, false},
		{"", args{"<![CDATA[wahaha]]]><![CDATA[]> wahaha]]>"}, false},

		{"", args{"<DIV>This is the first line <![CDATA[<div>]]></DIV>"}, true},
		{"", args{"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"}, true},

		{"", args{"<A>  <B> </A>   </B>"}, false},
		{"", args{"<DIV>  div tag is not closed  <DIV>"}, false},
		{"", args{"<DIV>  unmatched <  </DIV>"}, false},
		{"", args{"<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"}, false},
		{"", args{"<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>"}, false},
		{"", args{"<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid591(tt.args.code); got != tt.want {
				t.Errorf("isValid591() = %v, want %v", got, tt.want)
			}
		})
	}
}
