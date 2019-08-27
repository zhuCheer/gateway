package model

import (
	"database/sql"
	"fmt"
)

type Sites struct {
	ID        int64
	Name      string
	Scheme    string
	Domain    string
	Balance   string
	CreatedAt string
	UpdatedAt string
}

func (s *Sites) QueryOneById(id int) *Sites {
	row := DB.QueryRow("select `id`, `name`, `scheme`,`domain`, `balance`, `created_at`, `updated_at` from qi_sites where id=?", id)

	if err := row.Scan(&s.ID, &s.Name, &s.Scheme, &s.Domain, &s.Balance, &s.CreatedAt, &s.UpdatedAt); err != nil && err != sql.ErrNoRows {
		fmt.Printf("Sites->QueryOneById scan failed, err:%v", err)
		return nil
	}
	return s
}

func (s *Sites) QueryOneByDomain(domain string) *Sites {
	row := DB.QueryRow("select `id`, `name`, `scheme`,`domain`, `balance`, `created_at`, `updated_at` from qi_sites where domain=?", domain)

	if err := row.Scan(&s.ID, &s.Name, &s.Scheme, &s.Domain, &s.Balance, &s.CreatedAt, &s.UpdatedAt); err != nil && err != sql.ErrNoRows {
		fmt.Printf("Sites->QueryOneByDomain scan failed, err:%v", err)
		return nil
	}
	return s
}

func (s *Sites) GetAll() (list []*Sites) {
	rows, err := DB.Query("select `id`, `name`, `scheme`,`domain`, `balance`, `created_at`, `updated_at` from qi_sites where 1=1 LIMIT 100")
	if err != nil {
		fmt.Printf("Sites->GetAll exec failed, err:%v \n", err)
		return nil
	}

	for rows.Next() {
		var item Sites
		err := rows.Scan(&item.ID, &item.Name, &item.Scheme, &item.Domain, &item.Balance, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			fmt.Printf("GetAll rows fail err:%v \n", err)
			continue
		}
		list = append(list, &item)
	}
	return list
}
