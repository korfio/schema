package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
	"strings"
	"os"
	"unicode"
)

var f *os.File
var graph []interface{}
var properties map[string]interface{}
var classes map[string]interface{}
var root *schemaItem

type schemaItem struct {
	ID         string
	Type       string
	Comment    string
	Label      string
	Properties map[string]*schemaItem
	Children   map[string]*schemaItem
	Parents    []string
	Range      []string
}

func gatherSources() {

	sources := []string{
		"https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/releases/3.3/schema.jsonld",
		"https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/releases/3.3/ext-attic.jsonld",
		"https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/releases/3.3/ext-auto.jsonld",
		"https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/releases/3.3/ext-bib.jsonld",
		"https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/releases/3.3/ext-health-lifesci.jsonld",
		"https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/releases/3.3/ext-meta.jsonld",
		"https://raw.githubusercontent.com/schemaorg/schemaorg/master/data/releases/3.3/ext-pending.jsonld",
	}

	for _, source := range sources {
		res, err := http.Get(source)
		if err != nil {
			log.Fatal(err)
		}
		schemaRaw, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		var s interface{}
		json.Unmarshal(schemaRaw, &s)

		graph = append(graph, s.(map[string]interface{})["@graph"].([]interface{})...)
	}
}

func firstPass() {
	classes = make(map[string]interface{}, len(graph))
	properties = make(map[string]interface{}, len(graph))
	for _, item := range graph {

		// Get the type.
		itemType, ok := item.(map[string]interface{})["@type"].(string)
		if !ok || itemType != "rdf:Property" {
			itemType = "rdfs:Class"
		}

		itemID, ok := item.(map[string]interface{})["@id"].(string)

		// Skip if not valid.
		if !ok {
			continue
		}

		if itemType == "rdfs:Class" {
			classes[itemID] = item
		} else {
			properties[itemID] = item
		}
	}
}

func processClasses() {
	for id, item := range classes {
		nID := simpleID(id)

		itemComment, _ := item.(map[string]interface{})["rdfs:comment"].(string)
		itemLabel, _ := item.(map[string]interface{})["rdfs:label"].(string)
		itemType, _ := item.(map[string]interface{})["@type"].(string)

		var subClassOf []string
		subClass, ok := item.(map[string]interface{})["rdfs:subClassOf"].(map[string]interface{})
		if ok {
			subClassOf = append(subClassOf, simpleID(subClass["@id"].(string)))
		} else {
			subClassInterface, okay := item.(map[string]interface{})["rdfs:subClassOf"].([]interface{})
			if okay {
				for _, subClassItem := range subClassInterface {
					subClassOf = append(subClassOf, simpleID(subClassItem.(map[string]interface{})["@id"].(string)))
				}
			}
		}

		root.Children[nID] = &schemaItem{
			ID:         nID,
			Type:       itemType,
			Comment:    itemComment,
			Label:      itemLabel,
			Properties: make(map[string]*schemaItem, 100),
			Parents:    subClassOf,
		}
	}
}

func processProperties() {
	for id, item := range properties {
		nID := simpleID(id)
		itemComment, _ := item.(map[string]interface{})["rdfs:comment"].(string)
		itemLabel, _ := item.(map[string]interface{})["rdfs:label"].(string)

		var validValues []string
		rangeRaw, ok := item.(map[string]interface{})["http://schema.org/rangeIncludes"].([]interface{})
		if ok {
			validValues = make([]string, len(rangeRaw))
			for _, valueRaw := range rangeRaw {
				value, okay := valueRaw.(map[string]interface{})["@id"].(string)
				if ! okay {
					value = ""
				}
				value = simpleID(value)
				if value != "" {
					validValues = append(validValues, value)
				}
			}
			validValues = deleteEmpty(validValues)
		} else {
			rangeMapRaw, ok := item.(map[string]interface{})["http://schema.org/rangeIncludes"].(map[string]interface{})
			if ok {
				validValues = make([]string, len(rangeMapRaw))
				for _, valueRaw := range rangeMapRaw {
					value, okay := valueRaw.(string)
					if ! okay {
						value = ""
					}
					value = simpleID(value)
					if value != "" {
						validValues = append(validValues, value)
					}
				}
				validValues = deleteEmpty(validValues)
			}
		}

		prop := &schemaItem{
			ID:      nID,
			Type:    "Property",
			Comment: itemComment,
			Label:   itemLabel,
			Range:   validValues,
		}

		domainsRaw, ok := item.(map[string]interface{})["http://schema.org/domainIncludes"].([]interface{})
		if ok {
			for _, domainRaw := range domainsRaw {
				domain, _ := domainRaw.(map[string]interface{})["@id"].(string)
				domain = simpleID(domain)
				if root.Children[domain] != nil && root.Children[domain].Properties != nil {
					root.Children[domain].Properties[nID] = prop
				}
			}
		} else {
			domainsMapRaw, ok := item.(map[string]interface{})["http://schema.org/domainIncludes"].(map[string]interface{})
			if ok {
				validValues = make([]string, len(domainsMapRaw))
				for _, domainRaw := range domainsMapRaw {
					domain, _ := domainRaw.(string)
					domain = simpleID(domain)
					if root.Children[domain] != nil && root.Children[domain].Properties != nil {
						root.Children[domain].Properties[nID] = prop
					}
				}
			}
		}
	}
}

