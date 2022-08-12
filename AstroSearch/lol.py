import json
import string


class remLett:
	def __init__(self,keep=string.digits):
		self.comp = dict((ord(c),c) for c in keep)

	def __getitem__(self,k):
		return self.comp.get(k)

DD = remLett()


t = ''
with open('OG_Metadata.json', 'r') as f:

	data = f.read()
	t = t + data

	# print(data)
pf = json.loads(t)

Fdict = {}
AList =[]
llist = []
strbuffer = ''
for i in range(len(pf)):

	temp = pf[i]

	tr = temp['metadata']
	Name = tr['name']
	atts = tr['attributes']
	Numbers = Name.translate(DD)
	strbuffer = strbuffer + str(Numbers + ",")
	for i in atts:
		for key, value in i.items():
			strbuffer = strbuffer + str(value + ",")
	AList.append(strbuffer)
	strbuffer = ''


outfile = open('OGs.txt', 'w')
for i in AList:
	outfile.writelines(i + '\n')


