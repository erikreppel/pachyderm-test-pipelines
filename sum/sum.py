import os
import sys


if len(sys.argv) is not 3:
    print "Usage: python square.py <repopath> <outputpath>"
    sys.exit(1)

print 'starting'

input_path = sys.argv[1]
output_path = sys.argv[2]

print input_path, output_path

if not os.path.isdir(input_path):
    print "Invalid input path"
    sys.exit(1)
if not os.path.isdir(output_path) and output_path != '/pfs/out':
    print "Invalid output path"
    sys.exit(1)

print [f for f in os.walk(input_path)]

files = []
for directory in os.walk(input_path):
    for f in directory[2]:
        files.append([os.path.join(directory[0], f), f])

print 'Found %s files' % len(files)
print files

total = 0

file_write_path = os.path.join(output_path, 'sum')

for item in files:
    print "Opening", item[0]
    file_read_path = item[0]

    with open(file_read_path, 'r') as f:
        for line in f:
            try:
                total += int(line)
                print 'current sum:', total
            except Exception as e:
                print e

print 'The final total was', total

with open(file_write_path, 'w') as f:
    f.write(str(total))