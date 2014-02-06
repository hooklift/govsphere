package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
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
	if xsdType == "" {
		return ""
	}

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

	if strings.HasSuffix(type_, "[]") {
		type_ = type_[:len(type_)-2]
		if value, ok := xsd2GoTypes[type_]; ok {
			type_ = "[]" + value
		} else {
			type_ = "[]*" + type_
		}
	} else {
		type_ = "*" + type_
	}

	return type_
}

func stripns(xsdType string) string {
	r := strings.Split(xsdType, ":")
	type_ := r[0]

	if len(r) == 2 {
		type_ = r[1]
	}

	return type_
}

func makePublic(field_ string, public bool) string {
	field := []rune(field_)
	if len(field) == 0 {
		return field_
	}

	if public {
		field[0] = unicode.ToUpper(field[0])
	} else {
		field[0] = unicode.ToLower(field[0])
	}
	return string(field)
}

func comment(text string) string {
	lines := strings.Split(text, "\n")

	var output string
	if len(lines) == 1 && lines[0] == "" {
		return ""
	}

	// Helps to determine if
	// there is an actual comment
	// without screwing newlines
	// in real comments.
	hasComment := false

	for _, line := range lines {
		line = strings.TrimLeftFunc(line, unicode.IsSpace)
		if line != "" {
			hasComment = true
		}
		output += "\n// " + line
	}

	if hasComment {
		return output
	}
	return ""
}

//This is how we look for the package
//or namespace associated to one particular
//type. This is needed because 4 packages
//are being created such as: mo, enum, do and faults
//in order to make the API more idiomatic
//for users to use. Once a type is found
//this map is going to be used to find its
//namespace or package.
var objnsmap map[string]string

func lookUpNamespace(type_, currentNs string) string {
	//Embeddeds or extends are often times empty
	if type_ == "" {
		return type_
	}

	var prefix string
	if type_[0:1] == "*" {
		prefix = "*"
		type_ = type_[1:]
	} else if type_[0:3] == "[]*" {
		prefix = "[]*"
		type_ = type_[3:]
	} else if type_[0:2] == "[]" {
		prefix = "[]"
		type_ = type_[2:]
	}

	//strips out package name if it has it
	type_ = strings.TrimPrefix(type_, "mo.")
	type_ = strings.TrimPrefix(type_, "do.")
	type_ = strings.TrimPrefix(type_, "fault.")
	type_ = strings.TrimPrefix(type_, "enum.")

	targetNs := objnsmap[type_]
	if targetNs == "" || targetNs != "enum" || targetNs == currentNs {
		return prefix + type_
	}

	return prefix + targetNs + "." + type_
}

var funcMap = template.FuncMap{
	"toGoType":             toGoType,
	"stripns":              stripns,
	"replaceReservedWords": replaceReservedWords,
	"makePublic":           makePublic,
	"comment":              comment,
	"lookUpNamespace":      lookUpNamespace,
}

func generate(apiDefFile string) {
	apiDef, err := ioutil.ReadFile(apiDefFile)
	if err != nil {
		log.Fatalln(err)
	}

	var objects []Object
	err = json.Unmarshal(apiDef, &objects)
	if err != nil {
		log.Fatalln(err)
	}

	//Populates objnsmap
	objnsmap = make(map[string]string)
	for _, obj := range objects {
		objnsmap[obj.Name] = obj.Namespace
	}

	mainPkg := "vim"
	os.Mkdir(mainPkg, 0744)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		genCode(objects, mainPkg, doTmpl, "do")
		genCode(objects, mainPkg, moTmpl, "mo")
		genCode(objects, mainPkg, faultTmpl, "fault")
	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	genCode(objects, mainPkg, doTmpl, "do")
	// }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		genCode(objects, mainPkg, enumTmpl, "enum")
	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	genCode(objects, mainPkg, faultTmpl, "fault")
	// }()

	wg.Wait()
}

func genCode(objects []Object, mainPkg, tmpl, namespace string) {
	var fd *os.File
	pkg := mainPkg

	fileName := namespace
	if namespace != "enum" {
		fileName = pkg
	} else {
		pkg += "/" + namespace
	}

	if ok, err := exists(pkg); !ok && err == nil {
		os.Mkdir(pkg, 0744)
	}

	file := pkg + "/" + fileName + ".go"

	fd, err := os.Create(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer fd.Close()

	data := new(bytes.Buffer)
	data.WriteString(headerTmpl)

	if namespace != "enum" {
		data.WriteString("package " + mainPkg + "\n")
		data.WriteString(`
			import (
				"time"
				"github.com/c4milo/govsphere/vim/enum"
			)
		`)
	} else {
		data.WriteString("package " + namespace + "\n")
	}

	for _, obj := range objects {
		if namespace == "enum" {
			if obj.Namespace != namespace {
				continue
			}
		}

		log.Println("type: " + obj.Name)
		tmpl := template.Must(template.New(obj.Namespace).Funcs(funcMap).Parse(tmpl))
		err = tmpl.Execute(data, obj)
		if err != nil {
			log.Fatalln(err)
		}
	}

	source := data.Bytes()
	fsource, err := format.Source(source)
	if err != nil {
		fd.Write(source)
		log.Fatalf("There are errors in the generated source for %s: %s\n", file, err.Error())
	}
	fd.Write(fsource)

	log.Println("Done 💩")
}
