package routes

import "net/http"


// home godoc
// @Summary      homepage
// @Description  get homepage
// @Tags         medidores
// @Success      200  
// @Failure      400  
// @Failure      404  
// @Failure      500  
// @Router       / [get]
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}