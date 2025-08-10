package apperrors

import "errors"

var ErrSubjectHasFinalExam = errors.New("Subject already has a final exam")
