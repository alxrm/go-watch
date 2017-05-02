package main

import "gopkg.in/olahol/melody.v1"

var commands = map[string]func(*melody.Melody, string){
  "add": onAdd,
  "list": onList,
  "remove": onRemove,
  "clear": onClear,
  "default": onDefault,
}

func onAdd(channel *melody.Melody, data string) {

}

func onList(channel *melody.Melody, data string) {

}

func onRemove(channel *melody.Melody, data string) {

}

func onClear(channel *melody.Melody, data string) {

}

func onDefault(channel *melody.Melody, _ string) {
  channel.Broadcast([]byte("Wrong command"))
}