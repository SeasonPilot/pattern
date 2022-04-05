package main

type StrBuilder struct {
	result string
}

func (s *StrBuilder) Part1() {
	s.result += "1"
}

func (s *StrBuilder) Part2() {
	s.result += "2"

}

func (s *StrBuilder) Part3() {
	s.result += "3"
}

func (s *StrBuilder) GetResult() string {
	return s.result
}
