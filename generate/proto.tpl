syntax = "proto3";

option go_package = "./pb";

message Empty {}

{{range $i, $v := .Message}}{{if ne $v.Comment ""}}
{{$v.Comment}}{{end}}
message {{$v.Name}} { {{range $j, $v2 := $v.Item}}
    {{$v2.Type}} {{$v2.Name}} = {{add $j 1}};{{if ne $v2.Comment ""}}  //{{$v2.Comment}}{{end}}{{end}}
}
{{end}}

{{if ne .Service.Comment ""}}//{{.Service.Comment}}{{end}}
service {{.Service.Name}} {
{{range $i, $item := .Service.Item}}{{if ne $item.Comment ""}}
    //{{$item.Comment}}{{end}}
     rpc {{$item.FuncName}}({{$item.RequestName}}) returns ({{if eq $item.ResponseName ""}}Empty{{ else }}{{$item.ResponseName}}{{end}});
{{end}}
}