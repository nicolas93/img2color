# img2color

## Dependencies
* Pillow: ```pip install Pillow```

## Usage
```
usage: img2color.py [-h] [--output-format {image-palette,html-color-code}]
                    image

Find main colors in a given image.

positional arguments:
  image                 Image to be processed

optional arguments:
  -h, --help            show this help message and exit
  --output-format {image-palette,html-color-code}
                        Output-format
```


## Examples

### Testimage
This image is used for Tests. It was provided by https://www.pexels.com .
![test image](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg)


### Color-Palette output:
![test image with color-palette](https://raw.githubusercontent.com/nicolas93/img2color/master/testimage.jpeg_pallette.png)

### html-color-code
```
./img2color.py testimage.jpeg
#483a2b
#55cbda
#e5f8fa
```

## Kmeans-Algorithm