package main

import "os"
import "fmt"
import "log"
import "flag"
import "image"
import _ "image/jpeg"
import _ "image/png"
import _ "image/gif"



func assign_k(){
	var ks [1920][1080]int
	fmt.Println(ks)
}

func color_diff(a [3]int, b [3]int){
	var d [3]int
	d[0] = a[0]-b[0]
	d[1] = a[1]-b[1]
	d[2] = a[2]-b[2]
	if(d[0]<0){ d[0] *= -1}
	if(d[1]<0){ d[1] *= -1}
	if(d[2]<0){ d[2] *= -1}
	return d[0]+d[1]+d[2]
}

func kmeans(image image.Image, k int, t int){ 
	fmt.Println(image.At(k,k))
	fmt.Println(image.Bounds().Max.X)
	var k_med [k][3]int
	var k_mat [image.Bounds().Max.X][image.Bounds().Max.Y]int
	for x:=0; x < image.Bounds().Max.X; x++{
		for y:=0;  y< image.Bounds().Max.Y; y++{

		}
	}
}

func main() {
	k_ptr := flag.Int("k", 5, "Number of colors to find")
	t_ptr := flag.Int("t", 1, "Number of threads to use for computation")
	fast_ptr := flag.Bool("fast", false, "Activate fast mode.")
	var image_file string
	flag.StringVar(&image_file, "image", "", "Image to be processed")
	flag.Parse()

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
	kmeans(m, *k_ptr, *t_ptr)



	defer reader.Close()

	fmt.Println("k: ", *k_ptr)
	fmt.Println("t: ", *t_ptr)
	fmt.Println("fast?: ", *fast_ptr)
}