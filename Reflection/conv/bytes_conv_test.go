package conv

import (
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInt32Conv(t *testing.T) {
	v := int32(math.MinInt32)
	bytes := Int32ToByte(v)
	assert.Equal(t, v, ByteToInt32(bytes))
	v = math.MaxInt32
	bytes = Int32ToByte(v)
	assert.Equal(t, v, ByteToInt32(bytes))
}

func TestInt64Conv(t *testing.T) {
	v := int64(math.MinInt64)
	bytes := Int64ToByte(v)
	assert.Equal(t, v, ByteToInt64(bytes))
	v = math.MaxInt64
	bytes = Int64ToByte(v)
	assert.Equal(t, v, ByteToInt64(bytes))
}

func TestUint32Conv(t *testing.T) {
	v := uint32(0)
	bytes := Uint32ToByte(v)
	assert.Equal(t, v, ByteToUint32(bytes))
	v = math.MaxUint32
	bytes = Uint32ToByte(v)
	assert.Equal(t, v, ByteToUint32(bytes))
}

func TestUint64Conv(t *testing.T) {
	v := uint64(0)
	bytes := Uint64ToByte(v)
	assert.Equal(t, v, ByteToUint64(bytes))
	v = math.MaxUint64
	bytes = Uint64ToByte(v)
	assert.Equal(t, v, ByteToUint64(bytes))
}

func TestFloat64Conv(t *testing.T) {
	v := math.SmallestNonzeroFloat64
	bytes := Float64ToByte(v)
	assert.Equal(t, v, ByteToFloat64(bytes))
	v = math.MaxFloat64
	bytes = Float64ToByte(v)
	assert.Equal(t, v, ByteToFloat64(bytes))
}

func TestStringConv(t *testing.T) {
	v := `Paragraphs are the building blocks of papers. Many students define paragraphs in terms of length: a paragraph is a group of at least five sentences, a paragraph is half a page long, etc. In reality, though, the unity and coherence of ideas among sentences is what constitutes a paragraph. A paragraph is defined as “a group of sentences or a single sentence that forms a unit” (Lunsford and Connors 116). Length and appearance do not determine whether a section in a paper is a paragraph. For instance, in some styles of writing, particularly journalistic styles, a paragraph can be just one sentence long. Ultimately, a paragraph is a sentence or group of sentences that support one main idea. In this handout, we will refer to this as the “controlling idea,” because it controls what happens in the rest of the paragraph.`
	bytes := StringToByte(v)
	assert.Equal(t, v, ByteToString(bytes))
	v = ""
	bytes = StringToByte(v)
	assert.Equal(t, v, ByteToString(bytes))
}

type Config struct {
	s *Service
}


Third party Service looks something like this
type Service struct {
}
func NewService(...) (*Service) {
s := &Service{}
… // do some instantiation
return s;
}
func (s *Service) Operation(…) {
…
}

func NewConfig() {

}