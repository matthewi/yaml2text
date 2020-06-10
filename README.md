# yaml2text

This is YAML converter using golang template.

![](https://github.com/mashiike/yaml2text/workflows/Test/badge.svg)

I was tired of copying and pasting many similar files.
Do not use it to generate text files (or any code) that requires more advanced parsing.

## Simple usecase.
yaml file convert to any text file.

for example. following yaml file convert to sql file with go template file.
(testdata/basic.yaml)
```yaml
fields:
  - name: hoge
    sql: COALESCE(SUM(hoge), 0)
  - name: fuga
    sql: ARRAY_AGG(fuga)
  - name: piyo
    sql: AVG(piyo)

table: yamlyaml
```

go template file is this.
(testdata/basic.tpl)
```
SELECT{{ $length := len .fields}}{{range $i, $v := .fields }}
    {{$v.sql}} AS {{$v.name}}{{if lt $i (sub $length 1)}},{{end}}{{end}}
FROM {{ .table }}
```

and, execute following comand.
```shell
$ yaml2text -template testdata/basic.tpl testdata/basic.yaml
SELECT
    COALESCE(SUM(hoge), 0) AS hoge,
    ARRAY_AGG(fuga) AS fuga,
    AVG(piyo) AS piyo
FROM yamlyaml
```

It is convenient to use Makefile and shell redirection together.

## Install

### Homebrew (macOS only)

```
$ brew install mashiike/tap/mysqlbatch
```


### Binary packages

[Releases](https://github.com/mashiike/mysqlbatch/releases)
