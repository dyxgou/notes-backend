package student

func (r *Repository) GetStudentAverage(studentId, subjectId int64) (float64, error) {
	query := `
SELECT
  COALESCE(
    ROUND(
      AVG(
        CASE
          WHEN g.is_final_exam = FALSE THEN n.value * 0.7
        END
      ) + SUM(
        CASE
          WHEN g.is_final_exam = TRUE THEN n.value * 0.3
          ELSE 0
        END
      )
    ),
    10.0
  ) AS student_average
FROM
  grade g
  JOIN note n ON n.grade_id = g.id
  AND n.student_id = ?
WHERE
  subject_id = ?;`

	var studentAverage float64
	err := r.Db.QueryRow(query, studentId, subjectId).Scan(&studentAverage)
	if err != nil {
		return 0.0, err
	}

	return studentAverage, nil
}
