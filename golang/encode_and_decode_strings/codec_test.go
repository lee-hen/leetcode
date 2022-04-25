package encode_and_decode_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCodec_Encode(t *testing.T) {
	dummyInput := []string{"Hello","World"}
	codeC := new(Codec)
	require.Equal(t, "[72 101 108 108 111],[87 111 114 108 100],", codeC.Encode(dummyInput))
}

func TestCodec_Decode(t *testing.T) {
	codeC := new(Codec)
	require.Equal(t, []string{"Hello","World"}, codeC.Decode("[72 101 108 108 111],[87 111 114 108 100],"))
}

func TestSpecialCase(t *testing.T) {
	codeC := new(Codec)
	require.Equal(t, []string{""}, codeC.Decode(codeC.Encode([]string{""})))
}
