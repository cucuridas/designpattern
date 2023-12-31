package adaptor

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)



type XmlMembers struct {
}

func MakeXmlObj() *XmlMembers {
	return &XmlMembers{}
}

func (xm *XmlMembers)ConvertByte(path string) *Members {
	fp, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

	data, err := ioutil.ReadAll(fp)

	var members Members
	 
	xmlValue := xml.Unmarshal(data, &members)

	if xmlValue != nil {
		panic("Erorr")
	}
	
	fmt.Println(members)

	return &members

}	



func (xm *XmlMembers)RoadObject(datas *Members){
	const path = "./adaptor/convertMembers.xml"
	
	fp, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	xmlVluae,xmlErr := xml.Marshal(datas)

	if xmlErr != nil {
		panic(xmlErr)
	}



	ioutil.WriteFile(path,xmlVluae, 0666)

}
