import random
import string

import util

cnx = util.connect_db()
cursor = cnx.cursor()
insert = (
    "INSERT INTO `dorms` (`dorm_id`,`building_id`,`gender`,`available`,`bed_num`,`info`)"
    "VALUE (%s,%s,%s,%s,%s,%s);"
)

insert_data = []
for i in range(1, 1001):
    dorm_id = str(i)
    building_id = random.randint(1, 10)
    info = ''.join(random.choices(string.ascii_letters, k=6))
    odd = random.random()
    if odd < 0.5:
        gender = '男'
    else:
        gender = '女'
    bed_num = random.randint(4, 8)
    available = random.randint(1, bed_num)
    insert_data.append([dorm_id, building_id, gender, available, bed_num, info])

print(len(insert_data))
cursor.executemany(insert, insert_data)

cnx.commit()
cursor.close()
cnx.close()
print('finish')
