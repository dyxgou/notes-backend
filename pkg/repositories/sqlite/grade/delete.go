package grade

func (r *Repository) Delete(gradeId, subjectId int64) (int64, error) {
	q1 := "DELETE FROM grade WHERE id = ?;"
	q2 := `UPDATE subject AS s
	SET
	  grades = grades - 1,
	  has_final_exam = CASE g.is_final_exam
	    WHEN TRUE THEN FALSE
	    WHEN FALSE THEN FALSE
	  END
	FROM
	  grade AS g
	WHERE
	  s.id = ?;`

	tx, err := r.Db.Begin()
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec(q1, gradeId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}

		return 0, err
	}

	_, err = tx.Exec(q2, subjectId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}

		return 0, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return rows, nil
}
