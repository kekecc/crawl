package info

import (
	"regexp"
	"reptile/handle"
)

const express = `<a href="([^"]+)" class="tag">([^"]+)</a>` //匹配方式

func ParseContent(buf []byte) handle.ParseRes {
	exp := regexp.MustCompile(express)
	matches := exp.FindAllSubmatch(buf, -1)

	// for _, content := range matches {
	// 	fmt.Println("content: ", content[0])
	// }

	result := handle.ParseRes{}
	for _, content := range matches {
		result.Contents = append(result.Contents, content[0])
		result.Requests = append(result.Requests, handle.Request{
			Url:       "https://book.douban.com" + string(content[1]),
			ParseFunc: ParseBookList,
		})
	}
	return result
}

func ParseBookList(buf []byte) handle.ParseRes {
	exp := regexp.MustCompile(`<a href="([^"]+)" title="([^"]+)"`)
	matches := exp.FindAllSubmatch(buf, -1)

	result := handle.ParseRes{}
	for _, content := range matches {
		result.Contents = append(result.Contents, content[2])
		result.Requests = append(result.Requests, handle.Request{
			Url:       string(content[1]),
			ParseFunc: ParseNil,
		})
	}
	return result
}

func ParseNil(buf []byte) handle.ParseRes {
	return handle.ParseRes{}
}
