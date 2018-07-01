package xssfilter

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const (
	defNodeBlackList = "javascript,script"
	defAttWhiteList  = "style,class,id"
)

// Filter html过滤列表, 包含一个标签黑名单和一个属性白名单
type Filter struct {
	NodeBlackList []string
	AttrWhiteList map[string]struct{}
}

// NewFilter 通过nodes,attrs两个以","分割的字符串创建过滤器
func NewFilter(nodes, attrs string) *Filter {
	filter := &Filter{
		NodeBlackList: strings.Split(nodes, ","),
		AttrWhiteList: make(map[string]struct{}),
	}
	for _, attr := range strings.Split(attrs, ",") {
		filter.AttrWhiteList[attr] = struct{}{}
	}
	return filter
}

var defaultFilter = NewFilter(defNodeBlackList, defAttWhiteList)

// Clean 清理html的元素和属性
func (filter *Filter) Clean(str, root string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if err != nil {
		return "", errors.Wrap(err, "xssfilter")
	}
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, nodeName := range filter.NodeBlackList {
			// 如果s的名称等于nodeName,删除此节点
			if s.Is(nodeName) {
				s.Remove()
				return
			}
		}
		for _, node := range s.Nodes {
			for _, attr := range node.Attr {
				// 如果属性不存在于白名单, 则删除
				if _, ok := filter.AttrWhiteList[attr.Key]; !ok {
					s.RemoveAttr(attr.Key)
				}
			}
		}
	})
	if root == "" {
		return doc.Html()
	}
	return doc.Find(root).Html()
}
