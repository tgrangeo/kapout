package entity

type Quiz struct {
    ID         string          `bson:"_id" json:"id"`
    Name       string          `bson:"name"`
    Questions  []QuizQuestion `bson:"questions"`
}

type QuizQuestion struct {
    ID      string   `bson:"id" json:"id"`
    Name    string   `bson:"name" json:"name"`
    Choices []Choice `bson:"choices" json:"choices"`
}

type Choice struct {
    ID      string `bson:"id" json:"id"`
    Name    string `bson:"name" json:"name"`
    Correct bool   `bson:"correct" json:"correct"`
}
