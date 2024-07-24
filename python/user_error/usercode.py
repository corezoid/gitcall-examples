def handle(data):    
	
	if "error" in data:
		raise ValueError('My custom user ValueError')
	else:
		raise Exception('My custom user Exception')        
	
	return data