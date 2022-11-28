import bcrypt


def create_one(cursor, student_num, name, username):
    user_sql = (
        "INSERT INTO `users` (`student_num`,`name`,`gender`)"
        "VALUE (%s,%s,%s);"
    )
    gender = '男'
    cursor.execute(user_sql, [student_num, name, gender])
    user_id = cursor.lastrowid
    account_sql = (
        "INSERT INTO `accounts` (`user_id`,`username`,`password`)"
        "VALUE (%s,%s,%s)"
    )
    salt = bcrypt.gensalt(rounds=10)
    byte = username.encode('utf-8')
    password = bcrypt.hashpw(byte, salt)
    cursor.execute(account_sql, [user_id, username, password])


def create_all(cursor):
    create_one(cursor, '1', '测试用户', 'temp')
    create_one(cursor, '2', '测试队友1', 'temp1')
    create_one(cursor, '3', '测试队友2', 'temp2')
    create_one(cursor, '4', '测试队友3', 'temp3')
    print('temp finished')
