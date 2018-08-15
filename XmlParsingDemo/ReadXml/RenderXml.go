package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"fmt"
	"io"
	"github.com/pkg/errors"
)

var (
	ElementNotFound = errors.New("Element not found")
)


// a simple struct which contains all our
// social links
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

// the user struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

// our struct which contains the complete
// array of all Users in the file
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

func main() {

	// this is SAX parsing. once for all
	xmlFile, err := os.Open("XmlParsingDemo/ReadXml/users.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	data, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		fmt.Println(err)
	}
	var users Users

	xml.Unmarshal(data, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}


	// this is dom based parsing
	xmlFile1, err := os.Open("XmlParsingDemo/ReadXml/users.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile1.Close()

	d := xml.NewDecoder(xmlFile1)

	elem, err := GetNodeByAttrValue(d, "type", "admin")

	var b User
	if err == nil {
		d.DecodeElement(&b, &elem)
		fmt.Println(b.Name)
	} else {
		fmt.Println("error getting element with attr admin")
	}

	elem, err = GetNodeByAttrValue(d, "type", "admin")
	if err == nil {
		d.DecodeElement(&b, &elem)
		fmt.Println(b.Name)
	} else {
		fmt.Println("error getting element with attr admin")
	}

	for {
		t, tokenErr := d.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			} else {
				fmt.Println(tokenErr)
			}
		}
		switch t := t.(type) {
		case xml.StartElement:
			if t.Name.Local == "user" {
				var b User
				if err := d.DecodeElement(&b, &t); err != nil {
					// handle error
				}
				fmt.Println(b.Name)
			}
		}
	}
}

// get's first match
func GetNodeByAttrValue(d *xml.Decoder, attr string, value string) (xml.StartElement, error) {
	for {
		t, tokenErr := d.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			} else {
				fmt.Println(tokenErr)
				return xml.StartElement{}, tokenErr
			}
		}
		switch se := t.(type) {
		case xml.StartElement:
			attrs := se.Attr
			for i:=0; i< len(attrs); i++ {
				a := attrs[i]
				if a.Name.Local == attr && a.Value == value {
					return se, nil
				}
			}
		}
	}
	return xml.StartElement{}, ElementNotFound
}