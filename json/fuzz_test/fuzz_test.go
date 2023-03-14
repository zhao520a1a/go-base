package fuzz_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/bytedance/sonic"
)

func FuzzUnmarshalJSON_StdLib(f *testing.F) {
	f.Add([]byte(`{
					"object": {
						"slice": [
							1,
							2.0,
							"3",
							[4],
							{5: {}}
						]
					},
					"slice": [[]],
					"string": ":)",
					"int": 1e5,
					"float": 3e-9"
					}`))

	f.Fuzz(func(t *testing.T, b []byte) {
		for _, typ := range []func() interface{}{
			func() interface{} { return new(interface{}) },
			func() interface{} { return new(map[string]interface{}) },
			func() interface{} { return new([]interface{}) },
		} {
			i := typ()
			if err := json.Unmarshal(b, i); err != nil {
				return
			}

			encoded, err := json.Marshal(i)
			if err != nil {
				t.Fatalf("failed to marshal: %s", err)
			}

			if err := json.Unmarshal(encoded, i); err != nil {
				t.Fatalf("failed to roundtrip: %s", err)
			}
		}
	})
}

func FuzzDecoderToken_StdLib(f *testing.F) {
	f.Add([]byte(`{
					"object": {
						"slice": [
							1,
							2.0,
							"3",
							[4],
							{5: {}}
						]
					},
					"slice": [[]],
					"string": ":)",
					"int": 1e5,
					"float": 3e-9"
					}`))

	f.Fuzz(func(t *testing.T, b []byte) {
		r := bytes.NewReader(b)
		d := json.NewDecoder(r)
		for {
			_, err := d.Token()
			if err != nil {
				if err == io.EOF {
					break
				}
				return
			}
		}
	})
}

func FuzzUnmarshalJSON_Sonic(f *testing.F) {
	f.Add([]byte(`{
					"object": {
						"slice": [
							1,
							2.0,
							"3",
							[4],
							{5: {}}
						]
					},
					"slice": [[]],
					"string": ":)",
					"int": 1e5,
					"float": 3e-9"
					}`))

	f.Fuzz(func(t *testing.T, b []byte) {
		for _, typ := range []func() interface{}{
			func() interface{} { return new(interface{}) },
			func() interface{} { return new(map[string]interface{}) },
			func() interface{} { return new([]interface{}) },
		} {
			i := typ()
			if err := sonic.Unmarshal(b, i); err != nil {
				return
			}

			encoded, err := sonic.Marshal(i)
			if err != nil {
				t.Fatalf("failed to marshal: %s", err)
			}

			if err := sonic.Unmarshal(encoded, i); err != nil {
				t.Fatalf("failed to roundtrip: %s", err)
			}
		}
	})
}
