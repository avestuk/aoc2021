package day2

import "testing"

func TestForward(t *testing.T) {
	s := &sub{}
	s.aim = 1
	s.Forward(5)
	wantDepth := 5
	if wantDepth != s.depth {
		t.Errorf("gotDepth: %d != %d", s.depth, wantDepth)
	}

	s = &sub{}
	s.aim = 10
	s.depth = 10
	// Depth += Aim * X
	s.Forward(8)
	wantDepth = 90
	if wantDepth != s.depth {
		t.Errorf("gotDepth: %d != %d", s.depth, wantDepth)
	}

	//wantDepth = wantDepth * 2
	//s.Forward(2)
	//if wantDepth != s.depth {
	//	t.Errorf("gotDepth: %d != %d", s.depth, wantDepth)
	//}

	//got := s.horizontal
	//if got != 15 {
	//	t.Errorf("got != 15, got: %d", got)
	//}
}

func TestDown(t *testing.T) {
	s := &sub{}

	s.Down(5)
	wantAim := 5
	wantDepth := 0
	if s.depth != wantDepth {
		t.Errorf("depth != %d, got: %d", wantDepth, s.depth)
	}
	if s.aim != wantAim {
		t.Errorf("depth != %d, got: %d", wantAim, s.aim)
	}

	s.Down(10)
	wantAim += 10
	if s.depth != wantDepth {
		t.Errorf("depth != %d, got: %d", wantDepth, s.depth)
	}
	if s.aim != wantAim {
		t.Errorf("depth != %d, got: %d", wantAim, s.aim)
	}
}

func TestUp(t *testing.T) {
	s := &sub{}

	s.Up(5)
	wantDepth := 0
	wantAim := -5

	if s.depth != wantDepth {
		t.Errorf("depth != %d, got: %d", wantDepth, s.depth)
	}
	if s.aim != wantAim {
		t.Errorf("aim != %d, got: %d", wantAim, s.aim)
	}

	s.Up(10)
	wantAim = -15
	if s.depth != wantDepth {
		t.Errorf("depth != %d, got: %d", wantDepth, s.depth)
	}
	if s.aim != wantAim {
		t.Errorf("aim != %d, got: %d", wantAim, s.aim)
	}
}

func TestParseMovement(t *testing.T) {

	cmds := map[string][]int{"forward 1": {1, 0}, "down 1": {0, 1}, "up 1": {0, -1}}

	s := &sub{}
	for cmd, _ := range cmds {
		s.Parse(cmd)

	}
	wantHorizontal := 1
	wantAim := 0
	wantDepth := 0

	if s.horizontal != wantHorizontal {
		t.Errorf("%d != %d", s.horizontal, wantHorizontal)
	}

	if s.depth != wantDepth {
		t.Errorf("%d != %d", s.depth, wantDepth)
	}

	if s.aim != wantAim {
		t.Errorf("%d != %d", s.aim, wantAim)
	}
}

func TestTest(t *testing.T) {
	k := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	s := &sub{}
	for _, cmd := range k {
		s.Parse(cmd)
	}

	s.Multiply()

}

func TestMultiplyPosition(t *testing.T) {
	s := &sub{}

	s.Down(2)
	s.Forward(2)

	want := 8

	if s.Multiply() != want {
		t.Errorf("got %d != %d", s.Multiply(), want)
	}
}

func TestDay2P1(t *testing.T) {
	got := Day2P1("./day2_test.txt")
	want := 900

	if got != want {
		t.Errorf("got %d != %d", got, want)
	}
}
