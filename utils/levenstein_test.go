package utils

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	var a, b string
	var dist, trueDist int

	a = "kitten"
	b = "sitting"
	dist = LevenshteinDistance(&a, &b)
	trueDist = 3
	if dist != trueDist {
		t.Fatalf("Wrong distance: %d instead of %d", dist, trueDist)
	}

	a = "Среднее содержание hb в эритроците"
	b = "Среднее содержание Hb в эритроците"
	dist = LevenshteinDistance(&a, &b)
	trueDist = 1
	if dist != trueDist {
		t.Fatalf("Wrong distance: %d instead of %d", dist, trueDist)
	}

	a = "Среднее содержание НЬ в эритроците"
	b = "Среднее содержание Hb в эритроците"
	dist = LevenshteinDistance(&a, &b)
	trueDist = 2
	if dist != trueDist {
		t.Fatalf("Wrong distance: %d instead of %d", dist, trueDist)
	}

	a = "Среднее содержание НЬ в эритроците"
	b = "Среднее содержание Hb в эритроците hjdhэ78\n\t\r"
	dist = LevenshteinDistance(&a, &b)
	trueDist = 13
	if dist != trueDist {
		t.Fatalf("Wrong distance: %d instead of %d", dist, trueDist)
	}

	a = "Миелоциты"
	b = "Незрелые гранулоциты \t0.01 \t10*9/л \to - 0.09 \t\r"
	dist = LevenshteinDistance(&a, &b)
	trueDist = 40
	if dist != trueDist {
		t.Fatalf("Wrong distance: %d instead of %d", dist, trueDist)
	}
}
