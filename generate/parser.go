package generate

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"reflect"
	"strings"
)

func applyGenerate(p *plugin.Plugin) ([]protoMessage, protoService, error) {
	message := renderReplyAsDefinition(p.Api.Types)
	service := renderReplyAsService(p.Api.Service)
	return message, service, nil
}

func renderReplyAsDefinition(types []spec.Type) []protoMessage {
	protoMessages := make([]protoMessage, 0)

	for _, tp := range types {
		schema := protoMessage{}
		defineStruct, _ := tp.(spec.DefineStruct)

		schema.Name = defineStruct.Name()
		schema.Comment = strings.Join(defineStruct.Documents(), " ")

		for _, member := range defineStruct.Members {
			if hasPathParameters(member) {
				continue
			}

			messageM := schemaOfField(member)
			if tag, err := member.GetPropertyName(); err == nil {
				messageM.Name = tag
			}

			//首字母小写
			if len(messageM.Name) > 0 {
				messageM.Name = strings.TrimLeft(messageM.Name, "*")
				messageM.Name = strings.ToLower(string(messageM.Name[0])) + messageM.Name[1:]
			}

			schema.Item = append(schema.Item, messageM)
		}
		protoMessages = append(protoMessages, schema)
	}
	return protoMessages
}

func renderReplyAsService(p spec.Service) protoService {
	protoS := protoService{
		Name:    p.Name,
		Comment: "",
		Item:    make([]serviceMember, 0),
	}
	for _, group := range p.Groups {

		for _, route := range group.Routes {

			if len(route.Handler) > 0 {
				route.Handler = strings.ToUpper(string(route.Handler[0])) + route.Handler[1:]
			}
			schema := serviceMember{
				FuncName:     group.GetAnnotation("group") + route.Handler,
				RequestName:  route.RequestTypeName(),
				ResponseName: route.ResponseTypeName(),
				Comment:      route.JoinedDoc(),
			}

			protoS.Item = append(protoS.Item, schema)
		}

	}
	return protoS
}

func hasPathParameters(member spec.Member) bool {
	for _, tag := range member.Tags() {
		if tag.Key == "path" {
			return true
		}
	}

	return false
}
func schemaOfField(member spec.Member) messageMember {

	kind := protoMapTypes[member.Type.Name()]

	item := messageMember{
		Type:    "string",
		Name:    "",
		Comment: "",
	}

	comment := member.GetComment()
	item.Comment = strings.Replace(comment, "//", "", -1)

	switch ft := kind; ft {
	case reflect.Invalid: //[]Struct 也有可能是 Struct
		// []Struct
		// map[ArrayType:map[Star:map[StringExpr:UserSearchReq] StringExpr:*UserSearchReq] StringExpr:[]*UserSearchReq]
		refTypeName := strings.Replace(member.Type.Name(), "[", "", 1)
		refTypeName = strings.Replace(refTypeName, "]", "", 1)
		refTypeName = strings.Replace(refTypeName, "*", "", 1)
		refTypeName = strings.Replace(refTypeName, "{", "", 1)
		refTypeName = strings.Replace(refTypeName, "}", "", 1)
		// interface
		item.Name = refTypeName

		ftype, ok := primitiveSchema(ft, member.Type.Name())
		if ok {
			item.Type = ftype
			item.Name = refTypeName
		}
	case reflect.Slice:
		tempKind := protoMapTypes[member.Type.Name()]
		ftype, ok := primitiveSchema(tempKind, member.Type.Name())

		if ok {
			item.Type = ftype
		}
	default:
		ftype, ok := primitiveSchema(ft, member.Type.Name())
		if ok {
			item.Type = ftype
		}
	}

	for _, tag := range member.Tags() {
		if tag.Key != "json" {
			continue
		}
		item.Name = tag.Name
	}

	return item
}

// https://swagger.io/specification/ Data Types
func primitiveSchema(kind reflect.Kind, t string) (ftype string, ok bool) {
	switch kind {
	case reflect.Int8:
		return "int32", true
	case reflect.Uint8:
		return "int32", true
	case reflect.Int16:
		return "int32", true
	case reflect.Uint16:
		return "int32", true
	case reflect.Int32:
		return "int32", true
	case reflect.Int64:
		return "int64", true
	case reflect.Uint32:
		return "uint32", true
	case reflect.Uint64:
		return "uint64", true
	case reflect.Bool:
		return "bool", true
	case reflect.String:
		return "string", true
	case reflect.Float32:
		return "float", true
	case reflect.Float64:
		return "double", true
	case reflect.Slice:
		replace := strings.Replace(t, "[]", "", -1)
		ftype, ok = primitiveSchema(protoMapTypes[replace], replace)
		if ok {
			return "repeated " + ftype, true
		}
		return "repeated " + t, true
	case reflect.Struct:
		return t, true
	case reflect.Invalid:
		replace := strings.Replace(t, "[]", "", -1)
		if strings.HasPrefix(t, "[]") {
			return "repeated " + replace, true
		}
		return t, true
	default:
		return "", false
	}
}
