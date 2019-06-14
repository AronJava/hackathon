package paper

import(
	"net/http"

	"hackathon/utils"
)

func GenerateImg(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	img := vars.Get("img")
	if img == "" {
		utils.ResponseJson(w, -1, "img is invaild", nil)
		return
	}

	url := utils.GenerateImage(img)
	if url == "" {
		utils.ResponseJson(w, -1, "generate failed.", nil)
		return
	}

	utils.ResponseJson(w, 0, "generate success.", url)
}
