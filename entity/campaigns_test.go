package entity

import (
	"html/template"
	"testing"
)

func TestAttackType_validate(t *testing.T) {
	type fields struct {
		ID     string
		Params map[string]any
		Body   *template.Template
	}
	type args struct {
		params map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Identical",
			fields: fields{
				Params: map[string]any{"key1": "key1", "key2": "key2"},
			},
			args: args{
				params: map[string]any{"key1": "key1", "key2": "key2"},
			},
			want: true,
		},
		{
			name: "Missing 1 Key",
			fields: fields{
				Params: map[string]any{"key1": "key1", "key2": "key2"},
			},
			args: args{
				params: map[string]any{"key1": "key1"},
			},
			want: false,
		},
		{
			name: "Same Keys - Different Values",
			fields: fields{
				Params: map[string]any{"key1": "key1", "key2": "key2"},
			},
			args: args{
				params: map[string]any{"key1": "key1", "key2": "not-identical"},
			},
			want: true,
		},
		{
			name: "Different Keys - Same Values",
			fields: fields{
				Params: map[string]any{"key1": "key1", "key2": "key2"},
			},
			args: args{
				params: map[string]any{"key3": "key1", "key4": "key2"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attack := &AttackType{
				Params: tt.fields.Params,
			}
			if got := attack.validate(tt.args.params); got != tt.want {
				t.Errorf("AttackType.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
