# -*- coding: utf-8 -*-
# this script generate the csv file of mongodb data
import datetime
import pymongo
from config import config
import pandas as pd
import numpy as np


datetimetoint = lambda x:10000*x.year + 100*x.month + x.day

def mongo_cursor():
    conn = pymongo.MongoClient(config["mongodb_test"]["uri"])
    return conn

conn  = mongo_cursor()
db = conn["z3dbus"]

def gen_info_list(symbol):
    #time_list = db.Z3_EQUITY_HISTORY.find({'trade_date':{'$gte':start_time,'$lte':end_time}}).distinct('trade_date')
    #time_list = db.Z3_EQUITY_HISTORY.find({'trade_date':{'$gte':start_time,'$lte':end_time}}).distinct('trade_date')[:50]
    time_list = db.Z3_EQUITY_HISTORY.find({'trade_date':{'$gte':start_time,'$lte':end_time}}).distinct('trade_date')[:60]
    return [(symbol,y) for y in time_list]


def gen_symbol_close_px(info_list):
    # confirm it is not dead
    print(info_list[0])
    raw = map(find_day_inter,info_list)

    avg = round(np.mean(filter(lambda x:x!=0,raw)),2)
    maximum = max(raw)
    minimum = min(raw)
    data = [avg if x == 0 else x for x in raw]
    print(minimum,avg,len(data))
    return data


def find_day_inter(info):
    raw =  db.Z3_STK_MKT_DAY_INTER_NEW.find_one(
        {
            'end_date':datetimetoint(info[1]),
            'ex_status':1,
            'innerCode':info[0]
        },{
            'close_px':1
            },
        no_cursor_timeout=True)
    if raw:
        return raw.get('close_px',0)
    return 0

def gen_close_px_list(start_time,end_time):
    #stock_list = db.Z3_EQUITY_HISTORY.find({'trade_date':{'$gte':start_time,'$lte':end_time}}).distinct('innerCode')
    stock_list = db.Z3_EQUITY_HISTORY.find({'trade_date':{'$gte':start_time,'$lte':end_time}}).sort([('trade_date',1)]).distinct('innerCode')[:30]
    close_px_list = map(gen_symbol_close_px,map(gen_info_list,stock_list))
    for i in range(len(stock_list)):
        close_px_list[i] = [stock_list[i]] + close_px_list[i]
    return close_px_list

def gen_close_px_file(start_time,end_time):
    data = gen_close_px_list(start_time,end_time)
    df = pd.DataFrame(data)
    df = df.set_index([0])
    df.to_csv(str(start_time.date())+str(end_time.date())+".csv",header=False)

if __name__ == "__main__":
    start_time = datetime.datetime(2016,07,01)
    end_time = datetime.datetime(2017,07,01)
    #a = gen_close_px_list(start_time,end_time)
    gen_close_px_file(start_time,end_time)
    print("done")


