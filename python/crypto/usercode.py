from Cryptodome.Hash import SHA256

def handle(data):
	hash_object = SHA256.new(data=b'First')
	data["token"] = hash_object.hexdigest()
	return data