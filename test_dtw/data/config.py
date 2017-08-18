# coding:utf-8 # Created by qinlin.liu at 2017/3/14

config = {
    "mongodb_test": {
        "host": "117.121.98.91",
        "replicaset": "athena",
	    #"uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.37:27017,10.77.4.38:27017/z3dbus?authMechanism=SCRAM-SHA-1&replicaSet=zntytestdb"
	    #"uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.37:27017/z3dbus?authMechanism=SCRAM-SHA-1"
	    #"uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.37:27017/z3dbus?authMechanism=SCRAM-SHA-1"
	    "uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.45:27017/z3dbus?authMechanism=SCRAM-SHA-1"
	    #"uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.11:27017/z3dbus?authMechanism=SCRAM-SHA-1"
    },
    "mongodb_new": {
        "host": "117.121.98.91",
        "replicaset": "athena",
	    "uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.45:27017/z3dbus?authMechanism=SCRAM-SHA-1"
	    #"uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.37:27017,10.77.4.38:27017/z3dbus?authMechanism=SCRAM-SHA-1&replicaSet=zntytestdb"
	    #"uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.37:27017/z3dbus?authMechanism=SCRAM-SHA-1"
	    #"uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.37:27017/z3dbus?authMechanism=SCRAM-SHA-1"
	    #"uri": "mongodb://z3dbusadmin:z3dbusadmin@10.77.4.11:27017/z3dbus?authMechanism=SCRAM-SHA-1"
    },
    "mongodb_local": {
        "host": "127.0.0.1"
    },
}
