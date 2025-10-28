package report

func (r *Repository) GetSubjectAverage(
	studentId int64,
	name string,
	course byte,
) (float64, error) {
	q := `
SELECT
  COALESCE(ROUND(AVG(averages)), 0.0) as subject_average
FROM
  (
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
      ) as averages
    FROM
      subject s
      LEFT JOIN grade g ON g.subject_id = s.id
      LEFT JOIN note n ON n.grade_id = g.id
      AND n.student_id = ?
    WHERE
      s.name = ?
      AND s.course = ?
    GROUP BY
      s.id
    ORDER BY
      s.period
  )`

	var subjectAverage float64

	err := r.Db.QueryRow(q, studentId, name, course).Scan(&subjectAverage)
	if err != nil {
		return 0.0, err
	}

	return subjectAverage, nil
}
