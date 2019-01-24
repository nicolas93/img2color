package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
)

// color_diff_euklid is a function to calculate the distance between to colors.
// The parameters are a: a RGB-array with 16bit color values and b: a RGB-array with 8bit color values.
// It returns the euklidian distance between the colors.
func color_diff_euklid(a [3]int, b []int) int {
	return int(math.Sqrt(math.Pow(float64((a[0]>>8)-b[0]), 2) + math.Pow(float64((a[1]>>8)-b[1]), 2) + math.Pow(float64((a[2]>>8)-b[2]), 2)))
}

// assign_k is a function to assign each pixel of the image to one k.
// Since it is to be run multithreaded, the result is returned via a channeldoc: symbol img2color is not a type in package main installed in "."
// The start and stop of the assigned pixels are returned in the last field of the result matrix tmp_mat
// This is done to achieve deterministic results
func assign_k(image image.Image, k int, k_med [][]int, start int, stop int, ch chan [][]int) {
	tmp_mat := make([][]int, stop-start+1)
	for i := 0; i < len(tmp_mat); i++ {
		tmp_mat[i] = make([]int, image.Bounds().Max.Y)
	}
	for x := start; x < stop; x++ {
		for y := 0; y < (image).Bounds().Max.Y; y++ {
			minimum := k
			difference := 766
			for i := 0; i < k; i++ {
				R, G, B, _ := image.At(x, y).RGBA()
				new_diff := color_diff_euklid([3]int{int(R), int(G), int(B)}, k_med[i])
				if new_diff < difference {
					difference = new_diff
					minimum = i
				}
			}
			tmp_mat[x-start][y] = minimum
		}
	}
	tmp_mat[len(tmp_mat)-1][0] = start
	tmp_mat[len(tmp_mat)-1][1] = stop
	ch <- tmp_mat
}

// medium_k is used to calculate the medium color of all pixels that are assigned to one k
// It is used multithraded and the color is returned via a channel
func medium_k(image image.Image, j int, width int, height int, k_mat [][]int, ch_m chan []int) {
	k_m := make([]int, 4)
	count := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if k_mat[x][y] == j {
				count++
				r, g, b, _ := image.At(x, y).RGBA()
				k_m[0] += int(r) >> 8
				k_m[1] += int(g) >> 8
				k_m[2] += int(b) >> 8
			}
		}
	}
	if count > 0 {
		k_m[0] /= count
		k_m[1] /= count
		k_m[2] /= count
	}
	k_m[3] = j
	ch_m <- k_m
}

// kmeans is the function to calculate k mean colors
// image is the image to be processed
// k is the number of mean colors to be calculated
// t is the number of threads to be used
// n is the number of rounds the algorithm shall run. higher n achieves better results, but from a certain n up the mean colors will not change anymore
func kmeans(image image.Image, k int, t int, n int) [][]int {
	ch := make(chan [][]int)
	ch_m := make(chan []int)

	k_med := make([][]int, k)
	r, _, _, _ := image.At(0, 0).RGBA()
	rand.Seed(int64(r >> 8))
	for i := 0; i < k; i++ {
		R, G, B, _ := image.At(int(rand.Int31n(int32(image.Bounds().Max.X))), int(rand.Int31n(int32(image.Bounds().Max.Y)))).RGBA()
		k_med[i] = []int{int(R) >> 8, int(G) >> 8, int(B) >> 8}
	}
	k_mat := make([][]int, image.Bounds().Max.X)
	for i := 0; i < len(k_mat); i++ {
		k_mat[i] = make([]int, image.Bounds().Max.Y)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < t; j++ {
			start := int(math.Round(float64(j) * float64(image.Bounds().Max.X) / float64(t)))
			stop := int(math.Round((float64(j) + 1) * float64(image.Bounds().Max.X) / float64(t)))
			go assign_k(image, k, k_med, start, stop, ch)
		}
		for j := 0; j < t; j++ {
			re := <-ch
			copy(k_mat[re[len(re)-1][0]:re[len(re)-1][1]], re[0:len(re)-1])
		}
		for j := 0; j < k; j++ {
			go medium_k(image, j, image.Bounds().Max.X, image.Bounds().Max.Y, k_mat, ch_m)
		}
		for j := 0; j < k; j++ {
			k_m := <-ch_m
			k_med[k_m[3]] = k_m[:3]
		}

		fmt.Printf("\rProcessing:\t%.2f%%", (float64(i)*100)/float64(n))
	}
	fmt.Printf("\rProcessing:\t100.00%%")
	fmt.Println("\nDone.")
	return k_med
}

// main is used to interpret the parameters,
// start the algorithm and output the results
func main() {
	k_ptr := flag.Int("k", 5, "Number of colors to find")
	t_ptr := flag.Int("t", 1, "Number of threads to use for computation")
	n_ptr := flag.Int("n", 10, "Number of rounds for computation")
	//	fast_ptr := flag.Bool("fast", false, "Activate fast mode.")
	var image_file string
	flag.StringVar(&image_file, "image", "", "Image to be processed")
	var output_ptr string
	flag.StringVar(&output_ptr, "output", "palette", "Output option")
	flag.Parse()

	reader, err := os.Open(image_file)
	if err != nil {
		log.Fatal(err)
	}
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	k_med := kmeans(m, *k_ptr, *t_ptr, *n_ptr)

	for i := 0; i < len(k_med); i++ {
		fmt.Printf("#%02x%02x%02x\n", k_med[i][0], k_med[i][1], k_med[i][2])
	}

	fmt.Println(output_ptr)

	if strings.Compare(output_ptr, "silhouette") == 0 {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{m.Bounds().Max.X, m.Bounds().Max.Y}})
		for x := 0; x < m.Bounds().Max.X; x++ {
			for y := 0; y < m.Bounds().Max.Y; y++ {
				minimum := *k_ptr
				difference := 766
				for i := 0; i < *k_ptr; i++ {
					R, G, B, _ := m.At(x, y).RGBA()
					new_diff := color_diff_euklid([3]int{int(R), int(G), int(B)}, k_med[i])
					if new_diff < difference {
						difference = new_diff
						minimum = i
					}
				}
				img.Set(x, y, color.RGBA{uint8(k_med[minimum][0]), uint8(k_med[minimum][1]), uint8(k_med[minimum][2]), 0xff})
			}
		}
		f, _ := os.Create("image_s.png")
		png.Encode(f, img)
	}

	if strings.Compare(output_ptr, "palette") == 0 {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{m.Bounds().Max.X + 100, m.Bounds().Max.Y}})
		for x := 0; x < m.Bounds().Max.X; x++ {
			for y := 0; y < m.Bounds().Max.Y; y++ {
				img.Set(x, y, m.At(x, y))
			}
		}
		k := *k_ptr
		for x := m.Bounds().Max.X; x < m.Bounds().Max.X+100; x++ {
			for y := 0; y < m.Bounds().Max.Y; y++ {
				img.Set(x, y, color.RGBA{uint8(k_med[(y*k)/m.Bounds().Max.Y][0]), uint8(k_med[(y*k)/m.Bounds().Max.Y][1]), uint8(k_med[(y*k)/m.Bounds().Max.Y][2]), 0xff})
			}
		}
		f, _ := os.Create("image_p.png")
		png.Encode(f, img)
	}

	defer reader.Close()
}
