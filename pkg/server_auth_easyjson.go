// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package pkg

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson44dda824DecodeGithubComCoinsurfComProxyPkg(in *jlexer.Lexer, out *authMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "api_key":
			out.APIKey = string(in.String())
		case "user_agent":
			out.UserAgent = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson44dda824EncodeGithubComCoinsurfComProxyPkg(out *jwriter.Writer, in authMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"api_key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.APIKey))
	}
	{
		const prefix string = ",\"user_agent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserAgent))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v authMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson44dda824EncodeGithubComCoinsurfComProxyPkg(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v authMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson44dda824EncodeGithubComCoinsurfComProxyPkg(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *authMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson44dda824DecodeGithubComCoinsurfComProxyPkg(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *authMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson44dda824DecodeGithubComCoinsurfComProxyPkg(l, v)
}
