read = open('tmux.txt', 'r')
write = open('redis.conf', 'w')
in_lines = read.read().splitlines()
read.close()

i = 0
temp = ''
for line in in_lines:
    if i % 2 != 0:
        if len(line) == 0:
            print("line %d empty, %s" % (i, temp))
        else:
            temp += ' ' + line + '\n'
            write.writelines(temp)
    else:
        temp = line
    i += 1

write.close()
