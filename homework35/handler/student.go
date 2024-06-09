package handler

import (
	"encoding/json"
	"fmt"
	modul "mymod/module"
	"net/http"
	"strconv"
	"strings"
)

func (s *Handler3) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/getstudentbyid/")
	student := modul.Students{}

	students, err := s.PostStudent.GetAllStudent(student)
	if err != nil {
		panic(err)
	}
	check := false
	for _,v := range *students {
		num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
		if int(num) == v.Id {
			
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

func (s *Handler3) GetStudent(w http.ResponseWriter, r *http.Request) {
	student := modul.Students{}

	students, err := s.PostStudent.GetAllStudent(student)
	if err != nil {
		panic(err)
	}
	for _,v := range *students {
		json.NewEncoder(w).Encode(v)
	}
	if len(*students)==0{

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Empty database"))
		return
	}
	

}

func (s *Handler3) CreateStudent(w http.ResponseWriter, r *http.Request) {
	student := modul.Students{}

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ERROR DECODER"))
		return
	}

	s.PostStudent.CreateStudent(student)

	fmt.Println(student)
	json.NewDecoder(r.Body).Decode(&student)
	w.Write([]byte("Successfully created"))
}

func (s *Handler3) UpdateStudent(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/updatestudent/")
	newStudent := modul.Students{}

	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("ERROR DECODER"))
		return
	}
	Student := modul.Students{}

	err = s.PostStudent.UpdateStudent(Student,id)
	
		w.Write([]byte("Successfully updated"))
}

func (s *Handler3) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/deleteStudent/")

	student := modul.Students{}
	err := s.PostStudent.DeleteStudent(student,id)
	if err!= nil {
        panic(err)
    }
	
	w.Write([]byte("DELETE FROM DATABASE"))

}
