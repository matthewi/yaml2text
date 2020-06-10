SELECT{{ $length := len .fields}}{{range $i, $v := .fields }}
    {{$v.sql}} AS {{$v.name}}{{if lt $i (sub $length 1)}},{{end}}{{end}}
FROM {{ .table }}
