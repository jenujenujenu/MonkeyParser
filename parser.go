package MonkeyParser

import "regexp"

func (d *Document) ParseElement(el string) {
	d.Elements = append(d.Elements, Parse(el))
}

func ParseRegex(str string) interface{} {
	headingExpr := regexp.MustCompile(`^(#+) (.*)$`)
	if m := headingExpr.FindStringSubmatch(str); m != nil {
		return Heading{
			Level:   len(m[1]),
			Content: m[2],
		}
	}
	return Paragraph{str}
}

func Parse(str string) interface{} {
	if str[0] == '#'{
		level :=1
		for str[level] =='#'{
			level++
		}
		return Heading{
			Level:     level,
			Content:   str[level+1:],
		}
	}
	return Paragraph{str}
}
