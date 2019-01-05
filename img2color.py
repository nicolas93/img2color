#!/usr/bin/env python2

from PIL import Image
import sys
import random

def color_diff(a, b):
	d0 = a[0]-b[0]
	d1 = a[1]-b[1]
	d2 = a[2]-b[2]
	d0 = d0 if d0>0 else d0*-1
	d1 = d1 if d1>0 else d1*-1
	d2 = d2 if d2>0 else d2*-1
	return d0+d1+d2

def med(cluster):
	r,g,b = (0,0,0)
	for c in cluster:
		r = r + c[0]
		g = g + c[1]
		b = b + c[2]
	r = r/len(cluster)
	g = g/len(cluster)
	b = b/len(cluster)
	return [r,g,b]


def assign_cluster(k, cluster):
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


im = Image.open(sys.argv[1])

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
assign_cluster(k, cluster)
k[0] = med(cluster[0])
k[1] = med(cluster[1])
k[2] = med(cluster[2])
print k
assign_cluster(k, cluster1)
k[0] = med(cluster1[0])
k[1] = med(cluster1[1])
k[2] = med(cluster1[2])
print k
assign_cluster(k, cluster2)
k[0] = med(cluster2[0])
k[1] = med(cluster2[1])
k[2] = med(cluster2[2])
print k
assign_cluster(k, cluster3)
k[0] = med(cluster3[0])
k[1] = med(cluster3[1])
k[2] = med(cluster3[2])

print k
print len(cluster3[0])
print len(cluster3[1])
print len(cluster3[2])
print width
print height