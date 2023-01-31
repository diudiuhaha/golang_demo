package main

import "sync"

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex //读写互斥锁
}

// 实例化结构体（主要是初始化map）
func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)} //返回的是指向实例的指针
}

// 读的请求的方法，返回长url
func (s *URLStore) Get(key string) string { //这种写法学习了
	s.mu.RLock()         //上锁
	defer s.mu.RUnlock() //优化1，避免忘记解锁，以及快速的去追踪
	return s.urls[key]   //优化2，消除冗余的变量
}

// 写的请求的方法,返回该方法是否操作成果
func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock() //优化1
	_, p := s.urls[key]
	if p {
		return false
	} else {
		//如果不存在，便可以去写入数据(长url)
		s.urls[key] = url
		return true
	}
}

// 获取键值对的数量
func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

// 通过长URL生产短URL
func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count()) //调用算法生成一个键
		if ok := s.Set(key, url); ok {
			//一直去尝试写入，直到一个没有被使用过的url
			return key //返回这个键
		}
	}
}
