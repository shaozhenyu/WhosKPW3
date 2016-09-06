var http = require('http')
var urlp = require("url");
var qs = require('querystring');

var data = {}

var server = http.createServer(function (request, response) {
	response.writeHead(200, {'Content-Type': 'text/plain'});
	var url = urlp.parse(request.url);
	var query = qs.parse(url.query);
	if (url.pathname == '/get') {
		response.end('value=' + data[query['key']]);
	} else if (url.pathname == '/set') {
		data[query['key']] = query['value'];
		response.end('ok');
	}
});

server.listen(8080);
