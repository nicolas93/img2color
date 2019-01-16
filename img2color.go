package main

import "os"
import "fmt"
import "log"
import "flag"
import "image"
import "math/rand"
import _ "image/jpeg"
import _ "image/png"
import _ "image/gif"

func color_diff(a [3]int, b []int) int {
	var d [3]int
	d[0] = a[0]/257 - b[0]
	d[1] = a[1]/257 - b[1]
	d[2] = a[2]/257 - b[2]
	if d[0] < 0 {
		d[0] *= -1
	}
	if d[1] < 0 {
		d[1] *= -1
	}
	if d[2] < 0 {
		d[2] *= -1
	}
	return d[0] + d[1] + d[2]
}

func assign_k(image image.Image, k int, k_med [][]int, k_mat [][]int) {
	for x := 0; x < (image).Bounds().Max.X; x++ {
		for y := 0; y < (image).Bounds().Max.Y; y++ {
			minimum := k
			difference := 766
			for i := 0; i < k; i++ {
				R, G, B, _ := image.At(x, y).RGBA()
				new_diff := color_diff([3]int{int(R), int(G), int(B)}, k_med[i])
				if new_diff < difference {
					difference = new_diff
					minimum = i
				}
			}
			k_mat[x][y] = minimum
		}
	}
}

func medium(image image.Image, k int, width int, height int, k_med [][]int, k_mat [][]int) {
	for i := 0; i < k; i++ {
		count := 0
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if k_mat[x][y] == i {
					count++
					r, g, b, _ := image.At(x, y).RGBA()
					k_med[i][0] += int(r) / 257
					k_med[i][1] += int(g) / 257
					k_med[i][2] += int(b) / 257
				}
			}
		}
		if count > 0 {
			k_med[i][0] /= count
			k_med[i][1] /= count
			k_med[i][2] /= count
		}
	}
}

func kmeans(image image.Image, k int, t int) [][]int {
	k_med := make([][]int, k)
	r, _, _, _ := image.At(0, 0).RGBA()
	fmt.Println(r / 257)
	rand.Seed(int64(r / 257))
	for i := 0; i < k; i++ {
		k_med[i] = []int{int(rand.Int31n(255)), int(rand.Int31n(255)), int(rand.Int31n(255))}
	}
	fmt.Println(k_med)
	k_mat := make([][]int, image.Bounds().Max.X)
	for i:=0; i<len(k_mat);i++{
		k_mat[i] = make([]int, image.Bounds().Max.Y)
	}
	for i := 0; i < 5; i++ {
		assign_k(image, k, k_med, k_mat)
		medium(image, k, image.Bounds().Max.X, image.Bounds().Max.Y, k_med, k_mat)
	}

	return k_med
}

func main() {
	k_ptr := flag.Int("k", 5, "Number of colors to find")
	t_ptr := flag.Int("t", 1, "Number of threads to use for computation")
	fast_ptr := flag.Bool("fast", false, "Activate fast mode.")
	var image_file string
	flag.StringVar(&image_file, "image", "", "Image to be processed")
	flag.Parse()

	fmt.Println("k: ", *k_ptr)
	fmt.Println("t: ", *t_ptr)
	fmt.Println("fast?: ", *fast_ptr)

	reader, err := os.Open(image_file)
	if err != nil {
		log.Fatal(err)
	}
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds)
	k_med := kmeans(m, *k_ptr, *t_ptr)

	fmt.Println(k_med)
	for i:=0; i<len(k_med); i++{
		fmt.Printf("#%02x%02x%02x\n", k_med[i][0], k_med[i][1], k_med[i][2])
	}

	img := image.NewRGBA(image.Rectangle{image.Point{0,0}, image.Point{m.Bounds().Max.X, m.Bounds().Max.Y}})
	for x:=0; x<m.Bounds().Max.X; x++{
		for y:=0; y<m.Bounds().Max.Y; y++{
			minimum := *k_ptr
			difference := 766
			for i := 0; i < *k_ptr; i++ {
				R, G, B, _ := m.At(x, y).RGBA()
				new_diff := color_diff([3]int{int(R), int(G), int(B)}, k_med[i])
				if new_diff < difference {
					difference = new_diff
					minimum = i
				}
			}
			img.Set(x, y, color.RGBA{k_med[minimum][0],k_med[minimum][1],k_med[minimum][2],0xff})
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)

	defer reader.Close()
}
