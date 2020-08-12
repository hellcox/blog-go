package util

import "strconv"

func Page(page int, limit int, count int) string {
	pageUp := 0
	pageDown := 0
	if page > 1 {
		pageUp = 1
	}
	if count < limit*page {
		pageDown = 1
	}
	pageInfo := "<ul class=\"pager\">"
	if pageUp == 1 {
		pageInfo += "<li><a href=\"/p/" + strconv.Itoa(page-1) + "\"><i class=\"fa fa-chevron-left\"></i></a></li>"
	} else {
		pageInfo += "<li class=\"disabled\"><span><i class=\"fa fa-chevron-left\"></i></span></li> "
	}
	if pageDown == 1 {
		pageInfo += "<li class=\"disabled\"><span><i class=\"fa fa-chevron-right\"></i></span></li> "
	} else {
		pageInfo += "<li><a href=\"/p/" + strconv.Itoa(page+1) + "\"><i class=\"fa fa-chevron-right\"></i></a></li>"
	}
	pageInfo += "</ul>"
	return pageInfo
}
