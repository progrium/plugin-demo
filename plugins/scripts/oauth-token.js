
var oauthToken = "token123"

function implements() {
	return ["RequestFilter"]
}

function FilterRequest() {
	query = req.URL.RawQuery.split("&")
	for (i in query) {
		kvp = query[i].split("=")
		if (kvp[0] == "oauth_token" && kvp[1] == oauthToken) {
			return [true, "", 0]
		}
	}
	return [false, "missing or incorrect oauth_token", 403]
}