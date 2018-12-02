myfile = open('input.txt', 'r')
content = myfile.readlines()
myfile.close()
frequency = 0
for line in content:
    frequency += int(line) 
print(frequency)