package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func (l *level) makeJson(fileName string){
	levelData, err := json.Marshal(*l)
	if err!=nil{
		fmt.Printf("Creating json data from level\n")
		panic("rip")
	}

	f, err := os.Create("levels/"+fileName+".json")
	if err!=nil{
		fmt.Printf("Error creating file")
		panic("rip")
	}
	defer f.Close()

	if _, err = f.WriteString(string(levelData)); err!=nil{
		fmt.Printf("Error writting to file")
		panic("rip")
	}


	//fmt.Println(string(levelData))
}