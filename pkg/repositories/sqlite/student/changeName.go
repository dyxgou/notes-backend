package student

func (r *Repository) ChangeName(id int64, name string) (int64, error) {
	query := "UPDATE student SET name = ? WHERE id = ?;"

	res, err := r.Db.Exec(query, name, id)

	if err != nil {
		return 0, err
	}

	affectedId, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedId, err
}
