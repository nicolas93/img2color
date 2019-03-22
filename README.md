# img2color

## Usage

```
Usage of img2color.go:
  -image string
      Image to be processed
  -k int
      Number of colors to find (default 5)
  -mode string
      Output option (default "palette")
  -n int
      Number of rounds for computation (default 10)
  -o string
      Output file name (default "image.png")
  -t int
      Number of threads to use for computation (default 1)
```

## Examples

### Testimage
This image is used for tests. It was provided by https://www.pexels.com .
![test image](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg)


### Color-Palette output(k=5):
Main  colors are shown in a palette next to the image.

```
go run img2color.go -image testimage.jpeg -k 6 -t 10 -mode palette
```
![test image with color-palette](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg_palette_k6.png)

### Color-Silhouette output(k=6):
In this example every pixel is colored in its nearest main-color.

```
go run img2color.go -image testimage.jpeg -k 6 -t 10 -mode silhouette
```
![test image with color-silhouette](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg_silhouette_k6.png)

### Color-Silhouette output(k=12):
In this example every pixel is colored in its nearest main-color.

```
go run img2color.go -image testimage.jpeg -k 12 -t 10 -mode silhouette
```
![test image with color-silhouette](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg_silhouette_k12.png)


### html-color-code
```
go run img2color.go -image testimage.jpeg -k 6 -t 10 -mode html
Processing: 100.00%
Done.
#ba6223
#72dae8
#b9edf3
#f9fdfd
#09a4b8
#252827
html
```

## Kmeans-Algorithm

The kmeans algorithm is used to calculate k mean points of a set of points.
In each computation step every point is assigned to the nearest mean point. 
Then of every (k) subset a new mean point is calculated. The mean point does not have to be in the subset.

In this project we use the color of each pixel as a 3 dimensional point, and thus k mean (or dominant) colors are calculated. 


## Notes

### Python implementation
The python implementation (img2color.py) is no longer supported and discontinued.
It was much slower than the Go implementation.