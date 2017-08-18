# -*- coding: utf-8 -*-
import pandas as pd


def load_data(file_name):
	return clean(pd.read_csv(file_name))

def clean(df):
	return dict(map(lambda x:(x[0],x[1:]),df.values.tolist()))


if __name__ == "__main__":
	df = load_data('2016-07-012017-07-01.csv')
