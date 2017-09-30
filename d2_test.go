package d2

import "testing"
import "fmt"
import "encoding/json"

func Test_Map(t *testing.T) {
	data := NewD2()
	data.Add("hello", "world", 123)
	data.Add("hello", "key", "value")

	j, _ := json.Marshal(data)
	fmt.Println(string(j))

	data.Add("hello2", "world", 123)
	data.Add("hello2", "key", "value")

	mapmap:=data.GetMapMap()
	c,_:=json.Marshal(mapmap)
	fmt.Println("map-map:", string(c))

	data.RemoveKey("hello", "world")

	k, _ := json.Marshal(data)
	fmt.Println(string(k))

	data.RemoveKey("hello", "key")

	l, _ := json.Marshal(data)
	fmt.Println(string(l))

	fmt.Println(data.Get("hello2", "key"))

	data.RemoveSection("hello2")

	m, _ := json.Marshal(data)
	fmt.Println(string(m))

}
