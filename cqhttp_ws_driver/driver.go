package cqhttp_ws_driver

import (
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/http"
	"strconv"
	"sync"
)

type Driver struct {
	Name             string
	address          string
	port             int
	bots             sync.Map
	eventChan        chan []byte
	connectHandle    func(selfId int64, host string, clientRole string)
	disConnectHandle func(selfId int64)
}

func (d *Driver) OnConnect(f func(selfId int64, host string, clientRole string)) {
	d.connectHandle = f
}

func (d *Driver) OnDisConnect(f func(selfId int64)) {
	d.disConnectHandle = f
}

func (d *Driver) GetBots() map[int64]interface{} {
	m := make(map[int64]interface{})
	d.bots.Range(func(key, value interface{}) bool {
		m[key.(int64)] = value
		return true
	})

	return m
}

func (d *Driver) SetAddress(string2 string) {
	d.address = string2
}

// ws协议
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

func (d *Driver) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	selfId, err := strconv.ParseInt(request.Header.Get("X-Self-ID"), 10, 64)
	role := request.Header.Get("X-Client-Role")
	host := request.Header.Get("Host")
	conn, err := upgrade.Upgrade(writer, request, nil)
	if err != nil {
		return
	}
	b := new(Bot)
	b.conn = conn
	b.selfId = selfId
	b.responses = sync.Map{}
	_, ok := d.bots.Load(selfId)
	if ok {
		d.bots.LoadOrStore(selfId, b)
	} else {
		d.bots.Store(selfId, b)
	}
	b.disConnectHandle = d.disConnectHandle
	log.Infoln(fmt.Sprintf("the bot %v is connected", selfId))
	// 执行链接回调
	go d.connectHandle(selfId, host, role)
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				b.wsClose()
				log.Errorln("ws链接读取出现错误")
				log.Errorln(err)
			}
		}()
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				b.wsClose()
			}
			echo := gjson.GetBytes(data, "echo")
			if echo.Exists() {

				value, o := b.responses.Load(echo.String())
				if o {
					value.(chan []byte) <- data
				}
			} else {
				d.eventChan <- data
			}
		}
	}()

}

func (d *Driver) Run() {
	http.Handle("/"+d.Name+"/ws", d)
	if err := http.ListenAndServe(d.address, nil); err != nil {
		log.Panicln(err.Error())
	}
}

func (d *Driver) GetEvent() chan []byte {
	return d.eventChan
}

func (d *Driver) GetBot(i int64) interface{} {
	load, ok := d.bots.Load(i)
	if ok {
		return load
	}

	return nil
}

func NewDriver() *Driver {
	d := new(Driver)
	d.Name = "cqhttp"
	d.bots = sync.Map{}
	d.eventChan = make(chan []byte)
	return d
}
