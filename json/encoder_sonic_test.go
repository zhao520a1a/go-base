package json

import (
	"bytes"
	"encoding"
	"encoding/json"
	"os"
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/bytedance/sonic/encoder"
	"github.com/stretchr/testify/require"

	"github.com/zhao520a1a/go-base/json/testdata"
)

const TwitterJson = testdata.TwitterJson

var _GenericValue interface{}
var _BindingValue testdata.TwitterStruct
var (
	debugSyncGC = os.Getenv("JSON_SYNC_GC") != ""
)

func init() {
	_ = json.Unmarshal([]byte(TwitterJson), &_GenericValue)
	_ = json.Unmarshal([]byte(TwitterJson), &_BindingValue)
}

func TestGC(t *testing.T) {
	if debugSyncGC {
		return
	}
	out, err := encoder.Encode(_GenericValue, 0)
	if err != nil {
		t.Fatal(err)
	}
	n := len(out)
	wg := &sync.WaitGroup{}
	N := 10000
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, size int) {
			defer wg.Done()
			out, err := encoder.Encode(_GenericValue, 0)
			if err != nil {
				t.Fatal(err)
			}
			if len(out) != size {
				t.Fatal(len(out), size)
			}
			runtime.GC()
		}(wg, n)
	}
	wg.Wait()
}

type sample struct {
	M  map[string]interface{}
	S  []interface{}
	A  [0]interface{}
	MP *map[string]interface{}
	SP *[]interface{}
	AP *[0]interface{}
}

func TestOptionSliceOrMapNoNull(t *testing.T) {
	obj := sample{}
	out, err := encoder.Encode(obj, encoder.NoNullSliceOrMap)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, `{"M":{},"S":[],"A":[],"MP":null,"SP":null,"AP":null}`, string(out))

	obj2 := sample{}
	out, err = encoder.Encode(obj2, 0)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, `{"M":null,"S":null,"A":[],"MP":null,"SP":null,"AP":null}`, string(out))
}

func BenchmarkOptionSliceOrMapNoNull(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		obj := sample{}
		_, err := encoder.Encode(obj, encoder.NoNullSliceOrMap)
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = encoder.Encode(obj, encoder.NoNullSliceOrMap)
		}
	})

	b.Run("false", func(b *testing.B) {
		obj2 := sample{}
		_, err := encoder.Encode(obj2, 0)
		if err != nil {
			b.Fatal(err)
		}
		for i := 0; i < b.N; i++ {
			_, _ = encoder.Encode(obj2, 0)
		}
	})
}

func runEncoderTest(t *testing.T, fn func(string) string, exp string, arg string) {
	require.Equal(t, exp, fn(arg))
}

func TestEncoder_String(t *testing.T) {
	runEncoderTest(t, encoder.Quote, `""`, "")
	runEncoderTest(t, encoder.Quote, `"hello, world"`, "hello, world")
	runEncoderTest(t, encoder.Quote, `"hello啊啊啊aa"`, "hello啊啊啊aa")
	runEncoderTest(t, encoder.Quote, `"hello\\\"world"`, "hello\\\"world")
	runEncoderTest(t, encoder.Quote, `"hello\n\tworld"`, "hello\n\tworld")
	runEncoderTest(t, encoder.Quote, `"hello\u0000\u0001world"`, "hello\x00\x01world")
	runEncoderTest(t, encoder.Quote, `"hello\u0000\u0001world"`, "hello\x00\x01world")
	runEncoderTest(t, encoder.Quote, `"Cartoonist, Illustrator, and T-Shirt connoisseur"`, "Cartoonist, Illustrator, and T-Shirt connoisseur")
}

type StringStruct struct {
	X *int        `json:"x,string,omitempty"`
	Y []int       `json:"y"`
	Z json.Number `json:"z,string"`
	W string      `json:"w,string"`
}

func TestEncoder_FieldStringize(t *testing.T) {
	x := 12345
	v := StringStruct{X: &x, Y: []int{1, 2, 3}, Z: "4567456", W: "asdf"}
	r, e := encoder.Encode(v, 0)
	require.NoError(t, e)
	println(string(r))
}

