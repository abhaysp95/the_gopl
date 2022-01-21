// Modify the Lissajous server to read parameter values from the URL. For
// example, you might arrange it so that a URL like
// http://localhost:8000/?cycles=20 sets the number of cycles to 20 ins tead of
// the default 5. Use the strconv.Atoi function to convert the string parameter
// into an int eger. You can see its document ation with go doc strconv.Atoi.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

const (
	whiteIndex = 0  // first color in palette
	blackIndex = 1  // second color in palette
	redIndex = 2    // third color in palette
	greenIndex = 3  // fourth color in palette
	blueIndex = 4   // fifth color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	cycles := 5.0  // number of complete x oscillator revolutions
	for k, v := range r.Form {
		if k == "cycles" {
			cycles, err = strconv.ParseFloat(v[0], 32)
			if err != nil {
				log.Print(err)
				cycles = 5.0
			}
			break
		}
	}
	const (
		res = 0.001   // angular resolution
		size = 100    // image canvas covers [-size...+size]
		nframes = 64  // number of animation frames
		delay = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0  // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0  // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), uint8(i % (len(palette) - 1)) + 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)  // NOTE: ignoring encoding errors
}
