package model

import (
	"time"
)

const (
	READY_MSG_STATE = iota
	RESERVED_MSG_STATE
	DELAYED_MSG_STATE
)

const defaultPriority int = 5
const defaultReserveTimeoutInSec = 10

// Msg struct defines the structure of the msg that is enqueued/dequeued
type Msg struct {
	ID                  string
	Tube                *Tube
	Data                *Data
	ReserveTimeoutInSec int64
	Metadata            *MsgMetaData
	WaitingData         *Data
	ReceiptID           string
}

type Data struct {
	DelayInSec int64
	Priority   int
	Version    int64
	DataSlice  [][]byte
}

type MsgMetaData struct {
	State             int
	ReservedTimestamp *time.Time
	DelayedTimestamp  *time.Time
}

func NewData(data []byte, priority int, delayInSec int64) *Data {
	return &Data{
		DelayInSec: delayInSec,
		Priority:   priority,
		Version:    time.Now().Unix(),
		DataSlice:  [][]byte{data},
	}
}

func NewMsg(id string, dataBytes []byte, delayInSec int64, reserveTimeoutInSec int64, priority int, tube *Tube) *Msg {
	var dataSlice [][]byte
	dataSlice = append(dataSlice, dataBytes)
	var msgMetaData *MsgMetaData
	if delayInSec > 0 {
		delayedTimestamp := time.Now().Add(time.Duration(delayInSec) * time.Second)
		msgMetaData = &MsgMetaData{
			State:            DELAYED_MSG_STATE,
			DelayedTimestamp: &delayedTimestamp,
		}
	} else {
		msgMetaData = &MsgMetaData{
			State: READY_MSG_STATE,
		}
	}
	data := &Data{
		DelayInSec: delayInSec,
		Priority:   priority,
		Version:    time.Now().Unix(),
		DataSlice:  dataSlice,
	}
	msg := &Msg{
		ID:                  id,
		ReserveTimeoutInSec: reserveTimeoutInSec,
		Metadata:            msgMetaData,
		Tube:                tube,
		Data:                data,
	}
	return msg
}