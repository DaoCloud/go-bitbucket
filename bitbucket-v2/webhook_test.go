package bitbucket_v2

import "testing"

func Test_Webhook_CURD(T *testing.T) {
	webhook, _ := NewWebhook("http://baidu.com", "test", true, nil)

	if err := client.CreateUpdateWebHook("POST", "xuzhenglun", "test2", webhook); err != nil || webhook.Uuid == "" {
		T.Error(err)
		return
	}

	webhook2 := webhook
	webhook2.Url = "http://sina.com"
	if err := client.GetDeleteWebHook("GET", "xuzhenglun", "test2", webhook2); err != nil {
		T.Error(err)
		return
	}

	if webhook2.Url != webhook.Url {
		T.Error("GET FAILED")
		return
	}

	webhook.Active = false
	if err := client.CreateUpdateWebHook("PUT", "xuzhenglun", "test2", webhook); err != nil {
		T.Error(err)
		return
	}

	if err := client.GetDeleteWebHook("GET", "xuzhenglun", "test2", webhook2); err != nil {
		T.Error(err)
		return
	}

	if webhook2.Active != false {
		T.Error("PUT FAILED")
		return
	}

	if err := client.GetDeleteWebHook("DELETE", "xuzhenglun", "test2", webhook2); err != nil {
		T.Log("DELETE FAILED")
		T.Error(err)
		return
	}

}