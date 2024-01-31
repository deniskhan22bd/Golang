package testdata

type Book struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	PublishedYear int `json:"published year"`
}

var Books = []Book{
	{Id: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", PublishedYear: 1925},
	{Id: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", PublishedYear: 1960},
	{Id: 3, Title: "1984", Author: "George Orwell", PublishedYear: 1949},
	{Id: 4, Title: "The Catcher in the Rye", Author: "J.D. Salinger", PublishedYear: 1951},
	{Id: 5, Title: "Harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", PublishedYear: 1997},
	{Id: 6, Title: "The Hobbit", Author: "J.R.R. Tolkien", PublishedYear: 1937},
	{Id: 7, Title: "Pride and Prejudice", Author: "Jane Austen", PublishedYear: 1813},
	{Id: 8, Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", PublishedYear: 1954},
	{Id: 9, Title: "One Hundred Years of Solitude", Author: "Gabriel García Márquez", PublishedYear: 1967},
	{Id: 10, Title: "The Da Vinci Code", Author: "Dan Brown", PublishedYear: 2003},
	{Id: 11, Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", PublishedYear: 1979},
	{Id: 12, Title: "Brave New World", Author: "Aldous Huxley", PublishedYear: 1932},
	{Id: 13, Title: "The Shining", Author: "Stephen King", PublishedYear: 1977},
	{Id: 14, Title: "The Alchemist", Author: "Paulo Coelho", PublishedYear: 1988},
	{Id: 15, Title: "The Hunger Games", Author: "Suzanne Collins", PublishedYear: 2008},
	{Id: 16, Title: "Moby-Dick", Author: "Herman Melville", PublishedYear: 1851},
	{Id: 17, Title: "The Chronicles of Narnia", Author: "C.S. Lewis", PublishedYear: 1950},
	{Id: 18, Title: "A Tale of Two Cities", Author: "Charles Dickens", PublishedYear: 1859},
	{Id: 19, Title: "The Road", Author: "Cormac McCarthy", PublishedYear: 2006},
	{Id: 20, Title: "The Girl with the Dragon Tattoo", Author: "Stieg Larsson", PublishedYear: 2005},
}