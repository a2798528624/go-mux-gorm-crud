package route

import (
	"github.com/gorilla/mux"
	"go-mux-gorm-crud/internal/controller"
)

func RegisterPlaceRoute(r *mux.Router) {
	// 注册路由
	r.HandleFunc("/places", controller.GetAllPlaces).Methods("GET")
	r.HandleFunc("/place", controller.CreatePlace).Methods("POST")
	r.HandleFunc("/place/{id}", controller.GetPlace).Methods("GET")
	r.HandleFunc("/place/{id}", controller.UpdatePlace).Methods("PUT")
	r.HandleFunc("/place/{id}", controller.DeletePlace).Methods("DELETE")

}
