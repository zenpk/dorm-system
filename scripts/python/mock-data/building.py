def create(cursor):
    sql = (
        "INSERT INTO `buildings` (`num`,`enabled`,`info`)"
        "VALUES (%s,%s,%s);"
    )
    datas = [
        ['1', True, 'good'],
        ['2', True, 'big'],
        ['3', True, 'very big'],
        ['4', True, 'small'],
        ['5', True, 'medium'],
        ['6', False, 'very small']
    ]
    cursor.executemany(sql, datas)
    print("building finished")
