def create(cursor):
    sql = (
        "INSERT INTO `buildings` (`num`,`enabled`,`info`,`img_url`)"
        "VALUES (%s,%s,%s,%s);"
    )
    datas = [
        ['1', True, 'good', 'https://i.imgur.com/WDmvSab.jpeg'],
        ['2', True, 'big', 'https://i.imgur.com/RqoqseF.jpeg'],
        ['3', True, 'very big', 'https://i.imgur.com/fCM0iN2.jpeg'],
        ['4', True, 'small', 'https://i.imgur.com/NLfHBMS.jpeg'],
        ['5', True, 'medium', 'https://i.imgur.com/AzHs6II.jpeg'],
        ['6', False, 'very small', 'https://i.imgur.com/Uc8FNVB.jpeg']
    ]
    cursor.executemany(sql, datas)
    print("building finished")
