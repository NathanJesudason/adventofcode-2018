myfile = open('input.txt', 'r')
content = myfile.readlines()
myfile.close()
frequency = 0
#Originally used list, but set is more efficient.
frequencyList = {0}
noAns = True
while noAns:
    for line in content:
        frequency += int(line) 
        if frequency in frequencyList:
            print(frequency)
            noAns = False
            break
        else:
            frequencyList.add(frequency)