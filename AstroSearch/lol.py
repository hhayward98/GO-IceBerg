import json
import string


class remLett:
	def __init__(self,keep=string.digits):
		self.comp = dict((ord(c),c) for c in keep)

	def __getitem__(self,k):
		return self.comp.get(k)


def main():

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
	IMGlist = []
	strbuffer = ''
	for i in range(len(pf)):

		temp = pf[i]

		tr = temp['metadata']
		Name = tr['name']
		atts = tr['attributes']
		IMG = tr['image']

		Numbers = Name.translate(DD)
		IMGtuple = (Numbers, IMG)
		IMGlist.append(IMGtuple)

		strbuffer = strbuffer + str(Numbers + ",")
		for i in atts:
			for key, value in i.items():
				strbuffer = strbuffer + str(value + ",")
		AList.append(strbuffer)
		strbuffer = ''


	outfile = open('OGsIMG.txt', 'w')
	OneStr = ""
	for i in IMGlist:
		OneStr = str(i[0]) + ","+str(i[1])
		outfile.writelines(OneStr + '\n')


if __name__ == '__main__':
	main()
