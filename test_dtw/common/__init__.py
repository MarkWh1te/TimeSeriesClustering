# coding: utf-8
"""
# @Time    : 2017/8/13 17:22
# @Author  : Kylin
# @File    : __init__.py.py
# @Software: PyCharm
# @Descript:
"""
#
# from dtw import dtw
# import numpy as np
#
# x = np.array([1, 2, 3, 4, 5]).reshape(-1, 1)
# x = np.array([5, 4, 3, 2, 1]).reshape(-1, 1)
# y = np.array([0, 0, 0, 0, 0]).reshape(-1, 1)
# # dist, cost, acc, path = dtw(x, y, dist=lambda x, y: norm(x - y, ord=1))
# dist, cost, acc, path = dtw(x, y, dist=lambda x, y: x - y)
# print dist
#
#
#
# import os
# os.environ["R_HOME"] = r'C:\Program Files\R\R-3.4.1'
# os.environ["R_USER"] = 'Ky'
# import numpy as np
#
# import rpy2.robjects.numpy2ri
# from rpy2.robjects.packages import importr
#
# # Set up our R namespaces
# R = rpy2.robjects.r
# DTW = importr('dtw')
#
# # Generate our data
# # idx = np.linspace(0, 2*np.pi, 100)
# # template = np.array([5,4,3,2,1]).reshape(-1, 1)
# template = np.array([0,0,0,0,0]).reshape(-1, 1)
# # print template
# query = np.array([1,2,3,4,5]).reshape(-1, 1)
# query = np.array([5,4,3,2,1]).reshape(-1, 1)
# # print query
# # template = np.array([0, 0, 1, 1, 2, 4, 2, 1, 2, 0]).reshape(-1, 1)
# # query = np.array([1, 1, 1, 2, 2, 2, 2, 3, 2, 0]).reshape(-1, 1)
#
# # Calculate the alignment vector and corresponding distance
# alignment = R.dtw(template,query,  keep=True)
# dist = alignment.rx('distance')
#
# print(dist)

from math import sqrt
import numpy as np
import random

import matplotlib.pylab as plt


def DTWDistance(s1, s2, w=5):
    DTW = {}

    w = max(w, abs(len(s1) - len(s2)))

    for i in range(-1, len(s1)):
        for j in range(-1, len(s2)):
            DTW[(i, j)] = float('inf')
    DTW[(-1, -1)] = 0

    for i in range(len(s1)):
        for j in range(max(0, i - w), min(len(s2), i + w)):
            #             print s1,s2
            dist = (s1[i] - s2[j]) ** 2
            DTW[(i, j)] = dist + min(DTW[(i - 1, j)], DTW[(i, j - 1)], DTW[(i - 1, j - 1)])

    return sqrt(DTW[len(s1) - 1, len(s2) - 1])


def k_means_clust(data, num_clust, num_iter, w=5):
    print len(data)
    print type(num_clust)
    centroids = random.sample(data, num_clust)
    #     centroids=[data[0],data[1]]
    #     print centroids
    counter = 0
    for n in range(num_iter):
        counter += 1
        #         print counter
        assignments = {}
        # assign data points to clusters
        for ind, i in enumerate(data):
            min_dist = float('inf')
            closest_clust = None
            for c_ind, j in enumerate(centroids):
                if LB_Keogh(i, j, 5) < min_dist:
                    cur_dist = DTWDistance(i, j, w)
                    if cur_dist < min_dist:
                        min_dist = cur_dist
                        closest_clust = c_ind
                        #             if closest_clust in assignments:
                        #                 assignments[closest_clust].append(ind)
                        #             else:
                        #                 assignments[closest_clust]=[]

            assignments.setdefault(closest_clust, [])
            assignments[closest_clust].append(ind)

        # recalculate centroids of clusters
        for key in assignments:
            clust_sum = 0
            for k in assignments[key]:
                #                 print "clust_sum",clust_sum,data[k]
                clust_sum = clust_sum + data[k]
            # print clust_sum
            centroids[key] = [m / len(assignments[key]) for m in clust_sum]

    return centroids, assignments


def LB_Keogh(s1, s2, r):
    LB_sum = 0
    for ind, i in enumerate(s1):

        lower_bound = min(s2[(ind - r if ind - r >= 0 else 0):(ind + r)])
        upper_bound = max(s2[(ind - r if ind - r >= 0 else 0):(ind + r)])

        if i > upper_bound:
            LB_sum = LB_sum + (i - upper_bound) ** 2
        elif i < lower_bound:
            LB_sum = LB_sum + (i - lower_bound) ** 2

    return sqrt(LB_sum)


def get_cluster(d_list, types):
    random.uniform(1, 10)

   # d_list = [[round(np.mean(line),2) if x == 0 else x for x in line] for line in d_list]
    #avg = round(np.mean(d_list),2)
    #d_list = [avg if x == 0 else x for x in d_list]

    d_list = [map(lambda x: x - line[0], line) for line in d_list]

    data = np.array(d_list)
    # data = d_list

    centroids, assignments = k_means_clust(data, types, 30, 4)
    return centroids, assignments
    # print centroids
    # print assignments
    for i in centroids:
        plt.plot(i)

    # plt.plot(data[0])
    # plt.plot(data[1])
    plt.show()
    for i in assignments.values():
        for key in i:
            plt.plot(data[key])

        plt.show()
