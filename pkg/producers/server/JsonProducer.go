package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type JsonProducer struct {
	OutTemplate *template.Template
}

func (c *JsonProducer) Close() {
	// no need to close
}

func (c *JsonProducer) Produce(key []byte, value []byte, o interface{}) {
	var outBuffer bytes.Buffer
	var err error

	data := struct {
		K string
		V string
	}{string(key), string(value)}

	if err = c.OutTemplate.Execute(&outBuffer, data); err != nil {
		log.Println(err)
	}
	if o != nil {
		respWriter := o.(*http.ResponseWriter)
		(*respWriter).Write(outBuffer.Bytes())
	}
	fmt.Print(outBuffer.String())
}
