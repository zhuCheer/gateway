package model

import (
	"database/sql"
	"fmt"
)

type Nodes struct {
	ID        int64
	SiteId    int64
	Domain    string
	Addr      string
	Weight    uint32
	CreatedAt string
	UpdatedAt string
}

type LocalItem struct {
	Addr      string
	Weight    uint32
	CreatedAt string
	UpdatedAt string
}

func (n *Nodes) QueryOneById(id int) *Nodes {
	row := DB.QueryRow("select qi_nodes.id,qi_sites.id,qi_sites.domain,qi_nodes.addr, qi_nodes.weight, "+
		"qi_nodes.created_at, qi_nodes.updated_at"+
		" from qi_nodes LEFT JOIN qi_sites ON qi_sites.id=qi_nodes.site_id where qi_nodes.id=?", id)
	if err := row.Scan(&n.ID, &n.SiteId, &n.Domain, &n.Addr, &n.Weight, &n.CreatedAt, &n.UpdatedAt); err != nil && err != sql.ErrNoRows {
		fmt.Printf("Nodes->QueryOneById scan failed, err:%v", err)
		return nil
	}
	return n
}

func (n *Nodes) QueryListBySiteId(siteId int64) (list []*LocalItem) {
	rows, err := DB.Query("select `addr`, `weight`, `created_at`, `updated_at` FROM qi_nodes where site_id=?", siteId)

	if err != nil {
		fmt.Printf("Nodes->QueryListBySiteId exec failed, err:%v \n", err)
		return nil
	}

	for rows.Next() {
		var item LocalItem
		err := rows.Scan(&item.Addr, &item.Weight, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			fmt.Printf("QueryListBySiteId rows fail err:%v \n", err)
			continue
		}
		list = append(list, &item)
	}

	return list
}
