package ast


const _blankCharsMask = (1 << ' ') | (1 << '\t') | (1 << '\r') | (1 << '\n')

const (
    bytesNull   = "null"
    bytesTrue   = "true"
    bytesFalse  = "false"
    bytesEmptyObject = "{}"
    bytesArray  = "[]"
)

func isSpace(c byte) bool {
    return (int(1<<c) & _blankCharsMask) != 0
}

// UnmarshalJSON is just an adapter to json.Unmarshaler.
// If you want better performance, use Searcher.GetByPath() directly
func (self *Node) UnmarshalJSON(data []byte) (err error) {
	node, err := parseLazy(string(data), nil)
	if err != nil {
		return err
	}
	*self = node
    return nil
}
