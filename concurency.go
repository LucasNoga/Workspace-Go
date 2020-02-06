// voici le moyen de creer une routine go

package main

import(
	"fmt"
	"time"
	//"sync"
)

func main(){
	
	c := make(chan string)
	go count("sheep", c)
	for msg := range c {
		fmt.Println(msg)
	}


	// var wg sync.WaitGroup
	// wg.Add(1)

	//go count("sheep")
	//go count("fish")

	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()

	// wg.Wait()
}


func count(thing string,c chan string){
	for i := 1; i<=5; i++ {
		//fmt.Println(i, thing)
		msg :=  fmt.Sprint(i, ":", thing)
		c <- msg
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}