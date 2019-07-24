package middleware

//Auth 一个中间件，用于验证请求是否合法
// func Auth(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) (err error) {
// 		// 单获取时间戳
// 		//------------------------------------------------------------------------
// 		timestamp := c.QueryParam("timestamp")
// 		if timestamp == "" {
// 			return c.JSON(http.StatusBadRequest, &cusresponse.ResponseFmt{Ret: 0, Msg: "时间戳缺失"})
// 		}

// 		num, err := strconv.ParseInt(timestamp, 10, 64)

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, &cusresponse.ResponseFmt{Ret: 0, Msg: "时间戳格式错误", Data: err.Error()})
// 		}
// 		//---------------------------------------------------------------------------

// 		// 绑定并验证数据
// 		//---------------------------------------------------------------------------
// 		// requestCheck := &cusvalidator.CheckIndentify{
// 		// 	Key:  c.QueryParam("api_key"),
// 		// 	Time: num,
// 		// 	Sign: c.QueryParam("sign"),
// 		// }

// 		validate := validator.New()
// 		err = validate.Struct(requestCheck)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, &cusresponse.ResponseFmt{Ret: 0, Msg: "数据验证错误", Data: err.Error()})
// 		}
// 		//---------------------------------------------------------------------------

// 		client := model.VerifyClient(requestCheck.Key)

// 		if client.ClientID == 0 {
// 			return c.JSON(http.StatusBadRequest, &cusresponse.ResponseFmt{Ret: 0, Msg: "用户不存在"})
// 		}

// 		timeNow := int64(time.Now().Unix()) // 获取 64 位的当前时间戳

// 		// 设置请求时间戳必须在一小时内
// 		if (timeNow-requestCheck.Time) < 0 || (timeNow-requestCheck.Time) > 3600 {
// 			return c.JSON(http.StatusBadRequest, &cusresponse.ResponseFmt{Ret: 0, Msg: "请求已过期"})
// 		}

// 		// 对传入的参数进行真实性验证  // 此处签名曾出过问题 签名验证（兼容 base64 和 hex）
// 		//--------------------------------------------------------------------------
// 		mac := hmac.New(sha256.New, []byte(client.APISecret))
// 		mac.Write([]byte(requestCheck.Key + strconv.FormatInt(requestCheck.Time, 10)))
// 		expectedSign := mac.Sum(nil)

// 		requestSign, err := base64.URLEncoding.DecodeString(requestCheck.Sign)

// 		if len(requestCheck.Sign) == 64 { // 兼容 16 进制
// 			requestSign, err = hex.DecodeString(requestCheck.Sign)
// 		} else if len(requestCheck.Sign) != 44 {
// 			return c.JSON(http.StatusBadRequest, &cusresponse.ResponseFmt{Ret: 0, Msg: "签名长度错误"})
// 		}
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, &cusresponse.ResponseFmt{Ret: 0, Msg: "签名解码错误", Data: err.Error()})
// 		}

// 		expectedSignStr := base64.URLEncoding.EncodeToString(expectedSign)

// 		// mac := hmac.New(sha256.New, []byte(client.APISecret))
// 		// mac.Write([]byte(requestCheck.Key + strconv.FormatInt(requestCheck.Time, 10)))
// 		// expectedSign := mac.Sum(nil)
// 		// Sign := base64.URLEncoding.EncodeToString(expectedSign)

// 		log.Print("请求时间", strconv.FormatInt(requestCheck.Time, 10))
// 		log.Print("请求字符串：", requestCheck.Key+strconv.FormatInt(requestCheck.Time, 10))
// 		log.Print("请求签名：", requestCheck.Sign, "\n")
// 		log.Print("计算编码：", expectedSignStr, "\n")
// 		// log.Print("请求密钥", client.APISecret, "\n")
// 		// log.Print("请求密钥字节", []byte(client.APISecret), "\n")
// 		// log.Print("请求解码：", requestSign, "\n")
// 		// log.Print("请求hash", mac)
// 		// log.Print("计算签名：", expectedSign, "\n")

// 		if !hmac.Equal(requestSign, expectedSign) {
// 			return c.JSON(http.StatusBadRequest, &cusresponse.ResponseFmt{Ret: 0, Msg: "签名验证失败"})
// 		}
// 		//--------------------------------------------------------------------------

// 		// Set example variable
// 		c.Set("client", client)

// 		return next(c)

// 	}
// }
