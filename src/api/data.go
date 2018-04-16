package api

// Mock Data - @TODO implement DB
func loadMockData() {
	books = append(books, Book{
		ID: "1",
		Isbn: "44874",
		Title: "To Kill a Mocking Bird",
		Author: &Author{
			Firstname: "Harper",
			Lastname: "Lee",
		},
	})

	books = append(books, Book{
		ID: "2",
		Isbn: "45723",
		Title: "Pride and Prejudice",
		Author: &Author{
			Firstname: "Jane",
			Lastname: "Austin",
		},
	})
}
