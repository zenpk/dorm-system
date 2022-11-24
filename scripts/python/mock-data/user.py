import random
import account
from faker import Faker


def create(cursor):
    sql = (
        "INSERT INTO `users` (`student_num`,`name`,`gender`)"
        "VALUE (%s,%s,%s);"
    )
    fake = Faker("zh_CN")
    for i in range(1, 1001):
        student_num = str(i)
        odd = random.random()
        if odd < 0.5:
            gender = '男'
        else:
            gender = '女'
        name = fake.name()
        cursor.execute(sql, [student_num, name, gender])
        user_id = cursor.lastrowid
        account.create(cursor, user_id, 'stu' + str(i))
    print('users finished')
