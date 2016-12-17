package main

import (
	log "github.com/Sirupsen/logrus"
	"strconv"
	"regexp"
	"github.com/opesun/goquery"
	"time"
)

const (
	SAVE_PATH = "kproxy.orz"
	PROXY_URL = "http://www.kuaidaili.com/free/inha/"
)
var (
	IP_REGEXP = regexp.MustCompile(`[\d]+\.[\d]+\.[\d]+\.[\d]+\n\s+[\d]+`)
	IP_DETAIL_REGEXP = regexp.MustCompile(`[\d]+\.[\d]+\.[\d]+\.[\d]+`)
	INT_REGEXP = regexp.MustCompile(`\s[\d]+`)
)

func UrlGetter(num int) string {
	return PROXY_URL + strconv.Itoa(num)
}

func GetProxy(Url string) {
	nod, err := goquery.ParseUrl(Url)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	ret := nod.Text()
	ips := IP_REGEXP.FindAll([]byte(ret), -1)
	var port []string = make([]string, len(ips))
	var str string = ""
	for i := 0; i < len(ips); i++ {
		port[i] = string(INT_REGEXP.FindAll(ips[i], -1)[0])[1:]
		ips[i] = IP_DETAIL_REGEXP.FindAll(ips[i], -1)[0]
		str += string(ips[i])+":"+port[i]+"\n"
	}
	AppendFile("./", SAVE_PATH, str)
}

func main() {
	log.Infoln("Start getting proxy ...")
	SaveFile("./", SAVE_PATH, "")
	for i := 1; i <= 500; i++ {
		log.Println(UrlGetter(i))
		GetProxy(UrlGetter(i))
		time.Sleep(time.Second*5)
	}
}
