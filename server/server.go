package server

import (
	"net/http"

	"github.com/vinitparekh17/project-x/database"
)

func Init() {
	router := NewRouter()

	// file, err := os.OpenFile("/logss/main.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// logger := slog.New(slog.NewJSONHandler(file, nil))
	// logger.Info("this is working fine")
	db := database.Init()
	defer db.Close()
	router.Mount("/user", UserController{}.Routes())
	http.ListenAndServe(":3000", router)
}
