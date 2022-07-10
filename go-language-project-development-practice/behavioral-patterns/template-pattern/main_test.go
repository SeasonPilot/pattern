package template_pattern

import "testing"

func TestBaseCooker_Cooke(t *testing.T) {
	m := MyCooker{}
	DoCook(m)

	chaojidan := ChaoJiDan{}
	DoCook(chaojidan)
}
