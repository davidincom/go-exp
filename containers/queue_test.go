package containers

import (
	"container/list"
	"testing"
)

func Test_customQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		c       *customQueue
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "dequeue_test",
			c: &customQueue{
				queue: list.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Dequeue(); (err != nil) != tt.wantErr {
				t.Errorf("customQueue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_customQueue_Enqueue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		c    *customQueue
		args args
	}{
		{
			name: "enqueue_test",
			c: &customQueue{
				queue: list.New(),
			},
			args: args{value: "A"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Enqueue(tt.args.value)
			size := tt.c.Size()
			want := 1

			if size != want {
				t.Errorf("got size %q, wanted %q", size, want)
			}
		})
	}
}

func Test_customQueue_Front(t *testing.T) {
	tests := []struct {
		name    string
		c       *customQueue
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "front_test",
			c: &customQueue{
				queue: list.New(),
			},
			wantErr: false,
			want:    "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Enqueue("A")
			got, err := tt.c.Front()
			if (err != nil) != tt.wantErr {
				t.Errorf("customQueue.Front() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("customQueue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}
