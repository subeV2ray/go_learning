package _map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2}
	t.Log(m1["two"])

	m2 := map[string]int{}
	m2["one"] = 1
	t.Log(m2["one"])

	m3 := make(map[string]int, 5)
	t.Logf("len m3 = %d", len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[1] = 1

	if v, ok := m1[1]; ok {
		t.Log("key:", v)
	} else {
		t.Log("not exist key")
	}

	t.Log(m1[1])
}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"1": 1, "2": 2}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
