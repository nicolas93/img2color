#!/usr/bin/env python2

from PIL import Image
import sys
import random
import argparse

def color_diff(a, b):
	d0 = a[0]-b[0]
	d1 = a[1]-b[1]
	d2 = a[2]-b[2]
	d0 = d0 if d0>0 else d0*-1
	d1 = d1 if d1>0 else d1*-1
	d2 = d2 if d2>0 else d2*-1
	return d0+d1+d2

def med(cluster, k):
	if(len(cluster) == 0):
		return k
	r,g,b = (0,0,0)
	for c in cluster:
		r = r + c[0]
		g = g + c[1]
		b = b + c[2]
	r = r/len(cluster)
	g = g/len(cluster)
	b = b/len(cluster)
	return [r,g,b]


def assign_cluster(im, k, cluster, width, height):
	minimum = len(k)
	for i in range(0, width):
		for j in range(0, height):
			minimum = len(k)
			difference = 766
			for l in range(0, len(k)):
				new_diff = color_diff(im.getpixel((i, j)), k[l])
				if(new_diff < difference):
					difference = new_diff
					minimum = l
			cluster[minimum].append(im.getpixel((i, j)))

def kmeans(im):
	width, height = im.size
	k = []
	cluster = []
	cluster1 = []
	cluster2 = []
	cluster3 = []
	for i in range(0,3):
		k.append(im.getpixel((random.randint(0,width), random.randint(0,height))))
		cluster.append([])
		cluster1.append([])
		cluster2.append([])
		cluster3.append([])
	print k
	assign_cluster(im, k, cluster, width, height)
	k[0] = med(cluster[0], k[0])
	k[1] = med(cluster[1], k[1])
	k[2] = med(cluster[2], k[2])
	print k
	assign_cluster(im, k, cluster1, width, height)
	k[0] = med(cluster1[0], k[0])
	k[1] = med(cluster1[1], k[1])
	k[2] = med(cluster1[2], k[2])
	print k
	assign_cluster(im, k, cluster2, width, height)
	k[0] = med(cluster2[0], k[0])
	k[1] = med(cluster2[1], k[1])
	k[2] = med(cluster2[2], k[2])
	print k
	assign_cluster(im, k, cluster3, width, height)
	k[0] = med(cluster3[0], k[0])
	k[1] = med(cluster3[1], k[1])
	k[2] = med(cluster3[2], k[2])
	print k
	return k



def main():
	parser = argparse.ArgumentParser(description='Find main colors in a given image.')
	parser.add_argument("image", help="Image to be processed")
	parser.add_argument("--output-format",type=str, choices=['image-palette', 'html-color-code'], help="Output-format")
	args = parser.parse_args()
	print args
	image = Image.open(args.image)
	k = kmeans(image)
	if(args.output_format == "image-palette"):
		img = Image.new('RGB', (image.size[0]+100, image.size[1]))
		img.paste(image)
		img0 = Image.new('RGB', (100, image.size[1]/3), color = tuple(k[0]))
		img1 = Image.new('RGB', (100, image.size[1]/3), color = tuple(k[1]))
		img2 = Image.new('RGB', (100, image.size[1]/3), color = tuple(k[2]))
		
		img.paste(img0, box=(image.size[0], 0))
		img.paste(img1, box=(image.size[0], image.size[1]/3))
		img.paste(img2, box=(image.size[0], 2*image.size[1]/3))

		img.save(args.image +"_pallette.png", "PNG")	
	else:
		print "#%02x%02x%02x" % (k[0][0],k[0][1],k[0][2])
		print "#%02x%02x%02x" % (k[1][0],k[1][1],k[1][2])
		print "#%02x%02x%02x" % (k[2][0],k[2][1],k[2][2])



if __name__ == "__main__":
	main()