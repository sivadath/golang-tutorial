package HTTP

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HitCaseConverterServer(content, ip, port string) {
		resp, err := http.Get( "http://" + ip + ":" + port + "/" + content )
		if err != nil {
			log.Println("err",err.Error())
		}else {
			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body);err == nil {
				log.Println("CaseConverter response:\n",string(body))
			}
		}
}
