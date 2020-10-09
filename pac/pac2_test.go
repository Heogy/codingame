package main

import "testing"

func Test_fight3(t *testing.T) {
	i := fight3(CISSOR, ROOK)
	if i != -1{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, -1)

	}
	i = fight3(ROOK, CISSOR)
	if i != 1{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, 1)

	}

	i = fight3(CISSOR, PAPER)
	if i != 1{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, 1)

	}

	i = fight3(PAPER, CISSOR)
	if i != -1{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, -1)

	}

	i = fight3(ROOK, PAPER)
	if i != -1{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, -1)

	}

	i = fight3(PAPER, ROOK)
	if i != 1{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, 1)

	}

	i = fight3(ROOK, ROOK)
	if i != 0{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, 0)

	}

	i = fight3(CISSOR, CISSOR)
	if i != 0{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, 0)

	}
	i = fight3(PAPER, PAPER)
	if i != 0{
		t.Errorf("fight incorrect, got: %d, want: %d.", i, 0)

	}
}
