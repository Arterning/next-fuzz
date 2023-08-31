package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	log.Println("Starting simplefts")

	var docs []document
	_, err := os.Stat("doc.json")
	if os.IsNotExist(err) {
		//load documents
		start := time.Now()
		docs, err = loadDocuments(dumpPath)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

		//store the docs
		err = saveDocAsJSON(docs, "doc.json")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		docs, err = loadDocsFromJSON("doc.json")
		if err != nil {
			log.Fatal(err)
		}
	}

	idx := make(index)
	_, err = os.Stat("index.json")
	if os.IsNotExist(err) {
		//make index
		start := time.Now()
		idx.add(docs)
		log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

		//store the idx
		err = saveIndexAsJSON(idx, "index.json")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		idx, err = loadIndexFromJSON("index.json")
		if err != nil {
			log.Fatal(err)
		}
	}

	//do search
	doSearch(idx, query, docs)
}

func doSearch(idx index, query string, docs []document) {
	start := time.Now()
	matchedIDs := idx.search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}

func saveDocAsJSON(docs []document, filename string) error {
	// 将文档数组序列化为 JSON 格式
	jsonData, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		return err
	}

	// 将 JSON 数据写入文件
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadDocsFromJSON(filename string) ([]document, error) {
	// 从文件中读取 JSON 数据
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 反序列化 JSON 数据到 document 结构体数组
	var docs []document
	err = json.Unmarshal(jsonData, &docs)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func saveIndexAsJSON(idx index, filename string) error {
	// 将倒排索引序列化为 JSON 格式
	jsonData, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return err
	}

	// 将 JSON 数据写入文件
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadIndexFromJSON(filename string) (index, error) {
	// 从文件中读取 JSON 数据
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 反序列化 JSON 数据到 index 类型的倒排索引
	var idx index
	err = json.Unmarshal(jsonData, &idx)
	if err != nil {
		return nil, err
	}

	return idx, nil
}
