import util

cnx = util.connect_db()
cursor = cnx.cursor()

insert = (
    "INSERT INTO `buildings` (`building_id`,`is_available`,`info`)"
    "VALUE (%s,%s,%s);"
)

datas = [
    [1, True, 'good'],
    [2, True, 'big'],
    [3, True, 'very big'],
    [4, False, 'small'],
    [5, True, 'medium'],
    [6, True, 'near bus'],
    [7, True, 'near metro'],
    [8, False, 'sunshine'],
    [9, True, 'near building'],
    [10, True, 'very small']
]

cursor.executemany(insert, datas)

cnx.commit()
cursor.close()
cnx.close()
print("finish")
