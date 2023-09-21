package implement

//func (impl *implementation) makeActivateEmailData(email, activateToken *string, activateExpiredAt *int64) *inout.MsgBrokerSender {
//	var date = impl.DateTime.ConvertUnixToDateTimeWithFormat(*activateExpiredAt, "02")
//	var time = impl.DateTime.ConvertUnixToDateTimeWithFormat(*activateExpiredAt, "15.04")
//	var thaiMonth = impl.DateTime.ConvertUnixToDateTimeWithFormat(*activateExpiredAt, "01")
//	var year = impl.DateTime.ConvertUnixToDateTimeWithFormat(*activateExpiredAt, "2006")
//	var buddhaEra, _ = strconv.Atoi(year)
//	buddhaEra += 543
//	dateTimeString := fmt.Sprintf("%s %s %d เวลา %s น.", date, monthList[thaiMonth], buddhaEra, time)
//
//	msg := fmt.Sprintf("Activation Link ของคุณคือ %s/activation?token=%s     Link จะหมดอายุภายใน %s", impl.Config.ActivationDomain, *activateToken, dateTimeString)
//	var contentHtml string
//	contentHtml, err := impl.makeContentHTML(*activateToken, &dateTimeString, email)
//	if err != nil {
//		impl.Log.Error("Error make content html:", err.Error())
//		contentHtml = fmt.Sprintf("<p>%s</p>", msg)
//	}
//
//	return &inout.MsgBrokerSender{
//		Action:  msgBrokerIO.ActionCreate,
//		Contact: *email,
//		EmailData: domain.Email{
//			To:           *email,
//			ToEmail:      *email,
//			Subject:      "See It Live Activation Link",
//			ContentPlain: msg,
//			ContentHtml:  contentHtml,
//		},
//	}
//}
//
//func (impl *implementation) makeContentHTML(activateToken string, activateExpiredAt, email *string) (string, error) {
//	var buffer bytes.Buffer
//	t, err := template.ParseFiles("html/template/activation_link.html")
//	if err != nil {
//		return "", err
//	}
//
//	data := map[string]interface{}{
//		"activation_link": fmt.Sprintf("%s/activation?token=%s", impl.Config.ActivationDomain, activateToken),
//		"expired_at":      activateExpiredAt,
//		"email":           email,
//	}
//	if err = t.Execute(&buffer, data); err != nil {
//		return "", err
//	}
//	return buffer.String(), nil
//}
//
//func (impl *implementation) sendMsgRequest(input *inout.MsgBrokerSender) (err error) {
//	return impl.MsgSender(msgBrokerIO.TopicActivateLink, input)
//}
