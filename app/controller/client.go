package controller

/**
* @api {post}  /client/create
* @apiName create
* @apiGroup client
* @apiPermission admin
* @apiVersion 1.0.0
* @apiDescription SendEmail 用于创建发邮件权限用户
* @apiParam {String}   user_name     超级用户账号						 // 查询字符串			nletech
* @apiParam {Number}   timestamp     时间戳							    // 查询字符串   当前请求的时间戳，必须为10位整数
* @apiParam {String}   sign          签名							    // 查询字符串			(base64 or hex)(hmac(api_key+timestamp, secret))   secret: nletech
* @apiParam {String}   api_key       api 认证key					    // post_body
* @apiParam {String}   secret        密钥							    // post_body
* @apiParam {String}   email_pre     用于邮件前缀						 // post_body
* @apiExample {curl} Example usage:
			curl -v -X POST \
			127.0.0.1:8080/client/create?api_key=nletech&timestamp=1550631506&sign=dcd29dc779289d187d841030d1edd4a3af14d15fd4d774225a76f8e9207766bd \
			-H 'content-type: application/json' \
			-d '{"api_key":"example",\
				"secret":"example",
				"email_pre":"example"}'

* @apiSuccessExample {json}  返回示例
* HTTP/1.1 200 OK
* {
* "ret": 1,
* "msg":"用户添加成功"
* "data": []
* }
*/
