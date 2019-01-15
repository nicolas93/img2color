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



func color_diff(a [3]int, b [3]int) int{
	var d [3]int
	d[0] = a[0]/257-b[0]
	d[1] = a[1]/257-b[1]
	d[2] = a[2]/257-b[2]
	if(d[0]<0){ d[0] *= -1}
	if(d[1]<0){ d[1] *= -1}
	if(d[2]<0){ d[2] *= -1}
	return d[0]+d[1]+d[2]
}



func assign_k(image image.Image, k int, k_med *[5][3]int, k_mat *[1920][1080]int,){
	for x:=0; x < (image).Bounds().Max.X; x++{
		for y:=0;  y< (image).Bounds().Max.Y; y++{
			minimum := k
			difference := 766
			for i:=0; i<k; i++{
				R,G,B,_ := image.At(x,y).RGBA()
				new_diff := color_diff([3]int{int(R),int(G),int(B)}, k_med[i])
				if new_diff < difference{
					difference = new_diff
					minimum = i
				}
			}
			k_mat[x][y] = minimum
		}
	}
}



func kmeans(image image.Image, k int, t int){ 
	fmt.Println(image.At(k,k))
	fmt.Println(image.Bounds().Max.X)
	// CONSTANT ARRAY / K !!
	var k_med [5][3]int
	r,_,_,_ :=image.At(0,0).RGBA()
	fmt.Println(r/257)
	rand.Seed(int64(r/257))
	for i:=0; i < k; i++ {
		k_med[i]=[3]int{int(rand.Int31n(255)), int(rand.Int31n(255)), int(rand.Int31n(255))}
	}
	fmt.Println(k_med)
	//var k_mat [image.Bounds().Max.X][image.Bounds().Max.Y]int
	// CONSTANT ARRAY! hardcoded!
	var k_mat [1920][1080]int
	assign_k(image, k, &k_med, &k_mat)
	fmt.Println(k_mat[0])
	
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