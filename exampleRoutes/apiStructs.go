package exampleRoutes

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type Options struct {
	port int `help:"Port to listen on" short:"p" default:"8888"`
}

type ReviewInput struct {
	Body struct {
		Author  string `json:"author" maxLength:"10" doc:"Author of the review" example:"Mickey"`
		Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5" example:"5"`
		Message string `json:"message,omitempty" maxLength:"100" doc:"Review message" example:"Review"`
	}
}

type StringOutput struct {
	Body struct {
		Word string
	}
}

type SumNumbersStruct struct {
	Body struct {
		Num1 int `json:"number1" example:"1"`
		Num2 int `json:"number2" example:"2"`
	}
}

type SumNumbersOutput struct {
	Sum int
}
