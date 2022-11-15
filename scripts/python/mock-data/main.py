import building
import dorm
import temp
import user
import util


def main():
    cnx = util.connect_db()
    cursor = cnx.cursor()
    temp.create(cursor)
    building.create(cursor)
    dorm.create(cursor)
    user.create(cursor)
    cursor.close()
    cnx.commit()
    cnx.close()
    print('finished!')


if __name__ == '__main__':
    main()
