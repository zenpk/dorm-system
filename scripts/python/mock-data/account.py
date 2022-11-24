import bcrypt as bcrypt


def create(cursor, user_id, username):
    sql = (
        "INSERT INTO `accounts` (`user_id`,`username`,`password`)"
        "VALUE (%s,%s,%s)"
    )
    salt = bcrypt.gensalt(rounds=10)
    byte = username.encode('utf-8')
    password = bcrypt.hashpw(byte, salt)
    cursor.execute(sql, [user_id, username, password])
