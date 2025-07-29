package student

func (r *Repository) Delete(id int64) (int64, error) {
	query := "DELETE FROM student WHERE id = ?;"

	res, err := r.Db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	studentId, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return studentId, err
}
