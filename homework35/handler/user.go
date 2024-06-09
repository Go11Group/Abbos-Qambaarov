package handler

import (
	"encoding/json"
	"fmt"
	modul "mymod/module"
	"net/http"
	"strings"
)

func (u *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/getuserbyid/")
	user := modul.Users{}

	users, err := u.PostUser.GetAllUser(user)
	if err != nil {
		panic(err)
	}
	check := false
	for _,v := range *users {
		if id == v.Id {
			
			json.NewEncoder(w).Encode(v)
			check = true

		}
	}
	if check==false{

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
	

}

func (u *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user := modul.Users{}

	users, err := u.PostUser.GetAllUser(user)
	if err != nil {
		panic(err)
	}
	for _,v := range *users {
		json.NewEncoder(w).Encode(v)
	}
	if len(*users)==0{

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Empty database"))
		return
	}
	

}

func (u *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := modul.Users{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ERROR DECODER"))
		return
	}

	u.PostUser.CreateUser(user)

	fmt.Println(user)
	json.NewDecoder(r.Body).Decode(&user)
	w.Write([]byte("Successfully created"))
}

func (u *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/updateuser/")
	newuser := modul.Users{}

	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ERROR DECODER"))
		return
	}
	user := modul.Users{}

	err = u.PostUser.UpdateUser(user,id)
	
		w.Write([]byte("Successfully updated"))
}

func (u *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/deleteuser/")

	user := modul.Users{}
	err := u.PostUser.DeleteUser(user,id)
	if err!= nil {
        panic(err)
    }
	
	w.Write([]byte("DELETE FROM DATABASE"))

}
