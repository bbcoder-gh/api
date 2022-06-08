package endpoint

import "testing"

func TestMethodType_String(t *testing.T) {
	tests := []struct {
		name string
		m    methodType
		want string
	}{
		{"Get", Get, "GET"},
		{"Post", Post, "POST"},
		{"Put", Put, "PUT"},
		{"Delete", Delete, "DELETE"},
		{"Unsupported", 0, "Unsupported(0)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("MethodType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
