package eval

// An Expr is an arithmetic expression
type Expr interface{
  // Eval returns the value of this Expr in the given env
  Eval(env Env) float64
}

// A var identifies a variable eg., x
type Var string

// A literal is numeric constant
type literal float64

// A unary represents a unary operator expression, eg., -x
type unary struct {
  op rune  // one of '+', '-'
  x Expr
}

// A binary represents a binary operator expression, eg., x+y
type binary struct {
  op rune  // one of '+', '-', '*', '/'
  x, y Expr
}

// A call represents a function call expression, eg., sin(x)
type call struct {
  fn string  // one of 'sin', 'pow', 'sqrt'
  args []Expr
}
