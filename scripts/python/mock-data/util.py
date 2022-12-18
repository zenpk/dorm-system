import random

import mysql.connector


def connect_db():
    cnx = mysql.connector.connect(user='dorm',
                                  password='dorm',
                                  host='127.0.0.1',
                                  port=3306,
                                  database='dorm',
                                  charset='utf8mb4')
    return cnx


def get_from_list(item_list):
    return item_list[random.randint(0, len(item_list) - 1)]
