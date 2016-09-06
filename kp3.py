from wsgiref.simple_server import make_server
from urllib2 import urlparse

data = {}

def serve(env, start_response):
	path = env['PATH_INFO']
	args = urlparse.parse_qs(env['QUERY_STRING'])
	key = args['key'][0]

	start_response('200 OK', [('Content-type', 'text/plain')])

	if path == '/get':
		return data[key]
	if path == '/set':
		value = args['value'][0]
		data[key] = value
		return ['ok']
	return ['unknown error']

make_server('', 8080, serve).serve_forever()
