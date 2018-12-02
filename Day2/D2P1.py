myfile = open('input.txt', 'r')
contents = myfile.read().strip().splitlines()
myfile.close()

Count2 = 0
Count3 = 0
for line in contents:
    first2 = True
    first3 = True
    for char in line:
        charCount = line.count(char)
        if (charCount == 2) and (first2):
            first2 = False
            Count2 += 1
        elif (charCount == 3) and (first3):
            first3 = False
            Count3 += 1
print(Count2 * Count3)