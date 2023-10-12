NAME=bo

inst: $(NAME)
	sudo mv $(NAME) /usr/local/bin
	$(NAME)

$(NAME): main.go
	go build -o $(NAME) .