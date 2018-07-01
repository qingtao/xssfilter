package xssfilter

import "testing"

func TestFilter_Clean(t *testing.T) {
	type args struct {
		str  string
		root string
	}
	tests := []struct {
		name    string
		filter  *Filter
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "defaultFilter",
			filter: defaultFilter,
			args: args{
				str:  before,
				root: "body",
			},
			want:    after,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.filter.Clean(tt.args.str, tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("Filter.Clean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Filter.Clean() = \n%v\n want = \n%v\n", got, tt.want)
			}
		})
	}
}
