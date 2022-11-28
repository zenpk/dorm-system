def create(cursor):
    sql_team = (
        "INSERT INTO `teams` (`code`,`gender`,`owner_id`)"
        "VALUE (%s,%s,%s);"
    )
    cursor.execute(sql_team, ['test_code', '男', '1'])
    rels = [
        ['1', '2'],
        ['1', '3'],
        ['1', '4']
    ]
    sql_rel = (
        "INSERT INTO `team_users` (`team_id`,`user_id`)"
        "VALUE (%s,%s);"
    )
    cursor.executemany(sql_rel, rels)
    print('team finished')
