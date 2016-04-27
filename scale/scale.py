import os
import sys


if len(sys.argv) is not 4:
    print "Usage: python scale.py <repopath> <scalepath> <outputpath>"
    sys.exit(1)

input_path = sys.argv[1]
scale_path = sys.argv[2]
output_path = sys.argv[3]

print input_path, output_path

if not os.path.isdir(input_path):
    print "Invalid input path"
    sys.exit(1)
if not os.path.isdir(scale_path):
    print "Invalid scale path"
    sys.exit(1)
if not os.path.isdir(output_path) and output_path != '/pfs/out':
    print "Invalid output path"
    sys.exit(1)


input_files = []
for directory in os.walk(input_path):
    for f in directory[2]:
        input_files.append([os.path.join(directory[0], f), f])

scale_files = []
for directory in os.walk(scale_path):
    for f in directory[2]:
        scale_files.append([os.path.join(directory[0], f), f])


print 'Found %s files to scale by' % len(scale_files)
print 'scale files', scale_files
print 'input files', input_files

scale_factor = 0

with open(scale_files[0][0], 'r') as f:
    scale_factor = int(f.readline())

for f in input_files:
    print 'opening', f

    write_file_path = os.path.join(output_path, f[1])
    print write_file_path
    with open(write_file_path, 'w') as out:
        with open(f[0], 'r') as inputf:
            for line in inputf:
                try:
                    product = scale_factor * int(line)
                    out.write(str(product) + '\n')
                    print product
                except Exception as e:
                    print str(e)
