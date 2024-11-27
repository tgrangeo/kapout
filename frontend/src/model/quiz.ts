export interface QuizChoice {
  id: string;
  name: string;
  correct: boolean;
}

export interface Pplayer{
	id:string;
	name:string;
}

export interface QuizQuestion {
  id: string;
  name: string;
  choices: QuizChoice[];
}

export interface Quiz {
  id: string;
  Name: string;
  questions: QuizQuestion[];
}
