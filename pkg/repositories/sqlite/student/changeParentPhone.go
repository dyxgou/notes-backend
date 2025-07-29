package student

func (r *Repository) ChangeParentPhone(id int64, phone string) (int64, error) {
	query := "UPDATE student SET parent_phone = ? WHERE id = ?;"

	res, err := r.Db.Exec(query, phone, id)

	if err != nil {
		return 0, err
	}

	affectedId, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedId, err
}
