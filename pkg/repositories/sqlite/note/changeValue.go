package note

func (r *Repository) ChangeValue(id int64, value byte) error {
	query := "UPDATE note SET value = ? WHERE id = ?;"

	_, err := r.Db.Exec(query, value, id)

	return err
}
