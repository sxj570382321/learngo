package main

import "fmt"

type Books struct{
	title string
	author string
	subject string
	book_id int
}

func main(){
	var Book1,Book2 Books
	
	Book1.title = "go language"
	Book1.author = "www.runoob.com"
	Book1.subject = "go teaches book"
	Book1.book_id = 6495407

	Book2.title = "python language" 
        Book2.author = "www.runoob.com"
        Book2.subject = "python teaches book"
        Book2.book_id = 6495700

	fmt.Printf("Book1,title:%s,author:%s,subject:%s,book_id:%d\n",Book1.title,Book1.author,Book1.subject,Book1.book_id)

	fmt.Printf("Book2,title:%s,author:%s,subject:%s,book_id:%d\n",Book2.title,Book2.author,Book2.subject,Book2.book_id)

	fmt.Println("##########################")

	print(Book1)
	print(Book2)

	fmt.Println("##########################")

        print_ptr(&Book1)
        print_ptr(&Book2)
}

func print(book Books){
	fmt.Printf("Book,title:%s,author:%s,subject:%s,book_id:%d\n",book.title,book.author,book.subject,book.book_id)
}

func print_ptr(book *Books){
        fmt.Printf("Book,title:%s,author:%s,subject:%s,book_id:%d\n",book.title,book.author,book.subject,book.book_id)
}
