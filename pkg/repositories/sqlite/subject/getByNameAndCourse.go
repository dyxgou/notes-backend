package subject

func (r *Repository) GetByNameAndCourse(name string, course, period byte) (int64, error) {
	query := "SELECT id FROM subject WHERE name = ? AND course = ? AND period = ?;"

	row := r.Db.QueryRow(query, name, course, period)

	var id int64

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
