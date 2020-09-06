package imagecache

import (
	"bufio"
	"container/list"
	"fmt"
	"net/http"
	"os"
)

func Run() {
	//In Mem Cache
	var lruCache *LRUCache

	s := bufio.NewScanner(os.Stdin)
	var scannedLines, numberOfUrls int
	for s.Scan() {
		line := s.Text() //Get Line

		var cacheSize int
		if scannedLines == 0 {
			_, err := fmt.Sscanf(line, "%d", &cacheSize)
			if err != nil {
				fmt.Errorf("something went wrong %v", err)
				break
			}
			lruCache = new(LRUCache)
			lruCache.cacheSize = cacheSize
			lruCache.keys = make(map[string]*list.Element)
			lruCache.values = list.New()
			scannedLines++
			continue
		} else if scannedLines == 1 {
			_, err := fmt.Sscanf(line, "%d", &numberOfUrls)
			if err != nil {
				fmt.Errorf("something went wrong %v", err)
				continue
			}
			scannedLines++
			continue
		}

		var urlRequested string
		_, err := fmt.Sscanf(line, "%s", &urlRequested)
		if err != nil {
			fmt.Errorf("something went wrong %v", err)
			break
		}

		inCache, sizeOfImage := lruCache.Get(urlRequested)

		var isDownloadedStr string
		if inCache {
			isDownloadedStr = "IN_CACHE"
		} else {
			isDownloadedStr = "DOWNLOADED"
		}
		fmt.Printf("%s %s %d\n", urlRequested, isDownloadedStr, sizeOfImage)
		scannedLines++
	}
}

type LRUCache struct {
	cacheSize   int
	currentSize int
	keys        map[string]*list.Element
	values      *list.List
}

type Image struct {
	data []byte
	url  string
}

type ImagePtr *Image

func (c *LRUCache) Get(requested string) (bool, int) {
	value := c.keys[requested]
	if value != nil {
		c.values.Remove(value)
		c.values.PushFront(value.Value)
		return true, len(value.Value.(ImagePtr).data)
	}

	image, err := c.downloadFile(requested)
	if err != nil {

	}
	if len(image.data) > c.cacheSize {
		return false, len(image.data)
	}

	for len(image.data)+c.currentSize > c.cacheSize {
		val := c.values.Back()
		c.values.Remove(val)
		delete(c.keys, val.Value.(*Image).url)
		c.currentSize -= len(val.Value.(*Image).data)
	}

	c.values.PushFront(image)
	c.keys[image.url] = c.values.Front()
	c.currentSize += len(image.data)
	return false, len(image.data)
}

func (c *LRUCache) downloadFile(requested string) (ImagePtr, error) {
	response, err := http.Get(requested)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	image := &Image{
		data: make([]byte, 2*(c.cacheSize+1)),
		url:  requested,
	}
	readBytes, err := response.Body.Read(image.data)

	if err != nil {
		return nil, err
	}

	image.data = image.data[0:readBytes]
	return image, nil
}
