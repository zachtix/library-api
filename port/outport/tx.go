package outport

type TxRepos struct {
	Members MemberRepository
	Books   BookRepository
	Loans   LoanRepository
	Fines   FineRepository
}

type TxManager interface {
	Tx(fn func(r TxRepos) error) error
}
