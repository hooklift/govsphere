package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
	"unicode"
)

var reservedWords = map[string]string{
	"break":       "break_",
	"default":     "default_",
	"func":        "func_",
	"interface":   "interface_",
	"select":      "select_",
	"case":        "case_",
	"defer":       "defer_",
	"go":          "go_",
	"map":         "map_",
	"struct":      "struct_",
	"chan":        "chan_",
	"else":        "else_",
	"goto":        "goto_",
	"package":     "package_",
	"switch":      "switch_",
	"const":       "const_",
	"fallthrough": "fallthrough_",
	"if":          "if_",
	"range":       "range_",
	"type":        "type_",
	"continue":    "continue_",
	"for":         "for_",
	"import":      "import_",
	"return":      "return_",
	"var":         "var_",
}

func replaceReservedWords(identifier string) string {
	value := reservedWords[identifier]
	if value != "" {
		return value
	}
	return identifier
}

var xsd2GoTypes = map[string]string{
	"string":        "string",
	"token":         "string",
	"float":         "float32",
	"double":        "float64",
	"decimal":       "float64",
	"integer":       "int32",
	"int":           "int32",
	"short":         "int16",
	"byte":          "int8",
	"long":          "int64",
	"boolean":       "bool",
	"dateTime":      "time.Time",
	"date":          "time.Time",
	"time":          "time.Time",
	"base64Binary":  "[]byte",
	"hexBinary":     "[]byte",
	"unsignedInt":   "uint32",
	"unsignedShort": "uint16",
	"unsignedByte":  "byte",
	"unsignedLong":  "uint64",
	"anyType":       "interface{}",
}

func toGoType(xsdType string) string {
	//Handles name space, ie. xsd:string, xs:long
	r := strings.Split(xsdType, ":")

	type_ := r[0]

	if len(r) == 2 {
		type_ = r[1]
	}

	value := xsd2GoTypes[type_]

	if value != "" {
		return value
	}

	return "*" + type_
}

func stripns(xsdType string) string {
	r := strings.Split(xsdType, ":")
	type_ := r[0]

	if len(r) == 2 {
		type_ = r[1]
	}

	return type_
}

func makePublic(field_ string) string {
	field := []rune(field_)
	if len(field) == 0 {
		return field_
	}

	field[0] = unicode.ToUpper(field[0])
	return string(field)
}

func generate(apiDefFile string) {
	funcMap := template.FuncMap{
		"toGoType":             toGoType,
		"stripns":              stripns,
		"replaceReservedWords": replaceReservedWords,
		"makePublic":           makePublic,
	}

	apiDef, err := ioutil.ReadFile(apiDefFile)
	if err != nil {
		log.Fatalln(err)
	}

	//objects := make([]Object, 1)
	var objects []Object
	err = json.Unmarshal(apiDef, objects)
	if err != nil {
		log.Fatalln(err)
	}

	data := new(bytes.Buffer)
	tmpl := template.Must(template.New("api").Funcs(funcMap).Parse(apiTmpl))

	err = tmpl.Execute(data, objects)
	if err != nil {
		log.Fatalln(err)
	}

	source := data.Bytes()
	source, err = format.Source(source)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(source)
}
