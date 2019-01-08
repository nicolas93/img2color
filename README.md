# img2color

## Dependencies
* Pillow: ```pip install Pillow```

## Usage
```
usage: img2color.py [-h] [-k K] [-t T]
                    [--output-format {image-palette,silhouette,html-color-code}]
                    image

Find main colors in a given image.

positional arguments:
  image                 Image to be processed

optional arguments:
  -h, --help            show this help message and exit
  -k K                  Custom K for KMeans algorithm
  -t T                  Number of threads to use for computation
  --output-format {image-palette,silhouette,html-color-code}
                        Output-format
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