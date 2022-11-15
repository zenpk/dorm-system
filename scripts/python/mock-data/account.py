import random
import string

import bcrypt as bcrypt


def create(cursor, user_id):
    sql = (
        "INSERT INTO `accounts` (`user_id`,`username`,`password`)"
        "VALUE (%s,%s,%s)"
    )
    username = str(''.join(random.choices(string.ascii_letters, k=6)))
    salt = bcrypt.gensalt(rounds=10)
    byte = username.encode('utf-8')
    password = bcrypt.hashpw(byte, salt)
    cursor.execute(sql, [user_id, username, password])
