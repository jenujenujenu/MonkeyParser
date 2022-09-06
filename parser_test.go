package MonkeyParser

import "testing"

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("#### Heading bingo bonga bango")
	}
}

func TestParse(t *testing.T) {
	testHeading:= Parse("## Heading").(Heading)

	if testHeading.Level != 2{
		t.Error("Level should be 2, is ", testHeading.Level )
	}
	if testHeading.Content != "Heading"{
		t.Error("Content should be 'Heading', is ", testHeading.Content)
	}

	testParagraph:= Parse("This is a paragraph for testing monkeys").(Paragraph)

	if testParagraph.Content !="This is a paragraph for testing monkeys"{
		t.Error("Content schould be 'This is a paragraph for testing monkeys', is ", testParagraph.Content)
	}
}

func BenchmarkParseRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseRegex("#### Heading bingo bonga bango")
	}
}