def create(cursor):
    sql = (
        "INSERT INTO `buildings` (`num`,`enabled`,`info`,`img_url`)"
        "VALUES (%s,%s,%s,%s);"
    )
    datas = [
        ['1', True, 'good', 'https://imgur.com/gallery/BxBNP59'],
        ['2', True, 'big', 'https://imgur.com/gallery/RwmqDBX'],
        ['3', True, 'very big', 'https://imgur.com/gallery/9QHb2P0'],
        ['4', True, 'small', 'https://imgur.com/gallery/O7Km8nH'],
        ['5', True, 'medium', 'https://imgur.com/gallery/NBX3r'],
        ['6', False, 'very small', 'https://imgur.com/gallery/FJvcZ']
    ]
    cursor.executemany(sql, datas)
    print("building finished")
