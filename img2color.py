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


def assign_cluster(im, k, cluster, width, height, k_len):
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

def med_ks(k, cluster):
	for i in range(0, len(k)):
		k[i] = med(cluster[i], k[i])

def kmeans(im, k_len):
	width, height = im.size
	k = []
	cluster = []
	cluster1 = []
	cluster2 = []
	cluster3 = []
	for i in range(0,k_len):
		k.append(im.getpixel((random.randint(0,width), random.randint(0,height))))
		cluster.append([])
		cluster1.append([])
		cluster2.append([])
		cluster3.append([])
	print k
	assign_cluster(im, k, cluster, width, height, k_len)
	med_ks(k, cluster)
	print k
	assign_cluster(im, k, cluster1, width, height, k_len)
	med_ks(k, cluster1)
	print k
	assign_cluster(im, k, cluster2, width, height, k_len)
	med_ks(k, cluster2)
	print k
	assign_cluster(im, k, cluster3, width, height, k_len)
	med_ks(k, cluster3)
	print k
	return k



def main():
	parser = argparse.ArgumentParser(description='Find main colors in a given image.')
	parser.add_argument("image", help="Image to be processed")
	parser.add_argument("-k",type=int, help="Custom K for KMeans algorithm")
	parser.add_argument("--output-format",type=str, choices=['image-palette', 'html-color-code'], help="Output-format")
	args = parser.parse_args()
	print args
	image = Image.open(args.image)
	k_len = 0
	if(args.k == None):
		k_len = 3
	else:
		k_len = args.k
	k = kmeans(image, k_len)
	if(args.output_format == "image-palette"):
		img = Image.new('RGB', (image.size[0]+100, image.size[1]))
		img.paste(image)
		imgs = []
		for k_i in range(0, k_len):
			imgs.append(Image.new('RGB', (100, image.size[1]/k_len), color = tuple(k[k_i])))
		for k_i in range(0, k_len):
			img.paste(imgs[k_i], box=(image.size[0], k_i*(image.size[1]/k_len)))
		img.save(args.image +"_pallette.png", "PNG")	
	else:
		print "#%02x%02x%02x" % (k[0][0],k[0][1],k[0][2])
		print "#%02x%02x%02x" % (k[1][0],k[1][1],k[1][2])
		print "#%02x%02x%02x" % (k[2][0],k[2][1],k[2][2])



if __name__ == "__main__":
	main()