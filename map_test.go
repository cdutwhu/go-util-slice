package wrappers

import "testing"

func TestMap(t *testing.T) {
	m := map[string]string{"a": "b", "c": "d", "e": "f"}
	fPln(GetMapKeys(m).([]string))
	m1 := map[int]string{1: "B", 2: "D", 3: "F"}
	fPln(GetMapKeys(m1).([]int))
	fPln(GetMapKeys(EnCoDesc))

	fPln(" ---------------------------------------------- ")

	k, v := GetMapKVs(m)
	K, V := k.([]string), v.([]string)
	fPln(K)
	fPln(V)

	fPln(" ---------------------------------------------- ")

	k, v = GetMapKVs(m1)
	K1, V1 := k.([]int), v.([]string)
	fPln(K1)
	fPln(V1)

	fPln(" ---------------------------------------------- ")

	// m2 := map[string]string{"aa": "bb", "cc": "dd", "ee": "ff"}
	m3 := map[int]string{4: "BB", 5: "DD", 1: "FF"}
	m02 := MapsJoin(m1, m3).(map[int]string)
	fPln(m02)
	for k, v := range m02 {
		fPln(k, v)
	}

}
