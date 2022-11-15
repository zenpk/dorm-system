import random


def create(cursor):
    sql = (
        "INSERT INTO `dorms` (`num`,`building_id`,`gender`,`remain_cnt`,`bed_cnt`)"
        "VALUES (%s,%s,%s,%s,%s);"
    )
    insert_data = []
    for i in range(1, 101):
        num = str(i)
        building_id = random.randint(1, 5)
        odd = random.random()
        if odd < 0.5:
            gender = '男'
        else:
            gender = '女'
        bed_cnt = random.randint(4, 8)
        remain_cnt = random.randint(1, bed_cnt)
        insert_data.append([num, building_id, gender, remain_cnt, bed_cnt])
    cursor.executemany(sql, insert_data)
    print('dorm finished')