func TestEncodeErrorAndScratchBuf(t *testing.T) {
	var obj = map[string]interface{}{
		"a": json.RawMessage(" [} "),
	}
	buf := make([]byte, 0, 10)
	_ = encoder.EncodeInto(&buf, obj, 0)
	if len(buf) < 0 || len(buf) > 10 {
		println(buf)
		t.Fatal()
	}
}

type MarshalerImpl struct {
	X int
}

func (self *MarshalerImpl) MarshalJSON() ([]byte, error) {
	ret := []byte(strconv.Itoa(self.X))
	return append(ret, "    "...), nil
}

type MarshalerStruct struct {
	V MarshalerImpl
}

func TestEncoder_Marshaler(t *testing.T) {
	v := MarshalerStruct{V: MarshalerImpl{X: 12345}}
	ret, err := encoder.Encode(&v, 0)
	require.NoError(t, err)
	require.Equal(t, `{"V":12345    }`, string(ret))
	ret, err = encoder.Encode(v, 0)
	require.NoError(t, err)
	require.Equal(t, `{"V":{"X":12345}}`, string(ret))

	ret2, err2 := encoder.Encode(&v, 0)
	require.NoError(t, err2)
	require.Equal(t, `{"V":12345    }`, string(ret2))
	ret3, err3 := encoder.Encode(v, encoder.CompactMarshaler)
	require.NoError(t, err3)
	require.Equal(t, `{"V":{"X":12345}}`, string(ret3))
}

type MarshalerErrorStruct struct {
	V MarshalerImpl
}

func (self *MarshalerErrorStruct) MarshalJSON() ([]byte, error) {
	return []byte(`[""] {`), nil
}

func TestMarshalerError(t *testing.T) {
	v := MarshalerErrorStruct{}
	ret, err := encoder.Encode(&v, 0)
	require.EqualError(t, err, `invalid Marshaler output json syntax at 5: "[\"\"] {"`)
	require.Equal(t, []byte(nil), ret)
}

type RawMessageStruct struct {
	X json.RawMessage
}

func TestEncoder_RawMessage(t *testing.T) {
	rms := RawMessageStruct{
		X: json.RawMessage("123456    "),
	}
	ret, err := encoder.Encode(&rms, 0)
	require.NoError(t, err)
	require.Equal(t, `{"X":123456    }`, string(ret))

	ret, err = encoder.Encode(&rms, encoder.CompactMarshaler)
	require.NoError(t, err)
	require.Equal(t, `{"X":123456}`, string(ret))
}

type TextMarshalerImpl struct {
	X string
}

func (self *TextMarshalerImpl) MarshalText() ([]byte, error) {
	return []byte(self.X), nil
}

type TextMarshalerImplV struct {
	X string
}

func (self TextMarshalerImplV) MarshalText() ([]byte, error) {
	return []byte(self.X), nil
}

type TextMarshalerStruct struct {
	V TextMarshalerImpl
}

func TestEncoder_TextMarshaler(t *testing.T) {
	v := TextMarshalerStruct{V: TextMarshalerImpl{X: (`{"a"}`)}}
	ret, err := encoder.Encode(&v, 0)
	require.NoError(t, err)
	require.Equal(t, `{"V":"{\"a\"}"}`, string(ret))
	ret, err = encoder.Encode(v, 0)
	require.NoError(t, err)
	require.Equal(t, `{"V":{"X":"{\"a\"}"}}`, string(ret))

	ret2, err2 := encoder.Encode(&v, encoder.NoQuoteTextMarshaler)
	require.NoError(t, err2)
	require.Equal(t, `{"V":{"a"}}`, string(ret2))
	ret3, err3 := encoder.Encode(v, encoder.NoQuoteTextMarshaler)
	require.NoError(t, err3)
	require.Equal(t, `{"V":{"X":"{\"a\"}"}}`, string(ret3))
}

