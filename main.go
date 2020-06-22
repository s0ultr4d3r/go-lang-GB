package main

import (
	"bytes"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
	"text/template"
)

// Summarized func for sum, ,ade for tests
func Summarized(xs []float64) float64 {
	total := float64(0)
	for i := 0; i < len(xs); i++ {
		total += xs[i]
	}
	return total
}

// func for http

type IndexTemplate struct {
	Title string
	Name  string
}

func parse(in IndexTemplate) ([]byte, error) {
	tmpl, err := template.New("index.tmpl").ParseFiles("./assets/index.tmpl")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, in); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var Name string

func handler(w http.ResponseWriter, r *http.Request) {

	Name := r.URL.Query().Get("name")
	log.Println(Name)
	log.Println(string(Name))

}

// Starting drawing

var teal color.Color = color.RGBA{0, 200, 200, 255}
var red color.Color = color.RGBA{200, 30, 30, 255}
var green color.Color = color.RGBA{200, 30, 255, 30}

func main() {

	file, err := os.Create("someimage.png")

	if err != nil {
		fmt.Errorf("%s", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)
	// или draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)

	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)

	for x := 20; x < 380; x = x + 20 {
		for y := 1; y < 380; y++ {
			img.Set(x, y, red)
		}
	}

	for y := 20; y < 380; y = y + 20 {
		for x := 1; x < 380; x++ {
			img.Set(x, y, green)
		}
	}

	png.Encode(file, img)

	// Finished drawing

	xs := []float64{1, 2, 2, 2}
	fmt.Println(Summarized(xs))

	// Задание по Hello

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	IndexTemplate := IndexTemplate{
		Title: "Test",
		Name:  Name,
	}

	parsedTmpl, err := parse(IndexTemplate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(parsedTmpl))
}
