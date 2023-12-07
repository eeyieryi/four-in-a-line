package board_test

import (
	"reflect"
	"testing"

	"github.com/eeyieryi/four-in-a-row/board"
)

func TestNewBoard(t *testing.T) {
	t.Parallel()
	want := board.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	got := board.NewBoard(nil)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got: %v", want, got)
	}
}

func TestNewBoardWithState(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input board.Board
		want  board.Board
	}

	for _, tc := range []testCase{
		{
			input: board.Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 1, 2, 1, 2, 0, 0},
			},
			want: board.Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 1, 2, 1, 2, 0, 0},
			},
		},
		{
			input: board.Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 1, 2, 1, 2, 0, 1},
			},
			want: board.Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 2, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
				{0, 1, 2, 1, 2, 0, 1},
			},
		},
	} {
		got := board.NewBoard(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("want %v, got: %v", tc.want, got)
		}
	}
}

func TestAddPiece(t *testing.T) {
	t.Parallel()

	type input struct {
		state  board.Board
		player int
		column int
	}

	type testCase struct {
		input input
		want  board.Board
	}

	for _, tc := range []testCase{
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
				},
				player: 1,
				column: 5,
			},
			want: board.Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0},
			},
		},
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 0, 0},
				},
				player: 2,
				column: 4,
			},
			want: board.Board{
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 2, 1, 0, 0},
			},
		},
	} {
		got := board.AddPiece(tc.input.state, tc.input.player, tc.input.column)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("want %v, got: %v", tc.want, got)
		}
	}
}

func TestCheckValidMove(t *testing.T) {
	t.Parallel()

	type input struct {
		state  board.Board
		player int
		column int
	}

	type testCase struct {
		input input
	}

	for _, tc := range []testCase{
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
				},
				player: 1,
				column: 5,
			},
			// TODO: Add more cases
		},
	} {
		valid := board.IsValidMove(tc.input.state, tc.input.player, tc.input.column)
		if !valid {
			t.Errorf("want valid, got invalid")
		}
	}

}

func TestCheckValidMoveInvalid(t *testing.T) {
	t.Parallel()

	type input struct {
		state  board.Board
		player int
		column int
	}

	type testCase struct {
		input input
	}

	for _, tc := range []testCase{
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 1, 0, 0, 0},
					{0, 0, 0, 2, 0, 0, 0},
					{0, 0, 0, 1, 0, 0, 0},
					{0, 0, 0, 2, 0, 0, 0},
					{0, 0, 0, 1, 0, 0, 0},
					{0, 0, 0, 2, 1, 0, 0},
				},
				player: 2,
				column: 4,
			},
			// TODO: Add more cases
		},
	} {
		valid := board.IsValidMove(tc.input.state, tc.input.player, tc.input.column)
		if valid {
			t.Errorf("want invalid, got valid")
		}
	}

}

func TestCheckBoardState(t *testing.T) {
	t.Parallel()

	type input struct {
		state board.Board
	}

	type testCase struct {
		input input
		want  int
	}

	for _, tc := range []testCase{
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 1, 0, 0, 0},
					{0, 0, 0, 2, 0, 0, 0},
					{0, 1, 2, 1, 0, 0, 0},
					{0, 1, 1, 2, 0, 0, 0},
					{0, 2, 2, 1, 0, 0, 0},
					{0, 1, 2, 2, 1, 0, 0},
				},
			},
			want: 1,
		},
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 1, 0, 0, 0},
					{0, 0, 0, 2, 0, 0, 0},
					{0, 0, 2, 1, 0, 0, 0},
					{0, 1, 1, 2, 0, 0, 0},
					{0, 2, 2, 1, 0, 0, 0},
					{0, 1, 2, 2, 1, 0, 0},
				},
			},
			want: 0,
		},
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 1, 0, 0, 0},
					{0, 0, 1, 2, 0, 0, 0},
					{0, 0, 2, 1, 0, 0, 0},
					{0, 1, 1, 2, 0, 0, 0},
					{0, 2, 2, 1, 2, 0, 0},
					{0, 1, 2, 2, 1, 2, 1},
				},
			},
			want: 2,
		},
		{
			input: input{
				state: board.Board{
					{1, 1, 2, 1, 1, 2, 1},
					{2, 2, 1, 2, 2, 1, 1},
					{1, 2, 2, 1, 1, 2, 2},
					{2, 1, 1, 2, 2, 1, 1},
					{1, 2, 2, 1, 1, 2, 2},
					{2, 1, 2, 2, 1, 2, 1},
				},
			},
			want: -1,
		},
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 1, 0, 0, 0},
					{2, 1, 2, 2, 0, 0, 0},
					{1, 2, 1, 1, 0, 0, 2},
					{1, 2, 2, 2, 0, 0, 2},
					{1, 2, 1, 1, 0, 0, 2},
				},
			},
			want: 1,
		},
		{
			input: input{
				state: board.Board{
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{1, 1, 2, 2, 2, 2, 0},
					{2, 1, 2, 1, 1, 2, 0},
					{1, 2, 1, 2, 1, 1, 1},
					{1, 2, 1, 2, 2, 2, 1},
				},
			},
			want: 2,
		},
	} {
		got := board.GetBoardState(tc.input.state)
		if tc.want != got {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}
