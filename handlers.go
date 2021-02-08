package steganography

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"net/http"
	"os"
	"time"
)

func pingHandler(w http.ResponseWriter, _ *http.Request) {
	Log("pingHandler")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	Log("decodeHandler")

	// 1. Validate

	if err := r.ParseMultipartForm(128 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer func() {
		_ = f.Close()
	}()

	filePin := r.FormValue("pin")
	fileMessage := r.FormValue("message")

	if filePin == "" {
		http.Error(w, "filePin is not provided", http.StatusBadRequest)
		return
	}
	if fileMessage == "" {
		http.Error(w, "fileMessage is not provided", http.StatusBadRequest)
		return
	}

	// 2. Process

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Decode ", err)
	}

	cimg := image.NewRGBA(img.Bounds())
	draw.Draw(cimg, img.Bounds(), img, image.Point{}, draw.Over)

	msg, err := retrieveMessage(cimg)
	if err != nil {
		fmt.Println("retrieveMessage ", err)
	}

	msg = binToString(msg)
	expMsg := getHash(fileMessage, filePin)

	if expMsg == msg {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "no match", http.StatusBadRequest)
	}
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	Log("encodeHandler")

	// 1. Validate

	if err := r.ParseMultipartForm(128 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer f.Close()

	fileName := r.FormValue("name")
	filePin := r.FormValue("pin")
	fileMessage := r.FormValue("message")

	if fileName == "" {
		http.Error(w, "fileName is not provided", http.StatusBadRequest)
		return
	}
	if filePin == "" {
		http.Error(w, "filePin is not provided", http.StatusBadRequest)
		return
	}
	if fileMessage == "" {
		http.Error(w, "fileMessage is not provided", http.StatusBadRequest)
		return
	}

	// 2. Process

	img, imgType, err := image.Decode(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cimg := image.NewRGBA(img.Bounds())
	draw.Draw(cimg, img.Bounds(), img, image.Point{}, draw.Over)

	expMsg := stringToBin(getHash(fileMessage, filePin))
	encodedPhrase := fmt.Sprintf("%s%s", expMsg, ENCODING_END)
	encodeMessage(cimg, encodedPhrase)

	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", imgType))

	switch {
	case imgType == "png":
		err = png.Encode(w, cimg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case imgType == "jpeg" || imgType == "jpg":
		//err = jpeg.Encode(w, cimg, &jpeg.Options{Quality:100})
		//TODO: legit?
		err = png.Encode(w, cimg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	default:
		http.Error(w, "unsupported image format", http.StatusBadRequest)
		return
	}
}

func testServeImage(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("1.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	http.ServeContent(w, r, "test", time.Now(), f)
}
