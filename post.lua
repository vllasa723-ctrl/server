wrk.method = "POST"
wrk.body   = string.rep("A", 3 * 1024 * 1024) -- 3 МБ в памяти
wrk.headers["Content-Type"] = "application/octet-stream"
