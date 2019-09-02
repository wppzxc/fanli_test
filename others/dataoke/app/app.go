package app

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"k8s.io/klog"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	quanUrl     = "http://www.dataoke.com/qlist/?px=sell&page=%d"
	topUrl      = "http://www.dataoke.com/top_sell"
	getTblUrl   = "http://www.dataoke.com/gettpl?gid=%s&_=%s"
	goodsListId = ".goods-list"
	quanDataId  = "data-id"
	topId       = "id"
)

const (
	quanFile = "./quan.txt"
	topFile  = "./top.txt"
)

func StartQuan(begin int, end int) error {
	if begin > end {
		klog.Warningf("begin page can't bigger than end, use begin = end")
		begin = end
	}
	klog.Infof("execute url is : %s", quanUrl)
	var items []string
	for i := begin; i <= end; i++ {
		url := fmt.Sprintf(quanUrl,i)
		tmp, err := getQuanItems(url)
		if err != nil {
			klog.Errorf("Error in get quan Items page %d : %s ", i, err)
			continue
		}
		items = append(items, tmp...)
	}
	links, err := getTblinks(items)
	if err != nil {
		return fmt.Errorf("Error in get tbl : %s ", err)
	}
	if err := writeResult(links, quanFile); err != nil {
		return fmt.Errorf("Error in write links to file : %s ", err)
	}
	return nil
}

func StartTop() error {
	klog.Infof("execute url is : %s", topUrl)
	items, err := getTopItems(topUrl)
	if err != nil {
		return fmt.Errorf("Error in get Items : %s ", err)
	}
	links, err := getTblinks(items)
	if err != nil {
		return fmt.Errorf("Error in get tbl : %s ", err)
	}
	if err := writeResult(links, topFile); err != nil {
		return fmt.Errorf("Error in write links to file : %s ", err)
	}
	return nil
}

func getQuanItems(url string) ([]string, error) {
	var items []string
	dom, err := goquery.NewDocument(url)
	if err != nil {
		return nil, fmt.Errorf("Error in get html : %s ", err)
	}
	dom.Find(goodsListId).Children().Each(func(index int, content *goquery.Selection) {
		id, ok := content.Attr(quanDataId)
		if ok {
			items = append(items, id)
		}
	})
	return items, nil
}

func getTopItems(url string) ([]string, error) {
	var items []string
	dom, err := goquery.NewDocument(url)
	if err != nil {
		return nil, fmt.Errorf("Error in get html : %s ", err)
	}
	dom.Find(goodsListId).Children().Each(func(index int, content *goquery.Selection) {
		id, ok := content.Attr(topId)
		if ok {
			id = strings.Replace(id, "goods-items_", "", 1)
			items = append(items, id)
		}
	})
	return items, nil
}

func getTblinks(items []string) ([]string, error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	var links []string
	for _, i := range items {
		l, err := getTbl(i, now)
		if err != nil {
			continue
		}
		links = append(links, l)
	}
	if len(links) == 0 {
		return nil, fmt.Errorf("Error : the links is null ")
	}
	return links, nil
}

func getTbl(id string, timestamp string) (string, error) {
	url := fmt.Sprintf(getTblUrl, id, timestamp)
	dom, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}
	a := dom.Find("a")
	link := a.Nodes[1].Attr[1].Val
	return link, nil
}

func writeResult(links []string, filename string) error {
	os.Remove(filename)
	_, err := os.Create(filename)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, link := range links {
		data := append([]byte(link), '\n')
		//if err := ioutil.WriteFile(resultFile, data, 0755); err != nil {
		//	klog.Error(err)
		//}
		_, err := file.Write(data)
		if err != nil {
			klog.Error(err)
		}
	}
	return nil
}
