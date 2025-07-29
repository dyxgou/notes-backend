package grade

func (r *Repository) ChangeName(id int64, name string) error {
	query := "UPDATE grade SET name = ? WHERE id = ?;"

	_, err := r.Db.Exec(query, name, id)
	if err != nil {
		return err
	}

	return nil
}
