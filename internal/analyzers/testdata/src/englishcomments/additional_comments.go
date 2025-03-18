package englishcomments

// Empty comment
func EmptyComment() {}

// This is a good doc comment for a type
type DocumentedType struct {
	// This is a good field comment
	Field string
}

/* This is a good multi-line
 * English comment with stars
 * on each line
 */
func StarredMultilineComment() {}

// Це документація українською мовою
type BadDocumented struct { // want "comment should be in English"
	// Поле с русским комментарием
	Field string // want "comment should be in English"
}

/*
Многоредов коментар
на български език
тук пише нещо важно
*/
func MultilineWithoutStars() {} // want "comment should be in English"

/* Однорядковий коментар українською */
func SingleLineBlock() {} // want "comment should be in English"

func doSomething()     {}
func doSomethingElse() {}

func MixedComments() {
	// This is fine
	doSomething()
	// Това е лош коментар
	doSomethingElse() // want "comment should be in English"
}

// This comment has mixed languages: English and 中文
func MixedLanguages() {} // want "comment should be in English"

/*
 * Коментар на декілька
 * рядків українською
 * мовою з зірочками
 */
func MultilineStarredNonEnglish() {} // want "comment should be in English"

// Comment with numbers and symbols: 123 !@#$%^&*()
func SpecialCharacters() {}

// TODO: this is a valid English comment
func TodoComment() {}

// FIXME: Моля, оправете този код
func FixmeComment() {} // want "comment should be in English"
