package main

import (
	"fmt"
	"nomadcoder/banking"
)

type Dictionary map[interface{}]interface{}

func (dic Dictionary) Search(word interface{}) interface{} {
	val, is := dic[word]
	if !is {
		return nil
	}
	return val
}

func (dic Dictionary) Add(key interface{}, value interface{}) error {
	_, is := dic[key]
	if is {
		return fmt.Errorf("key is already exist in dict")
	}

	dic[key] = value
	return nil
}

func (dic Dictionary) Print(word interface{}) {
	val, is := dic[word]
	if !is {
		fmt.Println("no", word, "in dic")
	} else {
		fmt.Println(val, "is in dic")
	}

}

func (dic Dictionary) Update(key interface{}, value interface{}) error {
	_, is := dic[key]
	if !is {
		return fmt.Errorf("key is not in dic")
	}
	dic[key] = value
	fmt.Println(dic[key], "is update")
	return nil
}

func (dic Dictionary) Delete(key interface{}) error {
	_, is := dic[key]
	if !is {
		return fmt.Errorf("key is not in dic")
	}
	delete(dic, key)
	return nil
}
func main() {
	fmt.Println("hihi")
	account := banking.NewAccount("Lee")
	fmt.Println(account)
	account.Deposit(10000)
	fmt.Println(account.GetDeposit())
	account.AddDeposit(10000)
	fmt.Println(account.GetDeposit())

	dict := Dictionary{"name": "LEE"}
	dict.Print("name")
	dict.Print("fda")
	dict.Print("name")
	dict.Print("fda")

	err := dict.Add("fda", 123)
	if err != nil {
		fmt.Println("error!")
		return
	}

	dict.Print("fda")
	dict.Update("fda", 321)
	dict.Print("fda")
	dict.Delete("fda")
	dict.Print("fda")
}
