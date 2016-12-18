package main

type UnreadCounts struct {
	Unreadcounts []UnreadCount
}
type UnreadCount struct {
	Updated int
	Count int
	Id string
}
