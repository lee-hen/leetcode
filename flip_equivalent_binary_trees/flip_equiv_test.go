package flip_equivalent_binary_trees

import (
	SD "github.com/lee-hen/leecode/serialize_and_deserialize_binary_tree"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlipEquiv(t *testing.T) {
    sd := SD.Constructor()
	root1 := sd.Deserialize("1,2,3,4,5,6,N,N,N,7,8")
	root2 := sd.Deserialize("1,3,2,N,6,4,5,N,N,N,N,8,7")
	require.Equal(t, flipEquiv(root1, root2), FlipEquiv(root1, root2))
}
