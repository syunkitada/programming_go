// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  main

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( MyData ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* MyData ) UnmarshalJSON([]byte) error { return nil }
func ( MyData ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* MyData ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_MyData *MyData
