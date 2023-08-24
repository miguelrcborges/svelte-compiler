package lexer

import (
	"bytes"
)

func parseJs(b []byte, i, l int) (n_i, n_l int) { 
	for ; i < len(b); i++ {
		if r, s := isNewline(b, i); r {
			i += s
			l++
		} else if b[i] == '<' && bytes.Equal(b[i:i+9], []byte("</script>")) {
			return i + 8, l
		}
	}
	return i, l
}
