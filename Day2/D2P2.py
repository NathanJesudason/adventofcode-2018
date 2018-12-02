myfile = open('input.txt', 'r')
contents = myfile.read().strip().splitlines()
myfile.close()

for i in contents:
    for j in contents:
        diffChar = 0
        commonChar = ''
        for id, char in enumerate(i):
            if char != j[id]:
                diffChar += 1
            else:
                commonChar += char
        if diffChar == 1:
            print(commonChar)