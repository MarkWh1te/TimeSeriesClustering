# coding: utf-8
"""
# @Time    : 2017/8/13 13:29
# @Author  : Kylin
# @File    : main.py
# @Software: PyCharm
# @Descript:
"""
from flask import Flask, render_template, url_for, jsonify, request
import random
from common import get_cluster
from itertools import groupby
from data.data import  load_data


app = Flask(__name__)
app.data = load_data('data/2016-07-012017-07-01.csv')

@app.route('/')
def main():
    url_for('static', filename='style.css')
    return render_template('main.html',
                           # name=name)
                           )


@app.route('/cluster', methods=['POST'])
def cluster():
    days = int(request.form.get('days'))
    types = int(request.form.get('types'))
    print days, types


    s_data = {
        "000001.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000042.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000033.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000024.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000015.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000006.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000007.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000008.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000009.SZ": map(lambda x: random.uniform(1, 10), range(10)),
        "000010.SZ": map(lambda x: random.uniform(1, 10), range(10)),
    }

	# use mongo data

    s_data = app.data
    sort_keys = sorted(s_data.keys())
    d = [s_data[i] for i in sort_keys]
    print('fff')
    d = get_cluster(d, types, )
    print('dd')
    return jsonify({"source": s_data, "sort_keys": sort_keys, "cluster": d[1], "centers": d[0]})


@app.route('/stock_list')
def stock_list():
    return  # 股票列表


@app.route('/type_list')
def type_list():
    return  # 类型列表


if __name__ == '__main__':
	app.run(host='0.0.0.0',port=4444)
