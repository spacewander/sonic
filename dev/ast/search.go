package ast

type Searcher struct {
    parser Parser
}

func NewSearcher(src string) *Searcher {
    return &Searcher{
        parser: *NewParser(src),
    }
}

// GetByPathNoCopy search in depth from top json and returns a **Referenced** json node at the path location
//
// WARN: this search directly refer partial json from top json, which has faster speed,
// may consumes more memory.
func (self *Searcher) GetByPath(path ...interface{}) (Node, error) {
    return self.parser.getByPath(path...)
}

func Skip(json string, pos *int) (start int, err error ) {
	parser := NewParser(json)
	parser.pos = *pos

	start, err = parser.skip()
	*pos = parser.pos

    if err != nil {
		return -1, err
    }
	return start, nil
}
