package englishcomments

// This is a good English comment
func GoodComment() {}

// Ова је коментар на српском језику
func BadComment() {} // want "comment should be in English"

// 这是一个中文注释
func ChineseComment() {} // want "comment should be in English"

// Multiple lines of English comments
// describing what this function does
// in great detail
func MultilineEnglishComments() {}

/* Бұл қазақ тіліндегі көп жолды түсініктеме */
func MultilineNonEnglish() {} // want "comment should be in English"

// Special cases that should be ignored:
//
//go:generate echo "test"
//nolint:all
//revive:disable
func SpecialComments() {}