func TestTextMarshalTextKey_SortKeys(t *testing.T) {
	v := map[*TextMarshalerImpl]string{
		{"b"}: "b",
		{"c"}: "c",
		{"a"}: "a",
	}
	ret, err := encoder.Encode(v, encoder.SortMapKeys)
	require.NoError(t, err)
	require.Equal(t, `{"a":"a","b":"b","c":"c"}`, string(ret))

	v2 := map[TextMarshalerImplV]string{
		{"b"}: "b",
		{"c"}: "c",
		{"a"}: "a",
	}
	ret, err = encoder.Encode(v2, encoder.SortMapKeys)
	require.NoError(t, err)
	require.Equal(t, `{"a":"a","b":"b","c":"c"}`, string(ret))

	v3 := map[encoding.TextMarshaler]string{
		TextMarshalerImplV{"b"}: "b",
		&TextMarshalerImpl{"c"}: "c",
		TextMarshalerImplV{"a"}: "a",
	}
	ret, err = encoder.Encode(v3, encoder.SortMapKeys)
	require.NoError(t, err)
	require.Equal(t, `{"a":"a","b":"b","c":"c"}`, string(ret))
}

func TestEncoder_Marshal_EscapeHTML(t *testing.T) {
	v := map[string]TextMarshalerImpl{"&&": {"<>"}}
	ret, err := encoder.Encode(v, encoder.EscapeHTML)
	require.NoError(t, err)
	require.Equal(t, `{"\u0026\u0026":{"X":"\u003c\u003e"}}`, string(ret))
	ret, err = encoder.Encode(v, 0)
	require.NoError(t, err)
	require.Equal(t, `{"&&":{"X":"<>"}}`, string(ret))

	// “ is \xe2\x80\x9c, and ” is \xe2\x80\x9d,
	// similar as HTML escaped chars \u2028(\xe2\x80\xa8) and \u2029(\xe2\x80\xa9)
	m := map[string]string{"test": "“123”"}
	ret, err = encoder.Encode(m, encoder.EscapeHTML)
	require.Equal(t, string(ret), `{"test":"“123”"}`)
	require.NoError(t, err)

	m = map[string]string{"K": "\u2028\u2028\xe2"}
	ret, err = encoder.Encode(m, encoder.EscapeHTML)
	require.Equal(t, string(ret), "{\"K\":\"\\u2028\\u2028\xe2\"}")
	require.NoError(t, err)
}

func TestEncoder_EscapeHTML(t *testing.T) {
	// test data from libfuzzer
	test := []string{
		"&&&&&&&&&&&&&&&&&&&&&&&\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2\xe2&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&",
		"{\"\"\u2028\x94\xe2\x00\x00\x00\x00\x00\x00\x00\x00\u2028\x80\u2028\x80\u2028\xe2\u2028\x8a\u2028⑀\xa8\x8a\xa8\xe2\u2028\xe2\u2028\xe2\u2028\xe2\u2000\x8d\xe2\u2028\xe2\u2028\xe2\xe2\xa8\"}",
	}
	for _, s := range test {
		data := []byte(s)
		sdst := encoder.HTMLEscape(nil, data)
		var dst bytes.Buffer
		json.HTMLEscape(&dst, data)
		require.Equal(t, string(sdst), dst.String())
	}
}

func TestEncoder_Marshal_EscapeHTML_LargeJson(t *testing.T) {
	buf1, err1 := encoder.Encode(&_BindingValue, encoder.SortMapKeys|encoder.EscapeHTML)
	require.NoError(t, err1)
	buf2, err2 := json.Marshal(&_BindingValue)
	require.NoError(t, err2)
	require.Equal(t, buf1, buf2)
}

func TestEncoder_Generic(t *testing.T) {
	v, e := encoder.Encode(_GenericValue, 0)
	require.NoError(t, e)
	println(string(v))
}

func TestEncoder_Binding(t *testing.T) {
	v, e := encoder.Encode(_BindingValue, 0)
	require.NoError(t, e)
	println(string(v))
}

func TestEncoder_MapSortKey(t *testing.T) {
	m := map[string]string{
		"C": "third",
		"D": "forth",
		"A": "first",
		"F": "sixth",
		"E": "fifth",
		"B": "second",
	}
	v, e := encoder.Encode(m, encoder.SortMapKeys)
	require.NoError(t, e)
	require.Equal(t, `{"A":"first","B":"second","C":"third","D":"forth","E":"fifth","F":"sixth"}`, string(v))
}
