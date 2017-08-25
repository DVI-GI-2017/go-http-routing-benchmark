package parsetype

import (
	"github.com/kr/pretty"
	"os"
	"path"
	"testing"
)

func TestParse(t *testing.T) {
	p, _ := os.Getwd()
	dir := path.Join(p, "testparse")
	pretty.Println(dir)
	parser := NewParser()
	parser.ParseDir(dir)
	pretty.Println(parser.Types)
	// parser.ParseDir("/Users/witooh/dev/go/src/github.com/hyperworks/langfight/src/models")
}
