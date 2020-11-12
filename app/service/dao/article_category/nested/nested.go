package nested

import (
	. "gf-cms/app/model/article/article_category"
)

func isValidNode(n *ArticleCategory) bool {
	if n.ID == 0 || n.Lft == 0 || n.Rgt == 0 || n.Lft >= n.Rgt {
		return false
	}
	return true
}
