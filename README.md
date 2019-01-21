# img2color

## Dependencies
* Python
  * Pillow: ```pip install Pillow```

## Usage

### Python
```
usage: img2color.py [-h] [-k K] [--fast] [-t T]
                    [--output-format {image-palette,silhouette,html-color-code}]
                    image

Find main colors in a given image.

positional arguments:
  image                 Image to be processed

optional arguments:
  -h, --help            show this help message and exit
  -k K                  Custom K for KMeans algorithm
  --fast                Activate fast mode
  -t T                  Number of threads to use for computation
  --output-format {image-palette,silhouette,html-color-code}
                        Output-format
```

### Go
```
Usage of img2color.go:
  -fast
    	Activate fast mode.
  -image string
    	Image to be processed
  -k int
    	Number of colors to find (default 5)
  -output string
    	Output option (default "pallette")
  -t int
    	Number of threads to use for computation (default 1)
```

## Examples

### Testimage
This image is used for tests. It was provided by https://www.pexels.com .
![test image](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg)


### Color-Palette output(k=5):
Main  colors are shown in a palette next to the image.
![test image with color-palette](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg_pallette_k5.png)

### Color-Silhouette output(k=5):
In this example every pixel is colored in its nearest main-color.
![test image with color-silhouette](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg_silhouette_k5.png)

### html-color-code
```
./img2color.py testimage.jpeg
#483a2b
#55cbda
#e5f8fa
```

## Kmeans-Algorithm

The kmeans algorithm is used to calculate k mean points of a set of points.
In each computation step every point is assigned to the nearest mean point. 
Then of every (k) subset a new mean point is calculated. The mean point does not have to be in the subset.

In this project we use the color of each pixel as a 3 dimensional point, and thus k mean (or dominant) colors are calculated. 