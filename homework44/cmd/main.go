package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)


var Dictionary = map[string]string{
    "apple": "olma",
    "book": "kitob",
    "house": "uy",
    "car": "mashina",
    "tree": "daraxt",
    "water": "suv",
    "fire": "olov",
    "sky": "osmon",
    "earth": "yer",
    "sun": "quyosh",
    "moon": "oy",
    "star": "yulduz",
    "flower": "gul",
    "mountain": "tog'",
    "river": "daryo",
    "lake": "ko'l",
    "forest": "o'rmon",
    "animal": "hayvon",
    "bird": "qush",
    "fish": "baliq",
    "bread": "non",
    "milk": "sut",
    "cheese": "pishloq",
    "butter": "yog'",
    "egg": "tuxum",
    "meat": "go'sht",
    "vegetable": "sabzavot",
    "fruit": "meva",
    "school": "maktab",
    "teacher": "o'qituvchi",
    "student": "talaba",
    "pen": "ruchka",
    "pencil": "qalam",
    "notebook": "daftar",
    "computer": "kompyuter",
    "phone": "telefon",
    "window": "deraza",
    "door": "eshik",
    "chair": "stul",
    "table": "stol",
    "bed": "karavot",
    "cup": "fincan",
    "plate": "likopcha",
    "spoon": "qoshiq",
    "fork": "sanchqi",
    "knife": "pichoq",
    "clothes": "kiyim",
    "shoe": "oyoq kiyim",
    "hat": "shapka",
    "glasses": "ko'zoynak",
    "watch": "soat",
    "bag": "sumka",
    "road": "yo'l",
    "street": "ko'cha",
    "city": "shahar",
    "village": "qishloq",
    "country": "mamlakat",
    "world": "dunyo",
    "family": "oila",
    "friend": "do'st",
    "love": "sevgi",
    "happiness": "baxt",
    "sadness": "g'am",
    "anger": "g'azab",
    "fear": "qo'rquv",
    "surprise": "hayrat",
    "health": "salomatlik",
    "medicine": "dori",
    "doctor": "shifokor",
    "hospital": "kasalxona",
    "work": "ish",
    "money": "pul",
    "time": "vaqt",
    "day": "kun",
    "night": "tun",
    "morning": "ertalab",
    "afternoon": "tushdan keyin",
    "evening": "kechqurun",
    "week": "hafta",
    "month": "oy",
    "year": "yil",
    "spring": "bahor",
    "summer": "yoz",
    "autumn": "kuz",
    "winter": "qish",
}

type Translate struct {}

func (t *Translate) Tarjimon (word string, reply *string) error {
	if translation, ok := Dictionary[word]; ok {
		*reply = translation
	} else {
		*reply = "Tarjima topilmadi"
	}
	return nil
}
func main()  {

    translate := new(Translate)
    rpc.Register(translate)

    listener, err := net.Listen("tcp", ":8080")
    if err!= nil {
        log.Fatal("Error translate register: ", err)
    }

    defer listener.Close()
    fmt.Println("Server started on port :8080")

    for {
        conn, err := listener.Accept()
        if err!= nil {
            log.Printf("Error accepting: %v", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
    
}