func outputSchema(leaf schemaItem) {

	structDefStart := fmt.Sprintf("type %s struct {\n", leaf.ID)
	structDefEnd := "}\n\n"

	switch leaf.ID {
	case "root": fallthrough
	case "Text": fallthrough
	case "Number": fallthrough
	case "DataType": fallthrough
	case "Time": fallthrough
	case "DateTime": fallthrough
	case "Date": fallthrough
	case "Boolean":
		goto doChildren
	}

	// @todo How do we handle these?
	if strings.Contains(structDefStart, "#") || strings.Contains(structDefStart, ".") {
		goto doChildren
	}

	f.WriteString(structDefStart)

	for _, item := range leaf.Properties {
		itemType := "interface{}"
		itemLabel := item.Label

		if itemLabel == "" {
			itemLabel = lA(item.ID)
		}

		types := strings.Join(item.Range, ", ")

		if len(item.Range) == 1 {
			itemType = item.Range[0]
			f.WriteString(fmt.Sprintf("\t%s\t%s\t`json:\"%s,omitempty\"`\n", item.ID, itemType, itemLabel))
		} else {
			f.WriteString(fmt.Sprintf("\t%s\t%s\t`json:\"%s,omitempty\"`\t// %s\n", item.ID, itemType, itemLabel, types))
		}
	}

	for _, parent := range leaf.Parents {
		if parent == "Rdfs:Class" {
			continue
		}

		f.WriteString(fmt.Sprintf("\t%s\n", parent))
	}

	if strings.ToLower(leaf.Type) != "rdfs:class" {
		f.WriteString(fmt.Sprintf("\t%s\n", simpleID(leaf.Type)))
	}

	if leaf.ID == "Number" {
		f.WriteString(fmt.Sprintf("\t%s\n", "int"))
		f.WriteString(fmt.Sprintf("\t%s\n", "float64"))
	}

	if leaf.ID == "Text" {
		f.WriteString(fmt.Sprintf("\t%s\n", "string"))
	}

	f.WriteString(structDefEnd)

doChildren:
	for _, item := range leaf.Children {
		outputSchema(*item)
	}
}

func main() {
	var err error
	f, err = os.Create("../../schema.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gatherSources()
	firstPass()
	root = &schemaItem{
		ID:       "root",
		Children: make(map[string]*schemaItem, len(classes)),
	}

	processClasses()
	processProperties()

	f.WriteString("package schema\n\n")
	f.WriteString("import \"time\"\n\n")
	f.WriteString("type Text string\n")
	f.WriteString("type Boolean bool\n")
	f.WriteString("type Number struct {\n")
	f.WriteString("\tint\n")
	f.WriteString("\tfloat64\n")
	f.WriteString("}\n")
	f.WriteString("type Time time.Time\n")
	f.WriteString("type DateTime time.Time\n")
	f.WriteString("type Date time.Time\n")
	f.WriteString("\n")

	outputSchema(*root)

	//js, _ := json.Marshal(root)
	//fmt.Println(string(js))
	f.Sync()
}

func simpleID(id string) string {
	return strings.Title(strings.Split(id, "/")[len(strings.Split(id, "/"))-1])
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func lA(s string) string {
	if s == "" {
		return s
	}
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
