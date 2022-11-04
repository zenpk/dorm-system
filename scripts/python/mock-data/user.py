import random
import string
import util
from faker import Faker

cnx = util.connect_db()
cursor = cnx.cursor()
insert = (
    "INSERT INTO `user_infos` (`user_credential_id`,`username`,`student_id`,`gender`,`name`)"
    "VALUE (%s,%s,%s,%s,%s);"
)
fake = Faker("zh_CN")

insert_data = []
for i in range(1, 1001):
    credential_id = str(i)
    student_id = str(i)
    username = ''.join(random.choices(string.ascii_letters, k=6))
    odd = random.random()
    if odd < 0.5:
        gender = '男'
    else:
        gender = '女'
    name = fake.name()
    insert_data.append([credential_id, username, student_id, gender, name])

print(len(insert_data))
cursor.executemany(insert, insert_data)

cnx.commit()
cursor.close()
cnx.close()
print('finish')
