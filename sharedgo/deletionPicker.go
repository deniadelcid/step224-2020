// Picks what file from 10 to 50 (files 0-9 kept in the system for persistence) to delete

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func pick_file_to_delete() string {
	rand.Seed(time.Now().UnixNano())
	var file_ID int = rand.Intn(40) + 10;
	var file_name string = "Hermes_"+strconv.Itoa(file_ID)
	return file_name;	
}

func main(){
	fmt.Println(pick_file_to_delete());
}
