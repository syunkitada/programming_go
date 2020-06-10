package main

import "fmt"

type User struct {
	Age  int
	Name string
}

func (user *User) String() string {
	return fmt.Sprintf("Name=%s", user.Name)
}

func main() {
	NgRangePattern1()
	OkRangePattern1()
	OkRangePattern2()
}

func NgRangePattern1() {
	// ListをMapに変換する処理
	fmt.Printf("\nNG Pattern1\n")
	users := []User{
		User{Name: "User1"},
		User{Name: "User2"},
		User{Name: "User3"},
	}
	userMap := map[string]*User{}
	for i, user := range users {
		// range の返り値には同じアドレスが利用されているので参照してはならない(&は使ってはならない)
		// また、struct内のポインタに関しても同じアドレスが利用されるため参照してはならない
		fmt.Printf("i=%p, user=%p, user.Name=%s, &user.Name=%p\n", &i, &user, user.Name, &user.Name)
		// i=0xc000016188, user=0xc0000102c0, user.Name=User1, &user.Name=0xc0000102c0
		// i=0xc000016188, user=0xc0000102c0, user.Name=User2, &user.Name=0xc0000102c0
		// i=0xc000016188, user=0xc0000102c0, user.Name=User3, &user.Name=0xc0000102c0

		userMap[user.Name] = &user
	}

	for _, user := range userMap {
		fmt.Println(user)
	}
}

func OkRangePattern1() {
	// ListをMapに変換する処理
	fmt.Printf("\nOK Pattern1\n")
	users := []User{
		User{Name: "User1"},
		User{Name: "User2"},
		User{Name: "User3"},
	}
	userMap := map[string]*User{}
	for i, user := range users {
		fmt.Printf("i=%p, user=%p\n", &i, &user)
		userMap[user.Name] = &users[i]
	}

	for _, user := range userMap {
		fmt.Println(user)
	}
}

func OkRangePattern2() {
	// ListをMapに変換する処理
	fmt.Printf("\nOK Pattern2\n")
	users := []*User{
		&User{Name: "User1"},
		&User{Name: "User2"},
		&User{Name: "User3"},
	}
	userMap := map[string]*User{}
	for i, user := range users {
		fmt.Printf("i=%p, user=%p, &user.Name=%p\n", &i, &user, &user.Name)
		userMap[user.Name] = user
	}

	for _, user := range userMap {
		fmt.Println(user)
	}
}
