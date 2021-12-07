package logic

type Block struct {
    Try     func()
    Catch   func(Exception)
    Finally func()
}

type Exception interface{}

func Throw(ex Exception) {
    panic(ex)
}

func (that Block) Do() {
    if that.Finally != nil {

        defer that.Finally()
    }
    if that.Catch != nil {
        defer func() {
            if x := recover(); x != nil {
                that.Catch(x)
            }
        }()
    }
    that.Try()
}