package control

import "testing"

func TestControl(t *testing.T) {
	loadData("/Users/zdns/Desktop/Hulk", "param.json")
}

func TestRun(t *testing.T) {
	Run("/Users/zdns/Desktop/Hulk", "param.json")
}
