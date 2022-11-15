import bcrypt


def create(cursor):
    user_sql = (
        "INSERT INTO `users` (`student_num`,`name`,`gender`)"
        "VALUE (%s,%s,%s);"
    )
    student_num = str(0)
    gender = '男'
    cursor.execute(user_sql, [student_num, '测试用户', gender])
    user_id = cursor.lastrowid
    account_sql = (
        "INSERT INTO `accounts` (`user_id`,`username`,`password`)"
        "VALUE (%s,%s,%s)"
    )
    username = 'temp'
    salt = bcrypt.gensalt(rounds=10)
    byte = username.encode('utf-8')
    password = bcrypt.hashpw(byte, salt)
    cursor.execute(account_sql, [user_id, username, password])
    print('temp finished')
