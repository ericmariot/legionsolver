package main

import (
	"testing"
)

func TestSolver(t *testing.T) {
	t.Run("should create an empty 2d map", func(t *testing.T) {
		got := emptyGrid()
		want := [4][4]int{}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should set a single position", func(t *testing.T) {
		positions := [][]int{
			{0, 0},
		}
		got := solver(positions, []int{0})
		want := [4][4]int{
			{1, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should set multiple positions", func(t *testing.T) {
		positions := [][]int{
			{1, 1},
			{1, 2},
			{2, 1},
			{2, 2},
			{3, 2},
			{3, 3},
		}
		got := solver(positions, []int{0})
		want := [4][4]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 1, 1},
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should place a single piece in a position", func(t *testing.T) {
		grid := [4][4]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		}
		pieces := []int{1}
		got := setPieces(grid, pieces)
		want := [4][4]int{
			{0, 0, 0, 0},
			{0, 2, 1, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should place a single piece with size 3 in a position", func(t *testing.T) {
		grid := [4][4]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		}
		pieces := []int{3}
		got := setPieces(grid, pieces)
		want := [4][4]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 2, 2, 2},
			{0, 0, 0, 0},
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should place multiple pieces with size 1", func(t *testing.T) {
		grid := [4][4]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		}
		pieces := []int{1, 1}
		got := setPieces(grid, pieces)
		want := [4][4]int{
			{0, 0, 0, 0},
			{0, 2, 2, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should place multiple pieces with size 1 and 3", func(t *testing.T) {
		grid := [4][4]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		}
		pieces := []int{1, 3}
		got := setPieces(grid, pieces)
		want := [4][4]int{
			{0, 0, 0, 0},
			{0, 2, 1, 0},
			{0, 2, 2, 2},
			{0, 0, 0, 0},
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	// t.Run("should place multiple pieces", func(t *testing.T) {
	// 	grid := [4][4]int{
	// 		{0, 0, 1, 1},
	// 		{0, 1, 0, 0},
	// 		{0, 1, 1, 1},
	// 		{0, 0, 0, 1},
	// 	}
	// 	pieces := []int{1, 3, 1, 2}
	// 	got := setPieces(grid, pieces)
	// 	want := [4][4]int{
	// 		{0, 0, 2, 2},
	// 		{0, 2, 0, 0},
	// 		{0, 2, 2, 2},
	// 		{0, 0, 0, 2},
	// 	}

	// 	if got != want {
	// 		t.Errorf("got %v, want %v", got, want)
	// 	}
	// })
}
