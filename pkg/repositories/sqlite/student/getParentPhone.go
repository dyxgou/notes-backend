package student

func (r *Repository) GetParentPhone(id int64) (string, error) {
	q := "SELECT parent_phone FROM student WHERE id = ?;"

	var tel string

	if err := r.Db.QueryRow(q, id).Scan(&tel); err != nil {
		return tel, err
	}

	return tel, nil
}
