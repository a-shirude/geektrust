package service

import "testing"

func TestMetroCardServiceImpl_MetroCard(t *testing.T) {
	type args struct {
		input [][]string
	}
	tests := []struct {
		name string
		s    *MetroCardServiceImpl
		args args
	}{
		{
			name: "success",
			args: args{
				input: getIO(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MetroCardServiceImpl{}
			s.MetroCard(tt.args.input)
		})
	}
}

func getIO() [][]string {
	var io [][]string
	s1 := [3]string{"BALANCE", "MC1", "600"}
	s2 := [...]string{"CHECK_IN", "MC1", "ADULT", "CENTRAL"}
	s3 := [...]string{"PRINT_SUMMARY"}
	io = append(io, s1[:], s2[:], s3[:])
	return io
}
