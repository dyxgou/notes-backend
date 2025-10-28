package report

func (r *Repository) GetSubjectReport(
	studentId int64,
	name string,
	course byte,
) ([]float64, error) {
	q := `
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
  ) AS averages
FROM
  subject s
  JOIN grade g ON g.subject_id = s.id
  JOIN note n ON n.grade_id = g.id
  AND n.student_id = ?
WHERE
  s.name = ?
  AND s.course = ?
GROUP BY
  s.id
ORDER BY
  s.period`

	averages := make([]float64, 0, 4)

	rows, err := r.Db.Query(q, studentId, name, course)
	if err != nil {
		return averages, err
	}
	defer rows.Close()

	for rows.Next() {
		var avg float64

		if err := rows.Scan(&avg); err != nil {
			return averages, err
		}

		averages = append(averages, avg)
	}

	return averages, nil
}
