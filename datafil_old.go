// 河北金融学院 OA 疫情数据自动填报
// Powered By Luckykeeper <luckykeeper@luckykeeper.site | https://luckykeeper.site> 2022/09/03
package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

const (
	oaUsername = ""   // OA 账号
	oaPassword = ""   // OA 密码
	provinceId = ""   // 省 ID
	cityId     = ""   // 城市 ID
	areaId     = ""   // 区 ID
	streetId   = ""   // 街道 ID
	address    = ""   // 详细地址
	prove      = true //核酸检测证明，有填 true ，没有填 false
	fillTime   = 1    //没有填报的表数，用来补齐过去未填写的，如果没有未填写的就填1
)

func main() {
	for i := 0; i < fillTime; i++ {
		datafill()
	}
}

// 数据填报
func datafill() {
	var pic1 []byte // debug 使用
	var pic0, pic2, pic3, pic4 []byte
	// create context
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true)) // debug(false)|prod(true)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(
		allocCtx,
	)
	defer cancel()

	if err := chromedp.Run(ctx,
		// OA酱 登录页配置
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate("https://oa.hbfu.edu.cn/backstage/cas/login?service=https%3A%2F%2Foa.hbfu.edu.cn%2Fbackstage%2Fcas-proxy%2Fapp%2Fredirect"),
		chromedp.WaitVisible(`body > app-root > app-theme1 > div > main > div.right-container.box-shadow > app-login-panel > div > nz-tabset > div.ant-tabs-content.ant-tabs-top-content.ant-tabs-content-animated > div.ant-tabs-tabpane.ant-tabs-tabpane-active > app-userpass-panel > form > input:nth-child(1)`),
		chromedp.Sleep(2*time.Second),
		chromedp.FullScreenshot(&pic0, 90),
		chromedp.Click(`body > app-root > app-theme1 > div > main > div.right-container.box-shadow > app-login-panel > div > nz-tabset > div.ant-tabs-content.ant-tabs-top-content.ant-tabs-content-animated > div.ant-tabs-tabpane.ant-tabs-tabpane-active > app-userpass-panel > form > input:nth-child(1)`, chromedp.ByQuery),
		chromedp.SendKeys(`body > app-root > app-theme1 > div > main > div.right-container.box-shadow > app-login-panel > div > nz-tabset > div.ant-tabs-content.ant-tabs-top-content.ant-tabs-content-animated > div.ant-tabs-tabpane.ant-tabs-tabpane-active > app-userpass-panel > form > input:nth-child(1)`, oaUsername, chromedp.ByQuery),
		chromedp.Click(`body > app-root > app-theme1 > div > main > div.right-container.box-shadow > app-login-panel > div > nz-tabset > div.ant-tabs-content.ant-tabs-top-content.ant-tabs-content-animated > div.ant-tabs-tabpane.ant-tabs-tabpane-active > app-userpass-panel > form > input:nth-child(2)`, chromedp.ByQuery),
		chromedp.SendKeys(`body > app-root > app-theme1 > div > main > div.right-container.box-shadow > app-login-panel > div > nz-tabset > div.ant-tabs-content.ant-tabs-top-content.ant-tabs-content-animated > div.ant-tabs-tabpane.ant-tabs-tabpane-active > app-userpass-panel > form > input:nth-child(2)`, oaPassword, chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.SendKeys(`body > app-root > app-theme1 > div > main > div.right-container.box-shadow > app-login-panel > div > nz-tabset > div.ant-tabs-content.ant-tabs-top-content.ant-tabs-content-animated > div.ant-tabs-tabpane.ant-tabs-tabpane-active > app-userpass-panel > form > input:nth-child(2)`, kb.Enter, chromedp.ByQuery),
		chromedp.FullScreenshot(&pic1, 90),
		chromedp.Sleep(2*time.Second),
		// OA酱 - 跳转到首页完成登录流程
		chromedp.Navigate("https://oa.hbfu.edu.cn/new/angular-office-hall/#/angular-office-hall/index"),
		chromedp.Sleep(2*time.Second),
		chromedp.WaitVisible(`#hallBody > app-root > app-index > nz-layout > app-header > div > div > div:nth-child(2) > div > div > div > ul > li:nth-child(2)`, chromedp.ByQuery),
		chromedp.FullScreenshot(&pic2, 90),
		// OA酱 - 跳转到数据填报页面
		chromedp.Navigate("https://oa.hbfu.edu.cn/datafill/collect/usertask"),
		chromedp.WaitVisible(`#root > div > section > section > main > div > div.antd-pro-components-page-header-wrapper-index-content > div > div > div > div.ant-tabs.ant-tabs-top.ant-tabs-line > div.ant-tabs-content.ant-tabs-content-animated.ant-tabs-top-content > div.ant-tabs-tabpane.ant-tabs-tabpane-active > div.antd-pro-pages-collect-index-tableList > div.ant-table-wrapper > div > div > div > div > div > table > tbody > tr:nth-child(1) > td.ant-table-selection-column > span > label > span > input`, chromedp.ByQuery),
		chromedp.FullScreenshot(&pic3, 90),
		// OA酱 - 开始数据填报
		// 选择第一项
		chromedp.Click(`#root > div > section > section > main > div > div.antd-pro-components-page-header-wrapper-index-content > div > div > div > div.ant-tabs.ant-tabs-top.ant-tabs-line > div.ant-tabs-content.ant-tabs-content-animated.ant-tabs-top-content > div.ant-tabs-tabpane.ant-tabs-tabpane-active > div.antd-pro-pages-collect-index-tableList > div.ant-table-wrapper > div > div > div > div > div > table > tbody > tr:nth-child(1) > td.ant-table-selection-column > span > label > span > input`, chromedp.ByQuery),
		chromedp.WaitVisible(`#root > div > section > section > main > div > div.antd-pro-components-page-header-wrapper-index-content > div > div > div > div.ant-tabs.ant-tabs-top.ant-tabs-line > div.ant-tabs-content.ant-tabs-content-animated.ant-tabs-top-content > div.ant-tabs-tabpane.ant-tabs-tabpane-active > div.antd-pro-pages-collect-index-tableList > div.antd-pro-pages-collect-index-tableListOperator > span > span > button`, chromedp.ByQuery),
		chromedp.Click(`#root > div > section > section > main > div > div.antd-pro-components-page-header-wrapper-index-content > div > div > div > div.ant-tabs.ant-tabs-top.ant-tabs-line > div.ant-tabs-content.ant-tabs-content-animated.ant-tabs-top-content > div.ant-tabs-tabpane.ant-tabs-tabpane-active > div.antd-pro-pages-collect-index-tableList > div.antd-pro-pages-collect-index-tableListOperator > span > span > button`, chromedp.ByQuery),
		chromedp.WaitVisible(`body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div`, chromedp.ByQuery),
		// 居住地址
		chromedp.Sleep(1*time.Second),
		chromedp.Click(`/html/body/div[3]/div/div[2]/div/div[2]/div/div/div/form/div/fieldset/div[1]/div[1]/span/input`, chromedp.BySearch),
		chromedp.Sleep(1*time.Second),
		chromedp.Click("body > div:nth-child(9) > div > div > div > ul:nth-child(1) > li:nth-child("+provinceId+")", chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
		chromedp.Click("body > div:nth-child(9) > div > div > div > ul:nth-child(2) > li:nth-child("+cityId+")", chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
		chromedp.Click("body > div:nth-child(9) > div > div > div > ul:nth-child(3) > li:nth-child("+areaId+")", chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
		chromedp.Click("body > div:nth-child(9) > div > div > div > ul:nth-child(4) > li:nth-child("+streetId+")", chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
		chromedp.SendKeys(`body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(1) > div:nth-child(3) > input`, address, chromedp.ByQuery),
		// 体温
		chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(2) > div > div", kb.Enter, chromedp.ByQuery),
		chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(2) > div > div", kb.ArrowDown, chromedp.ByQuery),
		chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(2) > div > div", kb.Enter, chromedp.ByQuery),

		// chromedp.Click("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(2) > div > div", chromedp.ByQuery),
		// chromedp.Click(`//*[@id="2710d0cc-8305-48ee-b137-4dfa97b7e1c6"]/ul/li[2]`, chromedp.BySearch),
		// 隔离情况
		chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(3) > div > div", kb.Enter, chromedp.ByQuery),
		chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(3) > div > div", kb.ArrowDown, chromedp.ByQuery),
		chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(3) > div > div", kb.Enter, chromedp.ByQuery),
		// chromedp.Click("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(3) > div > div > div", chromedp.ByQuery),
		// chromedp.Click("#ac6c4775-f871-472a-c42b-18e3fbece1d6 > ul > li:nth-child(2)", chromedp.ByQuery),

	); err != nil {
		log.Fatal(err)
	}

	// 核酸情况
	if prove { // 有核酸证明
		if err := chromedp.Run(ctx,
			chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(4) > div > div", kb.Enter, chromedp.ByQuery),
			chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(4) > div > div", kb.ArrowDown, chromedp.ByQuery),
			chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(4) > div > div", kb.ArrowDown, chromedp.ByQuery),
			chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(4) > div > div", kb.Enter, chromedp.ByQuery),
			// chromedp.Click("#d7f4ff68-99de-43f3-b130-b5101fced23a > ul > li:nth-child("+prove+")", chromedp.ByQuery),); err != nil {
			chromedp.FullScreenshot(&pic4, 90),
		); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := chromedp.Run(ctx,
			chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(4) > div > div", kb.Enter, chromedp.ByQuery),
			chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(4) > div > div", kb.ArrowDown, chromedp.ByQuery),
			chromedp.SendKeys("body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > div > form > div > fieldset > div:nth-child(4) > div > div", kb.Enter, chromedp.ByQuery),
			// chromedp.Click("#d7f4ff68-99de-43f3-b130-b5101fced23a > ul > li:nth-child("+prove+")", chromedp.ByQuery),); err != nil {
			chromedp.FullScreenshot(&pic4, 90),
		); err != nil {
			log.Fatal(err)
		}
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`body > div:nth-child(8) > div > div.ant-modal-wrap > div > div.ant-modal-content > div > div > ul > li:nth-child(3) > span > button`, chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	} else {
		log.Println("顺利完成填报!")
	}

	ioutil.WriteFile("oaLoginPage.png", pic0, 0o644)
	ioutil.WriteFile("oaLoginPageWithAccount.png", pic1, 0o644)
	ioutil.WriteFile("oaOfficeHall.png", pic2, 0o644)
	ioutil.WriteFile("oaDataFillList.png", pic3, 0o644)
	ioutil.WriteFile("oaDataFilled.png", pic4, 0o644)
}
