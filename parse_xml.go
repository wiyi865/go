/*
golang中使用encoding/xml库解析xml文件
/tmp/1.xml
<?xml version="1.0" encoding="utf-8"?>
<server_config>
	<logpath>log</logpath>
	<eventpath>/home/datacenter/td_data/unieco/</eventpath>
	<common_redis_addr>10.10.98.99:6379</common_redis_addr>
</server_config>
*/
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	conf := xmlconfig{}
	fd, err := os.Open("/tmp/1.xml")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(content, &conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf.Logpath)
	fmt.Println(conf)
}

type xmlconfig struct {
	XMLName    string `xml:"server_config"`
	Logpath    string `xml:"logpath"`
	Eventpath  string `xml:"eventpath"`
	Redis_addr string `xml:"common_redis_addr"`
}
